package database

import (
	"time"
)

type Predicate[T any] func(item T) bool

type Repository[T any] interface {
	// Get the a single item by ID
	Get(id uint) (T, error)
	// Get all items where predicate is true
	Where(predicate Predicate[T]) ([]T, error)
	// Get all items
	All() ([]T, error)
	// Create a new item
	Create(T) (T, error)
	// Update an existing item
	Update(T) (T, error)
	// Delete an existing item
	Delete(id uint) error
}

type EntityBase struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
