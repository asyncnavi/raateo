package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	apiErrors "github.com/asyncnavi/raateo/pkg/errros"

	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) OrganizationProducts() gin.HandlerFunc {

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
