package http

import (
	"log"
	"net/http"
	"os"
	"time"
)


// Ping returns HTTP Status 200 with response "pong"
func Ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("pong"))
	if err != nil {
		log.Print(err.Error())
	}
}

// Version returns HTTP Status 200 with given version
func Version(version string) http.HandlerFunc {
	 return func(w http.ResponseWriter, r *http.Request) {
		 OK(w, struct {
			 Version     string    `json:"version"`
			 BuildNumber string    `json:"buildNumber"`
			 RequestDate time.Time `json:"requestDate"`
		 }{
			 Version:     version,
			 BuildNumber: os.Getenv("BUILD_NUMBER"),
			 RequestDate: time.Now().UTC(),
		 })
	}
}

// Welcome returns HTTP Status 200 with service info
func Welcome(service string, description string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		info := struct {
			Service     string
			Description string
		}{
			Service:     service,
			Description: description,
		}

		JSONResponse(w, 200, info)
	}
}

// NotImplementedEndpoint returns HTTP 501 with not implemented message
func NotImplementedEndpoint(w http.ResponseWriter, r *http.Request) {
	NotImplemented(w, "Not implemented yet")
}

// File servers a specific file
func File(filePath string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filePath)
	})
}
