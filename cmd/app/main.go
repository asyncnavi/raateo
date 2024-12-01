package main

import (
	"log"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/controller/home"
	"github.com/asyncnavi/raateo/controller/login"
	"github.com/asyncnavi/raateo/controller/organization"
	"github.com/asyncnavi/raateo/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(config.GetEnvPath())
	if err != nil {
		log.Fatal("ERROR : Could not load environment variables.", err)
	}

	server := setupRoutes(&cfg)

	log.Fatal(server.Run(":" + cfg.ServerPort))
}

func setupRoutes(cfg *config.Config) *gin.Engine {
	db := database.NewDatabase(cfg)

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	server.Use(cors.New(corsConfig))
	server.Static("/public", "./public")

	router := server.Group("").Use(controller.Authorize(db))

	{
		homeController := home.New(db)
		router.GET("", homeController.HandleIndex())
	}

	{
		loginController := login.New(db)
		router.GET("/login", loginController.HandleIndex())
	}
	{
		organizationController := organization.New(db)
		router.GET("/organization/:id", organizationController.HandleShow())
		router.GET("/organization/create", organizationController.HandleCreate())
		router.POST("/organization", organizationController.HandleCreate())
	}

	return server
}
