package services

import (
	"github.com/Fingann/Nibble/database"

	"github.com/Fingann/Nibble/models"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

var once sync.Once

type userService struct {
	database.Repository[models.User]
}

func NewUserService(repository database.Repository[models.User]) UserService {
	return &userService{
		Repository: repository,
	}
}

func (ur *userService) Register(email string, username string, password string) (uint, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user := models.User{
		Email:        email,
		Username:     username,
		PasswordHash: string(pass),
	}
	user, err = ur.Create(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (ur *userService) Login(username string, password string) (bool, error) {
	user := models.User{Username: username}
	result, err := ur.Repository.First(user)
	if err != nil {
		return false, err
	}
	bytePassword := []byte(password)
	err = bcrypt.CompareHashAndPassword([]byte(result.PasswordHash), bytePassword)
	if err != nil {
		return false, err
	}

	return true, nil
}
