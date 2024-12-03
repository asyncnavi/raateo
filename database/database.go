package database

import (
	"fmt"
	"log/slog"

	"github.com/asyncnavi/raateo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(cfg *config.Config) *Database {
	var err error
	slog.Info("Trying connecting to the Database")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		cfg.DBHost, cfg.DBUsername, cfg.DBUserPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("🥲 Failed to connect to the Database ", "error", err)
	}

	slog.Info("🚀 Connected Successfully to the Databases")
	return &Database{db: db}
}
