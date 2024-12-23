package controller

import (
	"log/slog"
	"strconv"

	apiErrors "github.com/asyncnavi/raateo/pkg/errros"
	"github.com/gin-gonic/gin"
)

func (pc *Controller) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("product_id")

		prodId, err := strconv.Atoi(id)

		if err != nil {
			slog.Error("failed to parse product Id")
			apiErrors.RespondWithError(c, err)
			return
		}

		product, err = pc.db.GetProduct(uint(prodId))
	}
}
