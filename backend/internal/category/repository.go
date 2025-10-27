package category

import (
	"context"
)

// Repository defines the interface for category data access
type Repository interface {
	Create(ctx context.Context, category *Category) error
	FindByID(ctx context.Context, id string) (*Category, error)
	FindByName(ctx context.Context, name string) (*Category, error)
	List(ctx context.Context) ([]*Category, error)
	Update(ctx context.Context, category *Category) error
	Delete(ctx context.Context, id string) error
}
