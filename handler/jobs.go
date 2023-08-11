package handler

import (
	"github.com/gofiber/fiber/v2"
)

func ListJobs(c *fiber.Ctx) error {
	return c.JSON("")
}

func GetJob(c *fiber.Ctx) error {
	return c.JSON("")
}

// fetch rss feed and request to convert. 202 を返したい
func CreateJob(c *fiber.Ctx) error {
	return c.JSON("")
}
