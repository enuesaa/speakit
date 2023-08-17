package controller

import (
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type WebController struct {
	env repository.Env
}

func NewWebController(env repository.Env) WebController {
	return WebController {
		env,
	}
}

func (ctl *WebController) Forward(c *fiber.Ctx) error {
	path := c.OriginalURL()
	url := "http://" + ctl.env.ADMIN_HOST + path
	return proxy.Forward(url)(c)
}
