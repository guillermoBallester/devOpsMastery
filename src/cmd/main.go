package main

import (
	"context"
	"github.com/guillermoBallester/devOpsMastery/src/internal/config"
	"github.com/guillermoBallester/devOpsMastery/src/internal/router"
	"github.com/guillermoBallester/devOpsMastery/src/internal/server"
	"github.com/guillermoBallester/devOpsMastery/src/internal/telemetry"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	tel, err := telemetry.Initialize(context.Background(), "devops-mastery", "0.1.0", "localhost:4317", true)
	if err != nil {
		return
	}

	r := router.NewRouter(tel)
	srv := server.NewServer(r, cfg.Server.HTTP)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
