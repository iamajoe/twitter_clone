package main

import (
	"context"
	"io"
	"twitter_clone/handlers"

	views_components "twitter_clone/views/components"
	views_routes "twitter_clone/views/routes"
	views_routes_feed "twitter_clone/views/routes/feed"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

type TemplViewEngine struct {
	views map[string](func(data fiber.Map) templ.Component)
}

func (e *TemplViewEngine) Load() error {
	e.views = make(map[string]func(data fiber.Map) templ.Component)

	// add your route components here
	e.views["routes/feed/index"] = views_routes_feed.Route
	e.views["routes/metrics"] = views_routes.RouteMetrics
	e.views["components/tweetOnly"] = views_components.TweetOnly

	return nil
}

func (e *TemplViewEngine) Render(w io.Writer, tmpl string, data interface{}, args ...string) error {
	template, ok := e.views[tmpl]
	if !ok {
		// TODO: should render a 500
		w.Write([]byte("500: template not found"))
		return nil
	}

	dataMap, ok := data.(fiber.Map)
	if !ok {
		// TODO: should render a 500
		w.Write([]byte("500: data not fiber.Map"))
		return nil
	}

	// render the template
	return template(dataMap).Render(context.Background(), w)
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
