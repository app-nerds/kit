package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
)

type SessionAuthConfig struct {
	AuthFailedHandler  func(w http.ResponseWriter, r *http.Request, err error)
	AuthSuccessHandler func(w http.ResponseWriter, r *http.Request, user goth.User)
	ErrorPath          string
	ExcludedPaths      []string
	HTMLResponsePaths  []string
	SessionName        string
	Store              sessions.Store
	UnapprovedPath     string
}
