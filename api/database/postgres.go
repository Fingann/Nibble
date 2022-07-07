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

var connectionString = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, password, dbname, port, TZ)

// Opening a database and save the reference to `Database` struct.
func NewGormPostgresDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	sql, err := db.DB()
	if err != nil {
		return nil, err
	}
	sql.SetMaxIdleConns(10)
	return db, nil
}
