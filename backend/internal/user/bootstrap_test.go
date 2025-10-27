package user

import (
	"context"
	"os"
	"testing"
	"time"

	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func TestService_BootstrapAdmin(t *testing.T) {
	tests := []struct {
		name          string
		existingAdmin bool
		envEmail      string
		envPassword   string
		envName       string
		wantErr       bool
		wantAdmin     bool
	}{
		{
			name:          "creates admin when none exists and credentials provided",
			existingAdmin: false,
			envEmail:      "admin@test.com",
			envPassword:   "SecureAdminPass123!",
			envName:       "Test Admin",
			wantErr:       false,
			wantAdmin:     true,
		},
		{
			name:          "skips when admin already exists",
			existingAdmin: true,
			envEmail:      "admin@test.com",
			envPassword:   "SecureAdminPass123!",
			envName:       "Test Admin",
			wantErr:       false,
			wantAdmin:     true,
		},
		{
			name:          "skips when no credentials provided",
			existingAdmin: false,
			envEmail:      "",
			envPassword:   "",
			envName:       "",
			wantErr:       false,
			wantAdmin:     false,
		},
		{
			name:          "fails with weak password",
			existingAdmin: false,
			envEmail:      "admin@test.com",
			envPassword:   "weak",
			envName:       "Test Admin",
			wantErr:       true,
			wantAdmin:     false,
		},
		{
			name:          "uses default name when not provided",
			existingAdmin: false,
			envEmail:      "admin@test.com",
			envPassword:   "SecureAdminPass123!",
			envName:       "",
			wantErr:       false,
			wantAdmin:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			repo := NewInMemoryRepository()
			jwtService := jwtPkg.NewService("test-secret", 15*time.Minute, 7*24*time.Hour)
			logger, _ := zap.NewDevelopment()
			service := NewService(repo, jwtService, logger)

			// Create existing admin if needed
			if tt.existingAdmin {
				existingAdmin := &User{
					ID:           "existing-admin",
					Email:        "existing@admin.com",
					PasswordHash: "hashed",
					Name:         "Existing Admin",
					Role:         "admin",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}
				repo.Create(context.Background(), existingAdmin)
			}

			// Set environment variables
			if tt.envEmail != "" {
				os.Setenv("ADMIN_EMAIL", tt.envEmail)
				defer os.Unsetenv("ADMIN_EMAIL")
			}
			if tt.envPassword != "" {
				os.Setenv("ADMIN_PASSWORD", tt.envPassword)
				defer os.Unsetenv("ADMIN_PASSWORD")
			}
			if tt.envName != "" {
				os.Setenv("ADMIN_NAME", tt.envName)
				defer os.Unsetenv("ADMIN_NAME")
			}

			// Execute
			err := service.BootstrapAdmin(context.Background())

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			// Check if admin was created
			hasAdmin, err := repo.HasAdmin(context.Background())
			require.NoError(t, err)
			assert.Equal(t, tt.wantAdmin, hasAdmin)

			// If admin should be created, verify it
			if tt.wantAdmin && !tt.existingAdmin && tt.envEmail != "" {
				admin, err := repo.FindByEmail(context.Background(), tt.envEmail)
				require.NoError(t, err)
				assert.Equal(t, "admin", admin.Role)
				assert.Equal(t, tt.envEmail, admin.Email)

				// Verify password was hashed
				assert.NotEqual(t, tt.envPassword, admin.PasswordHash)
				
				// Verify password can be validated
				if tt.envPassword != "weak" {
					err = bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(tt.envPassword))
					assert.NoError(t, err)
				}

				// Verify name
				expectedName := tt.envName
				if expectedName == "" {
					expectedName = "System Administrator"
				}
				assert.Equal(t, expectedName, admin.Name)
			}
		})
	}
}

func TestRepository_HasAdmin(t *testing.T) {
	tests := []struct {
		name      string
		setupRepo func(*InMemoryRepository)
		want      bool
	}{
		{
			name: "returns true when admin exists",
			setupRepo: func(r *InMemoryRepository) {
				admin := &User{
					ID:    "admin-1",
					Email: "admin@test.com",
					Role:  "admin",
				}
				r.Create(context.Background(), admin)
			},
			want: true,
		},
		{
			name: "returns false when only regular users exist",
			setupRepo: func(r *InMemoryRepository) {
				user := &User{
					ID:    "user-1",
					Email: "user@test.com",
					Role:  "user",
				}
				r.Create(context.Background(), user)
			},
			want: false,
		},
		{
			name: "returns false when repository is empty",
			setupRepo: func(r *InMemoryRepository) {
				// No setup needed
			},
			want: false,
		},
		{
			name: "returns true when multiple admins exist",
			setupRepo: func(r *InMemoryRepository) {
				admin1 := &User{
					ID:    "admin-1",
					Email: "admin1@test.com",
					Role:  "admin",
				}
				admin2 := &User{
					ID:    "admin-2",
					Email: "admin2@test.com",
					Role:  "admin",
				}
				r.Create(context.Background(), admin1)
				r.Create(context.Background(), admin2)
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := NewInMemoryRepository()
			tt.setupRepo(repo)

			got, err := repo.HasAdmin(context.Background())
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
