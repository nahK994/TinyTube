package security

import (
	"auth-service/pkg/app"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware is a Gin-compatible authentication middleware
func authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, "Authorization header missing")
			c.Abort()
			return
		}

		// Remove "Bearer " prefix from the token string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return app.GetConfig().App.JWT_secret_key, nil
		})

		// Check for parsing errors or invalid tokens
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// Extract claims and add user ID to context if claims are valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userId := int(claims["sub"].(float64))
			c.Set("userId", userId)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
		}
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
