// cmd/middleware/auth.middleware.go
package middleware

import (
	"fmt"
	"recorderis/cmd/services/auth/ports/drivens"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	tokenManager drivens.ForTokenManager
}

func NewAuthMiddleware(tokenManager drivens.ForTokenManager) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authorization header:", c.GetHeader("Authorization"))
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "no authorization header"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "invalid authorization format"})
			c.Abort()
			return
		}

		userID, err := m.tokenManager.ValidateToken(parts[1])
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		fmt.Println("Token validado, userID:", userID)

		c.Set("userID", userID)
		c.Next()
	}
}

func (m *AuthMiddleware) EnrichRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		userAgent := c.GetHeader("User-Agent")

		c.Set("ip_address", ip)
		c.Set("user_agent", userAgent)

		c.Next()
	}
}
