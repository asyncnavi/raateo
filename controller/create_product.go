package controller

import (
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			apiErrors.SendInvalidJSONError(c, err)
			return
		}

		if err := validate.Struct(input); err != nil {
			var validationErrors []string

			for _, err := range err.(validator.ValidationErrors) {
				validationErrors = append(validationErrors, err.Error())
			}

			apiErrors.SendValidationError(c, validationErrors)
		}

		ctx := c.Request.Context()

		user := UserFromContext(ctx)

	}
}
