package http

import (
	"log"
	"net/http"
)

// Ping returns HTTP Status 200 with response "pong"
func Ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Print(err.Error())
	}
}

// File servers a specific file
func File(filePath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})
}
