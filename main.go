package main

import (
	"github.com/enuesaa/speakit/handler"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	feedsController := handler.NewFeedsController(repository.Repos {
		Redis: &repository.RedisRepository{},
	})
	api.Get("/feeds", feedsController.ListFeeds)
	api.Get("/feeds/:id", feedsController.GetFeed)
	api.Post("/feeds", feedsController.CreateFeed)
	api.Delete("/feeds/:id", feedsController.DeleteFeed)

	api.Post("/jobs", handler.CreateJob)
	api.Get("/jobs", handler.ListJobs)
	api.Get("/jobs/:id", handler.GetJob)
	api.Get("/contents", handler.ListContents)
	api.Get("/contents/:id", handler.GetContent)

	api.Post("/voicevox/query", handler.CreateQuery)

	// - GET /storage/{id}  ... wav file

	app.Get("/*", func(c *fiber.Ctx) error {
		path := c.OriginalURL()
		return proxy.Forward("http://admin:3000" + path)(c)
	})

	app.Listen(":3000")
}
