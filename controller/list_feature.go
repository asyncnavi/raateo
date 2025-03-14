package controller

import (
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (ctrl *Controller) ListFeature() gin.HandlerFunc {

	return func(c *gin.Context) {

		orgId := c.Param("org_id")
		prodId := c.Param("product_id")

		organizationID, err := strconv.Atoi(orgId)

		if err != nil {
			slog.Error("Failed to parse organizationId", "organizationId", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		productID, err := strconv.Atoi(prodId)

		if err != nil {
			slog.Error("Failed to parse productId", "productId", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		features, err := ctrl.db.GetFeaturesByOrganization(uint(productID), uint(organizationID))

		var result []map[string]interface{}
		for _, feature := range features {
			result = append(result, map[string]interface{}{
				"id":              feature.ID,
				"organization_id": feature.OrganizationID,
				"product_id":      feature.ProductID,
				"name":            feature.Name,
				"description":     feature.Description,
				"video_url":       feature.VideoUrl,
				"status":          feature.Status,
				"thumbnail_url":   feature.ThumbnailUrl,
				"created_at":      feature.CreatedAt,
				"updated_at":      feature.UpdatedAt,
				"deleted_at":      feature.DeletedAt,
			})
		}

		c.JSON(http.StatusOK, result)
	}
}
