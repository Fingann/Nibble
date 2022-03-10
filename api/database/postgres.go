package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Should get from .env file
const user = "postgres"
const password = "postgres"
const host = "postgres"
const dbname = "postgres"
const port = "5432"
const TZ = "Europe/Rome"

var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, TZ)

type Database struct {
	*gorm.DB
}

// Opening a database and save the reference to `Database` struct.
func NewGormPostgresDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sql, err := db.DB()
	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	sql.SetMaxIdleConns(10)
	return db
}
