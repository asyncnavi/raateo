package controller

import (
	"context"
	"github.com/asyncnavi/raateo/internal/models"
	"gorm.io/gorm"

	"github.com/asyncnavi/raateo/internal/repository"
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
	db     *gorm.DB
	cdn    *cloudinary.Cloudinary
	repo   *repository.Repository
	cdnCtx context.Context
}

func New(db *gorm.DB, cdn *cloudinary.Cloudinary, cdnCtx context.Context, repo *repository.Repository) *Controller {
	return &Controller{db: db, cdn: cdn, cdnCtx: cdnCtx, repo: repo}
}

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func UserFromContext(ctx context.Context) *models.User {
	u, ok := ctx.Value(userContextKey).(*models.User)
	if !ok {
		return nil
	}
	return u
}
