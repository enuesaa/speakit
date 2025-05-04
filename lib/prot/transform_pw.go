package prot

import "github.com/playwright-community/playwright-go"

type PwTransformer struct {
	TransformFn func(*PwTransformer, *Record) error

	logger Logger
	pw *playwright.Playwright
	browser playwright.Browser
	page playwright.Page
}

func (g *PwTransformer) StartUp(app *App) error {
	g.logger = app.Logger("pw")

	if err := playwright.Install(); err != nil {
		return err
	}
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	g.pw = pw

	browser, err := pw.Firefox.Launch()
	if err != nil {
		return err
	}
	g.browser = browser

	page, err := browser.NewPage()
	if err != nil {
		return err
	}
	g.page = page

	return nil
}

func (g *PwTransformer) Transform(record *Record) error {
	return g.TransformFn(g, record)
}

func (g *PwTransformer) Close() error {
	if g.browser != nil {
		if err := g.browser.Close(); err != nil {
			g.logger.LogE(err)
		}
	}
	if g.pw != nil {
		if err := g.pw.Stop(); err != nil {
			g.logger.LogE(err)
		}
	}
	return nil
}
