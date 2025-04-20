package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Unsetenv("APP_SERVER_HTTP_PORT")
	os.Unsetenv("APP_SERVER_HTTP_READTIMEOUT")
	os.Unsetenv("CONFIG_PATH")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	// Check default values
	if cfg.Server.HTTP.Port != 8081 {
		t.Errorf("Expected default port 8081, got %d", cfg.Server.HTTP.Port)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Set environment variables
	os.Setenv("APP_SERVER_HTTP_PORT", "9000")
	defer func() {
		os.Unsetenv("APP_SERVER_HTTP_PORT")
	}()

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("Error loading config: %v", err)
	}

	if cfg.Server.HTTP.Port != 9000 {
		t.Errorf("Expected port 9000 from env, got %d", cfg.Server.HTTP.Port)
	}
}
