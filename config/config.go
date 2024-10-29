package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ServerConfig holds the server configuration details.
type ServerConfig struct {
	Host string
	Port int
}

// ClientConfig holds the client configuration details.
type ClientConfig struct {
	ServerAddress string
	Timeout       int
}

// Config is a struct that encapsulates the configuration management.
type Config struct {
	viper *viper.Viper
}

// NewConfig initializes and returns a new Config instance with default settings.
func NewConfig() *Config {
	v := viper.New()
	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", 8080)
	v.SetDefault("client.serverAddress", "http://localhost:8080")
	v.SetDefault("client.timeout", 30)
	return &Config{viper: v}
}

// LoadConfig loads the configuration from a file or environment variables.
func (c *Config) LoadConfig() error {
	c.viper.SetConfigName("config")
	c.viper.AddConfigPath(".")
	c.viper.AutomaticEnv()

	if err := c.viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	return nil
}

// GetServerConfig returns the server configuration.
func (c *Config) GetServerConfig() ServerConfig {
	return ServerConfig{
		Host: c.viper.GetString("server.host"),
		Port: c.viper.GetInt("server.port"),
	}
}

// GetClientConfig returns the client configuration.
func (c *Config) GetClientConfig() ClientConfig {
	return ClientConfig{
		ServerAddress: c.viper.GetString("client.serverAddress"),
		Timeout:       c.viper.GetInt("client.timeout"),
	}
}
