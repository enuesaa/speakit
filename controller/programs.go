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
	programsSrv := service.NewProgramService(ctl.repos)
	list := programsSrv.List()
	response := ListSchema[string] {
		Items: make([]string, 0),
	}
	for _, item := range list {
		response.Items = append(response.Items, item.Id)
	}

	return c.JSON(response)
}

func (ctl *ProgramsController) GetProgram(c *fiber.Ctx) error {
	return c.JSON(EmptySchema{})
}
