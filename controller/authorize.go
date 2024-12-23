package controller

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/asyncnavi/raateo/database"
	"github.com/clerk/clerk-sdk-go/v2/jwt"
	clerkusersdk "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ao *Controller) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		cookies := c.Request.Cookies()

		tkn := ""
		for _, cookie := range cookies {
			if cookie.Name == "__session" {
				tkn = cookie.Value
			}
		}

		if tkn == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "User is not logged",
				"details": "Cookie not found",
			})
			c.Abort()
			return
		}

		claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
			Token: tkn,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"details": "Failed to parse authentication token",
			})
			c.Abort()
			return
		}
		user, err := ao.db.FindByClerkID(claims.Subject)

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			clerkUser, err := clerkusersdk.Get(ctx, claims.Subject)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "User is not register",
					"details": "Unable to find user in clerk database",
				})
				c.Abort()
			}

			// saving the user
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

			if err := ao.db.SaveUser(user); err != nil {
				slog.Error("Failed to save user in database", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
					"details": "Failed to save user in database",
				})
				c.Abort()
				return
			}
		}

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Error("Failed to find user", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
				"details": "Error finding user in database" + " " + err.Error(),
			})
			c.Abort()
		}

		ctx = WithUser(ctx, user)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
