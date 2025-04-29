package main

import "github.com/enuesaa/speakit/cmd"

func main() {
	app := cmd.New()

	if err := app.Execute(); err != nil {
		panic(err)
	}
}
