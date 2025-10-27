package user

import (
	"context"
	"time"

	"github.com/google/uuid"
	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost = 12

// Service defines the interface for user business logic
type Service interface {
	Register(ctx context.Context, req RegisterRequest) (*User, error)
	Login(ctx context.Context, req LoginRequest) (*AuthResponse, error)
	GetProfile(ctx context.Context, userID string) (*User, error)
	UpdateProfile(ctx context.Context, userID string, req UpdateProfileRequest) (*User, error)
	RefreshToken(ctx context.Context, refreshToken string) (*AuthResponse, error)
}

// service implements Service
type service struct {
	repo       Repository
	jwtService *jwtPkg.Service
	logger     *zap.Logger
}

// NewService creates a new user service
func NewService(repo Repository, jwtService *jwtPkg.Service, logger *zap.Logger) Service {
	return &service{
		repo:       repo,
		jwtService: jwtService,
		logger:     logger,
	}
}

// Register registers a new user
func (s *service) Register(ctx context.Context, req RegisterRequest) (*User, error) {
	s.logger.Info("Registering new user", zap.String("email", req.Email))

	// Check if user already exists
	existingUser, err := s.repo.FindByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		s.logger.Warn("Registration attempt with existing email", zap.String("email", req.Email))
		return nil, ErrEmailAlreadyExists
	}
	if err != nil && err != ErrUserNotFound {
		s.logger.Error("Failed to check existing user", zap.Error(err))
		return nil, err
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcryptCost)
	if err != nil {
		s.logger.Error("Failed to hash password", zap.Error(err))
		return nil, err
	}

	// Create user
	now := time.Now()
	user := &User{
		ID:           uuid.New().String(),
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Name:         req.Name,
		Role:         "user",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		s.logger.Error("Failed to create user", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User registered successfully", zap.String("user_id", user.ID), zap.String("email", user.Email))
	return user, nil
}

// Login authenticates a user and returns JWT tokens
func (s *service) Login(ctx context.Context, req LoginRequest) (*AuthResponse, error) {
	s.logger.Info("User login attempt", zap.String("email", req.Email))

	// Find user by email
	user, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == ErrUserNotFound {
			s.logger.Warn("Login attempt with non-existent email", zap.String("email", req.Email))
			return nil, ErrInvalidCredentials
		}
		s.logger.Error("Failed to find user", zap.Error(err))
		return nil, err
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		s.logger.Warn("Login attempt with invalid password", zap.String("email", req.Email))
		return nil, ErrInvalidCredentials
	}

	// Generate tokens
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		s.logger.Error("Failed to generate access token", zap.Error(err))
		return nil, err
	}

	refreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		s.logger.Error("Failed to generate refresh token", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User logged in successfully", zap.String("user_id", user.ID))

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    900, // 15 minutes in seconds
		User:         user,
	}, nil
}

// GetProfile retrieves a user's profile
func (s *service) GetProfile(ctx context.Context, userID string) (*User, error) {
	s.logger.Debug("Getting user profile", zap.String("user_id", userID))

	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to get user profile", zap.String("user_id", userID), zap.Error(err))
		return nil, err
	}

	return user, nil
}

// UpdateProfile updates a user's profile
func (s *service) UpdateProfile(ctx context.Context, userID string, req UpdateProfileRequest) (*User, error) {
	s.logger.Info("Updating user profile", zap.String("user_id", userID))

	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.String("user_id", userID), zap.Error(err))
		return nil, err
	}

	user.Name = req.Name
	user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user); err != nil {
		s.logger.Error("Failed to update user", zap.Error(err))
		return nil, err
	}

	s.logger.Info("User profile updated successfully", zap.String("user_id", user.ID))
	return user, nil
}

// RefreshToken generates new tokens using a refresh token
func (s *service) RefreshToken(ctx context.Context, refreshToken string) (*AuthResponse, error) {
	s.logger.Info("Refreshing token")

	// Validate refresh token
	userID, err := s.jwtService.ValidateRefreshToken(refreshToken)
	if err != nil {
		s.logger.Warn("Invalid refresh token", zap.Error(err))
		return nil, jwtPkg.ErrInvalidToken
	}

	// Get user
	user, err := s.repo.FindByID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to find user", zap.String("user_id", userID), zap.Error(err))
		return nil, err
	}

	// Generate new tokens
	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Email, user.Role)
	if err != nil {
		s.logger.Error("Failed to generate access token", zap.Error(err))
		return nil, err
	}

	newRefreshToken, err := s.jwtService.GenerateRefreshToken(user.ID)
	if err != nil {
		s.logger.Error("Failed to generate refresh token", zap.Error(err))
		return nil, err
	}

	s.logger.Info("Token refreshed successfully", zap.String("user_id", user.ID))

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    900,
		User:         user,
	}, nil
}
