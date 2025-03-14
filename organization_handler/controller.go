package controller

import (
	"context"
	"github.com/asyncnavi/raateo/database"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/go-playground/validator/v10"
)

var (
	_ *validator.Validate
)

type contextKey string

const (
	userContextKey = contextKey("user")
)

func init() {
	_ = validator.New()
}

type Controller struct {
	db     *database.Database
	cdn    *cloudinary.Cloudinary
	cdnCtx context.Context
}

func New(db *database.Database, cdn *cloudinary.Cloudinary, cdnCtx context.Context) *Controller {
	return &Controller{db: db, cdn: cdn, cdnCtx: cdnCtx}
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
