package main

import (
	"github.com/guillermoBallester/devOpsMastery/src/internal/config"
	"github.com/guillermoBallester/devOpsMastery/src/internal/router"
	"github.com/guillermoBallester/devOpsMastery/src/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	r := router.NewRouter()
	srv := server.NewServer(r, cfg.Server.HTTP)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
