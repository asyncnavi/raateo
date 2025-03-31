package controller

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/asyncnavi/raateo/internal/models"

	"github.com/clerk/clerk-sdk-go/v2/jwt"
	clerkusersdk "github.com/clerk/clerk-sdk-go/v2/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ctrl *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization header missing"})
			c.Abort()
			return
		}

		// Ensure the header contains "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tkn := tokenParts[1]

		// Verify the JWT token
		claims, err := jwt.Verify(ctx, &jwt.VerifyParams{
			Token: tkn,
		})
		if err != nil {
			slog.Error("Failed to verify JWT token", "error", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid authentication token",
				"details": "Failed to parse or verify token",
			})
			c.Abort()
			return
		}

		user, err := ctrl.repo.FindByClerkID(claims.Subject)

		if err != nil {
			slog.Error("Error while finding the user.")
			if errors.Is(err, gorm.ErrRecordNotFound) {
				slog.Error("User does not exist in our database base ")
				slog.Info("Fetching user from clerk...")

				// If the user is not found, fetch details from Clerk
				clerkUser, err := clerkusersdk.Get(ctx, claims.Subject)
				if err != nil {
					slog.Error("Failed to fetch user from Clerk", "error", err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Internal Server Error",
						"details": "Unable to fetch user details from Clerk",
					})
					c.Abort()
					return
				}

				slog.Info("User fetched from clerk...")

				slog.Info("Trying saving user...")

				// Extract user details from Clerk
				var email, firstName, lastName string
				for _, addr := range clerkUser.EmailAddresses {
					if clerkUser.PrimaryEmailAddressID != nil && addr.ID == *clerkUser.PrimaryEmailAddressID {
						email = addr.EmailAddress
						break
					}
				}
				if clerkUser.FirstName != nil {
					firstName = *clerkUser.FirstName
				}
				if clerkUser.LastName != nil {
					lastName = *clerkUser.LastName
				}

				// Create a new user record
				user = &models.User{
					FirstName: firstName,
					LastName:  lastName,
					Email:     email,
					ClerkID:   clerkUser.ID,
				}

				// Save the user to the database
				if err := ctrl.repo.SaveUser(user); err != nil {
					slog.Error("Failed to save user in database", "error", err)
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "Internal Server Error",
						"details": "Failed to save user in database",
					})
					c.Abort()
					return
				}

				slog.Info("User saved")
			} else {
				// Handle other database errors
				slog.Error("Failed to find user in database", "error", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal Server Error",
					"details": "Error finding user in database",
				})
				c.Abort()
				return
			}
		}

		// Attach the user to the context
		ctx = WithUser(ctx, user)
		c.Request = c.Request.WithContext(ctx)

		// Proceed to the next middleware or handler
		c.Next()
	}
}
