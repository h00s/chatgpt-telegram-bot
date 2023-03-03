package services

import (
	"context"
	"fmt"

	"github.com/h00s-go/h00s-bot/config"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatGPT struct {
	Client   *gogpt.Client
	Messages []gogpt.ChatCompletionMessage
}

func NewChatGPT(c *config.ChatGPT) *ChatGPT {
	return &ChatGPT{
		Client: gogpt.NewClient(c.APIKey),
	}
}

func (c *ChatGPT) Reset() {
	c.Messages = []gogpt.ChatCompletionMessage{}
}

func (c *ChatGPT) Chat(message string) (string, error) {
	c.Messages = append(c.Messages, gogpt.ChatCompletionMessage{
		Role:    "user",
		Content: message,
	})
	resp, err := c.Client.CreateChatCompletion(context.Background(), gogpt.ChatCompletionRequest{
		Model:    gogpt.GPT3Dot5Turbo,
		Messages: c.Messages,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
