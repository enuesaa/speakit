package main

import (
	"flag"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var openapi bool
	flag.BoolVar(&openapi, "emit-openapi", false, "create openapi.yaml")
	if openapi {
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

	// api route
	feeds := controller.NewFeedsController(repos)
	app.Get("/api/feeds", feeds.List)
	app.Get("/api/feeds/:id", feeds.Get)
	app.Post("/api/feeds", feeds.Create)
	app.Delete("/api/feeds/:id", feeds.Delete)

	feedfetch := controller.NewFeedfetchController(repos)
	app.Post("/api/fetch", feedfetch.Create)
	
	convert := controller.NewConvertController(repos)
	app.Post("/api/convert", convert.Create)

	programs := controller.NewProgramsController(repos)
	app.Get("/api/programs", programs.List)
	app.Get("/api/programs/:id", programs.Get)
	app.Get("/api/programs/:id", programs.Delete)
	app.Get("/api/programs/:id/audio", programs.GetAudio)

	// web route
	web := controller.NewWebController(env)
	app.Get("/*", web.Forward)

	app.Listen(":3000")
}
