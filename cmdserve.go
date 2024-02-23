package main

import (
	"github.com/enuesaa/speakit/admin"
	"github.com/enuesaa/speakit/pkg/controller"
	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		repos := repository.NewRepos()

		app := fiber.New()
		app.Use(cors.New())
		app.Use(logger.New())
		app.Use(controller.HandleData)

		feeds := controller.NewFeedsController(repos)
		app.Get("/api/feeds", feeds.List)
		app.Get("/api/feeds/:id", feeds.Get)
		app.Post("/api/feeds", feeds.Create)
		app.Delete("/api/feeds/:id", feeds.Delete)
		app.Post("/api/feeds/:id/fetch", feeds.Fetch)

		programs := controller.NewProgramsController(repos)
		app.Get("/api/programs", programs.List)
		app.Get("/api/programs/:id", programs.Get)
		app.Delete("/api/programs/:id", programs.Delete)
		app.Post("/api/programs/:id/convert", programs.Convert)
		app.Get("/api/programs/:id/audio", programs.GetAudio)
		app.Get("/*", admin.Serve)

		return app.Listen(":3000")
	},
}
