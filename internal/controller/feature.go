package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/dtos"
	"github.com/asyncnavi/raateo/internal/models"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

func (ctrl *Controller) GetFeaturesForProduct() gin.HandlerFunc {

	return func(c *gin.Context) {
		var params ProductParams

		if err := c.ShouldBindUri(&params); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		features, err := ctrl.repo.GetFeaturesForProduct(params.ProductID, params.OrganizationID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Error("Features not found", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID))
				apiErrors.RespondWithError(c, errors.New("features not found"))
				return
			}
			slog.Error("Failed to get features", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID), slog.Any("error", err))
			apiErrors.RespondWithError(c, err)
			return
		}

		result := dtos.GenerateFeaturesResponse(features)
		c.JSON(http.StatusOK, result)
	}
}

type FeatureParams struct {
	ProductID      uint `uri:"product_id" binding:"required"`
	OrganizationID uint `uri:"organization_id" binding:"required"`
	FeatureID      uint `uri:"feature_id" binding:"required"`
}

func (ctrl *Controller) GetFeature() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params FeatureParams

		if err := c.ShouldBindUri(&params); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		feature, err := ctrl.repo.GetFeature(params.ProductID, params.OrganizationID, params.FeatureID)

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Error("Feature not found", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID), slog.Any("feature_id", params.FeatureID))
				apiErrors.RespondWithError(c, errors.New("feature not found"))
				return
			}
			slog.Error("Failed to get feature", slog.Any("product_id", params.ProductID), slog.Any("organization_id", params.OrganizationID), slog.Any("feature_id", params.FeatureID), slog.Any("error", err))
			apiErrors.RespondWithError(c, err)
			return
		}

		result := dtos.GenerateFeatureResponse(feature)
		c.JSON(http.StatusOK, result)
	}
}

func (ctrl *Controller) CreateFeature() gin.HandlerFunc {

	var requestInput dtos.CreateFeatureRequest

	return func(c *gin.Context) {

		if err := c.ShouldBindJSON(&requestInput); err != nil {
			slog.Error("Failed to parse json", "", err)
			apiErrors.RespondWithError(c, err)
		}

		feature := &models.Feature{
			ProductID:      requestInput.ProductID,
			OrganizationID: requestInput.OrganizationID,
			Name:           requestInput.Name,
			Description:    requestInput.Description,
			VideoUrl:       requestInput.VideoURL,
			ThumbnailUrl:   requestInput.ThumbnailURL,
		}

		if err := ctrl.repo.SaveFeature(feature); err != nil {
			slog.Error("Failed to save feature", "error", err)
			apiErrors.RespondWithError(c, err)
			return
		}

		result := dtos.GenerateFeatureResponse(feature)
		c.JSON(http.StatusOK, result)
	}
}
