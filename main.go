package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/database"
	"github.com/hjuraev31/blog_go/routes"
	"github.com/joho/godotenv"
)


func main() {

	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Couldn`t read .evn file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	log.Fatal(app.Listen(":" + port))
}
