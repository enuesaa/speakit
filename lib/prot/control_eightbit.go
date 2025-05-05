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
		c.log.Log("clicked: %s", kc)

		if kc == eightbitctl.KeyCodeA {
			if err := c.notify.Next(); err != nil {
				c.log.LogE(err)
			}
		}
		if kc == eightbitctl.KeyCodeB {
			if err := c.notify.Stop(); err != nil {
				c.log.LogE(err)
			}
		}

		if kc == eightbitctl.KeyCodeUP {
			if c.volume > 100 {
				c.log.Log("invalid volume value")
				return
			}
			c.volume += 10
			if err := c.applyVolume(); err != nil {
				c.log.LogE(err)
			}
		}
		if kc == eightbitctl.KeyCodeDOWN {
			if c.volume < 0 {
				c.log.Log("invalid volume value")
				return
			}
			c.volume -= 10
			if err := c.applyVolume(); err != nil {
				c.log.LogE(err)
			}
		}
	})

	if err := c.eightbit.Start(); err != nil {
		return err
	}
	return nil
}

func (c *EightbitController) applyVolume() error {
	target := fmt.Sprintf("%d%", c.volume)
	cmd := exec.Command("amixer", "sset", "PCM", target)

	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
