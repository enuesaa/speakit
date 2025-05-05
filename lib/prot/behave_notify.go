package prot

type NotifyBehavior struct {
	app App
}

func (a *NotifyBehavior) Next() error {
	a.app.wait = false

	return a.app.speaker.CancelWait()
}

func (a *NotifyBehavior) Stop() error {
	a.app.wait = true

	return a.app.speaker.CancelWait()
}
