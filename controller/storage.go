package controller

import (
	"github.com/enuesaa/speakit/repository"
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
	return c.JSON("")
}
