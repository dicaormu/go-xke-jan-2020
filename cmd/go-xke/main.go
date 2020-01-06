package main

import (
	"net"
	"net/http"
)

func main() {
	port := "8080"

	server := http.Server{
		Addr: net.JoinHostPort("", port),
	}

	server.ListenAndServe()
}
