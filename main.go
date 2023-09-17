package main

import (
	"flag"
	"os"

	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var task string
	flag.StringVar(&task, "task", "serve", "")
	flag.Parse()

	if task == "print-openapi" {
		PrintOpenapi()
		return
	}
	if task == "serve" {
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
		// app.Post("/api/prefetch", feedfetch.Create) // 少し頭のいいエンドポイント
		
		convert := controller.NewConvertController(repos)
		app.Post("/api/convert", convert.Create)
	
		programs := controller.NewProgramsController(repos)
		app.Get("/api/programs", programs.List)
		app.Get("/api/programs/:id", programs.Get)
		// app.Get("/api/programs/:id/audio", programs.GetAudio)
		// app.Get("/api/programs/:id", programs.Delete)

		// deperecated
		storage := controller.NewStorageController(repos)
		app.Get("/api/storage/:id", storage.GetItem)
	
		// web route
		web := controller.NewWebController(env)
		app.Get("/*", web.Forward)

		app.Listen(":3000")
		return
	}
}
