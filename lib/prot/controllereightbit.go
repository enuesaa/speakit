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

	c.eightbit.On(func(kc eightbitctl.KeyCode) {
		fmt.Printf("clicked: %s\n", kc)
	})

	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
