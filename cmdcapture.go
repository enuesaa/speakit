package main

import (
	"github.com/playwright-community/playwright-go"
	"github.com/spf13/cobra"
)

var captureUrl string

func init() {
	captureCmd.Flags().StringVar(&captureUrl, "url", "", "")
}

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "capture",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := playwright.Install(); err != nil {
			return err
		}

		pw, err := playwright.Run()
		if err != nil {
			return err
		}
		defer pw.Stop()

		browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
			Headless: playwright.Bool(true),
		})
		if err != nil {
			return err
		}
		defer browser.Close()

		page, err := browser.NewPage(playwright.BrowserNewPageOptions{
			UserAgent: playwright.String("Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/537.36 (KHTML, like Gecko) Version/16.0 Mobile/15E148 Safari/537.36"),
			Viewport: &playwright.Size{
				Width:  370,
				Height: 900,
			},
			DeviceScaleFactor: playwright.Float(3.0),
		})
		if err != nil {
			return err
		}

		if _, err := page.Goto(captureUrl); err != nil {
			return err
		}

		_, err = page.Screenshot(playwright.PageScreenshotOptions{
			Path: playwright.String("screenshot.png"),
		})
		if err != nil {
			return err
		}
		return nil
	},
}
