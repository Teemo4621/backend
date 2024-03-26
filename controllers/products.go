package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ProductController(c *gin.Context) {
	name := strings.TrimSpace(c.Query("name"))
	category := strings.TrimSpace(c.Query("category"))
	var query string
	var args []interface{}

	if name != "" && category != "" {
		query = "SELECT id, name, details, price, image, category FROM products WHERE name = $1 AND category = $2"
		args = append(args, name, category)
	} else if name != "" {
		query = "SELECT id, name, details, price, image, category FROM products WHERE name = $1"
		args = append(args, name)
	} else if category != "" {
		query = "SELECT id, name, details, price, image, category FROM products WHERE category = $1"
		args = append(args, category)
	} else {
		query = "SELECT id, name, details, price, image, category FROM products LIMIT 10"
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  "error",
		})
		return
	}
	defer rows.Close()

	var productArray []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Details, &product.Price, &product.Image, &product.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"status":  "error",
			})
			return
		}
		productArray = append(productArray, product)
	}

	if productArray == nil {
		productArray = []Product{}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Products successfully",
		"status":  "success",
		"data": gin.H{
			"products": productArray,
		},
	})
}
