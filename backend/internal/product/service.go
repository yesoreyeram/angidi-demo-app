package product

import (
	"context"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Service defines the interface for product business logic
type Service interface {
	Create(ctx context.Context, req CreateProductRequest) (*Product, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, filters ProductFilters) (*ProductList, error)
	Update(ctx context.Context, id string, req UpdateProductRequest) (*Product, error)
	Delete(ctx context.Context, id string) error
}

// service implements Service
type service struct {
	repo   Repository
	logger *zap.Logger
}

// NewService creates a new product service
func NewService(repo Repository, logger *zap.Logger) Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}

// Create creates a new product
func (s *service) Create(ctx context.Context, req CreateProductRequest) (*Product, error) {
	s.logger.Info("Creating new product", zap.String("name", req.Name))

	now := time.Now()
	product := &Product{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		CategoryID:  req.CategoryID,
		ImageURL:    req.ImageURL,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		s.logger.Error("Failed to create product", zap.Error(err))
		return nil, err
	}

	s.logger.Info("Product created successfully", zap.String("product_id", product.ID))
	return product, nil
}

// GetByID retrieves a product by ID
func (s *service) GetByID(ctx context.Context, id string) (*Product, error) {
	s.logger.Debug("Getting product", zap.String("product_id", id))

	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get product", zap.String("product_id", id), zap.Error(err))
		return nil, err
	}

	return product, nil
}

// List retrieves a list of products with filters
func (s *service) List(ctx context.Context, filters ProductFilters) (*ProductList, error) {
	s.logger.Debug("Listing products", zap.Any("filters", filters))

	// Set defaults
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PageSize < 1 {
		filters.PageSize = 10
	}
	if filters.PageSize > 100 {
		filters.PageSize = 100
	}

	products, totalCount, err := s.repo.List(ctx, filters)
	if err != nil {
		s.logger.Error("Failed to list products", zap.Error(err))
		return nil, err
	}

	totalPages := (totalCount + filters.PageSize - 1) / filters.PageSize

	return &ProductList{
		Products:   products,
		TotalCount: totalCount,
		Page:       filters.Page,
		PageSize:   filters.PageSize,
		TotalPages: totalPages,
	}, nil
}

// Update updates a product
func (s *service) Update(ctx context.Context, id string, req UpdateProductRequest) (*Product, error) {
	s.logger.Info("Updating product", zap.String("product_id", id))

	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to find product", zap.String("product_id", id), zap.Error(err))
		return nil, err
	}

	// Update only provided fields
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}
	if req.CategoryID != "" {
		product.CategoryID = req.CategoryID
	}
	if req.ImageURL != "" {
		product.ImageURL = req.ImageURL
	}

	product.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, product); err != nil {
		s.logger.Error("Failed to update product", zap.Error(err))
		return nil, err
	}

	s.logger.Info("Product updated successfully", zap.String("product_id", product.ID))
	return product, nil
}

// Delete deletes a product
func (s *service) Delete(ctx context.Context, id string) error {
	s.logger.Info("Deleting product", zap.String("product_id", id))

	if err := s.repo.Delete(ctx, id); err != nil {
		s.logger.Error("Failed to delete product", zap.String("product_id", id), zap.Error(err))
		return err
	}

	s.logger.Info("Product deleted successfully", zap.String("product_id", id))
	return nil
}
