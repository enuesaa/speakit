package controller

import (
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type ProgramSchema struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Converted bool   `json:"converted"`
}

func NewProgramsController(repos repository.Repos) ProgramsController {
	return ProgramsController{
		repos,
	}
}

type ProgramsController struct {
	repos repository.Repos
}

func (ctl *ProgramsController) List(c *fiber.Ctx) error {
	programSrv := service.NewProgramService(ctl.repos)

	items := make([]WithMetadata[ProgramSchema], 0)
	for _, program := range programSrv.List() {
		items = append(items, WithMetadata[ProgramSchema]{
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

	return WithItems(c, items)
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

	return WithData(c, res)
}

func (ctl *ProgramsController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	programSrv := service.NewProgramService(ctl.repos)
	programSrv.Delete(id)

	return WithData(c, EmptySchema{})
}

type ConvertSchema struct{}

func (ctl *ProgramsController) Convert(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ConvertSchema{}
	if err := Validate(c, &body); err != nil {
		return err
	}

	programSrv := service.NewProgramService(ctl.repos)
	programSrv.Convert(id)

	return WithData(c, EmptySchema{})
}

func (ctl *ProgramsController) GetAudio(c *fiber.Ctx) error {
	id := c.Params("id")
	programsSrv := service.NewProgramService(ctl.repos)

	obj, err := programsSrv.Download(id)
	if err != nil {
		return WithData(c, EmptySchema{})
	}

	c.Response().SetBodyRaw([]byte(obj))
	c.Response().Header.SetContentType("audio/wav")
	return nil
}
