package main

import (
	"context"
	"github.com/dicaormu/go-xke-jan-2020/internal/application"
	"github.com/dicaormu/go-xke-jan-2020/internal/diagnostics"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.Info("stating application...")

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error)

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("port not defined")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", application.HomeHandler(logger))
	router.HandleFunc("/readyz", diagnostics.ReadinessHandler(logger))
	router.HandleFunc("/healtz", diagnostics.LivenessHandler(logger))

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Error(err)
			shutdown <- err
		}
	}()

	select {
	case killsignal := <-interrupt:
		switch killsignal {
		case syscall.SIGTERM:
			logger.Info("sigterm")
		case os.Interrupt:
			logger.Info("interrupt")
		}
	case <-shutdown:
		logger.Info("error, shutdown")
	}

	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Info("error in shutdown")
	}

}
