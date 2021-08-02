package http

import (
	"log"
	"net/http"

	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/gorilla/mux"
)

const (
	defaultAccessControlAllowOrigin  = "*"
	defaultAccessControlAllowMethods = "POST, GET, OPTIONS, PUT, DELETE, PATCH"
	defaultAccessControlAllowHeaders = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Total-Count"
	defaultAccessControlExposeHeaders = "X-Total-Count"
)

// WithCORS register a middleware with public global CORS. Use env vars to override it:
// Variables: ACCESS_CONTROL_ALLOW_ORIGIN, ACCESS_CONTROL_ALLOW_METHODS and ACCESS_CONTROL_ALLOW_HEADERS
func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println("WithCORS middleware")

		w.Header().Set("Access-Control-Allow-Origin", common.GetenvOrDefault("ACCESS_CONTROL_ALLOW_ORIGIN", defaultAccessControlAllowOrigin))
		w.Header().Set("Access-Control-Allow-Methods", common.GetenvOrDefault("ACCESS_CONTROL_ALLOW_METHODS", defaultAccessControlAllowMethods))
		w.Header().Set("Access-Control-Allow-Headers", common.GetenvOrDefault("ACCESS_CONTROL_ALLOW_HEADERS", defaultAccessControlAllowHeaders))
		w.Header().Set("Access-Control-Expose-Headers", common.GetenvOrDefault("ACCESS_CONTROL_EXPOSE_HEADERS", defaultAccessControlExposeHeaders))

		log.Println("WithCORS completed")

		next.ServeHTTP(w, r)

	})
}

// AllowFullOptionsWithCORS set r.Use(WithCORS) and allow every request to use OPTION method
func AllowFullOptionsWithCORS(r *mux.Router) {
	r.Use(WithCORS)
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}