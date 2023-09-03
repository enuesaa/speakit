package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/gofiber/fiber/v2"
)

type ProgramSchema struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
}

type ProgramsController struct {
	repos repository.Repos
}

func NewProgramsController(repos repository.Repos) ProgramsController {
	return ProgramsController{
		repos,
	}
}

func (ctl *ProgramsController) ListPrograms(c *fiber.Ctx) error {
	programsSrv := service.NewProgramsService(ctl.repos)
	list := programsSrv.ListKeys()
	response := ListSchema[string] {
		Items: list,
	}

	return c.JSON(response)
}

func (ctl *ProgramsController) GetProgram(c *fiber.Ctx) error {
	return c.JSON(EmptySchema{})
}
