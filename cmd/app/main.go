package main

import (
	"fmt"
	"log"
	"twitter_clone/config"
	"twitter_clone/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	c := config.Get()

	app := fiber.New(fiber.Config{
		ErrorHandler:          handlers.ErrorHandler,
		DisableStartupMessage: true,
		PassLocalsToViews:     true,
		Views:                 &TemplViewEngine{},
	})

	viewApp := app.Group("/")
	initViewRoutes(viewApp)

	listenAddr := c.Host + c.Port
	fmt.Printf("app running in %s and listening on: %s\n", c.Env, listenAddr)
	log.Fatal(app.Listen(listenAddr))
}
