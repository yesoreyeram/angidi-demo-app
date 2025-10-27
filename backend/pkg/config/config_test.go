package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Test loading with defaults (no config file)
	os.Setenv("CONFIG_PATH", "nonexistent.yaml")
	defer os.Unsetenv("CONFIG_PATH")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if cfg.Server.Port != 8080 {
		t.Errorf("Expected default port 8080, got: %d", cfg.Server.Port)
	}

	if cfg.Logging.Level != "info" {
		t.Errorf("Expected default log level 'info', got: %s", cfg.Logging.Level)
	}
}

func TestLoadWithEnvOverride(t *testing.T) {
	os.Setenv("CONFIG_PATH", "nonexistent.yaml")
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("LOG_LEVEL", "debug")
	defer func() {
		os.Unsetenv("CONFIG_PATH")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("LOG_LEVEL")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if cfg.Server.Port != 9090 {
		t.Errorf("Expected port 9090, got: %d", cfg.Server.Port)
	}

	if cfg.Logging.Level != "debug" {
		t.Errorf("Expected log level 'debug', got: %s", cfg.Logging.Level)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "valid config",
			config:  newDefaultConfig(),
			wantErr: false,
		},
		{
			name: "invalid port - too low",
			config: &Config{
				Server: ServerConfig{
					Port:            0,
					ReadTimeout:     10,
					WriteTimeout:    10,
					ShutdownTimeout: 30,
				},
				Logging: LoggingConfig{Level: "info"},
			},
			wantErr: true,
		},
		{
			name: "invalid port - too high",
			config: &Config{
				Server: ServerConfig{
					Port:            70000,
					ReadTimeout:     10,
					WriteTimeout:    10,
					ShutdownTimeout: 30,
				},
				Logging: LoggingConfig{Level: "info"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateConfigPath(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "valid relative path",
			path:    "configs/local.yaml",
			wantErr: false,
		},
		{
			name:    "absolute path not allowed",
			path:    "/etc/config.yaml",
			wantErr: true,
		},
		{
			name:    "parent directory reference not allowed",
			path:    "../config.yaml",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfigPath(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateConfigPath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
