package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/getyourguide/simple-http-server/internal/pkg/config"
	"github.com/getyourguide/simple-http-server/internal/pkg/server"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"net/http"
	"time"
)

func main() {
	log.Info("Initializing")

	var httpSrv *http.Server

	go func() {
		conf := config.New()

		srv := &server.Server{Router: mux.NewRouter(), Config: conf}

		srv.Init()

		httpSrv = &http.Server{
			Handler:      srv.Router,
			Addr:         "0.0.0.0:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}

		log.Info("Server started")
		if err := httpSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Error(err, "Shutdown failed")
	} else {
		log.Info("Server stopped")
	}
}
