package main

import (
	"flag"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var taskOpenapi bool
func init() {
	flag.BoolVar(&taskOpenapi, "emit-openapi", false, "create openapi.yaml")
}

func main() {
	flag.Parse()
	if taskOpenapi {
		emitOpenapi()
		return
	}
	serve()
}

func serve() {
	// env
	env := repository.Env{
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST:   os.Getenv("MINIO_HOST"),
		REDIS_HOST:   os.Getenv("REDIS_HOST"),
		ADMIN_HOST:   os.Getenv("ADMIN_HOST"),
	}
	repos := repository.NewRepos(env)

	app := fiber.New()
	app.Use(cors.New())

	// api route
	feeds := controller.NewFeedsController(repos)
	app.Get("/api/feeds", feeds.List)
	app.Get("/api/feeds/:id", feeds.Get)
	app.Post("/api/feeds", feeds.Create)
	app.Delete("/api/feeds/:id", feeds.Delete)
	app.Delete("/api/feeds/:id/fetch", feeds.Fetch)

	programs := controller.NewProgramsController(repos)
	app.Get("/api/programs", programs.List)
	app.Get("/api/programs/:id", programs.Get)
	app.Get("/api/programs/:id", programs.Delete)
	app.Post("/api/programs/:id/convert", programs.Convert)
	app.Get("/api/programs/:id/audio", programs.GetAudio)

	app.Listen(":3000")
}
