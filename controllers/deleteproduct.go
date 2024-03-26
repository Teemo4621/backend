package controllers

import (
	"devmas_techcomp/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func DeleteProductController(c *gin.Context) {
	// Authorization header
	h_token := c.GetHeader("Authorization")
	token, err := services.JWTAuthService().ValidateToken(h_token[7:])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing authorization", "status": "error"})
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	user_id := claims["id"]

	var u User
	// check user role
	row := db.QueryRow("SELECT username, email, role, avatar FROM users WHERE id = $1", user_id)
	err = row.Scan(&u.Username, &u.Email, &u.Role, &u.Avatar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Get Data Error",
			"status":  "error",
		})
		return
	}

	if u.Role != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Permission",
			"status":  "error",
		})
		return
	}

	//delete product
	id := strings.TrimSpace(c.PostForm("id"))

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"status":  "error",
		})
		return
	}

	result, err := db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status": "error"})
		return
	}

	//check product status
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product not found",
			"status":  "error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
		"status":  "success",
	})
}
