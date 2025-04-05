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

		for i, realfeeditem := range realfeed.Items {
			prompt := fmt.Sprintf(`タイトルを読み上げ用に50文字程度で調整してください。基本はそのままで大丈夫です。もし頭に入りにくそうであれば、話し言葉に調整してください。出力された文章を読み上げます。要約ではなく「話し言葉に調整する」という視点が重要。
タイトル:
%s

詳細:
%s
`, realfeeditem.Title, realfeeditem.Description)

			fmt.Println(prompt)

			chatres, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo,
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: prompt,
					},
				},
			})
			if err != nil {
				return err
			}
			text := chatres.Choices[0].Message.Content
			fmt.Println(text)
	
			request := openai.CreateSpeechRequest{
				Model:          openai.TTSModel1,
				Input:          text,
				Voice:          openai.VoiceOnyx,
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
			break


			if i > 5 {
				break
			}
		}

		return nil
	},
}
