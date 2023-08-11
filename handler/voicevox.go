package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type QueryRequest struct {
	Text string `json:"text" validate:"required"`
}

func CreateQuery(c *fiber.Ctx) error {
	// body := new(QueryRequest)
	// if err := c.BodyParser(body); err != nil {
	// 	return err
	// }
    // c.Response().Header.Del(fiber.HeaderConnection)
    // c.Response().Header.Del(fiber.HeaderServer)

	// validate := validator.New()
	// if err := validate.Struct(body); err != nil {
	// 	return err.(validator.ValidationErrors)
	// }
	text := c.Query("text")

	return proxy.Forward("http://voicevox:50021/audio_query?speaker=1&text=" + text)(c)
}
