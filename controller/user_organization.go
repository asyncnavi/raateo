package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (oc *Controller) UserOrganization() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := UserFromContext(ctx)

		org, err := oc.db.FindOrganizationByUser(int(user.ID))

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
			"organization_id": org.ID,
			"name":            org.Name,
		})
	}
}
