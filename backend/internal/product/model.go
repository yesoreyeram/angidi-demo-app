package product

import (
	"errors"
	"time"
)

var (
	// ErrProductNotFound is returned when a product is not found
	ErrProductNotFound = errors.New("product not found")
)

// Product represents a product entity
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	CategoryID  string    `json:"category_id"`
	ImageURL    string    `json:"image_url,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateProductRequest represents a product creation request
type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=255"`
	Description string  `json:"description" validate:"max=2000"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
	CategoryID  string  `json:"category_id" validate:"required"`
	ImageURL    string  `json:"image_url" validate:"omitempty,url"`
}

// UpdateProductRequest represents a product update request
type UpdateProductRequest struct {
	Name        string  `json:"name" validate:"omitempty,min=3,max=255"`
	Description string  `json:"description" validate:"omitempty,max=2000"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock" validate:"omitempty,gte=0"`
	CategoryID  string  `json:"category_id" validate:"omitempty"`
	ImageURL    string  `json:"image_url" validate:"omitempty,url"`
}

// ProductFilters represents filters for listing products
type ProductFilters struct {
	CategoryID string  `json:"category_id,omitempty"`
	MinPrice   float64 `json:"min_price,omitempty"`
	MaxPrice   float64 `json:"max_price,omitempty"`
	Search     string  `json:"search,omitempty"`
	Page       int     `json:"page" validate:"min=1"`
	PageSize   int     `json:"page_size" validate:"min=1,max=100"`
}

// ProductList represents a paginated list of products
type ProductList struct {
	Products   []*Product `json:"products"`
	TotalCount int        `json:"total_count"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
	TotalPages int        `json:"total_pages"`
}
