package main

import (
	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
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

	corsConfig := config.SetupCors()
	cld, ctx := config.SetupStorage(cfg.CloudinaryURL)

	server.Use(cors.New(corsConfig))
	server.Static("/public", "./public")

	rc := controller.New(db, cld, ctx)

	public := server.Group("")
	withUser := server.Group("").Use(rc.AuthMiddleware())

	{
		withUser.POST("/org", rc.CreateOrganization())
		withUser.GET("/org/me", rc.UserOrganization())
		withUser.GET("/org/:org_id/products", rc.OrganizationProducts())
		withUser.POST("/org/products", rc.CreateProduct())
		withUser.GET("/org/:org_id/features/:product_id", rc.FeaturesByOrganization())
	}
	{
		public.GET("/products/:id", rc.SingleProduct())
		public.GET("/products/:id/features", rc.ListFeature())
		public.GET("/products", rc.Products())

	}
	{
		withUser.POST("/org/:organization_id/features", rc.CreateFeature())
	}
	{
		withUser.POST("/uploads", rc.UploadImage())
	}

	return server
}
