package main

import (
	"log"

	"github.com/asyncnavi/raateo/config"
	"github.com/asyncnavi/raateo/database"
	"github.com/asyncnavi/raateo/internal/controller"
	"github.com/asyncnavi/raateo/internal/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Ok",
		})
	}
}

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

	repo := repository.New(db)

	server.Use(cors.New(corsConfig))
	server.Static("/public", "./public")

	rc := controller.New(db, cld, ctx, repo)

	// router without authentication
	public := server.Group("")
	// router with authentication
	protected := server.Group("").Use(rc.AuthMiddleware())

	{
		public.GET("", HealthCheck())
	}
	{
		protected.POST("/org", rc.CreateOrganization())
		protected.GET("/org/me", rc.GetOrganizationForUser())
		protected.GET("/org/products", rc.GetProductsForOrganization())
		protected.DELETE("/org/products/:product_id", rc.DeleteProductForOrganization())
		protected.POST("/org/products", rc.CreateProductForOrganization())
	}
	{

		public.GET("/products", rc.GetProducts())
		public.GET("/products/search", rc.SearchProduct())
		public.GET("/products/:product_id/:organization_id", rc.GetProduct())
		public.GET("/products/:product_id/:organization_id/features", rc.GetFeaturesForProduct())
		public.GET("/products/:product_id/:organization_id/features/:feature_id", rc.GetFeature())
		public.GET("/products/:product_id/:organization_id/features/:feature_id/ratings", rc.GetRatings())
		public.GET("/product/:product_id/:organization_id", rc.GetRatingDetailForProduct())

	}
	{
		protected.POST("/ratings", rc.CreateRating())
		protected.POST("/features", rc.CreateFeature())

	}
	{
		protected.POST("/uploads", rc.UploadImage())
	}

	return server
}
