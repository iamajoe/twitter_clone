package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleBored(c *fiber.Ctx) error {
	return c.Render("routes/bored", fiber.Map{})
}
