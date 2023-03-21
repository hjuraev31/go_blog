package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
