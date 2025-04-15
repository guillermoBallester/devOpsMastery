package main

import (
	"fmt"
	"github.com/guillermoBallester/devOpsMastery/src/internal/config"
	"log"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v\n", err)
	}

	fmt.Printf("Loaded config: %+v\n", cfg)

}
