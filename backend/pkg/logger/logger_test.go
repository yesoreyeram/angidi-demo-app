package logger

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		level   string
		format  string
		wantErr bool
	}{
		{
			name:    "valid json format with info level",
			level:   "info",
			format:  "json",
			wantErr: false,
		},
		{
			name:    "valid console format with debug level",
			level:   "debug",
			format:  "console",
			wantErr: false,
		},
		{
			name:    "invalid level defaults to info",
			level:   "invalid",
			format:  "json",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger, err := New(tt.level, tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && logger == nil {
				t.Error("New() returned nil logger")
			}
			if logger != nil {
				// Test that sync doesn't panic
				_ = logger.Sync()
			}
		})
	}
}

func TestLogger_WithField(t *testing.T) {
	logger, err := New("info", "json")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	newLogger := logger.WithField("key", "value")
	if newLogger == nil {
		t.Fatal("WithField() returned nil")
	}
	if newLogger.SugaredLogger == nil {
		t.Error("WithField() returned logger with nil SugaredLogger")
	}
}

func TestLogger_Sync(t *testing.T) {
	tests := []struct {
		name    string
		logger  *Logger
		wantErr bool
	}{
		{
			name:    "valid logger",
			logger:  &Logger{SugaredLogger: nil},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.logger.Sync()
			if (err != nil) != tt.wantErr {
				t.Errorf("Sync() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	logger := Default()
	if logger == nil {
		t.Fatal("Default() returned nil")
	}
	if logger.SugaredLogger == nil {
		t.Error("Default() returned logger with nil SugaredLogger")
	}
}

func TestLogger_LogMethods(t *testing.T) {
	logger, err := New("debug", "json")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	// Test that log methods don't panic
	logger.Debug("debug message", "key", "value")
	logger.Info("info message", "key", "value")
	logger.Warn("warn message", "key", "value")
	logger.Error("error message", "key", "value")
}
