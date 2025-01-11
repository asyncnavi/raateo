package controller

import (
	"errors"
	"net/http"

	"github.com/asyncnavi/raateo/database"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (oc *Controller) CreateOrganization() gin.HandlerFunc {
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

		_, err := oc.db.FindOrganizationByUser(int(user.ID))

		if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
			apiErrors.InternalError()
			return
		}

		org := &database.Organization{
			Name:   input.Name,
			UserID: user.ID,
		}

		if err := oc.db.SaveOrganization(org); err != nil {
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
