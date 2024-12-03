package controller

import (
	"errors"
	"net/http"

	"github.com/asyncnavi/raateo/database"
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
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid Input Format",
				"details": "Failed to parse JSON" + err.Error(),
			})
			return
		}

		// TODO : Handle Field Errors properly
		if err := validate.Struct(input); err != nil {
			// var validationErrors []string
			// for _, err := range err.(validator.ValidationErrors) {
			// 	validationErrors = append(validationErrors, err.Error())
			// }
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Field Error",
				"details": "Invalid Field:",
			})

			return
		}

		user := UserFromContext(ctx)

		_, err := oc.db.FindOrganizationByUser(int(user.ID))

		if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server error",
				"details": err.Error(),
			})
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
