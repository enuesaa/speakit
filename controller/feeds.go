package controller

import (
	"fmt"

	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

type FeedSchema struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}

type FeedsController struct {
	repos repository.Repos
}

func NewFeedsController(repos repository.Repos) FeedsController {
	return FeedsController{
		repos,
	}
}

func (ctl *FeedsController) ListFeeds(c *fiber.Ctx) error {
	res := ListSchema[WithMetadata[FeedSchema]]{
		Items: make([]WithMetadata[FeedSchema], 0),
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	for _, feed := range feedSrv.List() {
		res.Items = append(res.Items, WithMetadata[FeedSchema] {
			Id: feed.Id,
			Data: FeedSchema {
				Name: feed.Name,
				Url: feed.Url,
			},
			Created: "",
			Modified: "",
		})
	}

	return c.JSON(res)
}

func (ctl *FeedsController) GetFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	feed := feedSrv.Get(id)
	res := WithMetadata[FeedSchema] {
		Id: feed.Id,
		Data: FeedSchema {
			Name: feed.Name,
			Url: feed.Url,
		},
		Created: "",
		Modified: "",
	}

	return c.JSON(res)
}

func (ctl *FeedsController) CreateFeed(c *fiber.Ctx) error {
	body := new(FeedSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	id := feedSrv.Create(service.Feed{
		Name: body.Name,
		Url:  body.Url,
	})

	return c.JSON(struct { Id string `json:"id"` } { Id: id })
}

func (ctl *FeedsController) DeleteFeed(c *fiber.Ctx) error {
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	feedSrv.Delete(id)

	return c.JSON(EmptySchema{})
}

// put job
func (ctl *FeedsController) FetchFeed(c *fiber.Ctx) error {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://gigazine.net/news/rss_2.0/")
	fmt.Println(feed.Title)

	return c.JSON(EmptySchema{})
}
