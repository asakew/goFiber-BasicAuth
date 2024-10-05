package database

import (
	"BasicAuth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	UsersDB *gorm.DB
	PostsDB *gorm.DB
)

func ConnectDB() {
	var err error
	// Database DSN
	userDSN := "host=localhost user=postgres password=Root dbname=BasicAuth_db port=5432 sslmode=disable timezone=Asia/Vladivostok"
	postDSN := "host=localhost user=postgres password=Root dbname=posts_db port=5432 sslmode=disable timezone=Asia/Vladivostok"

	// Connect to the database
	UsersDB, err = gorm.Open(postgres.Open(userDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	PostsDB, err = gorm.Open(postgres.Open(postDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	AutoMigrate() // Migrate the database

	log.Println("Connected to the database!")
}

func AutoMigrate() {
	err := UsersDB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = PostsDB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
