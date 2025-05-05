package prot

import (
	"context"
	"io"

	"github.com/sashabaranov/go-openai"
)

type AIVoiceTransformer struct {
	OpenAIKey string

	log    *LogBehavior
	client *openai.Client
}

func (g *AIVoiceTransformer) Inject(log *LogBehavior) {
	g.log = log
}

func (g *AIVoiceTransformer) StartUp() error {
	g.client = openai.NewClient(g.OpenAIKey)
	return nil
}

func (g *AIVoiceTransformer) Transform(record *Record) error {
	ctx := context.Background()

	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModelGPT4oMini,
		Input:          record.Text,
		Voice:          openai.VoiceAsh,
		Speed:          1.3,
		Instructions:   "穏やかに。ニュースのキャスターのように。抑揚をつけて。めちゃくちゃ早口で",
		ResponseFormat: openai.SpeechResponseFormatMp3,
	}
	res, err := g.client.CreateSpeech(ctx, request)
	if err != nil {
		return err
	}

	buf, err := io.ReadAll(res)
	if err != nil {
		return err
	}
	record.Voice = buf

	return nil
}

func (g *AIVoiceTransformer) Close() error {
	return nil
}
