package main

import (
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"time"
	"twitter_clone/config"
	"twitter_clone/handlers"

	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func cssStylesheet(name string) (res string) {
	filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == name {
			p := strings.Replace(path, "views/", "", 1)
			res = "<link rel=\"stylesheet\" href=\"/" + p + "\">"
		}

		return nil
	})

	return res
}

func jsScript(name string) (res string) {
	filepath.Walk("views/public", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == name {
			p := strings.Replace(path, "views/", "", 1)
			res = "<script type=\"module\" src=\"/" + p + "\" defer></script>"
		}

		return nil
	})

	return res
}

func createGlobalFunctions(engine *jet.Engine) {
	engine.AddFunc("css", func(name string) (res template.HTML) {
		sheet := cssStylesheet(name)
		if len(sheet) > 0 {
			res = template.HTML(sheet)
		}

		return
	})

	engine.AddFunc("js", func(name string) (res template.HTML) {
		script := jsScript(name)
		if len(script) > 0 {
			res = template.HTML(script)
		}

		return
	})

	engine.AddFunc("HumanizeFn", func(date time.Time) string {
		return humanize.Time(date)
	})
}

func createViewEngine() fiber.Views {
	engine := jet.New("./views", ".jet.html")

	engine.Reload(config.Get().Env != "production")
	createGlobalFunctions(engine)

	return engine
}

func initViewRoutes(app fiber.Router) {
	// handle assets like global css for example
	app.Static("/public", "./views/public")

	// handle views
	app.Get("/", handlers.HandleFeed)
	app.Get("/metrics", handlers.HandleMetrics)

	app.Post("/tweet/:id/like", handlers.HandleComponentTweetLike)

	app.Use(handlers.NotFoundMiddleware)
}
