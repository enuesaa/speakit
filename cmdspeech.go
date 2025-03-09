package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var rssfeed string

func init() {
	speechCmd.Flags().StringVar(&openaiApiKey, "openai", "", "OpenAI API Key")
	speechCmd.Flags().StringVar(&rssfeed, "rss", "", "")
}

var speechCmd = &cobra.Command{
	Use:   "speech",
	Short: "speech",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := openai.NewClient(openaiApiKey)

		feedsSrv := service.NewFeedSevice(repository.Repos{})
		realfeed, err := feedsSrv.TryFetch(rssfeed)
		if err != nil {
			return err
		}

		var texts []string

		for i, realfeeditem := range realfeed.Items {
			prompt := fmt.Sprintf(
				"tssで読み上げるので要約して。話し言葉で。ニュースのナレーターのように。全体で50文字程度。タイトルはあんまり変えないで。要約というよりポイントをまとめて: \nタイトル: %s\n概要: %s",
				realfeeditem.Title,
				realfeeditem.Description,
			)
			chatres, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
					Model: openai.GPT3Dot5Turbo,
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    openai.ChatMessageRoleUser,
							Content: prompt,
						},
					},
				},
			)
			if err != nil {
				return err
			}
			
			text := chatres.Choices[0].Message.Content
			fmt.Println(text)
			texts = append(texts, text)

			if i > 5 {
				break
			}	
		}

		text := ""
		for _, t := range texts {
			text += t
		}
		request := openai.CreateSpeechRequest{
			Model: openai.TTSModel1,
			Input: text,
			Voice: openai.VoiceShimmer,
			Speed: 1.3,
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
