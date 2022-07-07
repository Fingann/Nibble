package database

import (
	"gorm.io/gorm"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type gormRepository[T any] struct {
	*gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) Repository[T] {

	var t T
	db.AutoMigrate(&t)

	return &gormRepository[T]{
		DB: db,
	}
}

func (ur *gormRepository[T]) Get(id uint) (*T, error) {
	var item T
	result := ur.DB.First(&item, id)
	return &item, result.Error
}
func (ur *gormRepository[T]) Query(query any, args ...any) ([]T, error) {
	var items []T
	result := ur.DB.Find(&items, query, args)
	return items, result.Error
}

func (ur *gormRepository[T]) Where(item T) ([]T, error) {
	var items []T
	result := ur.DB.Find(&items, &item)
	return items, result.Error
}

func (ur *gormRepository[T]) All() ([]T, error) {
	var items []T
	result := ur.DB.Find(&items)
	return items, result.Error
}

func (ur *gormRepository[T]) Create(item T) (*T, error) {
	result := ur.DB.Create(&item)
	return &item, result.Error
}

func (ur *gormRepository[T]) Update(item T) (*T, error) {
	result := ur.DB.Save(&item)
	return &item, result.Error
}

func (ur *gormRepository[T]) Delete(id uint) error {
	var item T
	result := ur.DB.Delete(&item, id)
	return result.Error
}

func (ur *gormRepository[T]) First(usr T) (*T, error) {
	var item T
	result := ur.DB.First(&item, &usr)
	return &item, result.Error
}
