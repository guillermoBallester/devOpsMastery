package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server serverCfg `mapstructure:"server"`
}

type serverCfg struct {
	Port    int           `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

func LoadConfig() (*Config, error) {
	v := viper.New()

	setDefaults(v)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	configPath := os.Getenv("CONFIG_PATH")
	if configPath != "" {
		v.AddConfigPath(configPath)
	}
	v.AddConfigPath(".")

	v.AutomaticEnv()
	v.SetEnvPrefix("APP")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("could not load config: %v", err)
		}
		fmt.Println("Config file not found.")
	} else {
		fmt.Println("Using config file:", v.ConfigFileUsed())
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("could not unmarshal config: %v", err)
	}
	return &cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("server.port", 8081)
	v.SetDefault("server.timeout", "30s")
}
