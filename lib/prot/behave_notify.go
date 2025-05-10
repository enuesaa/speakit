package prot

import "time"

func newNotifyBehavior(speaker Speaker) *NotifyBehavior {
	return &NotifyBehavior{
		wait:    false,
		speaker: speaker,
	}
}

type NotifyBehavior struct {
	wait    bool
	speaker Speaker
}

func (a *NotifyBehavior) Next() error {
	a.wait = false
	return a.speaker.CancelWait()
}

func (a *NotifyBehavior) Stop() error {
	a.wait = true
	return a.speaker.CancelWait()
}

func (a *NotifyBehavior) Exit() {
	// TODO: SIGTERM should trigger similar behavior
	panic("exit")
}

func (a *NotifyBehavior) waitIfNeed() {
	if a.wait {
		for {
			if !a.wait {
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}
