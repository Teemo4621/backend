package main

import (
	"devmas_techcomp/controllers"
	"devmas_techcomp/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func GetEnv() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	//setup comtrollers and connect database
	controllers.SetupController()
	fmt.Println("Server Start http://localhost:" + GetEnv())

	r := gin.Default()

	api := r.Group("api")

	auth := api.Group("auth")
	auth.Use(middleware.MiddlewareUnauthorized())
	auth.POST("/register", controllers.RegisterController)
	auth.POST("/login", controllers.LoginController)

	users := api.Group("users")

	users.GET("/profile", controllers.ProfileController)
	users.GET("/me", middleware.MiddlewareAuthorized(), controllers.MeController)
	users.POST("/me/edit", middleware.MiddlewareAuthorized(), controllers.EditProfileController)

	products := api.Group("products")
	products.GET("/", controllers.ProductController)
	products.POST("/add", middleware.MiddlewareAuthorized(), controllers.AddProductController)
	products.POST("/edit", middleware.MiddlewareAuthorized(), controllers.EditProductController)
	products.POST("/delete", middleware.MiddlewareAuthorized(), controllers.DeleteProductController)

	r.Run(":" + GetEnv())
}
