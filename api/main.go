package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Fingann/notifyGame-api/database"
	"github.com/Fingann/notifyGame-api/models"
	"github.com/Fingann/notifyGame-api/services"

	"github.com/Fingann/notifyGame-api/endpoint"
)

func main() {

	db, err := database.NewGormPostgresDB()
	if err != nil {
		log.Fatalln("Unable to initialize Postges database, err:", err)
	}
	db.Migrator().AutoMigrate(&models.User{}, &models.Group{}, &models.Member{})
	userRepository := database.NewGormRepository[models.User](db)
	groupRepository := database.NewGormRepository[models.Group](db)
	userService := services.NewUserService(userRepository)
	groupService := services.NewGroupService(groupRepository)
	// change jwt secret for production
	jwtService := services.NewJWTAuthService("secret", "nibble")

	r := gin.Default()
	r.Use(gin.Recovery())
	v1 := r.Group("/api")
	users := v1.Group("/users")
	users.Use(JWTMiddleware(jwtService))

	SetupTestuser(userService, groupService)

	endpoint.SetupUserRoutes(userService, jwtService, users)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupTestuser(userService services.UserService, groupService services.GroupService) {
	user := models.User{
		Email:        "test@email.com",
		Username:     "test",
		PasswordHash: "testHash",
	}
	dbUser, err := userService.Register(user.Email, user.Username, user.PasswordHash)
	if err != nil {
		// error for everyting but User already exists
		if !errors.Is(err, services.ErrUserExists) {
			log.Fatalln(err)
		}
		// if user exists, we can get the user from the database
		dbUser, err = userService.Get(1)
		if err != nil {
			log.Fatalln(err)
		}
	}

	group := models.Group{
		Name:  "test",
		Owner: *dbUser,
		Members: []models.Member{
			{
				UserId: dbUser.ID,
			},
		}}
	createdGroup, err := groupService.Create(group)
	if err != nil {
		log.Fatalln("", err)
	}
	log.Println(createdGroup.Name)

}

func JWTMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}
		if len(authHeader) < len(BEARER_SCHEMA)+1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		if token.Valid {
			claims := token.Claims
			c.Set("claims", claims.(services.AuthCustomClaims))
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}
