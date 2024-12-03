package controller

import (
	"context"

	"github.com/asyncnavi/raateo/database"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

type contextKey string

const (
	userContextKey = contextKey("user")
)

func init() {
	validate = validator.New()
}

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
