package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/yesoreyeram/angidi-demo-app/backend/internal/gateway"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/product"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/user"
	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"go.uber.org/zap"
)

func TestHealthEndpoint(t *testing.T) {
	// Setup test dependencies
	zapLogger, _ := zap.NewDevelopment()
	jwtService := jwtPkg.NewService("test-secret", 15*time.Minute, 7*24*time.Hour)
	
	userRepo := user.NewInMemoryRepository()
	productRepo := product.NewInMemoryRepository()
	
	userService := user.NewService(userRepo, jwtService, zapLogger)
	productService := product.NewService(productRepo, zapLogger)
	
	userHandler := user.NewHandler(userService, zapLogger)
	productHandler := product.NewHandler(productService, zapLogger)
	
	router := gateway.Router(userHandler, productHandler, jwtService, zapLogger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.StatusCode)
	}

	var healthResp map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&healthResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	data, ok := healthResp["data"].(map[string]interface{})
	if !ok {
		t.Fatal("Expected 'data' field in response")
	}

	status, ok := data["status"].(string)
	if !ok || status != "healthy" {
		t.Errorf("Expected status 'healthy', got '%v'", status)
	}
}

func TestVersion(t *testing.T) {
	if version != "0.2.0" {
		t.Errorf("Expected version '0.2.0', got '%s'", version)
	}
}

