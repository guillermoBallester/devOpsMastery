package main

import (
	"github.com/guillermoBallester/devOpsMastery/src/internal/config"
	"github.com/guillermoBallester/devOpsMastery/src/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	srv := server.NewServer(cfg.Server.HTTP.Port)

	// Handle graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	// Start the server (this will block until the server exits)
	log.Fatal(srv.Start())

}
