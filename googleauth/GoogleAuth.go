package googleauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/app-nerds/kit/v6/auth"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sirupsen/logrus"
)

type GoogleAuthConfig struct {
	auth.SessionAuthConfig

	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURI  string
}

func Setup(router *mux.Router, config GoogleAuthConfig, logger *logrus.Entry) {
	gothic.Store = config.Store

	goth.UseProviders(
		google.New(config.GoogleClientID, config.GoogleClientSecret, config.GoogleRedirectURI, "email", "profile"),
	)

	router.HandleFunc("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		var (
			err     error
			user    goth.User
			session *sessions.Session
		)

		user, err = gothic.CompleteUserAuth(w, r)

		if err != nil {
			config.AuthFailedHandler(w, r, err)
			return
		}

		if session, err = config.Store.Get(r, config.SessionName); err != nil {
			logger.WithError(err).Error("error geting session")
			http.Redirect(w, r, config.ErrorPath, http.StatusTemporaryRedirect)
			return
		}

		session.Values["email"] = user.Email
		session.Values["firstName"] = user.FirstName
		session.Values["lastName"] = user.LastName
		session.Values["avatarURL"] = user.AvatarURL
		session.Values["approved"] = false

		if err = config.Store.Save(r, w, session); err != nil {
			logger.WithError(err).Error("error saving session")
			http.Redirect(w, r, config.ErrorPath, http.StatusTemporaryRedirect)
			return
		}

		config.AuthSuccessHandler(w, r, user)
	})

	router.HandleFunc("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	})

	setupMiddleware(router, config, logger)
}

func setupMiddleware(router *mux.Router, config GoogleAuthConfig, logger *logrus.Entry) {
	middleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				err     error
				session *sessions.Session
				ok      bool

				email string
			)

			/*
			 * If this path is excluded from auth, just keep going
			 */
			for _, excludedPath := range config.ExcludedPaths {
				if excludedPath == "/" && r.URL.Path == "/" {
					next.ServeHTTP(w, r)
					return
				}

				if strings.HasPrefix(r.URL.Path, excludedPath) && excludedPath != "/" {
					next.ServeHTTP(w, r)
					return
				}
			}

			/*
			 * If not, let's verify we have a cookie
			 */
			if session, err = config.Store.Get(r, config.SessionName); err != nil {
				logger.WithError(err).Error("error getting session information")
				http.Redirect(w, r, config.ErrorPath, http.StatusOK)
				return
			}

			email, ok = session.Values["email"].(string)

			if !ok {
				sendResponse(w, r, config)
				return
			}

			if email == "" {
				sendResponse(w, r, config)
				return
			}

			approved, _ := session.Values["approved"].(bool)

			if !approved {
				http.Redirect(w, r, config.UnapprovedPath, http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	router.Use(middleware)
}

func sendResponse(w http.ResponseWriter, r *http.Request, config GoogleAuthConfig) {
	for _, path := range config.HTMLResponsePaths {
		if strings.HasPrefix(r.URL.Path, path) {
			http.Redirect(w, r, config.ErrorPath, http.StatusTemporaryRedirect)
			return
		}
	}

	result := map[string]interface{}{
		"success": false,
		"error":   "User unauthorized",
	}

	b, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, b)
}
