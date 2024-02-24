package config

import (
	"gorm.io/driver/postgres"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

// LoadConfig loads application configuration
func LoadConfig() {
	err := godotenv.Load() // Load environment variables from .env
	if err != nil {
		panic("Failed to load .env file")
	}

	DB = setupDatabaseConnection()
}

func setupDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	return db
}
