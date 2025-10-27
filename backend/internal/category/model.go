package category

import (
	"errors"
	"time"
)

var (
	// ErrCategoryNotFound is returned when a category is not found
	ErrCategoryNotFound = errors.New("category not found")
	// ErrCategoryNameExists is returned when category name already exists
	ErrCategoryNameExists = errors.New("category name already exists")
)

// Category represents a product category
type Category struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	ParentID    *string    `json:"parent_id,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// CreateCategoryRequest represents a request to create a category
type CreateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Description string  `json:"description,omitempty" validate:"max=500"`
	ParentID    *string `json:"parent_id,omitempty"`
}

// UpdateCategoryRequest represents a request to update a category
type UpdateCategoryRequest struct {
	Name        string  `json:"name" validate:"required,min=1,max=100"`
	Description string  `json:"description,omitempty" validate:"max=500"`
	ParentID    *string `json:"parent_id,omitempty"`
}
