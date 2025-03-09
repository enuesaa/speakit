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

		prompt := `TTSで読み上げ用のニュース原稿を作ってください。
出力された原稿をそのまま読み上げます。
「ニュース番号」はあなたのために振ってます。そのため番号を出力しないでください。

ニュース番組のナレーションのように、自然な話し言葉にしてください。
1つのニュースにつき50文字程度でポイントをまとめてください。
要約ではなく、読み上げやすい原稿にしてください。
タイトルは変えず、内容の要点をまとめてください。
読み終わったら、次のニュースに続けてください。
`

		for i, realfeeditem := range realfeed.Items {
			prompt += fmt.Sprintf(
				"\nニュース%d:%s\n%s\n\n",
				i,
				realfeeditem.Title,
				realfeeditem.Description,
			)
			if i > 5 {
				break
			}	
		}
		fmt.Println(prompt)

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
		fmt.Printf("\n\n")
		fmt.Println(text)

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
