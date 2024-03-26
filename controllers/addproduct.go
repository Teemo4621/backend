package controllers

import (
	"devmas_techcomp/services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AddProductController(c *gin.Context) {
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

	name := strings.TrimSpace(c.PostForm("name"))
	details := strings.TrimSpace(c.PostForm("details"))
	image := strings.TrimSpace(c.PostForm("image"))
	category := strings.TrimSpace(c.PostForm("category"))
	priceStr := strings.TrimSpace(c.PostForm("price"))

	if name == "" || details == "" || priceStr == "" || image == "" || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
			"status":  "error",
		})
		return
	}
	//convert string to float32
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid price format",
			"status":  "error",
		})
		return
	}

	product := Product{
		Name:     name,
		Details:  details,
		Image:    image,
		Category: category,
		Price:    float32(price),
	}

	_, err = db.Exec("INSERT INTO products(name, price, details, image, category) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Price, product.Details, product.Image, product.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "status": "error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Add Product Successfully",
		"status":  "success",
		"data":    product,
	})
}
