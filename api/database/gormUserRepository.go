package database

import (
	"github.com/Fingann/Nibble/models"
	"sync"

	"gorm.io/gorm"
)

var once sync.Once

type gormUserRepository struct {
	Repository[models.User]
	*gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *gormUserRepository {
	once.Do(func() {
		db.AutoMigrate(&models.User{})
	})
	return &gormUserRepository{
		DB: db,
	}
}

func (ur *gormUserRepository) Get(id uint) (models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	return user, result.Error
}

func (ur *gormUserRepository) Where(usr models.User) ([]models.User, error) {
	var users []models.User
	result := ur.DB.Find(&users, &usr)
	return users, result.Error
}

func (ur *gormUserRepository) All() ([]models.User, error) {
	var users []models.User
	result := ur.DB.Find(&users)
	return users, result.Error
}

func (ur *gormUserRepository) Create(user models.User) (models.User, error) {
	result := ur.DB.Create(&user)
	return user, result.Error
}

func (ur *gormUserRepository) Update(user models.User) (models.User, error) {
	result := ur.DB.Save(&user)
	return user, result.Error
}

func (ur *gormUserRepository) Delete(id uint) error {
	result := ur.DB.Delete(&models.User{}, id)
	return result.Error
}

func (ur *gormUserRepository) First(usr models.User) (models.User, error) {
	var user models.User
	result := ur.DB.First(&user, &usr)
	return user, result.Error
}
