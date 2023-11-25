package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"twitter_clone/config"
	"twitter_clone/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

func createGlobalFunctions(engine *django.Engine) {
	engine.AddFunc("css", func(name string) (res template.HTML) {
		filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Name() == name {
				p := strings.Replace(path, "views/", "", 1)
				res = template.HTML("<link rel=\"stylesheet\" href=\"/" + p + "\">")
			}

			return nil
		})

		return
	})

	engine.AddFunc("js", func(name string) (res template.HTML) {
		filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Name() == name {
				p := strings.Replace(path, "views/", "", 1)
				res = template.HTML("<script type=\"module\" src=\"/" + p + "\" defer>")
			}

			return nil
		})

		return
	})
}

func createViewEngine() fiber.Views {
	engine := django.New("./views", ".html")

	engine.Reload(config.Get().Env != "production")
	createGlobalFunctions(engine)

	return engine
}

func initViewRoutes(app fiber.Router) {
	// handle assets like global css for example
	app.Static("/public", "./views/public")

	// handle views
	app.Get("/", handlers.HandleFeed)
	app.Get("/bored", handlers.HandleBored)

	app.Use(handlers.NotFoundMiddleware)
}
