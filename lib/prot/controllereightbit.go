package prot

type EightbitController struct {
	app *App
}

func (c *EightbitController) StartUp(app *App) error {
	c.app = app

	// c.app.Next()

	return nil
}

func (c *EightbitController) Close() error {
	return nil
}
