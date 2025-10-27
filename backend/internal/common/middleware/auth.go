package middleware

import (
	"context"
	"net/http"
	"strings"

	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/response"
	"go.uber.org/zap"
)

// Authentication middleware validates JWT tokens
func Authentication(jwtService *jwtPkg.Service, logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract token from Authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				response.WriteError(w, http.StatusUnauthorized, "MISSING_TOKEN", "Authorization header is required", "")
				return
			}

			// Check Bearer token format
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				response.WriteError(w, http.StatusUnauthorized, "INVALID_TOKEN_FORMAT", "Authorization header must be Bearer token", "")
				return
			}

			token := parts[1]

			// Validate token
			claims, err := jwtService.ValidateToken(token)
			if err != nil {
				if err == jwtPkg.ErrExpiredToken {
					response.WriteError(w, http.StatusUnauthorized, "EXPIRED_TOKEN", "Token has expired", "")
					return
				}
				response.WriteError(w, http.StatusUnauthorized, "INVALID_TOKEN", "Invalid token", "")
				return
			}

			// Add user information to context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "user_email", claims.Email)
			ctx = context.WithValue(ctx, "user_role", claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRole middleware checks if user has required role
func RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value("user_role").(string)
			if !ok {
				response.WriteError(w, http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized", "")
				return
			}

			if role != requiredRole {
				response.WriteError(w, http.StatusForbidden, "FORBIDDEN", "Insufficient permissions", "")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
