package prot

import "github.com/enuesaa/speakit/internal/eightbitctl"

type EightbitController struct {
	notify   *NotifyBehavior
	logger   Logger
	eightbit eightbitctl.Eightbit
}

func (c *EightbitController) Inject(logger Logger) {
	c.logger = logger
}

func (c *EightbitController) StartUp(logger Logger) error {
	c.eightbit = eightbitctl.New()

	c.eightbit.On(func(kc eightbitctl.KeyCode) {
		c.logger.Log("clicked: %s", kc)

		if kc == eightbitctl.KeyCodeA {
			if err := c.notify.Next(); err != nil {
				c.logger.LogE(err)
			}
		}
		if kc == eightbitctl.KeyCodeB {
			if err := c.notify.Stop(); err != nil {
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
