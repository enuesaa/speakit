package prot

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)


type AITransformer struct {
	OpenAIKey  string
	PromptTmpl string // include {text}

	client *openai.Client
}

func (g *AITransformer) StartUp() error {
	g.client = openai.NewClient(g.OpenAIKey)

	return nil
}

func (g *AITransformer) Transform(record *Record) error {
	ctx := context.Background()

	prompt := strings.ReplaceAll(g.PromptTmpl, "{text}", record.Text)

	res, err := g.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
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
	fmt.Printf("before: %s\n", prompt)
	record.Text = res.Choices[0].Message.Content
	fmt.Printf("ai: %s\n", record.Text)

	return nil
}

func (g *AITransformer) Close() error {
	return nil
}
