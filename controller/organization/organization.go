package organization

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/asyncnavi/raateo/controller"
	"github.com/asyncnavi/raateo/database"
	"github.com/asyncnavi/raateo/views"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	db *database.Database
}

func New(db *database.Database) *Controller {
	return &Controller{db: db}
}

func (oc *Controller) HandleShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		t := controller.NewTemplate("Raateo | Organization")
		user := controller.UserFromContext(ctx)
		if user == nil {
			controller.MissingSession(c)
			return
		}
		t.User = user

		orgID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			slog.Error("failed to parse org id", "error", err)
			t.AddErrorMessage("Invalid organization id")
			oc.RenderOrg(ctx, c, t, nil)
			return
		}

		org, err := oc.db.FindOrganization(orgID)
		if err != nil {
			slog.Error("failed to parse org id", "error", err)
			t.AddErrorMessage("Organization not found")
			oc.RenderOrg(ctx, c, t, nil)
			return
		}

		t.AddTitle("Raateo | " + org.Name)

		oc.RenderOrg(ctx, c, t, org)
	}
}

func (oc *Controller) RenderOrg(ctx context.Context, c *gin.Context, t *controller.Template, org *database.Organization) {
	views.Organization(t, org).Render(ctx, c.Writer)
}

func (oc *Controller) HandleCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		t := controller.NewTemplate("Raateo | Create Organization")

		user := controller.UserFromContext(ctx)
		if user == nil {
			controller.MissingSession(c)
			return
		}
		t.User = user

		form := &views.CreateOrgForm{}
		if c.Request.Method == http.MethodPost {
			if err := c.ShouldBind(form); err != nil {
				slog.Error("failed to bind form", "error", err)
				controller.InternalError(c)
				return
			}

			if form.Name == "" {
				t.AddFieldError("name", "Name is required")
			}
			if len(form.Name) > 255 {
				t.AddFieldError("name", "Name should be maxiumum 255 characters")
			}

			if t.HasErrors() {
				views.CreateOrganization(t, form).Render(ctx, c.Writer)
				return
			}

			org := &database.Organization{
				Name:   form.Name,
				UserID: user.ID,
			}

			if err := oc.db.SaveOrganization(org); err != nil {
				slog.Error("failed to save organization", "error", err)
				controller.InternalError(c)
				return
			}
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("/organization/%d", org.ID))
			return
		}

		views.CreateOrganization(t, form).Render(ctx, c.Writer)
	}
}
