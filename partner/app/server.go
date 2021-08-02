package app

import (
	"log"
	"net/http"

	"github.com/Ralphbaer/ze-delivery/common"
)

// Server represents the http server for partner service
type Server struct {
	httpServer *http.Server
}

// NewServer creates an instance of Server
func NewServer(cfg *Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    cfg.ServerAddress,
			Handler: handler,
		},
	}
}

// Run runs the server
func (a *Server) Run(l *common.Launcher) error {
	log.Printf("Server started listen on %s\n", a.httpServer.Addr)
	return a.httpServer.ListenAndServe()
}
