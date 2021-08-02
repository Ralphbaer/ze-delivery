package http

import (
	"crypto/subtle"
	"net/http"
)

// BasicAuthFunc represents a func which returns if a username and password was authenticated or not.
// It's return a (realm, true) if authenticated, whereas when not authenticated (nil, and false).
type BasicAuthFunc func(username, password string) bool

// FixedBasicAuthFunc is a fixed username and password to use as BasicAuthFunc
func FixedBasicAuthFunc(username, password string) BasicAuthFunc {
	return func(user, pass string) bool {
		if subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			return true
		}
		return false
	}
}

// WithBasicAuth creates a basic authentication middleware
func WithBasicAuth(f BasicAuthFunc, realm string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()

			if ok {
				if ok := f(user, pass); !ok {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			JSONResponse(w, 401, ResponseError{
				Code:    401,
				Message: "Unauthorized request",
			})
		})
	}
}
