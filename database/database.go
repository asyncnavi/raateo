package database

import (
	"fmt"
	"log"

	"github.com/asyncnavi/raateo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(cfg *config.Config) *Database {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		cfg.DBHost, cfg.DBUsername, cfg.DBUserPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the Database: %v", err)
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	return &Database{db: db}
}
