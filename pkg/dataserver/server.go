package dataserver

import (
	"context"
	"net/http"
	"time"
)

// Server base structure for dataserver
type Server struct {
	dataServer *http.Server
}

// Run the dataserver instance
func (s *Server) Run(mux *http.ServeMux, port string) error {
	s.dataServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		Handler:        mux,
	}
	return s.dataServer.ListenAndServe()
}

// Shutdown the dataserver instance
func (s *Server) Shutdown(ctx context.Context) error {
	return s.dataServer.Shutdown(ctx)
}
