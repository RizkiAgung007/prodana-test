package config

import (
	"backend/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	
	err := godotenv.Load()
	if err != nil {
		log.Println("Gagal meload file .env: ", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", 
		host, user, pass, dbName, port)

	database, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal connect ke database: ", err)
	}

	log.Println("Dataase berhasil connect")

	err = database.AutoMigrate(&models.Role{}, &models.User{}, &models.Product{})
	if err != nil {
		log.Fatal("Gagal melakukan auto-migrate: ", err)
	}

	log.Println("Auto Migrate Berhasil")

	DB = database
}