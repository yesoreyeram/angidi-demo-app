package gateway

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	jwtPkg "github.com/yesoreyeram/angidi-demo-app/backend/pkg/jwt"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/response"
	"go.uber.org/zap"

	"github.com/yesoreyeram/angidi-demo-app/backend/internal/common/middleware"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/product"
	"github.com/yesoreyeram/angidi-demo-app/backend/internal/user"
)

// Router sets up the HTTP router with all routes and middleware
func Router(
	userHandler *user.Handler,
	productHandler *product.Handler,
	jwtService *jwtPkg.Service,
	logger *zap.Logger,
) http.Handler {
	r := chi.NewRouter()

	// Global middleware
	r.Use(chiMiddleware.Compress(5))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.RateLimit(100)) // 100 requests per minute

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"X-Request-ID"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		response.WriteSuccess(w, http.StatusOK, map[string]interface{}{
			"status": "healthy",
			"timestamp": "2025-10-27T03:00:00Z",
		})
	})

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Public user routes
		r.Post("/users/register", userHandler.Register)
		r.Post("/users/login", userHandler.Login)
		r.Post("/users/refresh-token", userHandler.RefreshToken)

		// Public product routes
		r.Get("/products", productHandler.List)
		r.Get("/products/{id}", productHandler.GetByID)

		// Protected routes (require authentication)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Authentication(jwtService, logger))

			// User profile routes
			r.Get("/users/me", userHandler.GetProfile)
			r.Put("/users/me", userHandler.UpdateProfile)

			// Admin-only product routes
			r.Group(func(r chi.Router) {
				r.Use(middleware.RequireRole("admin"))

				r.Post("/products", productHandler.Create)
				r.Put("/products/{id}", productHandler.Update)
				r.Delete("/products/{id}", productHandler.Delete)
			})
		})
	})

	return r
}
