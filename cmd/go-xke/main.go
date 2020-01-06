package main

import (
	"github.com/dicaormu/go-xke-jan-2020/internal/application"
	"github.com/gorilla/mux"
	"net"
	"net/http"
)

func main() {
	port := "8080"

	router := mux.NewRouter()
	router.HandleFunc("/", application.HomeHandler)

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	server.ListenAndServe()
}
