package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GenerateAccessToken(t *testing.T) {
	service := NewService("test-secret-key", 15*time.Minute, 7*24*time.Hour)

	token, err := service.GenerateAccessToken("user-123", "test@example.com", "user")
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the generated token
	claims, err := service.ValidateToken(token)
	require.NoError(t, err)
	assert.Equal(t, "user-123", claims.UserID)
	assert.Equal(t, "test@example.com", claims.Email)
	assert.Equal(t, "user", claims.Role)
}

func TestService_GenerateRefreshToken(t *testing.T) {
	service := NewService("test-secret-key", 15*time.Minute, 7*24*time.Hour)

	token, err := service.GenerateRefreshToken("user-123")
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate the generated token
	userID, err := service.ValidateRefreshToken(token)
	require.NoError(t, err)
	assert.Equal(t, "user-123", userID)
}

func TestService_ValidateToken_InvalidToken(t *testing.T) {
	service := NewService("test-secret-key", 15*time.Minute, 7*24*time.Hour)

	_, err := service.ValidateToken("invalid-token")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
}

func TestService_ValidateToken_ExpiredToken(t *testing.T) {
	service := NewService("test-secret-key", -1*time.Hour, 7*24*time.Hour)

	token, err := service.GenerateAccessToken("user-123", "test@example.com", "user")
	require.NoError(t, err)

	// Wait a moment to ensure token is expired
	time.Sleep(100 * time.Millisecond)

	_, err = service.ValidateToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrExpiredToken, err)
}

func TestService_ValidateToken_WrongSecret(t *testing.T) {
	service1 := NewService("secret-key-1", 15*time.Minute, 7*24*time.Hour)
	service2 := NewService("secret-key-2", 15*time.Minute, 7*24*time.Hour)

	token, err := service1.GenerateAccessToken("user-123", "test@example.com", "user")
	require.NoError(t, err)

	// Try to validate with different secret
	_, err = service2.ValidateToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
}

func TestService_ValidateRefreshToken_InvalidToken(t *testing.T) {
	service := NewService("test-secret-key", 15*time.Minute, 7*24*time.Hour)

	_, err := service.ValidateRefreshToken("invalid-token")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidToken, err)
}

func TestService_ValidateRefreshToken_ExpiredToken(t *testing.T) {
	service := NewService("test-secret-key", 15*time.Minute, -1*time.Hour)

	token, err := service.GenerateRefreshToken("user-123")
	require.NoError(t, err)

	// Wait a moment to ensure token is expired
	time.Sleep(100 * time.Millisecond)

	_, err = service.ValidateRefreshToken(token)
	assert.Error(t, err)
	assert.Equal(t, ErrExpiredToken, err)
}
