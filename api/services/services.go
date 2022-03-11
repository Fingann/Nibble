package services

import (
	"fileslut/database"
	"fileslut/models"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

const (
	USER_SERVICE_KEY = "user_service"
	JWT_SERVICE_KEY  = "jwt_service"
)

func Get[T any](c *gin.Context, key string) T {
	return c.MustGet(key).(T)
}

type UserService interface {
	database.Repository[models.User]
	Login(username string, password string) bool
}

type Request[T any] interface {
	Populate(c *gin.Context) error
}

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}
