package main

import (
	"context"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"fileslut/common"
	"fileslut/database"
	"fileslut/models"
	"fileslut/routers"
	"fileslut/services"

	"fileslut/services/auth"
	"fileslut/services/user"
)

func service[T any](service T, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(key, service)
		c.Next()
	}
}
func main() {

	db := database.NewGormPostgresDB()
	userService := user.NewGormUserService(db)
	jwtService := auth.NewJWTAuthService()

	r := gin.Default()
	r.Use(service(jwtService, services.JWT_SERVICE_KEY))

	v1 := r.Group("/api")
	users := v1.Group("/users")
	users.Use(service(userService, services.USER_SERVICE_KEY))

	routers.UsersRegister(users)
	v1.Use(common.AuthMiddleware(false))

	v1.Use(common.AuthMiddleware(true))
	routers.UserRegister(v1.Group("/user"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userA := models.User{
		Username:     "userB",
		PasswordHash: "HAAASH",
		Email:        "email@test.com",
	}
	user, err := userService.Create(userA)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	// test 1 to 1

	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func something() {

	minioClient, err := minio.New(
		"example.com:9000",
		&minio.Options{
			Creds:  credentials.NewStaticV4("minio", "minio123", ""),
			Secure: false,
		},
	)
	exists, err := minioClient.BucketExists(context.Background(), "test")
	if err != nil {
		panic(err)
	}
	if exists {
		object, err := minioClient.GetObject(context.Background(), "test", "payload.c", minio.GetObjectOptions{})
		if err != nil {
			panic(err)
		}
		defer object.Close()
		stat, err := minioClient.StatObject(context.Background(), "test", "payload.c", minio.StatObjectOptions{})
		if err != nil {
			panic(err)
		}
		minioClient.PutObject(context.Background(), "test", "payload.c", object, stat.Size, minio.PutObjectOptions{})
	}
	println("Hello, World!", minioClient, err)
}
