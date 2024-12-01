package controller

import (
	"context"
	"net/http"

	"github.com/asyncnavi/raateo/database"
	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	userContextKey = contextKey("user")
)

type Controller struct {
	db *database.Database
}

func NewController(db *database.Database) *Controller {
	return &Controller{db: db}
}

func WithUser(ctx context.Context, user *database.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) *database.User {
	u, ok := ctx.Value(userContextKey).(*database.User)
	if !ok {
		return nil
	}
	return u
}

func InternalError(c *gin.Context) {
	c.Status(http.StatusInternalServerError)
	c.Abort()
}

func MissingSession(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/login")
	c.Abort()
}
