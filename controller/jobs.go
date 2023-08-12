package controller

import (
	"fmt"

	"github.com/enuesaa/speakit/repository"
	"github.com/enuesaa/speakit/service"
	"github.com/go-playground/validator/v10"
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

type JobRequest struct {
	Text string `json:"text" validate:"required"`
}
func (ctl *JobsController) CreateJob(c *fiber.Ctx) error {
	body := new(JobRequest)
	if err := c.BodyParser(body); err != nil {
		return err
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return err.(validator.ValidationErrors)
	}

	voicevoxSrv := service.NewVoicevoxService(ctl.repos)
	programsSrv := service.NewProgramsService(ctl.repos)
	query, err := voicevoxSrv.AudioQuery(body.Text)
	if err != nil {
		fmt.Println(err)
		return c.JSON("")
	}
	converted, err := voicevoxSrv.Synthesis(query)
	if err != nil {
		fmt.Println(err)
		return c.JSON("")
	}

	programsSrv.Create("aaaa.wav", converted)

	return c.JSON("")
}
