package main

import (
	"context"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"fileslut/database"
	"fileslut/models"
	"fileslut/user"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
}
func main() {

	db := database.NewGormPostgresDB()

	r := gin.Default()

	v1 := r.Group("/api")
	user.UsersRegister(v1.Group("/users"))
	v1.Use(user.AuthMiddleware(false))

	v1.Use(user.AuthMiddleware(true))
	user.UserRegister(v1.Group("/user"))
	user.ProfileRegister(v1.Group("/profiles"))

	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userService := models.NewGormUserRepository(db)
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
