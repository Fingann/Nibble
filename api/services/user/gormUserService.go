package user

import (
	"fileslut/database"

	"fileslut/models"
	"fileslut/services"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var once sync.Once

type gormUserService struct {
	database.Repository[models.User]
	*gorm.DB
}

func NewGormUserService(db *gorm.DB) services.UserService {
	once.Do(func() {
		db.AutoMigrate(&models.User{})
	})
	return &gormUserService{
		DB: db,
	}
}

func (ur *gormUserService) Get(id uint) (models.User, error) {
	var user models.User
	result := ur.DB.First(&user, id)
	return user, result.Error
}

func (ur *gormUserService) Where(predicate database.Predicate[models.User]) ([]models.User, error) {
	var users []models.User
	result := ur.DB.Where(predicate).Find(&users)
	return users, result.Error
}

func (ur *gormUserService) All() ([]models.User, error) {
	var users []models.User
	result := ur.DB.Find(&users)
	return users, result.Error
}

func (ur *gormUserService) Create(user models.User) (models.User, error) {
	result := ur.DB.Create(&user)
	return user, result.Error
}

func (ur *gormUserService) Update(user models.User) (models.User, error) {
	result := ur.DB.Save(&user)
	return user, result.Error
}

func (ur *gormUserService) Delete(id uint) error {
	result := ur.DB.Delete(&models.User{}, id)
	return result.Error
}

func (ur *gormUserService) Login(username string, password string) bool {
	var user models.User
	result := ur.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return false
	}
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), bytePassword)
	if err != nil {
		return false
	}

	return true
}
