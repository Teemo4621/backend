package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func SetupDatabase() *sql.DB {
	godotenv.Load(".env")
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}