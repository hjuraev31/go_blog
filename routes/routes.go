package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/controller"
	"github.com/hjuraev31/blog_go/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
}
