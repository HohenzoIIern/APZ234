package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/HohenzoIIern/APZ234/Lab3/server/server"
)

type HttpPortNumber int

// ChatApiServer configures necessary handlers and starts listening on a configured port.
type ChatApiServer struct {
	Port HttpPortNumber

	ServerHandler server.HttpHandlerFunc
	DiskHandler   server.HttpHandlerFunc

	server *http.Server
}

// Start will set all handlers and start listening.
// If this methods succeeds, it does not return until server is shut down.
// Returned error will never be nil.
func (s *ChatApiServer) Start() error {
	if s.ServerHandler == nil {
		return fmt.Errorf("server HTTP handler is not defined - cannot start")
	}
	if s.DiskHandler == nil {
		return fmt.Errorf("disk HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/server", s.ServerHandler)
	handler.HandleFunc("/disk", s.DiskHandler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

// Stops will shut down previously started HTTP server.
func (s *ChatApiServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
