package controller

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

func (ctl *FeedsController) List(c *fiber.Ctx) error {
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

func (ctl *FeedsController) Get(c *fiber.Ctx) error {
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

func (ctl *FeedsController) Create(c *fiber.Ctx) error {
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

func (ctl *FeedsController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	feedSrv.Delete(id)

	return c.JSON(EmptySchema{})
}


type FeedfetchSchema struct {}

func (ctl *FeedsController) Fetch(c *fiber.Ctx) error {
	body := new(FeedfetchSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}
	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	programSrv := service.NewProgramService(ctl.repos)
	realfeed, err := feedSrv.Refetch(id)
	if err != nil {
		return err
	}
	
	for _, realfeeditem := range realfeed.Items {
		programSrv.Create(service.Program{
			Title: realfeeditem.Title,
			Content: realfeeditem.Content,
			Converted: false,
		})
	}

	return c.JSON(EmptySchema{})
}
