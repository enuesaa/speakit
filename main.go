package main

import (
	"strings"

	"github.com/enuesaa/speakit/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	app.Get("/feeds", handler.ListFeeds)
	app.Get("/feeds/:id", handler.GetFeed)
	app.Post("/feeds", handler.CreateFeed)
	app.Delete("/feeds/:id", handler.DeleteFeed)

	// - POST /jobs ... fetch rss feed and request to convert. 202 を返したい
	// - GET /jobs
	// - GET /jobs/{id}

	// - GET /contents ... 一覧
	// - GET /contents/{id} ... asset id を返す

	// - GET /assets/{id}  ... wav file

	// - GET /_admin
	// - GET /_admin/feeds
	// - GET /_admin/player
	app.Get("/_admin/*", func(c *fiber.Ctx) error {
		path := c.OriginalURL()
		path = strings.TrimLeft(path, "/_admin")
		return proxy.Forward("http://admin:3000" + path)(c)
	})

	app.Listen(":3000")
}
