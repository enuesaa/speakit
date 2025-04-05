package main

import (
	"fmt"
	// "io"
	"time"

	"github.com/spf13/cobra"
)

var sonosCmd = &cobra.Command{
	Use:   "sonos",
	Short: "sonos",
	RunE: func(cmd *cobra.Command, args []string) error {
		go Serve()
		time.Sleep(2 * time.Second)

		sonosIpAddr, err := Discover()
		if err != nil {
			return err
		}

		sonos := NewSonos(sonosIpAddr)
		res, err := sonos.makeSetUriRequest()
		if err != nil {
			return err
		}
		fmt.Printf("res: %+v\n", res)
		res, err = sonos.makePlayRequest()
		if err != nil {
			return err
		}
		fmt.Printf("res: %+v\n", res)

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
