package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ConvertSchema struct {
	Id string `json:"id" validate:"required"`
}

type ConvertController struct {
	repos repository.Repos
}

func NewConvertController(repos repository.Repos) ConvertController {
	return ConvertController{
		repos,
	}
}

func (ctl *ConvertController) Create(c *fiber.Ctx) error {
	body := new(ConvertSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	programSrv := service.NewProgramService(ctl.repos)
	program := programSrv.Get(body.Id)

	voicevoxSrv := service.NewVoicevoxService(ctl.repos)
	audioquery, _ := voicevoxSrv.AudioQuery(program.Title)
	converted, _ := voicevoxSrv.Synthesis(audioquery)

	programSrv.Upload(body.Id, converted)

	return c.JSON(EmptySchema{})
}
