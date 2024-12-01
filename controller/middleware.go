package controller

import (
	"errors"
	"log"
	"log/slog"
	"net/http"

	"github.com/asyncnavi/raateo/database"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	clerkusersdk "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Authorize(db *database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		cookies := c.Request.Cookies()
		sessionToken := ""
		for _, cookie := range cookies {
			if cookie.Name == "__session" {
				sessionToken = cookie.Value
			}
		}

		if sessionToken == "" {
			slog.Info("User is not logged in passing handle to controller")
			c.Next()
			return
		}

		claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
			Token: sessionToken,
		})
		if err != nil {

			return
		}

		user, err := db.FindByClerkID(claims.Subject)

		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("failed to find user: %v", err)
				c.AbortWithError(http.StatusInternalServerError, errors.New("internal server error"))
				return
			}

			clerkUser, err := clerkusersdk.Get(ctx, claims.Subject)
			if err != nil {
				slog.Info("Error is happening here")
				c.AbortWithError(http.StatusUnauthorized, errors.New("user is not logged in please login"))
			}

			slog.Info("User is authenticated")
			slog.Info("Trying passing user to context...")

			var (
				email     string
				firstName string
				lastName  string
			)
			for _, addr := range clerkUser.EmailAddresses {
				if addr.ID == *clerkUser.PrimaryEmailAddressID {
					email = addr.EmailAddress
				}
			}
			if clerkUser.FirstName != nil {
				firstName = *clerkUser.FirstName
			}
			if clerkUser.LastName != nil {
				lastName = *clerkUser.LastName
			}

			user = &database.User{
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				ClerkID:   clerkUser.ID,
			}

			if err := db.SaveUser(user); err != nil {
				log.Printf("failed to save user: %v", err)
				c.AbortWithError(http.StatusInternalServerError, errors.New("internal server error"))
				return
			}
		}

		ctx = WithUser(ctx, user)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
