package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/dtos"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl *Controller) CreateRating() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx := c.Request.Context()

		user := UserFromContext(ctx)
		var reqInput dtos.CreateRatingRequest

		if err := c.ShouldBindJSON(&reqInput); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		if user.ID != reqInput.UserID {
			apiErrors.RespondWithError(c, errors.New("user id does not match with authenticated user"))
			return
		}

		org := dtos.GenerateCreateRatingDTO(&reqInput)

		if err := ctrl.repo.SaveRating(org); err != nil {
			apiErrors.InternalError()
			return
		}
		result := dtos.GenerateRatingResponse(org)

		c.JSON(http.StatusOK, result)
	}

}

func (ctrl *Controller) GetRatings() gin.HandlerFunc {

	var RatingParams struct {
		ProductID      uint `uri:"product_id" binding:"required"`
		OrganizationID uint `uri:"organization_id" binding:"required"`
		FeatureID      uint `uri:"feature_id" binding:"required"`
	}

	return func(c *gin.Context) {
		if err := c.ShouldBindUri(&RatingParams); err != nil {
			apiErrors.RespondWithError(c, errors.New("invalid Params"))
		}

		ratings, err := ctrl.repo.GetRatings(RatingParams.OrganizationID, RatingParams.ProductID, RatingParams.FeatureID)

		if err != nil {
			apiErrors.RespondWithError(c, errors.New("could not fetch ratings"))
			return
		}

		result := dtos.GenerateRatingsResponse(ratings)

		c.JSON(http.StatusOK, result)
	}

}

func (ctrl *Controller) GetRatingDetailForProduct() gin.HandlerFunc {

	var params struct {
		ProductID      uint `uri:"product_id" binding:"required"`
		OrganizationID uint `uri:"organization_id" binding:"required"`
	}

	return func(c *gin.Context) {
		if err := c.ShouldBindUri(&params); err != nil {
			apiErrors.RespondWithError(c, errors.New("invalid params"))
		}

		ratings, err := ctrl.repo.GetRatingsByProduct(params.OrganizationID, params.ProductID)

		if err != nil {
			apiErrors.RespondWithError(c, errors.New("could Not fetch ratings"))
		}
		result := dtos.GenerateRatingDetailForProduct(ratings)

		c.JSON(http.StatusOK, result)

	}
}
