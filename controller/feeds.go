package controller

import (
	"fmt"

	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

type FeedsController struct {
	repos repository.Repos
}

func NewFeedsController(repos repository.Repos) FeedsController {
	return FeedsController {
		repos,
	}
}

func (ctl *FeedsController) ListFeeds(c *fiber.Ctx) error {
	return c.JSON("")
}

func (ctl *FeedsController) GetFeed(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Printf("%s", id)

	return c.JSON("")
}

type FeedRequest struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}
type FeedResponse struct {}

func (ctl *FeedsController) CreateFeed(c *fiber.Ctx) error {
	body := new(FeedRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	feedSrv.Create(service.Feed {
		Name: body.Name,
		Url: body.Url,
	})

	return c.JSON(FeedResponse{})
}

func (ctl *FeedsController) DeleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Printf("%s", id)

	return c.JSON("")
}

func (ctl *FeedsController) FetchFeed(c *fiber.Ctx) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://gigazine.net/news/rss_2.0/")
	fmt.Println(feed.Title)

	return c.JSON("")
}
