package controllers

import (
	"devmas_techcomp/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func EditProfileController(c *gin.Context) {
	h_token := c.GetHeader("Authorization")
	username := c.PostForm("username")
	avatar := c.PostForm("avatar")

	token, err := services.JWTAuthService().ValidateToken(h_token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing authorization", "status": "error"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	user_email := claims["email"]

	_, err = db.Exec("UPDATE users SET username = $1, avatar = $2 WHERE email = $3", username, avatar, user_email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update data err", "status": "error"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update data success", "status": "success"})
}
