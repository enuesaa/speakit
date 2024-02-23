package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Validate(c *fiber.Ctx, body interface{}) error {
	if err := c.BodyParser(body); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
