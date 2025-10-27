package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

// Config holds all configuration for the application
type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Logging LoggingConfig `yaml:"logging"`
	CORS    CORSConfig    `yaml:"cors"`
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	Output string `yaml:"output"`
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
}

// Load loads configuration from file
func Load() (*Config, error) {
	// Default configuration
	cfg := newDefaultConfig()

	// Try to load from config file
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "configs/local.yaml"
	}

	// Validate and sanitize the config path
	cleanPath := filepath.Clean(configPath)
	if err := validateConfigPath(cleanPath); err != nil {
		return nil, fmt.Errorf("invalid config path: %w", err)
	}

	data, err := os.ReadFile(cleanPath) // #nosec G304 - path is validated above
	if err != nil {
		// If config file doesn't exist, use defaults but still apply env overrides
		if os.IsNotExist(err) {
			// Apply environment variable overrides to defaults
			if err := applyEnvOverrides(cfg); err != nil {
				return nil, fmt.Errorf("failed to apply env overrides: %w", err)
			}
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	// Validate configuration
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// Override with environment variables
	if err := applyEnvOverrides(cfg); err != nil {
		return nil, fmt.Errorf("failed to apply env overrides: %w", err)
	}

	return cfg, nil
}

// newDefaultConfig returns a config with sensible defaults
func newDefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:            "localhost",
			Port:            8080,
			ReadTimeout:     10 * time.Second,
			WriteTimeout:    10 * time.Second,
			ShutdownTimeout: 30 * time.Second,
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
			Output: "stdout",
		},
		CORS: CORSConfig{
			AllowedOrigins: []string{"http://localhost:3000"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Authorization"},
		},
	}
}

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.Server.Port < 1 || c.Server.Port > 65535 {
		return fmt.Errorf("invalid port: %d", c.Server.Port)
	}
	if c.Server.ReadTimeout <= 0 {
		return fmt.Errorf("read timeout must be positive")
	}
	if c.Server.WriteTimeout <= 0 {
		return fmt.Errorf("write timeout must be positive")
	}
	if c.Server.ShutdownTimeout <= 0 {
		return fmt.Errorf("shutdown timeout must be positive")
	}
	if c.Logging.Level == "" {
		return fmt.Errorf("log level cannot be empty")
	}
	return nil
}

// validateConfigPath validates the configuration file path to prevent directory traversal
func validateConfigPath(path string) error {
	// Check for directory traversal attempts
	if filepath.IsAbs(path) {
		return fmt.Errorf("absolute paths are not allowed")
	}
	// Check for parent directory references
	if filepath.Dir(path) == ".." || filepath.Base(path) == ".." {
		return fmt.Errorf("parent directory references are not allowed")
	}
	return nil
}

// applyEnvOverrides applies environment variable overrides to the configuration
func applyEnvOverrides(cfg *Config) error {
	if host := os.Getenv("SERVER_HOST"); host != "" {
		cfg.Server.Host = host
	}
	if port := os.Getenv("SERVER_PORT"); port != "" {
		p, err := strconv.Atoi(port)
		if err != nil {
			return fmt.Errorf("invalid SERVER_PORT: %w", err)
		}
		if p > 0 && p <= 65535 {
			cfg.Server.Port = p
		} else {
			return fmt.Errorf("SERVER_PORT out of valid range: %d", p)
		}
	}
	if level := os.Getenv("LOG_LEVEL"); level != "" {
		cfg.Logging.Level = level
	}
	return nil
}
