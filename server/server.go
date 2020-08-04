package server

import (
	"github.com/cainelli/go-ha-demo/config"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
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

	// ctrl endpoints
	server.Router.HandleFunc("/_set", server.SetState)

	// service endpoint
	server.Router.HandleFunc("/", server.IndexHandler)
}
