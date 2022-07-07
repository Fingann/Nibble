package endpoint

import (
	"net/http"

	"github.com/Fingann/notifyGame-api/models"
	"github.com/Fingann/notifyGame-api/services"
	"github.com/gin-gonic/gin"
)

func SetupGroupRoutes(repository services.GroupService, jwtService services.JWTService, users *gin.RouterGroup) {
	users.POST("/create", makeGroupCreateEndpoint(repository))
	users.GET("/:Id", makeGetEndpoint(repository))
	//users.PUT("/:Id", makeUserUpdateEndpoint(repository))
	//users.DELETE("/:Id", makeUserDeleteEndpoint(repository))
}

func makeGetEndpoint(groupService services.GroupService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Id uint `json:"id" form:"Id" binding:"required"`
		}

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		group, err := groupService.Get(request.Id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"id":      group.ID,
			"name":    group.Name,
			"members": group.Members,
		})
	}
}
func makeGroupCreateEndpoint(groupService services.GroupService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			Name string `json:"name"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		group := models.Group{
			Name: request.Name,
		}
		createdGroup, err := groupService.Create(group)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK,
			gin.H{
				"id":       createdGroup.ID,
				"Username": createdGroup.Name,
				"Email":    createdGroup.Name,
			})
	}
}
