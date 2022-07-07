package services

import (
	"github.com/Fingann/notifyGame-api/database"
	"github.com/Fingann/notifyGame-api/models"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Services struct {
	UserService  UserService
	GroupService GroupService
	GameService  GameService
	JWTService   JWTService
}

type UserService interface {
	Database() *gorm.DB
	Login(username string, password string) (*models.User, error)
	Register(email string, username string, password string) (*models.User, error)
}

type GameService interface {
	Database() *gorm.DB
}
type GroupService interface {
	Database() *gorm.DB
}

//jwt service
type JWTService interface {
	GenerateToken(userId uint, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
