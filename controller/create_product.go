package controller

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/asyncnavi/raateo/database"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
)

func (rc *Controller) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			OrganizationID string `json:"organization_id" validate:"required"`
			Name           string `json:"name" validate:"required, min=2, max=256"`
			Description    string `json:"description" validate:"max=2000"`
			LogoURL        string `json:"logo_url"`
			ThumbnailURL   string `json:"thumbnail_url" `
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			slog.Error("failed to parse json", "", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		// Parse Organization ID to uint
		orgID, err := strconv.Atoi(input.OrganizationID)

		if err != nil {
			slog.Error("failed to parse organization_id", "", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		prod := &database.Product{
			OrganizationID: uint(orgID),
			Name:           input.Name,
			Description:    input.Description,
			LogoURL:        input.LogoURL,
			ThumbnailURL:   input.ThumbnailURL,
		}
		if err := rc.db.SaveProduct(prod); err != nil {
			slog.Error("failed to save product")
			apiErrors.RespondWithError(c, err)
			return
		}

		resultProduct := gin.H{
			"organization_id": uint(orgID),
			"name":            input.Name,
			"description":     input.Description,
			"logo_url":        input.LogoURL,
			"thumbnail_url":   input.ThumbnailURL,
		}

		c.JSON(http.StatusOK, resultProduct)
	}
}
