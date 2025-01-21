package utils

import (
	"MiniWallet/internal/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB initializes and returns a GORM database connection
func ConnectDB() *gorm.DB {
	// Using SQLite for simplicity; replace with your desired database configuration
	db, err := gorm.Open(sqlite.Open("MiniWallet.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate models (run this only in development; avoid in production)
	err = db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
