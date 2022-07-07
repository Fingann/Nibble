package endpoint

import (
	"net/http"

	"github.com/Fingann/notifyGame-api/models"
	"github.com/Fingann/notifyGame-api/services"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupUserRoutes(userService services.UserService, jwtService services.JWTService, users *gin.RouterGroup) {
	users.POST("/login", makeLoginEndpoint(userService, jwtService))
	users.POST("/register", makeRegistrationEndpoint(userService))
	users.GET("/:Id", makeUserGetEndpoint(userService))
	users.PUT("/:Id", makeUserUpdateEndpoint(userService))
	users.DELETE("/:Id", makeUserDeleteEndpoint(userService))
}

func makeRegistrationEndpoint(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required" `
			Email    string `json:"email" binding:"required,email" `
			Password string `json:"password" binding:"required" `
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := userService.Register(request.Email, request.Username, request.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func makeLoginEndpoint(userService services.UserService, jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Username string `json:"username" binding:"required" `
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userService.Login(request.Username, request.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := jwtService.GenerateToken(user.ID, true)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	}

}

func makeUserGetEndpoint(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Id uint `json:"id" form:"Id" binding:"required"`
		}

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userService.Get(request.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":       user.ID,
			"Username": user.Username,
			"Email":    user.Email,
		})
	}
}

func makeUserUpdateEndpoint(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Id       uint   `json:"id" form:"Id" binding:"required"`
			Username string `json:"username"`
			Email    string `json:"email" binding:"email"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user := models.User{
			Model:    gorm.Model{ID: request.Id},
			Username: request.Username,
			Email:    request.Email,
		}
		result, err := userService.Update(user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{
				"id":       result.ID,
				"Username": result.Email,
				"Email":    result.Email,
			})
	}
}

func makeUserDeleteEndpoint(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			ID uint `json:"id" binding:"required"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userService.Delete(request.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id": request.ID,
		})
	}
}
