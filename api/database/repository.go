package database

type Repository[T any] interface {
	// Get the a single item by ID
	Get(id uint) (*T, error)
	//
	Where(T) ([]T, error)
	// Get all items
	All() ([]T, error)
	// Create a new item
	Create(T) (*T, error)
	// Update an existing item
	Update(T) (*T, error)
	// Delete an existing item
	Delete(id uint) error
	// Get the first item where predicate is true
	First(T) (*T, error)
	// Query the database with a custom query
	Query(query any, args ...any) ([]T, error)
}
