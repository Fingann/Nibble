package services

import (
	"errors"

	"github.com/Fingann/notifyGame-api/database"
	"github.com/Fingann/notifyGame-api/models"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

var ErrUserExists = errors.New("user already exists")

type userService struct {
	database *gorm.DB
}

func NewUserService(database *gorm.DB) UserService {
	return &userService{
		database: database,
	}
}

func (us *userService) Database() *gorm.DB {
	return us.database
}

func (ur *userService) Register(email string, username string, password string) (*models.User, error) {
	// check if user exists by email or username
	tx := ur.Database().Where(models.User{Username: username}).Or(models.User{Email: email})
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, ErrUserExists
		}

		return nil, tx.Error
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userToCreate := models.User{
		Email:        email,
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	tx = ur.Database().Create(&userToCreate)
	if tx.Error != nil {
		return nil, err
	}
	return &userToCreate, nil
}

func (ur *userService) Login(username string, password string) (*models.User, error) {
	user := &models.User{}

	tx := ur.database.First(user, &models.User{Username: username})
	if tx.Error != nil {
		return nil, tx.Error
	}
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), bytePassword)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func emailOrUserNameExists(username string, email string, ur *userService) (bool, error) {
	return tx.Error != gorm.ErrRecordNotFound

}

var ErrEmailOrUserNameExists = errors.New("Email or username already exists")
var ErrUsernameExists = errors.New("Username already exists")
var ErrEmailExists = errors.New("Email already exists")
