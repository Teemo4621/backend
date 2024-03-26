package controllers

import (
	"devmas_techcomp/database"
)

type User struct {
	Username string
	Email    string
	Role     string
	Avatar   string
}

type Product struct {
	Id       int
	Name     string
	Details  string
	Price    float32
	Category string
	Image    string
}

var db = database.SetupDatabase()

func SetupController() {
	defer database.SetupDatabase().Close()
}
