package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost         string
	DBUsername     string
	DBUserPassword string
	DBName         string
	DBPort         string

	ServerPort   string
	ClientOrigin string

	Secret string
}

func LoadConfig(path string) (config Config, err error) {
	fmt.Println("Attempting to load config from:", path)
	err = godotenv.Load(path)

	if err != nil {
		log.Println("Error opening config file:", err)
		return config, err
	}

	config = Config{
		DBHost:         os.Getenv("DB_HOST"),
		DBUsername:     os.Getenv("DB_USER"),
		DBUserPassword: os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
		DBPort:         os.Getenv("DB_PORT"),

		ServerPort: os.Getenv("SERVER_ADDR"),

		Secret: os.Getenv("SECRET"),
	}

	return config, nil
}

// GetEnvPath fetches env file path. Note: Call this function in main
func GetEnvPath() string {
	_, filename, _, _ := runtime.Caller(1)
	pwd := filepath.Dir(filename)

	return filepath.Join(pwd, "../../.env")
}
