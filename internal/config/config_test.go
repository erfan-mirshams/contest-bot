package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	origConfigFile string
	tempConfigFile string
}

func (s *ConfigTestSuite) SetupTest() {
	// Save original path
	s.origConfigFile = "config.yaml"

	// Create a temporary config file for testing
	s.tempConfigFile = "config_test.yaml"

	// Write test configuration
	err := os.WriteFile(s.tempConfigFile, []byte(`
server:
  port: 8081
log:
  level: debug
`), 0644)
	s.Require().NoError(err)
}

func (s *ConfigTestSuite) TearDownTest() {
	os.Remove(s.tempConfigFile)
}

func (s *ConfigTestSuite) TestLoadConfig() {
	// Test successful config loading
	s.Run("Loads config from file", func() {
		// Temporarily set env var to use test config file
		os.Setenv("VIPER_CONFIG_NAME", "config_test")
		defer os.Unsetenv("VIPER_CONFIG_NAME")

		cfg, err := Load()
		s.NoError(err)
		s.NotNil(cfg)
		s.Equal(8081, cfg.Server.Port)
		s.Equal("debug", cfg.Log.Level)
	})

	// Test environment variable override
	s.Run("Environment variables override config file", func() {
		os.Setenv("VIPER_CONFIG_NAME", "config_test")
		os.Setenv("CONTEST_BOT_SERVER_PORT", "9000")
		defer func() {
			os.Unsetenv("VIPER_CONFIG_NAME")
			os.Unsetenv("CONTEST_BOT_SERVER_PORT")
		}()

		cfg, err := Load()
		s.NoError(err)
		s.NotNil(cfg)
		s.Equal(9000, cfg.Server.Port)
		s.Equal("debug", cfg.Log.Level)
	})
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
