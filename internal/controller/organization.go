package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/internal/dtos"
	"github.com/asyncnavi/raateo/internal/models"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func (ctrl *Controller) CreateOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var reqInput dtos.CreateOrganizationRequest

		if err := c.ShouldBindJSON(&reqInput); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		// gets the user from the context
		user := UserFromContext(ctx)

		// checks if the user already has an organization
		_, err := ctrl.repo.OrganizationFromUser(user.ID)

		if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
			apiErrors.RespondWithError(c, errors.New("an Organization already exists with the user"))
			return
		}

		org := &models.Organization{
			Name:        reqInput.Name,
			UserID:      user.ID,
			Description: reqInput.Description,
		}

		if err := ctrl.repo.SaveOrganization(org); err != nil {
			apiErrors.InternalError()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Organization is saved",
		})
	}

}

func (ctrl *Controller) GetOrganizationForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := UserFromContext(ctx)

		org, err := ctrl.repo.OrganizationFromUser(user.ID)

		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {
				apiErrors.RespondWithError(c, errors.New("user does not have any organization in the database"))
				return
			} else {
				apiErrors.InternalError()
				return
			}
		}

		result := dtos.GenerateOrganizationResponse(org)
		c.JSON(http.StatusOK, result)
	}
}
