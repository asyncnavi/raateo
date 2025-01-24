package controller

import (
	"github.com/asyncnavi/raateo/database"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (fc *Controller) CreateFeature() gin.HandlerFunc {

	var requestInput struct {
		OrganizationID string `json:"organization_id"`
		ProductID      string `json:"product_id"`
		Name           string `json:"name" validate:"required, min=2, max=256"`
		Description    string `json:"description" validate:"max=2000"`
		VideoURL       string `json:"video_url"`
		ThumbnailURL   string `json:"thumbnail_url" `
	}

	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&requestInput); err != nil {
			slog.Error("Failed to parse json", "", err)
			apiErrors.RespondWithError(c, err)
		}

		//  parse organization ID and Product ID
		organizationID, err := strconv.Atoi(requestInput.OrganizationID)

		if err != nil {
			slog.Error("Failed to parse organizationId", "organizationId", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		productID, err := strconv.Atoi(requestInput.ProductID)

		if err != nil {
			slog.Error("Failed to parse productId", "productId", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		featureToCreate := &database.Feature{
			ProductID:      uint(productID),
			OrganizationID: uint(organizationID),
			Name:           requestInput.Name,
			Description:    requestInput.Description,
			VideoUrl:       requestInput.VideoURL,
			ThumbnailUrl:   requestInput.ThumbnailURL,
		}

		if err := fc.db.SaveFeature(featureToCreate); err != nil {
			slog.Error("Failed to save feature", "error", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		featureToResult := gin.H{
			"organizationId": organizationID,
			"productId":      productID,
			"name":           featureToCreate.Name,
			"description":    featureToCreate.Description,
			"videoUrl":       featureToCreate.VideoUrl,
			"thumbnailUrl":   featureToCreate.ThumbnailUrl,
		}
		c.JSON(http.StatusOK, featureToResult)
	}
}
