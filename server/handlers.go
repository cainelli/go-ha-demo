package server

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// IndexHandler handles /
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Duration(s.Config.Delay) * time.Second)

	if !s.Config.Healthy {
		time.Sleep(120 * time.Second)
	}

	w.Write([]byte("Hello World!"))

	log.Info("host: ", r.Host, " uri: ", r.RequestURI, " status: ", 200)

}

// HealthCheckHandler is the health check endpoint handler
func (s *Server) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(s.Config.Delay) * time.Second)
	status := 200
	if !s.Config.Healthy {
		status = 500
	}
	w.WriteHeader(status)
	log.Info("host: ", r.Host, " uri: ", r.RequestURI, " status: ", status)

}

// ReadinessHandler is the health check endpoint handler
func (s *Server) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(s.Config.Delay) * time.Second)

	status := 200
	if !s.Config.Ready {
		if !s.Config.WarmingUp {
			go func(s *Server) {
				log.Info("warming up func")
				s.Config.WarmingUp = true
				time.Sleep(10 * time.Second)
				s.Config.Ready = true
			}(s)
		}
		status = 500
	}
	w.WriteHeader(status)
	log.Info("host: ", r.Host, " uri: ", r.RequestURI, " status: ", status)
}
