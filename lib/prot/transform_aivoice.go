package prot

import (
	"context"
	"io"

	"github.com/sashabaranov/go-openai"
)

type AIVoiceTransformer struct {
	OpenAIKey string

	logger Logger
	client *openai.Client
}

func (g *AIVoiceTransformer) StartUp(logger Logger) error {
	g.logger = logger
	g.client = openai.NewClient(g.OpenAIKey)

	return nil
}

func (g *AIVoiceTransformer) Transform(record *Record) error {
	ctx := context.Background()

	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModelGPT4oMini,
		Input:          record.Text,
		Voice:          openai.VoiceAsh,
		Speed:          1.7,
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
