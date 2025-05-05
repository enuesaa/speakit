package prot

import "github.com/playwright-community/playwright-go"

type PwTransformer struct {
	Logger  Logger
	Pw      *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

func (g *PwTransformer) Inject(logger Logger) {
	g.Logger = logger
}

func (g *PwTransformer) StartUp(logger Logger) error {
	if err := playwright.Install(&playwright.RunOptions{Browsers: []string{"firefox"}}); err != nil {
		return err
	}
	pw, err := playwright.Run()
	if err != nil {
		return err
	}
	g.Pw = pw

	browser, err := pw.Firefox.Launch()
	if err != nil {
		return err
	}
	g.Browser = browser

	page, err := browser.NewPage()
	if err != nil {
		return err
	}
	g.Page = page

	return nil
}

func (g *PwTransformer) Transform(record *Record) error {
	return nil
}

func (g *PwTransformer) Close() error {
	if g.Browser != nil {
		if err := g.Browser.Close(); err != nil {
			g.Logger.LogE(err)
		}
	}
	if g.Pw != nil {
		if err := g.Pw.Stop(); err != nil {
			g.Logger.LogE(err)
		}
	}
	return nil
}
