package user

import (
	"context"
	"testing"
	"time"

	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func setupTestService() Service {
	repo := NewInMemoryRepository()
	jwtService := jwtPkg.NewService("test-secret", 15*time.Minute, 7*24*time.Hour)
	logger, _ := zap.NewDevelopment()
	return NewService(repo, jwtService, logger)
}

func TestService_Register(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	tests := []struct {
		name    string
		request RegisterRequest
		wantErr bool
	}{
		{
			name: "successful registration",
			request: RegisterRequest{
				Email:    "test@example.com",
				Password: "SecurePass123!",
				Name:     "Test User",
			},
			wantErr: false,
		},
		{
			name: "duplicate email",
			request: RegisterRequest{
				Email:    "test@example.com",
				Password: "AnotherPass123!",
				Name:     "Another User",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.Register(ctx, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, user)
				if err != nil {
					assert.Equal(t, ErrEmailAlreadyExists, err)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, user)
				assert.NotEmpty(t, user.ID)
				assert.Equal(t, tt.request.Email, user.Email)
				assert.Equal(t, tt.request.Name, user.Name)
				assert.Equal(t, "user", user.Role)
				assert.NotEmpty(t, user.PasswordHash)
				assert.NotEqual(t, tt.request.Password, user.PasswordHash)
			}
		})
	}
}

func TestService_Login(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Register a user first
	registerReq := RegisterRequest{
		Email:    "test@example.com",
		Password: "SecurePass123!",
		Name:     "Test User",
	}
	_, err := service.Register(ctx, registerReq)
	require.NoError(t, err)

	tests := []struct {
		name    string
		request LoginRequest
		wantErr bool
	}{
		{
			name: "successful login",
			request: LoginRequest{
				Email:    "test@example.com",
				Password: "SecurePass123!",
			},
			wantErr: false,
		},
		{
			name: "wrong password",
			request: LoginRequest{
				Email:    "test@example.com",
				Password: "WrongPassword",
			},
			wantErr: true,
		},
		{
			name: "non-existent user",
			request: LoginRequest{
				Email:    "nonexistent@example.com",
				Password: "SomePassword",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authResp, err := service.Login(ctx, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, authResp)
				if err != nil {
					assert.Equal(t, ErrInvalidCredentials, err)
				}
			} else {
				require.NoError(t, err)
				assert.NotNil(t, authResp)
				assert.NotEmpty(t, authResp.AccessToken)
				assert.NotEmpty(t, authResp.RefreshToken)
				assert.Equal(t, 900, authResp.ExpiresIn)
				assert.NotNil(t, authResp.User)
				assert.Equal(t, tt.request.Email, authResp.User.Email)
			}
		})
	}
}

func TestService_GetProfile(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Register a user first
	registerReq := RegisterRequest{
		Email:    "test@example.com",
		Password: "SecurePass123!",
		Name:     "Test User",
	}
	user, err := service.Register(ctx, registerReq)
	require.NoError(t, err)

	tests := []struct {
		name    string
		userID  string
		wantErr bool
	}{
		{
			name:    "existing user",
			userID:  user.ID,
			wantErr: false,
		},
		{
			name:    "non-existent user",
			userID:  "non-existent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			profile, err := service.GetProfile(ctx, tt.userID)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, profile)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, profile)
				assert.Equal(t, tt.userID, profile.ID)
			}
		})
	}
}

func TestService_UpdateProfile(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Register a user first
	registerReq := RegisterRequest{
		Email:    "test@example.com",
		Password: "SecurePass123!",
		Name:     "Test User",
	}
	user, err := service.Register(ctx, registerReq)
	require.NoError(t, err)

	tests := []struct {
		name    string
		userID  string
		request UpdateProfileRequest
		wantErr bool
	}{
		{
			name:   "successful update",
			userID: user.ID,
			request: UpdateProfileRequest{
				Name: "Updated Name",
			},
			wantErr: false,
		},
		{
			name:   "non-existent user",
			userID: "non-existent-id",
			request: UpdateProfileRequest{
				Name: "Updated Name",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updatedUser, err := service.UpdateProfile(ctx, tt.userID, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, updatedUser)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, updatedUser)
				assert.Equal(t, tt.request.Name, updatedUser.Name)
			}
		})
	}
}

func TestService_RefreshToken(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Register and login a user first
	registerReq := RegisterRequest{
		Email:    "test@example.com",
		Password: "SecurePass123!",
		Name:     "Test User",
	}
	_, err := service.Register(ctx, registerReq)
	require.NoError(t, err)

	loginReq := LoginRequest{
		Email:    "test@example.com",
		Password: "SecurePass123!",
	}
	authResp, err := service.Login(ctx, loginReq)
	require.NoError(t, err)

	tests := []struct {
		name         string
		refreshToken string
		wantErr      bool
	}{
		{
			name:         "valid refresh token",
			refreshToken: authResp.RefreshToken,
			wantErr:      false,
		},
		{
			name:         "invalid refresh token",
			refreshToken: "invalid-token",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add a delay to ensure new tokens have different timestamps (at least 1 second for JWT)
			time.Sleep(1 * time.Second)
			
			newAuthResp, err := service.RefreshToken(ctx, tt.refreshToken)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, newAuthResp)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, newAuthResp)
				assert.NotEmpty(t, newAuthResp.AccessToken)
				assert.NotEmpty(t, newAuthResp.RefreshToken)
				// Tokens should be valid (we can't easily compare them due to timing)
			}
		})
	}
}
