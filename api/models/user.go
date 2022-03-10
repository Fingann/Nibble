package models

import (
	"fileslut/database"
	"sync"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var once sync.Once

type User struct {
	gorm.Model
	Username     string
	Email        string
	PasswordHash string
}

type UserService interface {
	database.Repository[User]
	Login(username string, password string) bool
}

type gormUserRepository struct {
	database.Repository[User]
	*gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserService {
	once.Do(func() {
		db.AutoMigrate(&User{})
	})
	return &gormUserRepository{
		DB: db,
	}
}

func (ur *gormUserRepository) Get(id uint) (User, error) {
	var user User
	result := ur.DB.First(&user, id)
	return user, result.Error
}

func (ur *gormUserRepository) Where(predicate database.Predicate[User]) ([]User, error) {
	var users []User
	result := ur.DB.Where(predicate).Find(&users)
	return users, result.Error
}

func (ur *gormUserRepository) All() ([]User, error) {
	var users []User
	result := ur.DB.Find(&users)
	return users, result.Error
}

func (ur *gormUserRepository) Create(user User) (User, error) {
	result := ur.DB.Create(&user)
	return user, result.Error
}

func (ur *gormUserRepository) Update(user User) (User, error) {
	result := ur.DB.Save(&user)
	return user, result.Error
}

func (ur *gormUserRepository) Delete(id uint) error {
	result := ur.DB.Delete(&User{}, id)
	return result.Error
}

func (ur *gormUserRepository) Login(username string, password string) bool {
	var user User
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
