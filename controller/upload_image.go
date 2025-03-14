package controller

import (
	"errors"
	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (ctrl *Controller) UploadImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, header, err := c.Request.FormFile("image")

		if err != nil {
			slog.Error("failed to upload file")
			apiErrors.RespondWithError(c, err)
			return
		}

		filePath := header.Filename

		result, err := ctrl.cdn.Upload.Upload(ctrl.cdnCtx, file, uploader.UploadParams{
			PublicID: filePath,
		})

		if err != nil {
			slog.Error("failed to upload file")
			apiErrors.RespondWithError(c, errors.New("failed to upload to cloudinary"))
			return
		}

		imageURL := result.SecureURL

		c.JSON(http.StatusOK, gin.H{
			"url": imageURL,
		})

	}
}
