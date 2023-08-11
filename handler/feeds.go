package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

func ListFeeds(c *fiber.Ctx) error {
	return c.JSON("")
}

func GetFeed(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Printf("%s", id)

	return c.JSON("")
}

type FeedRequest struct {
    Name string `json:"name" validate:"required"`
	Url string `json:"url" validate:"required"`
}
func CreateFeed(c *fiber.Ctx) error {
	body := new(FeedRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err
	}

	fmt.Printf("%+v", body)

	return c.JSON("")
}

func DeleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Printf("%s", id)

	return c.JSON("")
}

func FetchFeed(c *fiber.Ctx) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://gigazine.net/news/rss_2.0/")
	fmt.Println(feed.Title)
	
	return c.JSON("")
}
