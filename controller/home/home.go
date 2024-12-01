package home

import (
	"net/http"

	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/database"
	"github.com/asyncnavi/raateo/views"
	"github.com/gin-gonic/gin"
)

type Home struct {
	db *database.Database
}

func New(db *database.Database) *Home {
	return &Home{db: db}
}

func (h *Home) HandleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		t := &controller.Template{
			Title: "Raateo | Home",
		}

		user := controller.UserFromContext(ctx)

		t.User = user

		if err := views.Home(t).Render(c.Request.Context(), c.Writer); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}
