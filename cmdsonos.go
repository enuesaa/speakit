package main

import (
	"fmt"
	"os"
	"strings"

	// "io"
	"time"

	"github.com/enuesaa/speakit/internal/sonosctl"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

var sonosCmd = &cobra.Command{
	Use:   "sonos",
	Short: "sonos",
	RunE: func(cmd *cobra.Command, args []string) error {
		go sonosctl.Serve()
		time.Sleep(2 * time.Second)

		entries, err := os.ReadDir("speechdata")
		if err != nil {
			return err
		}

		sonos, err := sonosctl.New()
		if err != nil {
			return err
		}

		go func ()  {
			app := fiber.New()
			app.Static("/", "./speechdata")
			app.Listen(":3000")
		}()

		for _, entry := range entries {
			filename := entry.Name()
			if !strings.HasSuffix(filename, "mp3") {
				continue
			}

			url := fmt.Sprintf("http://:3000/%s", filename)
			fmt.Println(url)

			res, err := sonos.SetUriRequest(url)
			if err != nil {
				return err
			}
			fmt.Printf("res: %+v\n", res)
			res, err = sonos.PlayRequest()
			if err != nil {
				return err
			}
			fmt.Printf("res: %+v\n", res)

			time.Sleep(5 * time.Second)
		}

		// res, err := sonos.subscribeSonos()
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("res: %+v\n", res)
		// resbody, err := io.ReadAll(res.Body)
		// if err != nil {
		// 	return err
		// }
		// fmt.Printf("resbody: %s\n", string(resbody))

		time.Sleep(100 * time.Second)

		return nil
	},
}
