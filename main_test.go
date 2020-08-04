package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cainelli/go-ha-demo/config"
	"github.com/cainelli/go-ha-demo/server"
	log "github.com/sirupsen/logrus"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func TestAPM(t *testing.T) {
	os.Setenv("DD_PROPAGATION_STYLE_EXTRACT", "B3")
	os.Setenv("DD_PROPAGATION_STYLE_INJECT", "B3")
	tracer.Start(
		tracer.WithServiceName(os.Getenv("APM_APPLICATION")),
		tracer.WithPropagator(tracer.NewPropagator(&tracer.PropagatorConfig{})),
	)

	defer tracer.Stop()

	conf := config.Load()

	srv := &server.Server{Router: mux.NewRouter(), Config: conf}

	srv.Init()

	server := httptest.NewServer(srv.Router)

	url := fmt.Sprintf("%s?call=%s", server.URL, server.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
		return
	}

	req.Header.Add("ot-baggage-e2e", `true`)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
