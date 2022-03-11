package routers

import (
	"fileslut/services"
	"fileslut/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func UsersRegistration(c *gin.Context) {
	var request user.UserRegistrationRequest
	err := request.Populate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.Get[services.UserService](c, services.USER_SERVICE_KEY)
	createdUser, err := userService.Create(request.User)

	c.JSON(http.StatusCreated, user.UserRegistrationResponse{ID: createdUser.ID})
}

func UsersLogin(c *gin.Context) {
	var request user.UserLoginRequest
	err := request.Populate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userService := services.Get[services.UserService](c, services.USER_SERVICE_KEY)
	exists := userService.Login(request.Username, request.Password)
	if exists {
		jwtService := services.Get[services.JWTService](c, services.JWT_SERVICE_KEY)
		token, err := jwtService.GenerateToken(request.Username, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return

	}
	c.JSON(http.StatusCreated, gin.H{"login": "WRONG CREDENTIALS"})

}

func UserRetrieve(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}

func UserUpdate(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}
