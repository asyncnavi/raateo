package main

import (
	"github.com/asyncnavi/raateo/controller"
	"log"
	"net/http"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {

	cfg, err := config.LoadConfig(config.GetEnvPath())

	if err != nil {
		log.Fatal("ERROR : Could not load environment variables.", err)
	}

	database.InitDB(&cfg)

	server = gin.Default()
}

func main() {

	cfg, err := config.LoadConfig(config.GetEnvPath())

	if err != nil {
		log.Fatal("ERROR : Could not load environment variables.", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}

	server.Use(cors.New(corsConfig))
	server.Static("/public", "./public")
	router := server.Group("")
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "PONG")
	})
	router.GET("", controller.HomeController)

	log.Fatal(server.Run(":" + cfg.ServerPort))
}
