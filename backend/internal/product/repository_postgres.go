package product

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository implements Repository using PostgreSQL
type PostgresRepository struct {
	pool *pgxpool.Pool
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		pool: pool,
	}
}

// Create creates a new product
func (r *PostgresRepository) Create(ctx context.Context, product *Product) error {
	query := `
		INSERT INTO products (id, name, description, price, stock, category_id, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := r.pool.Exec(ctx, query,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CategoryID,
		product.ImageURL,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

// FindByID finds a product by ID
func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*Product, error) {
	query := `
		SELECT id, name, description, price, stock, category_id, image_url, created_at, updated_at
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`

	var product Product
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Stock,
		&product.CategoryID,
		&product.ImageURL,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProductNotFound
		}
		return nil, fmt.Errorf("failed to find product: %w", err)
	}

	return &product, nil
}

// List lists products with filters and pagination
func (r *PostgresRepository) List(ctx context.Context, filters ProductFilters) ([]*Product, int, error) {
	// Build WHERE clause dynamically
	var whereClauses []string
	var args []interface{}
	argPosition := 1

	// Always filter out soft-deleted products
	whereClauses = append(whereClauses, "deleted_at IS NULL")

	// Category filter
	if filters.CategoryID != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("category_id = $%d", argPosition))
		args = append(args, filters.CategoryID)
		argPosition++
	}

	// Price filters
	if filters.MinPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("price >= $%d", argPosition))
		args = append(args, filters.MinPrice)
		argPosition++
	}
	if filters.MaxPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("price <= $%d", argPosition))
		args = append(args, filters.MaxPrice)
		argPosition++
	}

	// Search filter (name and description)
	if filters.Search != "" {
		searchPattern := "%" + strings.ToLower(filters.Search) + "%"
		whereClauses = append(whereClauses, fmt.Sprintf("(LOWER(name) LIKE $%d OR LOWER(description) LIKE $%d)", argPosition, argPosition))
		args = append(args, searchPattern)
		argPosition++
	}

	whereClause := ""
	if len(whereClauses) > 0 {
		whereClause = "WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Get total count
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM products %s", whereClause)
	var totalCount int
	err := r.pool.QueryRow(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// Set defaults for pagination
	if filters.Page < 1 {
		filters.Page = 1
	}
	if filters.PageSize < 1 {
		filters.PageSize = 10
	}

	// Calculate pagination
	offset := (filters.Page - 1) * filters.PageSize

	// Get paginated results
	query := fmt.Sprintf(`
		SELECT id, name, description, price, stock, category_id, image_url, created_at, updated_at
		FROM products
		%s
		ORDER BY created_at DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argPosition, argPosition+1)
	
	args = append(args, filters.PageSize, offset)

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list products: %w", err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		var product Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.CategoryID,
			&product.ImageURL,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("error iterating products: %w", err)
	}

	return products, totalCount, nil
}

// Update updates a product
func (r *PostgresRepository) Update(ctx context.Context, product *Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, stock = $4, category_id = $5, image_url = $6, updated_at = $7
		WHERE id = $8 AND deleted_at IS NULL
	`

	result, err := r.pool.Exec(ctx, query,
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
		product.CategoryID,
		product.ImageURL,
		time.Now(),
		product.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}

// Delete soft-deletes a product
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE products
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	result, err := r.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrProductNotFound
	}

	return nil
}
