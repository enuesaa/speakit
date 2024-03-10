package controller

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type FeedSchema struct {
	Name string `json:"name" validate:"required"`
	Url  string `json:"url" validate:"required"`
}

func NewFeedsController(repos repository.Repos) FeedsController {
	return FeedsController{
		repos,
	}
}

type FeedsController struct {
	repos repository.Repos
}

func (ctl *FeedsController) List(c *fiber.Ctx) error {
	feedSrv := service.NewFeedSevice(ctl.repos)

	items := make([]WithMetadata[FeedSchema], 0)
	for _, feed := range feedSrv.List() {
		items = append(items, WithMetadata[FeedSchema]{
			Id: feed.Id,
			Data: FeedSchema{
				Name: feed.Name,
				Url:  feed.Url,
			},
			Created:  "",
			Modified: "",
		})
	}
	return WithItems(c, items)
}

func (ctl *FeedsController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	feed := feedSrv.Get(id)
	res := WithMetadata[FeedSchema]{
		Id: feed.Id,
		Data: FeedSchema{
			Name: feed.Name,
			Url:  feed.Url,
		},
		Created:  "",
		Modified: "",
	}

	return WithData(c, res)
}

func (ctl *FeedsController) Create(c *fiber.Ctx) error {
	body := FeedSchema{}
	if err := Validate(c, &body); err != nil {
		return err
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	id := feedSrv.Create(service.Feed{
		Name: body.Name,
		Url:  body.Url,
	})

	return WithData(c, IdSchema{Id: id})
}

func (ctl *FeedsController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	feedSrv.Delete(id)

	return WithData(c, EmptySchema{})
}

type FeedfetchSchema struct{}

func (ctl *FeedsController) Fetch(c *fiber.Ctx) error {
	id := c.Params("id")
	body := FeedfetchSchema{}
	if err := Validate(c, &body); err != nil {
		return err
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	programSrv := service.NewProgramService(ctl.repos)
	realfeed, err := feedSrv.Refetch(id)
	if err != nil {
		return err
	}

	for _, realfeeditem := range realfeed.Items {
		programSrv.Create(service.Program{
			Title:     realfeeditem.Title,
			Content:   realfeeditem.Content,
			Converted: false,
		})
	}

	return WithData(c, EmptySchema{})
}
