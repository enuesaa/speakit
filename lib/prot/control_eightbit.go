package prot

import (
	"github.com/enuesaa/speakit/internal/eightbitctl"
	"github.com/itchyny/volume-go"
)

type EightbitController struct {
	notify   *NotifyBehavior
	log      *LogBehavior
	eightbit *eightbitctl.Eightbit
}

func (c *EightbitController) Inject(log *LogBehavior, notify *NotifyBehavior) {
	c.log = log
	c.notify = notify
}

func (c *EightbitController) StartUp() error {
	c.eightbit = eightbitctl.New()

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

		if kc == eightbitctl.KeyCodeUP || kc == eightbitctl.KeyCodeDOWN {
			vol, err := volume.GetVolume()
			if err != nil {
				c.log.LogE(err)
				return
			}
			if kc == eightbitctl.KeyCodeUP {
				vol += 10
			} else {
				vol -= 10
			}
			if 0 <= vol && vol <= 100 {
				if err := volume.SetVolume(10); err != nil {
					c.log.LogE(err)
				}
				return
			}
			c.log.Log("invalid volume value: %d", vol)
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
