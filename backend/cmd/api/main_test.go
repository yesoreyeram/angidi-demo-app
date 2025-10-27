package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/logger"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a test logger
	testLogger, err := logger.New("info", "json")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	h := newHandler(testLogger)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	h.healthCheckHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var healthResp HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&healthResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if healthResp.Status != "healthy" {
		t.Errorf("Expected status 'healthy', got '%s'", healthResp.Status)
	}

	if healthResp.Timestamp == "" {
		t.Error("Expected non-empty timestamp")
	}
}

func TestWelcomeHandler(t *testing.T) {
	// Create a test logger
	testLogger, err := logger.New("info", "json")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	h := newHandler(testLogger)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.welcomeHandler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	var welcomeResp WelcomeResponse
	if err := json.NewDecoder(resp.Body).Decode(&welcomeResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if welcomeResp.Message != "Welcome to Angidi API" {
		t.Errorf("Expected message 'Welcome to Angidi API', got '%s'", welcomeResp.Message)
	}

	if welcomeResp.Version != version {
		t.Errorf("Expected version '%s', got '%s'", version, welcomeResp.Version)
	}
}

func TestNewHandler(t *testing.T) {
	testLogger, err := logger.New("info", "json")
	if err != nil {
		t.Fatalf("Failed to create logger: %v", err)
	}

	h := newHandler(testLogger)
	if h == nil {
		t.Fatal("newHandler() returned nil")
	}
	if h.logger == nil {
		t.Error("newHandler() returned handler with nil logger")
	}
}
