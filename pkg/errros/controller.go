package errros

import (
	"errors"
	"net/http"

	"github.com/asyncnavi/raateo/pkg/app"
	"github.com/asyncnavi/raateo/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RespondWithError(c *gin.Context, err error) {
	appErr := &app.AppError{}
	if errors.As(err, &appErr) {
		c.AbortWithStatusJSON(appErr.Status(), APIError{
			Error:     appErr.Error(),
			ErrorCode: appErr.Code(),
			Details:   appErr.Details(),
		})
	} else if verrors, ok := err.(validator.ValidationErrors); ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, &ValidationAPIError{
			APIError: *Errorf("One or more fields are invalid").WithCode("invalid_payload"),
			Fields:   validate.ToFieldErrors(verrors).FieldsMap(),
		})
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
	}
}
