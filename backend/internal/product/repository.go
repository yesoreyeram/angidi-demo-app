package product

import (
	"context"
	"strings"
	"sync"
)

// Repository defines the interface for product data access
type Repository interface {
	Create(ctx context.Context, product *Product) error
	FindByID(ctx context.Context, id string) (*Product, error)
	List(ctx context.Context, filters ProductFilters) ([]*Product, int, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

// InMemoryRepository implements Repository using in-memory storage
type InMemoryRepository struct {
	products map[string]*Product
	mutex    sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		products: make(map[string]*Product),
	}
}

// Create creates a new product
func (r *InMemoryRepository) Create(ctx context.Context, product *Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.products[product.ID] = product
	return nil
}

// FindByID finds a product by ID
func (r *InMemoryRepository) FindByID(ctx context.Context, id string) (*Product, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	product, exists := r.products[id]
	if !exists {
		return nil, ErrProductNotFound
	}

	return product, nil
}

// List lists products with filters and pagination
func (r *InMemoryRepository) List(ctx context.Context, filters ProductFilters) ([]*Product, int, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	// Collect all products
	allProducts := make([]*Product, 0, len(r.products))
	for _, product := range r.products {
		allProducts = append(allProducts, product)
	}

	// Filter products
	filtered := make([]*Product, 0)
	for _, product := range allProducts {
		// Category filter
		if filters.CategoryID != "" && product.CategoryID != filters.CategoryID {
			continue
		}

		// Price filters
		if filters.MinPrice > 0 && product.Price < filters.MinPrice {
			continue
		}
		if filters.MaxPrice > 0 && product.Price > filters.MaxPrice {
			continue
		}

		// Search filter (search in name and description)
		if filters.Search != "" {
			searchLower := strings.ToLower(filters.Search)
			nameLower := strings.ToLower(product.Name)
			descLower := strings.ToLower(product.Description)
			if !strings.Contains(nameLower, searchLower) && !strings.Contains(descLower, searchLower) {
				continue
			}
		}

		filtered = append(filtered, product)
	}

	totalCount := len(filtered)

	// Pagination
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PageSize < 1 {
		filters.PageSize = 10
	}

	start := (filters.Page - 1) * filters.PageSize
	end := start + filters.PageSize

	if start >= totalCount {
		return []*Product{}, totalCount, nil
	}

	if end > totalCount {
		end = totalCount
	}

	paginated := filtered[start:end]
	return paginated, totalCount, nil
}

// Update updates a product
func (r *InMemoryRepository) Update(ctx context.Context, product *Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[product.ID]; !exists {
		return ErrProductNotFound
	}

	r.products[product.ID] = product
	return nil
}

// Delete deletes a product
func (r *InMemoryRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.products[id]; !exists {
		return ErrProductNotFound
	}

	delete(r.products, id)
	return nil
}
