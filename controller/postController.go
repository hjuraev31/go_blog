package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/database"
	"github.com/hjuraev31/blog_go/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to pass body!")
	}

	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid paylod",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Successfully posted!",
	})
}

func AllPost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to pass body")
	}
	//---------------to be continued------------------
	return nil
}
