//go:build integration
// +build integration

package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/yesoreyeram/angidi-demo-app/backend/internal/gateway"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/product"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/user"
	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
)

func setupTestServer(t *testing.T) *httptest.Server {
	zapLogger, err := zap.NewDevelopment()
	require.NoError(t, err)

	jwtService := jwtPkg.NewService("test-secret-key", 15*time.Minute, 7*24*time.Hour)

	userRepo := user.NewInMemoryRepository()
	productRepo := product.NewInMemoryRepository()

	userService := user.NewService(userRepo, jwtService, zapLogger)
	productService := product.NewService(productRepo, zapLogger)

	userHandler := user.NewHandler(userService, zapLogger)
	productHandler := product.NewHandler(productService, zapLogger)

	router := gateway.Router(userHandler, productHandler, jwtService, zapLogger)

	return httptest.NewServer(router)
}

func TestHealthEndpoint_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err)

	data, ok := result["data"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "healthy", data["status"])
}

func TestUserRegistrationFlow_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Test registration
	registerPayload := map[string]string{
		"email":    "integration@test.com",
		"password": "SecurePass123!",
		"name":     "Integration Test User",
	}
	body, _ := json.Marshal(registerPayload)

	resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var registerResult map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&registerResult)
	require.NoError(t, err)

	userData, ok := registerResult["data"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, "integration@test.com", userData["email"])
	assert.Equal(t, "Integration Test User", userData["name"])

	// Test duplicate registration fails
	resp, err = http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestUserLoginFlow_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Register user first
	registerPayload := map[string]string{
		"email":    "login@test.com",
		"password": "SecurePass123!",
		"name":     "Login Test User",
	}
	body, _ := json.Marshal(registerPayload)
	resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	resp.Body.Close()

	// Test login
	loginPayload := map[string]string{
		"email":    "login@test.com",
		"password": "SecurePass123!",
	}
	body, _ = json.Marshal(loginPayload)

	resp, err = http.Post(server.URL+"/api/v1/users/login", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var loginResult map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&loginResult)
	require.NoError(t, err)

	authData, ok := loginResult["data"].(map[string]interface{})
	require.True(t, ok)
	assert.NotEmpty(t, authData["access_token"])
	assert.NotEmpty(t, authData["refresh_token"])
	assert.Equal(t, float64(900), authData["expires_in"])

	// Test wrong password fails
	wrongLoginPayload := map[string]string{
		"email":    "login@test.com",
		"password": "WrongPassword",
	}
	body, _ = json.Marshal(wrongLoginPayload)
	resp, err = http.Post(server.URL+"/api/v1/users/login", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestProtectedEndpoints_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Register and login to get token
	registerPayload := map[string]string{
		"email":    "protected@test.com",
		"password": "SecurePass123!",
		"name":     "Protected Test User",
	}
	body, _ := json.Marshal(registerPayload)
	resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginPayload := map[string]string{
		"email":    "protected@test.com",
		"password": "SecurePass123!",
	}
	body, _ = json.Marshal(loginPayload)
	resp, err = http.Post(server.URL+"/api/v1/users/login", "application/json", bytes.NewReader(body))
	require.NoError(t, err)

	var loginResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&loginResult)
	resp.Body.Close()

	authData := loginResult["data"].(map[string]interface{})
	accessToken := authData["access_token"].(string)

	// Test accessing protected endpoint with token
	req, _ := http.NewRequest("GET", server.URL+"/api/v1/users/me", nil)
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err = client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var profileResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&profileResult)
	profileData := profileResult["data"].(map[string]interface{})
	assert.Equal(t, "protected@test.com", profileData["email"])

	// Test accessing protected endpoint without token fails
	req, _ = http.NewRequest("GET", server.URL+"/api/v1/users/me", nil)
	resp, err = client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestProductCRUD_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Test listing products (public endpoint)
	resp, err := http.Get(server.URL + "/api/v1/products")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&listResult)
	listData := listResult["data"].(map[string]interface{})
	assert.NotNil(t, listData["products"])
	assert.Equal(t, float64(0), listData["total_count"])

	// Note: Create, Update, Delete are admin-only endpoints
	// Would need to create an admin user to test those
}

func TestProductFiltering_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Test with query parameters
	resp, err := http.Get(server.URL + "/api/v1/products?page=1&page_size=5&search=test")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var listResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&listResult)
	listData := listResult["data"].(map[string]interface{})
	assert.Equal(t, float64(1), listData["page"])
	assert.Equal(t, float64(5), listData["page_size"])
}

func TestRefreshToken_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Register and login
	registerPayload := map[string]string{
		"email":    "refresh@test.com",
		"password": "SecurePass123!",
		"name":     "Refresh Test User",
	}
	body, _ := json.Marshal(registerPayload)
	resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	resp.Body.Close()

	loginPayload := map[string]string{
		"email":    "refresh@test.com",
		"password": "SecurePass123!",
	}
	body, _ = json.Marshal(loginPayload)
	resp, err = http.Post(server.URL+"/api/v1/users/login", "application/json", bytes.NewReader(body))
	require.NoError(t, err)

	var loginResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&loginResult)
	resp.Body.Close()

	authData := loginResult["data"].(map[string]interface{})
	refreshToken := authData["refresh_token"].(string)

	// Wait a moment to ensure new tokens will have different timestamps
	time.Sleep(1 * time.Second)

	// Test refresh token
	refreshPayload := map[string]string{
		"refresh_token": refreshToken,
	}
	body, _ = json.Marshal(refreshPayload)
	resp, err = http.Post(server.URL+"/api/v1/users/refresh-token", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var refreshResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&refreshResult)
	newAuthData := refreshResult["data"].(map[string]interface{})
	assert.NotEmpty(t, newAuthData["access_token"])
	assert.NotEmpty(t, newAuthData["refresh_token"])
}

func TestInputValidation_Integration(t *testing.T) {
	server := setupTestServer(t)
	defer server.Close()

	// Test with invalid email
	registerPayload := map[string]string{
		"email":    "invalid-email",
		"password": "SecurePass123!",
		"name":     "Test User",
	}
	body, _ := json.Marshal(registerPayload)
	resp, err := http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	// Test with short password
	registerPayload = map[string]string{
		"email":    "test@example.com",
		"password": "short",
		"name":     "Test User",
	}
	body, _ = json.Marshal(registerPayload)
	resp, err = http.Post(server.URL+"/api/v1/users/register", "application/json", bytes.NewReader(body))
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
