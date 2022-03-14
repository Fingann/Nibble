package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/Fingann/Nibble/database"
	"github.com/Fingann/Nibble/services"

	"github.com/Fingann/Nibble/endpoint"
)

func GinUriHandler[T any, K any](endpoint endpoint.Endpoint[T, K]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request T
		if err := c.ShouldBindUri(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := endpoint(c.Request.Context(), request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}
func GinJsonHandler[T any, K any](endpoint endpoint.Endpoint[T, K]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request T
		if err := c.ShouldBindUri(&request); err != nil {

		}

		if err := c.Copy().ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result, err := endpoint(c.Request.Context(), request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, result)

	}
}

func main() {

	db := database.NewGormPostgresDB()
	userRepository := database.NewGormUserRepository(db)

	userService := services.NewUserService(userRepository)
	jwtService := services.NewJWTAuthService()
	r := gin.Default()
	r.Use(gin.Recovery())
	v1 := r.Group("/api")

	setupUserRoutes(userService, jwtService, v1.Group("/users"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func setupUserRoutes(userService services.UserService, jwtService services.JWTService, users *gin.RouterGroup) {
	loginEndpoint := endpoint.MakeLoginEndpoint(userService, jwtService)
	registrationEndpoint := endpoint.MakeRegistrationEndpoint(userService)
	retrievalEndpoint := endpoint.MakeUserRetrieveEndpoint(userService)
	updateEndpoint := endpoint.MakeUserUpdateEndpoint(userService)
	deleteEndpoint := endpoint.MakeUserDeleteEndpoint(userService)

	users.POST("/login", GinJsonHandler(loginEndpoint))
	users.POST("/register", GinJsonHandler(registrationEndpoint))
	users.Use(AuthorizeJWT(jwtService))
	users.GET("/:Id", GinUriHandler(retrievalEndpoint))
	users.PUT("/:Id", GinJsonHandler(updateEndpoint))
	users.DELETE("/:Id", GinUriHandler(deleteEndpoint))
}

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}
		if len(authHeader) < len(BEARER_SCHEMA)+1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("claims", claims)
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
