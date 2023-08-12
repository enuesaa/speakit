package handler

import (
	"fmt"
	"os"

	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/gofiber/fiber/v2"
)

type JobsController struct {
	repos repository.Repos
}

func NewJobsController(repos repository.Repos) JobsController {
	return JobsController{
		repos,
	}
}

func (ctl *JobsController) ListJobs(c *fiber.Ctx) error {
	return c.JSON("")
}

func (ctl *JobsController) GetJob(c *fiber.Ctx) error {
	return c.JSON("")
}

// fetch rss feed and request to convert. 202 を返したい
func (ctl *JobsController) CreateJob(c *fiber.Ctx) error {
	voicevoxSrv := service.NewVoicevoxService(ctl.repos)
	query, err := voicevoxSrv.AudioQuery("こんにちは")
	if err != nil {
		fmt.Println(err)
		return c.JSON("")
	}
	converted, err := voicevoxSrv.Synthesis(query)
	if err != nil {
		fmt.Println(err)
		return c.JSON("")
	}

	f, err := os.Create("example.wav")
	if err != nil {
		return c.JSON("")
	}
	defer f.Close()
	f.Write([]byte(converted))

	return c.JSON("")
}
