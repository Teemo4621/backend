package middleware

import (
	"devmas_techcomp/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func MiddlewareUnauthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "" {
			if strings.HasPrefix(token, "Bearer ") {
				token = token[7:]
				_, err := services.JWTAuthService().ValidateToken(token)
				if err == nil {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "๊User Is Logged in", "status": "error"})
					c.Abort()
					return
				}

				fmt.Println(err)
			}
		}

		c.Next()
	}
}
