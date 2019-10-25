package rest

import (
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

type ServerConfig struct {
	Port int
}

// requires http.Handler in arguments
//   which is returned in router.go
func NewServer(cfg *ServerConfig, handler Router) *Server {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: handler,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}
