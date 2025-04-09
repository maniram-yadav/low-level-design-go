package middleware

import (
	"lld/stackoverflow/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userServices *services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		// Mock user ID - in real app, we'd decode from JWT
		userID := uint(1)
		c.Set("userID", userID)
		c.Next()
	}
}
