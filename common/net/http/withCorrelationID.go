package http

import (
	"net/http"

	gid "github.com/google/uuid"
)

// WithCorrelationID creates a correlation id
func WithCorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := gid.New().String()
		w.Header().Add(headerCorrelationID, cid)
		next.ServeHTTP(w, r)
	})
}
