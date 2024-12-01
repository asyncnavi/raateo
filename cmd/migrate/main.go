package main

import (
	"context"
	"log"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/database"
)

func main() {
	cfg, err := config.LoadConfig(config.GetEnvPath())
	if err != nil {
		log.Fatal("ERROR : Could not load environment variables.", err)
	}

	ctx := context.TODO()

	db := database.NewDatabase(&cfg)

	if err := db.Migrate(ctx); err != nil {
		log.Fatal("ERROR : failed to migrate.", err)
	}

	log.Printf("Migration went succesffully.")
}
