package prot

type Controller interface {
	StartUp(app *App) error
	Close() error
}

type SampleController struct {
	app *App
}

func (c *SampleController) StartUp(app *App) error {
	c.app = app

	// c.app.Next()

	return nil
}

func (c *SampleController) Close() error {
	return nil
}
