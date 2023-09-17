package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FeedfetchSchema struct {
	Id string `json:"id" validate:"required"` // feed id
}

type FeedfetchController struct {
	repos repository.Repos
}

func NewFeedfetchController(repos repository.Repos) FeedfetchController {
	return FeedfetchController{
		repos,
	}
}

func (ctl *FeedfetchController) Create(c *fiber.Ctx) error {
	body := new(FeedfetchSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	feedSrv := service.NewFeedSevice(ctl.repos)
	programSrv := service.NewProgramService(ctl.repos)
	realfeed, err := feedSrv.Refetch(body.Id)
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
