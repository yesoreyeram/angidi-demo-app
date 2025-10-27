package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/config"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/logger"
)

const version = "0.1.0"

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// WelcomeResponse represents the welcome message response
type WelcomeResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	appLogger, err := logger.New(cfg.Logging.Level, cfg.Logging.Format)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func() {
		if syncErr := appLogger.Sync(); syncErr != nil {
			log.Printf("Failed to sync logger: %v", syncErr)
		}
	}()

	appLogger.Info("Starting Angidi API server",
		"version", version,
		"port", cfg.Server.Port,
	)

	// Setup HTTP server with dependency injection
	handler := newHandler(appLogger)
	mux := http.NewServeMux()

	// Register endpoints
	mux.HandleFunc("/health", handler.healthCheckHandler)
	mux.HandleFunc("/", handler.welcomeHandler)

	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      mux,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		appLogger.Info("Server listening", "address", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for interrupt signal or server error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			appLogger.Error("Server failed", "error", err)
			os.Exit(1)
		}
	case <-quit:
		appLogger.Info("Shutting down server...")

		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			appLogger.Error("Server forced to shutdown", "error", err)
			os.Exit(1)
		}

		appLogger.Info("Server exited")
	}
}

// handler holds dependencies for HTTP handlers
type handler struct {
	logger *logger.Logger
}

// newHandler creates a new handler with dependencies
func newHandler(l *logger.Logger) *handler {
	return &handler{
		logger: l,
	}
}

// healthCheckHandler handles health check requests
func (h *handler) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Failed to encode health response", "error", err)
	}
}

// welcomeHandler handles welcome requests
func (h *handler) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	response := WelcomeResponse{
		Message: "Welcome to Angidi API",
		Version: version,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Failed to encode welcome response", "error", err)
	}
}
