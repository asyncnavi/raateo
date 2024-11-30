package main

import (
	"fmt"
	"log"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/database"
)

func init() {

	envPath := config.GetEnvPath()

	cfg, err := config.LoadConfig(envPath)

	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	database.InitDB(&cfg)
}

func main() {
	database.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	err := database.DB.AutoMigrate()
	if err != nil {
		log.Fatal("Migrate Failed:", err)
		return
	}
	fmt.Println("Migration Complete")
}
