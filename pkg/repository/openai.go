package repository

import (
	"context"
	"io"

	"github.com/sashabaranov/go-openai"
)

type OpenAIRepositoryInterface interface {
	Speech(text string) (io.Reader, error)
}

type OpenAIRepository struct {
	APIKey string
}

func (repo *OpenAIRepository) Speech(text string) (io.Reader, error) {
	client := openai.NewClient(repo.APIKey)

	request := openai.CreateSpeechRequest{
		Model:          openai.TTSModel1,
		Input:          text,
		Voice:          openai.VoiceAlloy,
		Speed:          1.7,
		ResponseFormat: openai.SpeechResponseFormatMp3,
	}
	return client.CreateSpeech(context.Background(), request)
}
