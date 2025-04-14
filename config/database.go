package config

import (
	"fiber-auth-app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	log.Println("👉 DSN:", dsn) // Debug DSN

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check connection
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get generic DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf(" Failed to ping database: %v", err)
	}

	// Migrate
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf(" Error migrating database: %v", err)
	}

	log.Println("✅ Database connected and migration successful")
}
