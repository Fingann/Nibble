package database

import (
	"github.com/Fingann/Nibble/models"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once

type gormUserRepository[T any] struct {
	Repository[T]
	*gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *gormUserRepository[T] {
	once.Do(func() {
		var t T
		db.AutoMigrate(&t)
	})
	return &gormUserRepository[T]{
		DB: db,
	}
}

func (ur *gormUserRepository[T]) Get(id uint) (models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	return user, result.Error
}

func (ur *gormUserRepository[T]) Where(item T) ([]T, error) {
	var items []T
	result := ur.DB.Find(&items, &item)
	return items, result.Error
}

func (ur *gormUserRepository[T]) All() ([]T, error) {
	var items []T
	result := ur.DB.Find(&items)
	return items, result.Error
}

func (ur *gormUserRepository[T]) Create(item T) (T, error) {
	result := ur.DB.Create(&item)
	return item, result.Error
}

func (ur *gormUserRepository[T]) Update(item T) (T, error) {
	result := ur.DB.Save(&item)
	return item, result.Error
}

func (ur *gormUserRepository[T]) Delete(id uint) error {
	var item T
	result := ur.DB.Delete(&item, id)
	return result.Error
}

func (ur *gormUserRepository[T]) First(usr T) (T, error) {
	var item T
	result := ur.DB.First(&item, &usr)
	return item, result.Error
}
