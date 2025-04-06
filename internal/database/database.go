package database

import (
	"log"

	user "github.com/linzheng99/go-backend-demo/features/user/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&user.User{})

	if err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return nil
}
