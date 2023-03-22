package database

import (
	"fmt"
	"log"
	"os"

	"github.com/hjuraev31/blog_go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	fmt.Println("connecting...")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while reading .evn file")
	}

	dsn := os.Getenv("DSN")
	db, errorr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errorr != nil {
		panic("Couldn`t connect to database!")
	} else {
		fmt.Println("Connected!")
		log.Println("Connected successfully!")
	}

	DB = db
	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)
	return db

}
