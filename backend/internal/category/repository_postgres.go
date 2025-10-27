package category

import (
	"context"
	"errors"
	"fmt"

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

// Create creates a new category
func (r *PostgresRepository) Create(ctx context.Context, category *Category) error {
	query := `
		INSERT INTO categories (id, name, description, parent_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.pool.Exec(ctx, query,
		category.ID,
		category.Name,
		category.Description,
		category.ParentID,
		category.CreatedAt,
		category.UpdatedAt,
	)

	if err != nil {
		// Check for unique constraint violation
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"categories_name_key\" (SQLSTATE 23505)" {
			return ErrCategoryNameExists
		}
		return fmt.Errorf("failed to create category: %w", err)
	}

	return nil
}

// FindByID finds a category by ID
func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*Category, error) {
	query := `
		SELECT id, name, description, parent_id, created_at, updated_at
		FROM categories
		WHERE id = $1
	`

	var category Category
	var parentID *string
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&parentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	category.ParentID = parentID
	return &category, nil
}

// FindByName finds a category by name
func (r *PostgresRepository) FindByName(ctx context.Context, name string) (*Category, error) {
	query := `
		SELECT id, name, description, parent_id, created_at, updated_at
		FROM categories
		WHERE name = $1
	`

	var category Category
	var parentID *string
	err := r.pool.QueryRow(ctx, query, name).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&parentID,
		&category.CreatedAt,
		&category.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrCategoryNotFound
		}
		return nil, fmt.Errorf("failed to find category: %w", err)
	}

	category.ParentID = parentID
	return &category, nil
}

// List lists all categories
func (r *PostgresRepository) List(ctx context.Context) ([]*Category, error) {
	query := `
		SELECT id, name, description, parent_id, created_at, updated_at
		FROM categories
		ORDER BY name ASC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}
	defer rows.Close()

	var categories []*Category
	for rows.Next() {
		var category Category
		var parentID *string
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&parentID,
			&category.CreatedAt,
			&category.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		category.ParentID = parentID
		categories = append(categories, &category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating categories: %w", err)
	}

	return categories, nil
}

// Update updates a category
func (r *PostgresRepository) Update(ctx context.Context, category *Category) error {
	query := `
		UPDATE categories
		SET name = $1, description = $2, parent_id = $3, updated_at = $4
		WHERE id = $5
	`

	result, err := r.pool.Exec(ctx, query,
		category.Name,
		category.Description,
		category.ParentID,
		category.UpdatedAt,
		category.ID,
	)

	if err != nil {
		// Check for unique constraint violation
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"categories_name_key\" (SQLSTATE 23505)" {
			return ErrCategoryNameExists
		}
		return fmt.Errorf("failed to update category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrCategoryNotFound
	}

	return nil
}

// Delete deletes a category
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM categories
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrCategoryNotFound
	}

	return nil
}
