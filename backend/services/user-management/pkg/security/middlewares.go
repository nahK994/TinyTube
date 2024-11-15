package security

import (
	"context"
	"net/http"
	"strings"
	"user-management/pkg/app"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthenticateMiddleware validates JWT tokens and injects user information into the context.
func authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort() // Stop further processing
			return
		}

		// Remove "Bearer " prefix from token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return app.GetConfig().App.JWT_secret_key, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract claims and set userId in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userId, ok := claims["sub"].(float64); ok {
				// Attach userId to context
				ctx := context.WithValue(c.Request.Context(), "userId", int(userId))
				c.Request = c.Request.WithContext(ctx)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
	}
}

// MiddlewareManager applies multiple middlewares in sequence.
func MiddlewareManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define your middlewares in order
		middlewares := []gin.HandlerFunc{
			authenticateMiddleware(),
			// Add other middlewares here if necessary
		}

		// Execute middlewares sequentially
		for _, middleware := range middlewares {
			middleware(c)
			if c.IsAborted() {
				return // Stop processing if a middleware aborts
			}
		}

		c.Next()
	}
}
