package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/spf13/cobra"
)

var speakCmd = &cobra.Command{
	Use:   "speak",
	Short: "speak <filename>",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("please specify filename to speak")
		}

		filename := args[0]

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		streamer, format, err := wav.Decode(f)
		if err != nil {
			return err
		}
		defer streamer.Close()
	
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			return err
		}

	    done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done

		return nil
	},
}
