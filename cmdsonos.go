package main

import (
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
		if err := sonos.subscribeSonos(); err != nil {
			return err
		}

		time.Sleep(100 * time.Second)

		return nil
	},
}
