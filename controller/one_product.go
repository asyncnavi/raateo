package controller

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

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
