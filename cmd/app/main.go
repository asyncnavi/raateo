package main

import (
	"log"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig(config.GetEnvPath())
	if err != nil {
		log.Fatal("ERROR : Could not load environment variables.", err)
	}

	server := initRoutes(&cfg)

	log.Fatal(server.Run(":" + cfg.ServerPort))
}

func initRoutes(cfg *config.Config) *gin.Engine {
	db := database.NewDatabase(cfg)

	server := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173", "http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Set-Cookie"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))
	server.Static("/public", "./public")

	rc := controller.NewController(db)

	router := server.Group("").Use(rc.Authorize())
	{
		router.GET("/organization/me", rc.UserOrganization())
		router.GET("/organization/products/:org_id", rc.ListProduct())
		router.POST("/organization", rc.CreateOrganization())
	}
	{
		router.POST("/product", rc.CreateProduct())
		router.GET("/product/:id", rc.SingleProduct())
	}
	return server
}
