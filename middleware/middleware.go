package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hjuraev31/blog_go/util"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.Parsejwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "User is not authenticated!",
		})
	}
	return c.Next()
}
