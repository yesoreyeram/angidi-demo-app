package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/response"
	"go.uber.org/zap"
)

// Handler handles HTTP requests for user operations
type Handler struct {
	service   Service
	validator *validator.Validate
	logger    *zap.Logger
}

// NewHandler creates a new user handler
func NewHandler(service Service, logger *zap.Logger) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
		logger:    logger,
	}
}

// Register handles user registration
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", "")
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		validationErrors := make([]response.ValidationError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, response.ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}
		response.WriteValidationError(w, validationErrors, "")
		return
	}

	user, err := h.service.Register(r.Context(), req)
	if err != nil {
		if err == ErrEmailAlreadyExists {
			response.WriteError(w, http.StatusConflict, "EMAIL_EXISTS", "Email already registered", "")
			return
		}
		h.logger.Error("Failed to register user", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusCreated, user)
}

// Login handles user authentication
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", "")
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		validationErrors := make([]response.ValidationError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, response.ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}
		response.WriteValidationError(w, validationErrors, "")
		return
	}

	authResp, err := h.service.Login(r.Context(), req)
	if err != nil {
		if err == ErrInvalidCredentials {
			response.WriteError(w, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid email or password", "")
			return
		}
		h.logger.Error("Failed to login user", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, authResp)
}

// GetProfile handles getting user profile
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", "")
		return
	}

	user, err := h.service.GetProfile(r.Context(), userID)
	if err != nil {
		if err == ErrUserNotFound {
			response.WriteError(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found", "")
			return
		}
		h.logger.Error("Failed to get user profile", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, user)
}

// UpdateProfile handles updating user profile
func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by auth middleware)
	userID, ok := r.Context().Value("user_id").(string)
	if !ok {
		response.WriteError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", "")
		return
	}

	var req UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", "")
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		validationErrors := make([]response.ValidationError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, response.ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}
		response.WriteValidationError(w, validationErrors, "")
		return
	}

	user, err := h.service.UpdateProfile(r.Context(), userID, req)
	if err != nil {
		if err == ErrUserNotFound {
			response.WriteError(w, http.StatusNotFound, "USER_NOT_FOUND", "User not found", "")
			return
		}
		h.logger.Error("Failed to update user profile", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, user)
}

// RefreshToken handles token refresh
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body", "")
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		validationErrors := make([]response.ValidationError, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, response.ValidationError{
				Field:   err.Field(),
				Message: err.Tag(),
			})
		}
		response.WriteValidationError(w, validationErrors, "")
		return
	}

	authResp, err := h.service.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, "INVALID_TOKEN", "Invalid or expired refresh token", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, authResp)
}
