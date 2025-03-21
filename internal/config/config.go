package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Config holds all application configuration
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Log    LogConfig    `mapstructure:"log"`
}

// ServerConfig holds HTTP server configuration
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// LogConfig holds logging configuration
type LogConfig struct {
	Level string `mapstructure:"level"`
}

// Load reads configuration from config.yaml and environment variables
func Load() (*Config, error) {
	v := viper.New()

	// Determine config file name - allow override for testing
	configName := os.Getenv("VIPER_CONFIG_NAME")
	if configName == "" {
		configName = "config"
	}

	// Set up config file
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	// Read from config file
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Set up environment variables
	v.SetEnvPrefix("CONTEST_BOT")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
