package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/enuesaa/speakit/controller"
	"github.com/enuesaa/speakit/repository"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
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
	env := repository.Env{
		MINIO_BUCKET: os.Getenv("MINIO_BUCKET"),
		MINIO_HOST:   os.Getenv("MINIO_HOST"),
		REDIS_HOST:   os.Getenv("REDIS_HOST"),
		ADMIN_HOST:   os.Getenv("ADMIN_HOST"),
	}
	repos := repository.NewRepos(env)

	app := fiber.New()
	// app.Use(logger())
	app.Use(adaptor.HTTPMiddleware(monitorSentry))
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

func monitorSentry(next http.Handler) http.Handler {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("SENTRY_DSN"),
		AttachStacktrace: true,
		EnableTracing: true,
		TracesSampleRate: 1.0,
		ProfilesSampleRate: 1.0,
	})
	if err != nil {
		return next
	}
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("It works!")

	sentryHandler := sentryhttp.New(sentryhttp.Options{})
	return sentryHandler.Handle(next)
}
