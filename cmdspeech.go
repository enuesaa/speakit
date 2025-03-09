package main

import (
	"context"
	"io"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

func init() {
	speechCmd.Flags().StringVar(&openaiApiKey, "openai", "", "OpenAI API Key")
}

var speechCmd = &cobra.Command{
	Use:   "speech",
	Short: "speech",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := openai.NewClient(openaiApiKey)

		request := openai.CreateSpeechRequest{
			Model: openai.TTSModel1,
			Input: "",
			Voice: openai.VoiceEcho,
			Speed: 1.7,
			ResponseFormat: openai.SpeechResponseFormatMp3,
		}
		res, err := client.CreateSpeech(context.Background(), request)
		if err != nil {
			return err
		}
		f, err := os.Create("a.mp3")
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := io.Copy(f, res); err != nil {
			return err
		}
		return nil
	},
}
