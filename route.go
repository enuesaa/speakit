package main

import (
	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
)

func createRoute(app *fiber.App, repos repository.Repos, env repository.Env) {
	// api route
	feeds := controller.NewFeedsController(repos)
	app.Get("/api/feeds", feeds.ListFeeds)
	app.Get("/api/feeds/:id", feeds.GetFeed)
	app.Post("/api/feeds", feeds.CreateFeed)
	app.Delete("/api/feeds/:id", feeds.DeleteFeed)

	jobs := controller.NewJobsController(repos)
	app.Post("/api/jobs", jobs.CreateJob)

	programs := controller.NewProgramsController(repos)
	app.Get("/api/programs", programs.ListPrograms)
	app.Get("/api/programs/:id", programs.GetProgram)

	storage := controller.NewStorageController(repos)
	app.Get("/api/storage/:id", storage.GetItem)

	// web route
	web := controller.NewWebController(env)
	app.Get("/*", web.Forward)
}
