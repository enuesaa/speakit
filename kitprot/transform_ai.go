package kitprot

// import (
// 	"bytes"
// 	"context"
// 	"text/template"

// 	"github.com/sashabaranov/go-openai"
// )

// type AITransformer struct {
// 	OpenAIKey  string
// 	PromptTmpl string // include {text}

// 	log    *LogBehavior
// 	client *openai.Client
// }

// func (g *AITransformer) Inject(log *LogBehavior) {
// 	g.log = log
// 	g.client = openai.NewClient(g.OpenAIKey)
// }

// func (g *AITransformer) StartUp() error {
// 	return nil
// }

// func (g *AITransformer) Transform(record *Record) error {
// 	ctx := context.Background()

// 	tmpl, err := template.New("prompt").Parse(g.PromptTmpl)
// 	if err != nil {
// 		return err
// 	}
// 	tmplvars := map[string]any{
// 		"text": record.Segments[0],
// 		"meta": record.Meta,
// 	}

// 	var buf bytes.Buffer
// 	if err = tmpl.Execute(&buf, tmplvars); err != nil {
// 		return err
// 	}
// 	prompt := buf.String()

// 	res, err := g.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
// 		Model: openai.GPT3Dot5Turbo,
// 		Messages: []openai.ChatCompletionMessage{
// 			{
// 				Role:    openai.ChatMessageRoleUser,
// 				Content: prompt,
// 			},
// 		},
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	g.log.Info("prompt: %s", prompt)
// 	record.Text = res.Choices[0].Message.Content
// 	g.log.Info("ai: %s", record.Text)

// 	return nil
// }

// func (g *AITransformer) Close() error {
// 	return nil
// }
