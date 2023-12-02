package handlers

import (
	"fmt"

	"twitter_clone/config"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	conf := config.Get()
	if conf.Env != "production" {
		fmt.Println("ERR:", err)
	}

	return c.Render("error/500", nil)
}
