package product

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func setupTestService() Service {
	repo := NewInMemoryRepository()
	logger, _ := zap.NewDevelopment()
	return NewService(repo, logger)
}

func TestService_Create(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	req := CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       100,
		CategoryID:  "category-1",
		ImageURL:    "https://example.com/image.jpg",
	}

	product, err := service.Create(ctx, req)
	require.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, req.Name, product.Name)
	assert.Equal(t, req.Description, product.Description)
	assert.Equal(t, req.Price, product.Price)
	assert.Equal(t, req.Stock, product.Stock)
	assert.Equal(t, req.CategoryID, product.CategoryID)
	assert.Equal(t, req.ImageURL, product.ImageURL)
}

func TestService_GetByID(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Create a product first
	req := CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       100,
		CategoryID:  "category-1",
	}
	created, err := service.Create(ctx, req)
	require.NoError(t, err)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "existing product",
			id:      created.ID,
			wantErr: false,
		},
		{
			name:    "non-existent product",
			id:      "non-existent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := service.GetByID(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, product)
				assert.Equal(t, ErrProductNotFound, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.id, product.ID)
			}
		})
	}
}

func TestService_List(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Create multiple products
	products := []CreateProductRequest{
		{Name: "Product 1", Description: "Description 1", Price: 10.00, Stock: 100, CategoryID: "cat1"},
		{Name: "Product 2", Description: "Description 2", Price: 20.00, Stock: 50, CategoryID: "cat1"},
		{Name: "Product 3", Description: "Description 3", Price: 30.00, Stock: 75, CategoryID: "cat2"},
		{Name: "Electronics", Description: "Phone", Price: 500.00, Stock: 10, CategoryID: "cat3"},
	}

	for _, req := range products {
		_, err := service.Create(ctx, req)
		require.NoError(t, err)
	}

	tests := []struct {
		name        string
		filters     ProductFilters
		wantCount   int
		minProducts int
	}{
		{
			name: "all products - page 1",
			filters: ProductFilters{
				Page:     1,
				PageSize: 10,
			},
			minProducts: 4,
		},
		{
			name: "filter by category",
			filters: ProductFilters{
				CategoryID: "cat1",
				Page:       1,
				PageSize:   10,
			},
			minProducts: 2,
		},
		{
			name: "filter by price range",
			filters: ProductFilters{
				MinPrice: 15.00,
				MaxPrice: 35.00,
				Page:     1,
				PageSize: 10,
			},
			minProducts: 2,
		},
		{
			name: "search by name",
			filters: ProductFilters{
				Search:   "Electronics",
				Page:     1,
				PageSize: 10,
			},
			minProducts: 1,
		},
		{
			name: "pagination - page 2",
			filters: ProductFilters{
				Page:     2,
				PageSize: 2,
			},
			minProducts: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.List(ctx, tt.filters)
			require.NoError(t, err)
			assert.NotNil(t, result)
			assert.GreaterOrEqual(t, len(result.Products), tt.minProducts)
			assert.Equal(t, tt.filters.Page, result.Page)
			assert.Equal(t, tt.filters.PageSize, result.PageSize)
		})
	}
}

func TestService_Update(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Create a product first
	req := CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       100,
		CategoryID:  "category-1",
	}
	created, err := service.Create(ctx, req)
	require.NoError(t, err)

	tests := []struct {
		name    string
		id      string
		request UpdateProductRequest
		wantErr bool
	}{
		{
			name: "successful update",
			id:   created.ID,
			request: UpdateProductRequest{
				Name:  "Updated Product",
				Price: 149.99,
				Stock: 150,
			},
			wantErr: false,
		},
		{
			name: "non-existent product",
			id:   "non-existent-id",
			request: UpdateProductRequest{
				Name: "Updated Product",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updated, err := service.Update(ctx, tt.id, tt.request)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, updated)
				assert.Equal(t, ErrProductNotFound, err)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, updated)
				if tt.request.Name != "" {
					assert.Equal(t, tt.request.Name, updated.Name)
				}
				if tt.request.Price > 0 {
					assert.Equal(t, tt.request.Price, updated.Price)
				}
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	service := setupTestService()
	ctx := context.Background()

	// Create a product first
	req := CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       99.99,
		Stock:       100,
		CategoryID:  "category-1",
	}
	created, err := service.Create(ctx, req)
	require.NoError(t, err)

	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{
			name:    "delete existing product",
			id:      created.ID,
			wantErr: false,
		},
		{
			name:    "delete non-existent product",
			id:      "non-existent-id",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.Delete(ctx, tt.id)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, ErrProductNotFound, err)
			} else {
				require.NoError(t, err)

				// Verify product is deleted
				_, err := service.GetByID(ctx, tt.id)
				assert.Error(t, err)
				assert.Equal(t, ErrProductNotFound, err)
			}
		})
	}
}
