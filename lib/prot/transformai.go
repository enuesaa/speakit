package prot

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

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

	tmpl, err := template.New("prompt").Parse(g.PromptTmpl)
	if err != nil {
		return err
	}
	tmplvars := map[string]any{
		"text": record.Text,
		"meta": record.Meta,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, tmplvars); err != nil {
		return err
	}
	prompt := buf.String()

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
