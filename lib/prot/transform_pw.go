package prot

import "github.com/playwright-community/playwright-go"

type PwTransformer struct {
	TransformFn func(*PwTransformer, *Record) error

	logger  Logger
	Pw      *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

func (g *PwTransformer) StartUp(logger Logger) error {
	g.logger = logger

	if err := playwright.Install(); err != nil {
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
	return g.TransformFn(g, record)
}

func (g *PwTransformer) Close() error {
	if g.Browser != nil {
		if err := g.Browser.Close(); err != nil {
			g.logger.LogE(err)
		}
	}
	if g.Pw != nil {
		if err := g.Pw.Stop(); err != nil {
			g.logger.LogE(err)
		}
	}
	return nil
}
