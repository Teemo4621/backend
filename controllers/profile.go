package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProfileController(c *gin.Context) {
	id := strings.TrimSpace(c.Query("id"))

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"status":  "error",
		})
		return
	}

	var u User
	row := db.QueryRow("SELECT username, email, role, avatar FROM users WHERE id = $1", id)
	err := row.Scan(&u.Username, &u.Email, &u.Role, &u.Avatar)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"status":  "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"status":  "success",
		"data":    u,
	})
}
