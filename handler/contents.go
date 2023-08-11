package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ListContents(c *fiber.Ctx) error {
	return c.JSON("")
}

// storage の id を返す
func GetContent(c *fiber.Ctx) error {
	return c.JSON("")
}
