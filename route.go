package main

import (
	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func createApiRoute(app *fiber.App, repos repository.Repos) {
	router := app.Group("/api")

	feeds := controller.NewFeedsController(repos)
	router.Get("/feeds", feeds.ListFeeds)
	router.Get("/feeds/:id", feeds.GetFeed)
	router.Post("/feeds", feeds.CreateFeed)
	router.Delete("/feeds/:id", feeds.DeleteFeed)

	jobs := controller.NewJobsController(repos)
	router.Post("/jobs", jobs.CreateJob)

	programs := controller.NewProgramsController(repos)
	router.Get("/programs", programs.ListPrograms)
	router.Get("/programs/:id", programs.GetProgram)

	storage := controller.NewStorageController(repos)
	router.Get("/storage/:id", storage.GetItem)
}

func createWebRoute(app *fiber.App, repos repository.Repos) {
	app.Get("/*", func(c *fiber.Ctx) error {
		path := c.OriginalURL()
		return proxy.Forward("http://admin:3000" + path)(c)
	})
}
