package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UserRegister(router *gin.RouterGroup) {
	router.GET("/", UserRetrieve)
	router.PUT("/", UserUpdate)
}

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/:username", ProfileRetrieve)

}

func ProfileRetrieve(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}

func UsersRegistration(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})
}

func UsersLogin(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}

func UserRetrieve(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}

func UserUpdate(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"user": "NOT IMPLEMENTED"})

}
