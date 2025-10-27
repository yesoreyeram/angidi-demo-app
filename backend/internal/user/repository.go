package user

import (
	"context"
	"sync"
)

// Repository defines the interface for user data access
type Repository interface {
	Create(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
}

// InMemoryRepository implements Repository using in-memory storage
type InMemoryRepository struct {
	users     map[string]*User
	usersByEmail map[string]*User
	mutex     sync.RWMutex
}

// NewInMemoryRepository creates a new in-memory repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		users:     make(map[string]*User),
		usersByEmail: make(map[string]*User),
	}
}

// Create creates a new user
func (r *InMemoryRepository) Create(ctx context.Context, user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Check if email already exists
	if _, exists := r.usersByEmail[user.Email]; exists {
		return ErrEmailAlreadyExists
	}

	r.users[user.ID] = user
	r.usersByEmail[user.Email] = user
	return nil
}

// FindByID finds a user by ID
func (r *InMemoryRepository) FindByID(ctx context.Context, id string) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// FindByEmail finds a user by email
func (r *InMemoryRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	user, exists := r.usersByEmail[email]
	if !exists {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// Update updates a user
func (r *InMemoryRepository) Update(ctx context.Context, user *User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return ErrUserNotFound
	}

	r.users[user.ID] = user
	r.usersByEmail[user.Email] = user
	return nil
}

// Delete deletes a user
func (r *InMemoryRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	user, exists := r.users[id]
	if !exists {
		return ErrUserNotFound
	}

	delete(r.users, id)
	delete(r.usersByEmail, user.Email)
	return nil
}
