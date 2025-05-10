package kitprot

import "github.com/playwright-community/playwright-go"

type PwBehavior struct {
	Pw      *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

func (g *PwBehavior) Launch() error {
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

func (g *PwBehavior) Close() error {
	var err error
	if g.Browser != nil {
		err = g.Browser.Close()
	}
	if g.Pw != nil {
		err = g.Pw.Stop()
	}
	return err
}
