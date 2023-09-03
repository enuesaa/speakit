package main

import (
	"flag"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	var task string
	flag.StringVar(&task, "task", "serve", "")
    flag.Parse()
	if task == "print-openapi" {
		PrintOpenapi()
		return
	}

	// env
	env := repository.Env {
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST: os.Getenv("MINIO_HOST"),
		REDIS_HOST: os.Getenv("REDIS_HOST"),
		ADMIN_HOST: os.Getenv("ADMIN_HOST"),
	}
	repos := repository.NewRepos(env)

	app := fiber.New()
	app.Use(logger())
	createRoute(app, repos, env)

	app.Listen(":3000")
}

func createRoute(app *fiber.App, repos repository.Repos, env repository.Env) {
	// api route
	feeds := controller.NewFeedsController(repos)
	app.Get("/api/feeds", feeds.ListFeeds)
	app.Get("/api/feeds/:id", feeds.GetFeed)
	app.Post("/api/feeds", feeds.CreateFeed)
	app.Delete("/api/feeds/:id", feeds.DeleteFeed)

	app.Post("/api/feeds/:id/fetch", feeds.RefetchFeed)

	programs := controller.NewProgramsController(repos)
	app.Get("/api/programs", programs.ListPrograms)
	app.Get("/api/programs/:id", programs.GetProgram)

	storage := controller.NewStorageController(repos)
	app.Get("/api/storage/:id", storage.GetItem)

	// web route
	web := controller.NewWebController(env)
	app.Get("/*", web.Forward)
}

func logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		logfile, err := os.Create("./tmp/app.log")
		if err != nil {
			return err
		}
		defer logfile.Close()
		log.SetOutput(logfile)

		c.Next()

		return nil
	}
}
