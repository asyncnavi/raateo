package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/database"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strconv"
)

func (ctrl *Controller) CreateProduct() gin.HandlerFunc {
	var reqArgs struct {
		OrganizationID string `json:"organization_id" validate:"required"`
		Name           string `json:"name" validate:"required, min=2, max=256"`
		Description    string `json:"description" validate:"max=2000"`
		LogoURL        string `json:"logo_url"`
		ThumbnailURL   string `json:"thumbnail_url" `
	}
	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&reqArgs); err != nil {
			slog.Error("failed to parse json", "", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		// Parse Organization ID to uint
		orgID, err := strconv.Atoi(reqArgs.OrganizationID)

		if err != nil {
			slog.Error("failed to parse organization_id", "", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		prod := &database.Product{
			OrganizationID: uint(orgID),
			Name:           reqArgs.Name,
			Description:    reqArgs.Description,
			LogoURL:        reqArgs.LogoURL,
			ThumbnailURL:   reqArgs.ThumbnailURL,
		}
		if err := ctrl.db.SaveProduct(prod); err != nil {
			slog.Error("failed to save product")
			apiErrors.RespondWithError(c, err)
			return
		}

		resultProduct := gin.H{
			"organization_id": uint(orgID),
			"name":            reqArgs.Name,
			"description":     reqArgs.Description,
			"logo_url":        reqArgs.LogoURL,
			"thumbnail_url":   reqArgs.ThumbnailURL,
		}

		c.JSON(http.StatusOK, resultProduct)
	}
}

func (ctrl *Controller) SingleProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Product ID
		idParam := c.Param("id")

		productId, err := strconv.Atoi(idParam)

		if err != nil {
			slog.Error("failed to parse organization Id")
			apiErrors.RespondWithError(c, err)

			return
		}

		product, err := ctrl.db.GetProduct(uint(productId))

		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				slog.Error("failed to find product with id")
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Product not found",
					"details": "User does have product in the database associated to this id",
				})

				return
			} else {
				slog.Error("Error querying database", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
					"details": "Something went wrong at backend",
				})
				return
			}

		}

		resultProduct := map[string]any{
			"id":              product.ID,
			"organization_id": product.OrganizationID,
			"name":            product.Name,
			"description":     product.Description,
			"logo_url":        product.LogoURL,
			"thumbnail_url":   product.ThumbnailURL,
			"created_at":      product.CreatedAt,
			"updated_at":      product.UpdatedAt,
			"deleted_at":      product.DeletedAt,
		}

		c.JSON(http.StatusOK, resultProduct)

	}
}

func (ctrl *Controller) ListProducts() gin.HandlerFunc {

	return func(c *gin.Context) {
		products, err := ctrl.db.GetAllProducts()
		if err != nil {
			slog.Error("failed to fetch products")
			apiErrors.RespondWithError(c, err)
			return
		}

		var resultProducts []map[string]interface{}

		for _, product := range products {
			resultProducts = append(resultProducts, map[string]interface{}{
				"id":              product.ID,
				"organization_id": product.OrganizationID,
				"name":            product.Name,
				"description":     product.Description,
				"logo_url":        product.LogoURL,
				"thumbnail_url":   product.ThumbnailURL,
				"created_at":      product.CreatedAt,
				"updated_at":      product.UpdatedAt,
				"deleted_at":      product.DeletedAt,
			})
		}

		c.JSON(http.StatusOK, resultProducts)
	}
}

func (ctrl *Controller) ListOrganizationProducts() gin.HandlerFunc {

	return func(c *gin.Context) {
		// get the id
		id := c.Param("org_id")

		// parsing product id
		orgId, err := strconv.Atoi(id)

		if err != nil {
			slog.Error("failed to parse organization Id")
			apiErrors.RespondWithError(c, err)
			return
		}

		products, err := ctrl.db.GetProductsByOrganization(uint(orgId))

		if err != nil {
			slog.Error("failed to fetch products")
			apiErrors.RespondWithError(c, err)
			return
		}

		var resultProducts []map[string]interface{}

		for _, product := range products {
			resultProducts = append(resultProducts, map[string]interface{}{
				"id":              product.ID,
				"organization_id": product.OrganizationID,
				"name":            product.Name,
				"description":     product.Description,
				"logo_url":        product.LogoURL,
				"thumbnail_url":   product.ThumbnailURL,
				"created_at":      product.CreatedAt,
				"updated_at":      product.UpdatedAt,
				"deleted_at":      product.DeletedAt,
			})
		}

		c.JSON(http.StatusOK, resultProducts)
	}
}
