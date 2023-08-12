package main

import (
	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	api := app.Group("/api")
	repos := repository.Repos {
		Redis: &repository.RedisRepository{},
		Httpcall: &repository.HttpcallRepository{},
		Minio: &repository.MinioRepository{},
	}
	feeds := controller.NewFeedsController(repos)
	api.Get("/feeds", feeds.ListFeeds)
	api.Get("/feeds/:id", feeds.GetFeed)
	api.Post("/feeds", feeds.CreateFeed)
	api.Delete("/feeds/:id", feeds.DeleteFeed)

	jobs := controller.NewJobsController(repos)
	api.Post("/jobs", jobs.CreateJob)
	api.Get("/jobs", jobs.ListJobs)
	api.Get("/jobs/:id", jobs.GetJob)

	programs := controller.NewProgramsController(repos)
	api.Get("/programs", programs.ListPrograms)
	api.Get("/programs/:id", programs.GetProgram)

	// - GET /storage/{id}  ... wav file

	app.Get("/*", func(c *fiber.Ctx) error {
		path := c.OriginalURL()
		return proxy.Forward("http://admin:3000" + path)(c)
	})

	app.Listen(":3000")
}
