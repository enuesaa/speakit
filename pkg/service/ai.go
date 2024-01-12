package service

import (
	"context"
	"fmt"

	"github.com/enuesaa/speakit/pkg/repository"
	openai "github.com/sashabaranov/go-openai"
)

type AiService struct {
	repos repository.Repos
}

func NewAiService(repos repository.Repos) *AiService {
	return &AiService{
		repos: repos,
	}
}

func (srv *AiService) Call(token string, message string) (string, error) {
	// see https://github.com/sashabaranov/go-openai
	client := openai.NewClient(token)
	res, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return res.Choices[0].Message.Content, nil
}

func (srv *AiService) Speak(token string, path string) (string, error) {
	client := openai.NewClient(token)
	ctx := context.Background()

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: path,
	}
	res, err := client.CreateTranscription(ctx, req)
	if err != nil {
		return "", err
	}
	fmt.Printf("res: %+v\n", res)

	return res.Text, nil
}
