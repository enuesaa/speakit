package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
)

type ProgramsController struct {
	repos repository.Repos
}

func NewProgramsController(repos repository.Repos) ProgramsController {
	return ProgramsController {
		repos,
	}
}

func (ctl *ProgramsController) ListPrograms(c *fiber.Ctx) error {
	return c.JSON("")
}

// storage の id を返す
func (ctl *ProgramsController) GetProgram(c *fiber.Ctx) error {
	return c.JSON("")
}
