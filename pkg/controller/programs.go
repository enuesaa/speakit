package controller

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProgramSchema struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Converted bool   `json:"converted"`
}

type ProgramsController struct {
	repos repository.Repos
}

func NewProgramsController(repos repository.Repos) ProgramsController {
	return ProgramsController{
		repos,
	}
}

func (ctl *ProgramsController) List(c *fiber.Ctx) error {
	res := ListSchema[WithMetadata[ProgramSchema]]{
		Items: make([]WithMetadata[ProgramSchema], 0),
	}

	programSrv := service.NewProgramService(ctl.repos)
	for _, program := range programSrv.List() {
		res.Items = append(res.Items, WithMetadata[ProgramSchema]{
			Id: program.Id,
			Data: ProgramSchema{
				Title:     program.Title,
				Content:   program.Content,
				Converted: program.Converted,
			},
			Created:  "",
			Modified: "",
		})
	}

	return c.JSON(res)
}

func (ctl *ProgramsController) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	programSrv := service.NewProgramService(ctl.repos)
	program := programSrv.Get(id)
	res := WithMetadata[ProgramSchema]{
		Id: program.Id,
		Data: ProgramSchema{
			Title:     program.Title,
			Content:   program.Content,
			Converted: program.Converted,
		},
		Created:  "",
		Modified: "",
	}

	return c.JSON(res)
}

func (ctl *ProgramsController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	programSrv := service.NewProgramService(ctl.repos)
	programSrv.Delete(id)

	return c.JSON(EmptySchema{})
}

type ConvertSchema struct{}

func (ctl *ProgramsController) Convert(c *fiber.Ctx) error {
	body := new(ConvertSchema)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}
	id := c.Params("id")

	programSrv := service.NewProgramService(ctl.repos)
	programSrv.Convert(id)

	return c.JSON(EmptySchema{})
}

func (ctl *ProgramsController) GetAudio(c *fiber.Ctx) error {
	id := c.Params("id")
	programsSrv := service.NewProgramService(ctl.repos)

	obj, err := programsSrv.Download(id)
	if err != nil {
		return c.JSON(EmptySchema{})
	}

	c.Response().SetBodyRaw([]byte(obj))
	c.Response().Header.SetContentType("audio/wav")
	return nil
}
