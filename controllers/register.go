package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c *gin.Context) {
	username := strings.TrimSpace(c.PostForm("username"))
	email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	cg_password := strings.TrimSpace(c.PostForm("cf_password"))

	if username == "" || password == "" || cg_password == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "BadRequest", "status": "error"})
		return
	}

	if password != cg_password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match", "status": "error"})
		return
	}

	var c_email string
	row := db.QueryRow("SELECT email FROM users WHERE email = $1", email)
	err := row.Scan(&c_email)
	if err != nil {
		//hash password bycrypt
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

		//create User row
		var id int
		err = db.QueryRow(`INSERT INTO users(username, email, password, role, avatar) VALUES($1, $2, $3, $4, $5) RETURNING id;`, username, email, hash, "user", "image/user").Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "status": "error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "register successfully",
			"status":  "success",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "register failed (email already registered)",
			"status":  "error",
		})
		return
	}
}
