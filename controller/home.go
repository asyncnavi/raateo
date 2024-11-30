package controller

import (
	"github.com/asyncnavi/raateo/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomeController(ctx *gin.Context) {
	if err := views.Home().Render(ctx.Request.Context(), ctx.Writer); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}
