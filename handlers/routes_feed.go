package handlers

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
}

type Tweet struct {
	Author    User
	Message   string
	Likes     int
	IsLiked   bool
	CreatedAt time.Time
}

func HumanizeFn(date time.Time) string {
	return "there was a date here"
}

func HandleFeed(c *fiber.Ctx) error {
	tweets := []Tweet{
		{
			Author:    User{Username: "ajoesantos"},
			Message:   "With something so strong, a little bit can go a long way. It's hard to see things when you're too close. Take a step back and look. You can spend all day playing with mountains.",
			Likes:     10,
			IsLiked:   true,
			CreatedAt: time.Now().Add(-time.Minute * 10),
		},
		{
			Author:    User{Username: "anthdm"},
			Message:   "Just relax and let it flow. That easy. Let's build some happy little clouds up here. Trees get lonely too, so we'll give him a little friend.",
			Likes:     1,
			IsLiked:   false,
			CreatedAt: time.Now().Add(-time.Hour * 7),
		},
	}

	return c.Render("routes/feed/index", fiber.Map{
		"humanizeFn": humanize.Time,
		"tweets":     tweets,
	})
}
