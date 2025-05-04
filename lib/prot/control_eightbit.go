package prot

import (
	"fmt"

	"github.com/enuesaa/speakit/internal/eightbitctl"
)

type EightbitController struct {
	app *App
	eightbit eightbitctl.Eightbit
}

func (c *EightbitController) StartUp(app *App) error {
	c.app = app
	c.eightbit = eightbitctl.New()

	fmt.Println("startup eightbit controller")

	c.eightbit.On(func(kc eightbitctl.KeyCode) {
		fmt.Printf("clicked: %s\n", kc)
	})

	if err := c.eightbit.Start(); err != nil {
		return err
	}
	fmt.Println("startup eightbit controller completed")

	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
