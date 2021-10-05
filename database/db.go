package database

import (
	"fifth_test/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", username, password, dbHost, dbPort, dbName)

	DB, err = gorm.Open(postgres.Open(dbURI))
	if err != nil {
		log.Fatal("Failed to connect database.")
	}

	log.Println("Connection opened")

	if err := DB.AutoMigrate(&models.Image{}); err != nil {
		log.Panic("Image migration failed")
	}

	if err := DB.AutoMigrate(&models.Post{}); err != nil {
		log.Panic("Post migration failed")
	}
}
