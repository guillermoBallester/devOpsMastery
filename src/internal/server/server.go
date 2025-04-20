package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/guillermoBallester/devOpsMastery/src/internal/config"
	"github.com/guillermoBallester/devOpsMastery/src/internal/connection"
	"github.com/guillermoBallester/devOpsMastery/src/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	httpServer *http.Server
	config     config.HTTPConfig
	connMgr    *connection.Manager
}

func NewServer(r *router.Router, cfg config.HTTPConfig) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			Handler:      r.Handler(),
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			IdleTimeout:  cfg.IdleTimeout,
		},
		connMgr: r.GetConnectionManager(),
	}
}

// Start starts the HTTP server and handles graceful shutdown
func (s *Server) Start() error {
	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %d", s.config.Port)
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	return s.Shutdown()
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown() error {
	log.Println("Starting graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
	defer cancel()

	// stop accepting new connections
	log.Println("Stopping HTTP server...")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %w", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
