package controllers

import (
	"devmas_techcomp/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func MeController(c *gin.Context) {
	h_token := c.GetHeader("Authorization")

	token, err := services.JWTAuthService().ValidateToken(h_token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing authorization", "status": "error"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	user_id := claims["id"]

	var u User
	row := db.QueryRow("SELECT username, email, role, avatar FROM users WHERE id = $1", user_id)
	err = row.Scan(&u.Username, &u.Email, &u.Role, &u.Avatar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Get Data Error",
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "token is valid", "status": "success", "data": u})
}
