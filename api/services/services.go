package services

import (
	"context"

	"github.com/Fingann/Nibble/database"
	"github.com/Fingann/Nibble/models"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	database.Repository[models.User]
	Login(username string, password string) bool
	Register(email string, username string, password string) (uint, error)
}

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type MinioService interface {
	GetBucket(ctx context.Context, bucketName string) ([]string, error)
	GetObject(ctx context.Context, bucketName string, objectName string) ([]byte, error)
	PutObject(ctx context.Context, bucketName string, objectName string, data []byte) error
	DeleteObject(ctx context.Context, bucketName string, objectName string) error
}
