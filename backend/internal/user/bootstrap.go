package user

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

// BootstrapAdmin creates the initial admin user if no admin exists.
// This is called on application startup to ensure there's always a way to access admin functions.
// Credentials are read from environment variables: ADMIN_EMAIL, ADMIN_PASSWORD, ADMIN_NAME
func (s *service) BootstrapAdmin(ctx context.Context) error {
	// Check if any admin user exists
	hasAdmin, err := s.repo.HasAdmin(ctx)
	if err != nil {
		s.logger.Error("Failed to check for existing admin", zap.Error(err))
		return err
	}

	if hasAdmin {
		s.logger.Info("Admin user already exists, skipping bootstrap")
		return nil
	}

	// Get admin credentials from environment
	email := os.Getenv("ADMIN_EMAIL")
	password := os.Getenv("ADMIN_PASSWORD")
	name := os.Getenv("ADMIN_NAME")

	// If no credentials provided, skip bootstrap (development mode)
	if email == "" || password == "" {
		s.logger.Warn("No admin credentials provided in environment, skipping bootstrap. Set ADMIN_EMAIL and ADMIN_PASSWORD to create initial admin.")
		return nil
	}

	// Default name if not provided
	if name == "" {
		name = "System Administrator"
	}

	// Validate password strength
	if len(password) < 12 {
		s.logger.Error("Admin password does not meet minimum requirements (12 characters)")
		return errors.New("admin password must be at least 12 characters")
	}

	s.logger.Info("Creating initial admin user", zap.String("email", email))

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		s.logger.Error("Failed to hash admin password", zap.Error(err))
		return err
	}

	// Create admin user
	now := time.Now()
	admin := &User{
		ID:           uuid.New().String(),
		Email:        email,
		PasswordHash: string(hashedPassword),
		Name:         name,
		Role:         "admin",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.repo.Create(ctx, admin); err != nil {
		s.logger.Error("Failed to create admin user", zap.Error(err))
		return err
	}

	s.logger.Info("Initial admin user created successfully",
		zap.String("admin_id", admin.ID),
		zap.String("email", email))

	// Clear password from environment for security
	os.Unsetenv("ADMIN_PASSWORD")

	return nil
}
