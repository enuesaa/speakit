package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/gofiber/fiber/v2"
)

type ProgramsController struct {
	repos repository.Repos
}

func NewProgramsController(repos repository.Repos) ProgramsController {
	return ProgramsController{
		repos,
	}
}

type ListProgramsResponse struct {
	Items []string `json:"items"`
}

func (ctl *ProgramsController) ListPrograms(c *fiber.Ctx) error {
	programsSrv := service.NewProgramsService(ctl.repos)
	list := programsSrv.List()
	response := ListProgramsResponse{
		Items: list,
	}

	return c.JSON(response)
}

func (ctl *ProgramsController) GetProgram(c *fiber.Ctx) error {
	return c.JSON("")
}
