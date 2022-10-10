# Google Authentication

```go
router := mux.NewRouter()

key := "my-secret-key"
sessionStorage := sessions.NewCookieStore([]byte(key))
sessionStorage.MaxAge(86400 * 2)
sessionStorage.Options.Path = "/"
sessionStorage.Options.HttpOnly = true
sessionStorage.Options.Secure = false

googleAuthConfig := googleauth.GoogleAuthConfig{
  SessionAuthConfig: auth.SessionAuthConfig{
    AuthFailedHandler: func(w http.ResponseWriter, r *http.Request, err error) {
      nerdweb.WriteJSON(logger, w, http.StatusUnauthorized, map[string]interface{}{
        "success": false,
        "error":   err.Error(),
      })
    },
    AuthSuccessHandler: func(w http.ResponseWriter, r *http.Request, user goth.User) {
      logger.WithField("user", user).Info("Successful login")
      http.Redirect(w, r, "/view-logs", http.StatusTemporaryRedirect)
    },
    ErrorPath:         "/unauthorized",
    ExcludedPaths:     []string{"/", "/unauthorized", "/static", "/auth", "/main.js", "/index.html", "/version"},
    HTMLResponsePaths: []string{"/view-logs", "/manage-servers", "/edit-server"},
    SessionName:       "fireplacelogging",
    Store:             sessionStorage,
  },
  GoogleClientID:     config.GoogleClientID,
  GoogleClientSecret: config.GoogleClientSecret,
  GoogleRedirectURI:  config.GoogleRedirectURI,
}

googleauth.Setup(router, googleAuthConfig, logger)
```
