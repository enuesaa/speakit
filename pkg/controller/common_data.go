package controller

import (
	"github.com/gofiber/fiber/v2"
)

type ListSchema struct {
	Items interface{} `json:"items"`
}
type DataSchema struct {
	Data interface{} `json:"data"`
}

func WithItems(c *fiber.Ctx, items interface{}) error {
	c.Locals("items", items)
	return nil
}
func WithData(c *fiber.Ctx, data interface{}) error {
	c.Locals("data", data)
	return nil
}

func HandleData(c *fiber.Ctx) error {
	if err := c.Next(); err != nil {
		return err
	}

	items := c.Locals("items")
	if items != nil {
		res := ListSchema{
			Items: items,
		}
		return c.JSON(res)
	}
	data := c.Locals("data")
	if data != nil {
		res := DataSchema{
			Data: data,
		}
		return c.JSON(res)
	}
	return nil
}
