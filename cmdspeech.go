package main

import (
	"context"
	"encoding/json"
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

type SpeechItem struct {
	Site        string `json:"site"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type SpeechItems struct {
	News []SpeechItem `json:"news"`
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

		prompt := `読み上げ原稿を作って下さい。

下記はニュース記事の一覧です
これをニュースキャスターが読み上げるふうに文章を構成して下さい。
この文章を TTS で読み上げるので、自然と読み上げられるよう接続詞などに気を配って下さい。
口語調にして下さい。

フォーマット:
では、ニュースをお伝えします。xxで、、続いて xx で、、最後に xx で。ニュースをお伝えしました

データ:
`

		items := SpeechItems{
			News: make([]SpeechItem, 0),
		}

		for i, realfeeditem := range realfeed.Items {
			items.News = append(items.News, SpeechItem{
				Site:  realfeeditem.Link,
				Title: realfeeditem.Title,
			})
			if i > 5 {
				break
			}
		}
		itemsjson, err := json.Marshal(items)
		if err != nil {
			return err
		}
		prompt += "\n"
		prompt += string(itemsjson)
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
			Model:          openai.TTSModel1,
			Input:          text,
			Voice:          openai.VoiceShimmer,
			Speed:          1.3,
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
