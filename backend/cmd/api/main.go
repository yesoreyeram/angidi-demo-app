package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yesoreyeram/angidi-demo-app/backend/internal/database"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/gateway"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/product"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/user"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/config"
	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/logger"
	"go.uber.org/zap"
)

const version = "0.2.0"

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

	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize zap logger: %v", err)
	}
	defer zapLogger.Sync()

	appLogger.Info("Starting Angidi API server",
		"version", version,
		"port", cfg.Server.Port,
	)

	// Initialize database connection pool
	dbConfig := database.Config{
		Host:            cfg.Database.Host,
		Port:            cfg.Database.Port,
		User:            cfg.Database.User,
		Password:        cfg.Database.Password,
		Database:        cfg.Database.Database,
		SSLMode:         cfg.Database.SSLMode,
		MaxConns:        cfg.Database.MaxConns,
		MinConns:        cfg.Database.MinConns,
		MaxConnLifetime: cfg.Database.MaxConnLifetime,
		MaxConnIdleTime: cfg.Database.MaxConnIdleTime,
	}

	ctx := context.Background()
	pool, err := database.NewPool(ctx, dbConfig)
	if err != nil {
		zapLogger.Fatal("Failed to create database pool", zap.Error(err))
	}
	defer pool.Close()

	zapLogger.Info("Database connection established",
		zap.String("host", cfg.Database.Host),
		zap.Int("port", cfg.Database.Port),
		zap.String("database", cfg.Database.Database),
	)

	// Initialize services
	jwtService := jwtPkg.NewService(
		getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		15*time.Minute,  // access token duration
		7*24*time.Hour,  // refresh token duration
	)

	// Initialize repositories (using PostgreSQL)
	userRepo := user.NewPostgresRepository(pool)
	productRepo := product.NewPostgresRepository(pool)

	// Initialize services
	userService := user.NewService(userRepo, jwtService, zapLogger)
	productService := product.NewService(productRepo, zapLogger)

	// Bootstrap admin user if needed
	if err := userService.BootstrapAdmin(context.Background()); err != nil {
		zapLogger.Fatal("Failed to bootstrap admin user", zap.Error(err))
	}

	// Initialize handlers
	userHandler := user.NewHandler(userService, zapLogger)
	productHandler := product.NewHandler(productService, zapLogger)

	// Setup router
	router := gateway.Router(userHandler, productHandler, jwtService, zapLogger)

	// Setup HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	// Start server in a goroutine
	serverErrors := make(chan error, 1)
	go func() {
		appLogger.Info("Server listening", "address", server.Addr)
		zapLogger.Info("API Gateway ready",
			zap.String("address", server.Addr),
			zap.String("version", version),
		)
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for interrupt signal or server error
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != nil && err != http.ErrServerClosed {
			appLogger.Error("Server failed", "error", err)
			zapLogger.Error("Server failed", zap.Error(err))
			os.Exit(1)
		}
	case <-quit:
		appLogger.Info("Shutting down server...")
		zapLogger.Info("Shutting down server...")

		// Graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			appLogger.Error("Server forced to shutdown", "error", err)
			zapLogger.Error("Server forced to shutdown", zap.Error(err))
			os.Exit(1)
		}

		appLogger.Info("Server exited")
		zapLogger.Info("Server exited")
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
