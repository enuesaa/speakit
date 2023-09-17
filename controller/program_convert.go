package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProgramConvertSchema struct {
	Id string `json:"id" validate:"required"`
}

type ProgramConvertController struct {
	repos repository.Repos
}

func NewProgramConvertController(repos repository.Repos) ProgramConvertController {
	return ProgramConvertController{
		repos,
	}
}

func (ctl *ProgramConvertController) Create(c *fiber.Ctx) error {
	body := new(ProgramConvertSchema)
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
