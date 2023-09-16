package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FeedfetchSchema struct {
	Id string `json:"id" validate:"required"`
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

	id := c.Params("id")

	feedSrv := service.NewFeedSevice(ctl.repos)
	programSrv := service.NewProgramService(ctl.repos)
	voicevoxSrv := service.NewVoicevoxService(ctl.repos)
	realfeed := feedSrv.Refetch(id)
	
	for _, realfeeditem := range realfeed.Items {
		audioquery, _ := voicevoxSrv.AudioQuery(realfeeditem.Content)
		converted, _ := voicevoxSrv.Synthesis(audioquery)
		
		programSrv.Create(service.Program{
			Title: realfeeditem.Title,
			Content: converted, // todo: save to storage
		})
	}

	return c.JSON(EmptySchema{})
}
