package prot

import "github.com/enuesaa/speakit/internal/eightbitctl"

type EightbitController struct {
	app      *App
	logger   Logger
	eightbit eightbitctl.Eightbit
}

func (c *EightbitController) StartUp(logger Logger, app *App) error {
	c.app = app
	c.logger = logger
	c.eightbit = eightbitctl.New()

	c.eightbit.On(func(kc eightbitctl.KeyCode) {
		c.logger.Log("clicked: %s", kc)

		if kc == eightbitctl.KeyCodeA {
			if err := app.Next(); err != nil {
				c.logger.LogE(err)
			}
		}
		if kc == eightbitctl.KeyCodeB {
			if err := app.Stop(); err != nil {
				c.logger.LogE(err)
			}
		}
	})

	if err := c.eightbit.Start(); err != nil {
		return err
	}
	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
