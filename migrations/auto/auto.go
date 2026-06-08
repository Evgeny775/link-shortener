package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"link-shortener/internal/link"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(".env file not found")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&link.Link{})
}
