package controllers

import (
	"devmas_techcomp/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c *gin.Context) {
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))

	if email == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest", "status": "error"})
		return
	}

	var c_email string
	row := db.QueryRow("SELECT email FROM users WHERE email = $1", email)
	err := row.Scan(&c_email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "You have entered an invalid username or password", "status": "error"})
		return
	}

	var c_password string
	var id int
	row = db.QueryRow("SELECT password, id FROM users WHERE email = $1", email)
	err = row.Scan(&c_password, &id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "You have entered an invalid username or password", "status": "error"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(c_password), []byte(password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid password", "status": "error"})
		return
	}

	token, err := services.JWTAuthService().GenerateToken(id, email, true, 6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT Gen Token Error", "status": "error"})
		return
	}

	refreshToken, err := services.JWTAuthService().GenerateToken(id, email, true, 24)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JWT Gen Ref Error", "status": "error"})
		return
	}

	_, err = db.Exec("UPDATE users SET refresh_token = $1 WHERE email = $2", refreshToken, email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err, "status": "error"})
		return
	}

	data := gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"status":  "success",
		"data":    data,
	})
}
