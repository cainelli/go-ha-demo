package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SetState sets the state of server.Config
func (s *Server) SetState(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(body, s.Config)
	if err != nil {
		log.Error(err)
	}

	data, err := json.Marshal(s.Config)
	if err != nil {
		log.Error(err)
	}

	_, err = w.Write(data)
	if err != nil {
		log.Error(err)
	}
}
