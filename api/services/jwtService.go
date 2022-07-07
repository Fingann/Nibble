package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthCustomClaims struct {
	UserId uint `json:"userId"`
	IsUser bool `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func NewJWTAuthService(secret, issuer string) JWTService {
	return &jwtServices{
		secretKey: secret,
		issure:    issuer,
	}
}

func (service *jwtServices) GenerateToken(userId uint, isUser bool) (string, error) {
	claims := &AuthCustomClaims{
		userId,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}
func (service *jwtServices) GetClaims(token *jwt.Token) (*AuthCustomClaims, error) {
	if claims, ok := token.Claims.(*AuthCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid token")
}
