package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var IsProd bool

var DB *gorm.DB

func LoadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    IsProd = os.Getenv("ENV") == "prod"
}

func ConnectDatabase() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
    )
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    DB = database
    log.Println("Database connected")
}


