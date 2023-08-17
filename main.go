package main

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/enuesaa/speakit/repository"
)

func main() {
	app := fiber.New()

	env := repository.Env {
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST: os.Getenv("MINIO_HOST"),
		REDIS_HOST: os.Getenv("REDIS_HOST"),
		ADMIN_HOST: os.Getenv("ADMIN_HOST"),
	}
	repos := repository.NewRealRepos(env)

	createApiRoute(app, repos)
	createWebRoute(app, env)

	app.Listen(":3000")
}
