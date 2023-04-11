package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/controller"
	"github.com/hjuraev31/blog_go/middleware"
	"github.com/hjuraev31/blog_go/models"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Get("/api/test/:param", func(c *fiber.Ctx) error {
		var blog models.Blog
		if err := c.BodyParser(&blog); err != nil {
			fmt.Println("error while parsing body")
		}
		fmt.Println(blog.Title)
		c.Append("Link", "http://google.com", "http://localhost")
		return c.SendString("param: " + c.Params("param"))
	})

	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allposts", controller.AllPost)
	app.Get("/api/get_post_detail/:id", controller.GetPostDetail)
	app.Put("/api/update_post/:id", controller.UpdatePost)
	app.Get("/api/unique_post", controller.UniquePost)
	app.Get("/api/delete_post/:id", controller.DeletePost)
	app.Post("/api/upload_image", controller.UploadImage)
	app.Static("/api/uploads", "./uploads")
}
