package aictl

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/enuesaa/speakit/pkg/repository"
	"github.com/enuesaa/speakit/pkg/service"
	"github.com/sashabaranov/go-openai"
)

var Data map[string][]byte

func Run(openaiApiKey string, rssfeed string) error {
	Data = make(map[string][]byte)

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
			Model:          openai.TTSModelGPT4oMini,
			Input:          text,
			Voice:          openai.VoiceAsh,
			Speed:          1.7,
			Instructions:   "抑揚をつけて。めちゃくちゃ早口で",
			ResponseFormat: openai.SpeechResponseFormatMp3,
		}
		res, err := client.CreateSpeech(context.Background(), request)
		if err != nil {
			return err
		}
		buf, err := io.ReadAll(res)
		if err != nil {
			return err
		}
		filename := fmt.Sprintf("%d.mp3", i)
		Data[filename] = buf

		a, err := os.Create(filename)
		if err != nil {
			break
		}
		a.Write(buf)

		if i > 2 {
			break
		}
	}

	return nil
}
