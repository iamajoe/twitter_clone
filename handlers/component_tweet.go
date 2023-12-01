package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HandleComponentTweetLike(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	tweet := Tweet{
		ID:        id,
		Author:    User{Username: "ajoesantos"},
		Message:   "With something so strong, a little bit can go a long way. It's hard to see things when you're too close. Take a step back and look. You can spend all day playing with mountains.",
		Likes:     11,
		IsLiked:   true,
		CreatedAt: time.Now().Add(-time.Minute * 10),
	}

	return c.Render("components/tweetOnly", fiber.Map{
		"Item": tweet,
	})
}
