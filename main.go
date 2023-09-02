package main

import (
	"flag"
	"os"

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

	app := fiber.New()

	// env
	env := repository.Env {
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST: os.Getenv("MINIO_HOST"),
		REDIS_HOST: os.Getenv("REDIS_HOST"),
		ADMIN_HOST: os.Getenv("ADMIN_HOST"),
	}
	repos := repository.NewRepos(env)

	// route
	createRoute(app, repos, env)

	app.Listen(":3000")
}
