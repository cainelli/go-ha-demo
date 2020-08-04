package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	wrapperHTTP "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"

	log "github.com/sirupsen/logrus"
)

// IndexHandler handles /
func (s *Server) IndexHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Expose-Headers", "x-some-header")

	time.Sleep(time.Duration(s.Config.Delay) * time.Second)

	if !s.Config.Healthy {
		time.Sleep(120 * time.Second)
	}

	url := r.URL.Query().Get("call")
	if url != "" {
		client := http.Client{}
		wrapperHTTP.WrapRoundTripper()
		wrapClient := wrapperHTTP.WrapClient(&client)
		res, err := wrapClient.Get(url)

		if err != nil {
			log.Error(err)
			return
		}
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Error(err)
			return
		}

		w.Write(bodyBytes)
	}

	hostname, _ := os.Hostname()

	w.Write([]byte(fmt.Sprintf("------------------------\nFrom server: %s\n", hostname)))
	for k, v := range r.Header {
		header := fmt.Sprintf("%s: %v\n", k, v)
		w.Write([]byte(header))
	}

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

// ReadinessHandler is the readiness check endpoint handler
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
