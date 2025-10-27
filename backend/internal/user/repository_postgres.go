package user

import (
	"context"
	"errors"
	"fmt"
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

// Create creates a new user
func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, email, password_hash, name, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.pool.Exec(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.Name,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		// Check for unique constraint violation
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)" ||
			err.Error() == "SQLSTATE 23505" {
			return ErrEmailAlreadyExists
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// FindByID finds a user by ID
func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, email, password_hash, name, role, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`

	var user User
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}

// FindByEmail finds a user by email
func (r *PostgresRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, password_hash, name, role, created_at, updated_at
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`

	var user User
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	return &user, nil
}

// Update updates a user
func (r *PostgresRepository) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users
		SET name = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
	`

	result, err := r.pool.Exec(ctx, query,
		user.Name,
		time.Now(),
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}

// Delete soft-deletes a user
func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE users
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL
	`

	result, err := r.pool.Exec(ctx, query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrUserNotFound
	}

	return nil
}

// HasAdmin checks if any admin user exists in the repository
func (r *PostgresRepository) HasAdmin(ctx context.Context) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 FROM users
			WHERE role = 'admin' AND deleted_at IS NULL
		)
	`

	var exists bool
	err := r.pool.QueryRow(ctx, query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check admin existence: %w", err)
	}

	return exists, nil
}
