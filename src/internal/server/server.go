package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/guillermoBallester/devOpsMastery/src/internal/connection"
	"github.com/guillermoBallester/devOpsMastery/src/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	httpServer *http.Server
	port       int
	connMgr    *connection.Manager
}

func NewServer(r *router.Router, port int) *Server {
	return &Server{
		port: port,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      r.Handler(),
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		connMgr: r.GetConnectionManager(),
	}
}

// Start starts the HTTP server and handles graceful shutdown
func (s *Server) Start() error {
	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %d", s.port)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
