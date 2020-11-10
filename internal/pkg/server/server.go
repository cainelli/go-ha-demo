package server

import (
	"github.com/getyourguide/simple-http-server/internal/pkg/config"
	"github.com/gorilla/mux"
)

// Server is where external dependencies live. This makes it easy to mock them in unit tests and keeps the code organized.
type Server struct {
	Router *mux.Router
	Config *config.Config
}

// Init the HTTP endpoints
func (server *Server) Init() {

	// lifecycle endpoints
	server.Router.HandleFunc("/_health", server.HealthCheckHandler)
	server.Router.HandleFunc("/_ready", server.ReadinessHandler)

	// service endpoint
	server.Router.HandleFunc("/", server.IndexHandler)
}
