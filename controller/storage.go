package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/gofiber/fiber/v2"
)

type StorageController struct {
	repos repository.Repos
}

func NewStorageController(repos repository.Repos) StorageController {
	return StorageController{
		repos,
	}
}

func (ctl *StorageController) GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	programsSrv := service.NewProgramsService(ctl.repos)

	obj, err := programsSrv.Download(id)
	if err != nil {
		return c.JSON(EmptySchema{})
	}

	c.Response().SetBodyRaw([]byte(obj))
	c.Response().Header.SetContentType("audio/wav")
	return c.JSON(EmptySchema{})
}
