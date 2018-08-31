package server

import (
	"context"
	"log"
	"net/http"

	"github.com/PeterBooker/webview/internal/config"
	"github.com/go-chi/chi"
)

// Server holds all the data the App needs
type Server struct {
	Logger *log.Logger
	Config *config.Config
	router *chi.Mux
	http   *http.Server
}

// New returns a pointer to the main server struct
func New(log *log.Logger, config *config.Config) *Server {
	s := &Server{
		Config: config,
		Logger: log,
	}

	return s
}

// Setup starts the HTTP Server
func (s *Server) Setup() {
	s.startHTTP()
}

// Shutdown will release resources and stop the server.
func (s *Server) Shutdown(ctx context.Context) {
	s.http.Shutdown(ctx)
}
