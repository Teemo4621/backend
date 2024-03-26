package middleware

import (
	"devmas_techcomp/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization", "status": "error"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization", "status": "error"})
			c.Abort()
			return
		}

		token = token[7:]
		_, err := services.JWTAuthService().ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token is expired", "status": "error"})
			c.Abort()
			return
		}

		c.Next()
	}
}
