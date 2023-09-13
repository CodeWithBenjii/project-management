package db

import (
	"fmt"
	"log"

	"benjaminlee.dev/ProjectManagement/config"
	"benjaminlee.dev/ProjectManagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(cfg *config.Config) {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=5438 sslmode=disable", cfg.DatabaseUsername, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabaseURL)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	err = db.AutoMigrate(models.User{})
	if err != nil {
		log.Printf("Failed to migrate: %v", err)
	}
	log.Printf("Connected to the database!")
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to get DB instance: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Failed to close DB: %v", err)
	}
}

func Migrate(models ...interface{}) {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}
}
