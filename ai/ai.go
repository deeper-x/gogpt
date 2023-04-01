package ai

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

// Engine is the actor object
type Engine struct {
	Client openai.Client
	ctx    context.Context
}

// NewEngine client builder
func NewEngine(authToken string) Engine {
	c := openai.NewClient(authToken)

	return Engine{
		Client: *c,
		ctx:    context.Background(),
	}
}

// Query ask a question to the engine
func (e Engine) Query(m string) (string, error) {
	resp, err := e.Client.CreateChatCompletion(
		e.ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: m,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	output := resp.Choices[0].Message.Content
	return output, nil
}
