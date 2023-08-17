package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/enuesaa/speakit/repository"
)

func main() {
	app := fiber.New()

	repos := repository.NewRealRepos()
	createApiRoute(app, repos)
	createWebRoute(app, repos)

	app.Listen(":3000")
}
