package controller

import (
	"errors"
	"github.com/asyncnavi/raateo/database"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ctrl *Controller) UserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := UserFromContext(ctx)

		org, err := ctrl.db.FindOrganizationByUser(int(user.ID))

		if err != nil {
			if errors.Is(gorm.ErrRecordNotFound, err) {

				c.JSON(http.StatusNotFound, gin.H{
					"message": "Organization not found",
					"details": "User does not have any organization in the database",
				})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server error",
					"details": err.Error(),
				})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"id":   org.ID,
			"name": org.Name,
		})
	}
}

func (ctrl *Controller) CreateOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var input struct {
			Name string `json:"name" validate:"required,min=2"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			apiErrors.RespondWithError(c, err)
			return
		}

		user := UserFromContext(ctx)

		_, err := ctrl.db.FindOrganizationByUser(int(user.ID))

		if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
			apiErrors.InternalError()
			return
		}

		org := &database.Organization{
			Name:   input.Name,
			UserID: user.ID,
		}

		if err := ctrl.db.SaveOrganization(org); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server error",
				"details": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Organization is saved",
		})
	}

}
