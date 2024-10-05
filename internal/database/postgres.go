package database

import (
	"BasicAuth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var UsersDB *gorm.DB

func ConnectDB() {
	var err error
	dsn := "host=localhost user=postgres password=Root dbname=BasicAuth_db port=5432 sslmode=disable" // Update with your credentials
	UsersDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Auto migrate the User model
	err = UsersDB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	log.Println("Connected to the database!")
}
