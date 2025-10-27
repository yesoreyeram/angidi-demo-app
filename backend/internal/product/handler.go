package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/yesoreyeram/angidi-demo-app/backend/pkg/response"
	"go.uber.org/zap"
)

// Handler handles HTTP requests for product operations
type Handler struct {
	service   Service
	validator *validator.Validate
	logger    *zap.Logger
}

// NewHandler creates a new product handler
func NewHandler(service Service, logger *zap.Logger) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
		logger:    logger,
	}
}

// Create handles product creation
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateProductRequest
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

	product, err := h.service.Create(r.Context(), req)
	if err != nil {
		h.logger.Error("Failed to create product", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusCreated, product)
}

// GetByID handles getting a product by ID
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Product ID is required", "")
		return
	}

	product, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if err == ErrProductNotFound {
			response.WriteError(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", "")
			return
		}
		h.logger.Error("Failed to get product", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, product)
}

// List handles listing products with filters
func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	filters := ProductFilters{
		CategoryID: r.URL.Query().Get("category_id"),
		Search:     r.URL.Query().Get("search"),
		Page:       1,
		PageSize:   10,
	}

	if pageStr := r.URL.Query().Get("page"); pageStr != "" {
		if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
			filters.Page = page
		}
	}

	if pageSizeStr := r.URL.Query().Get("page_size"); pageSizeStr != "" {
		if pageSize, err := strconv.Atoi(pageSizeStr); err == nil && pageSize > 0 {
			filters.PageSize = pageSize
		}
	}

	if minPriceStr := r.URL.Query().Get("min_price"); minPriceStr != "" {
		if minPrice, err := strconv.ParseFloat(minPriceStr, 64); err == nil && minPrice >= 0 {
			filters.MinPrice = minPrice
		}
	}

	if maxPriceStr := r.URL.Query().Get("max_price"); maxPriceStr != "" {
		if maxPrice, err := strconv.ParseFloat(maxPriceStr, 64); err == nil && maxPrice >= 0 {
			filters.MaxPrice = maxPrice
		}
	}

	productList, err := h.service.List(r.Context(), filters)
	if err != nil {
		h.logger.Error("Failed to list products", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, productList)
}

// Update handles updating a product
func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Product ID is required", "")
		return
	}

	var req UpdateProductRequest
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

	product, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		if err == ErrProductNotFound {
			response.WriteError(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", "")
			return
		}
		h.logger.Error("Failed to update product", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	response.WriteSuccess(w, http.StatusOK, product)
}

// Delete handles deleting a product
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		response.WriteError(w, http.StatusBadRequest, "INVALID_REQUEST", "Product ID is required", "")
		return
	}

	err := h.service.Delete(r.Context(), id)
	if err != nil {
		if err == ErrProductNotFound {
			response.WriteError(w, http.StatusNotFound, "PRODUCT_NOT_FOUND", "Product not found", "")
			return
		}
		h.logger.Error("Failed to delete product", zap.Error(err))
		response.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error", "")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
