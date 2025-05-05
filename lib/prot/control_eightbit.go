package prot

import (
	"fmt"
	"os/exec"

	"github.com/enuesaa/speakit/internal/eightbitctl"
)

type EightbitController struct {
	volume   int
	notify   *NotifyBehavior
	log      *LogBehavior
	eightbit *eightbitctl.Eightbit
}

func (c *EightbitController) Inject(log *LogBehavior, notify *NotifyBehavior) {
	c.log = log
	c.notify = notify
	c.eightbit = eightbitctl.New()
}

func (c *EightbitController) StartUp() error {
	c.volume = 50

	c.eightbit.On(func(kc eightbitctl.KeyCode) {
		c.log.Info("clicked: %s", kc)

		switch kc {
		case eightbitctl.KeyCodeA:
			if err := c.notify.Next(); err != nil {
				c.log.Err(err)
			}
		case eightbitctl.KeyCodeB:
			if err := c.notify.Stop(); err != nil {
				c.log.Err(err)
			}
		case eightbitctl.KeyCodeUP:
			if c.volume > 100 {
				c.log.Info("invalid volume value")
				return
			}
			c.volume += 10
			if err := c.applyVolume(); err != nil {
				c.log.Err(err)
			}
		case eightbitctl.KeyCodeDOWN:
			if c.volume < 0 {
				c.log.Info("invalid volume value")
				return
			}
			c.volume -= 10
			if err := c.applyVolume(); err != nil {
				c.log.Err(err)
			}
		case eightbitctl.KeyCodeL:
			panic("exit from eightbit controller")
		}
	})

	return c.eightbit.Start()
}

func (c *EightbitController) applyVolume() error {
	c.log.Info("volume changes to %d%%", c.volume)

	target := fmt.Sprintf("%d%%", c.volume)
	cmd := exec.Command("amixer", "sset", "PCM", target)

	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
