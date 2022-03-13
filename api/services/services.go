package services

import (
	"fileslut/database"
	"fileslut/models"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	database.Repository[models.User]
	Login(username string, password string) bool
	Register(email string, username string, password string) (uint, error)
}

type Request[T any] interface {
	Populate(c *gin.Context) error
}

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
