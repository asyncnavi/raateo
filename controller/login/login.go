package login

import (
	"net/http"

	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/database"
	"github.com/asyncnavi/raateo/views"
	"github.com/gin-gonic/gin"
)

type Login struct {
	db *database.Database
}

func New(db *database.Database) *Login {
	return &Login{db: db}
}

func (l *Login) HandleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		user := controller.UserFromContext(ctx)
		if user != nil {
			c.Redirect(http.StatusConflict, "/")
			return
		}

		views.Login().Render(ctx, c.Writer)
	}
}
