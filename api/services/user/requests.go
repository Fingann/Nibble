package user

import (
	"fileslut/common"
	"fileslut/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserRegistrationRequest struct {
	common.Validator[UserRegistrationRequest]
	Username string      `json:"username" binding:"required" validate:"required"`
	Email    string      `json:"email" binding:"required" validate:"email"`
	Password string      `json:"password" binding:"required" validate:"required"`
	User     models.User `json:"-"`
}

func (ur *UserRegistrationRequest) Populate(c *gin.Context) error {
	err := ur.Validate(c, ur)
	if err != nil {
		return err
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	ur.User = models.User{
		Username:     ur.Username,
		Email:        ur.Email,
		PasswordHash: string(pass),
	}
	return nil
}

type UserLoginRequest struct {
	common.Validator[UserLoginRequest]
	Username string `json:"username" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"required"`
}

func (lr *UserLoginRequest) Populate(c *gin.Context) error {
	return lr.Validate(c, lr)
}
