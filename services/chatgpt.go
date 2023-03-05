package services

import (
	"context"
	"fmt"

	"github.com/h00s-go/h00s-bot/config"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type ChatGPT struct {
	Client   *gogpt.Client
	Messages map[int64][]gogpt.ChatCompletionMessage
}

func NewChatGPT(c *config.OpenAI) *ChatGPT {
	return &ChatGPT{
		Client:   gogpt.NewClient(c.APIKey),
		Messages: map[int64][]gogpt.ChatCompletionMessage{},
	}
}

func (c *ChatGPT) Reset(user int64) {
	c.Messages[user] = []gogpt.ChatCompletionMessage{}
}

func (c *ChatGPT) Chat(user int64, message string) (string, error) {
	if _, ok := c.Messages[user]; !ok {
		c.Reset(user)
	}
	c.Messages[user] = append(c.Messages[user], gogpt.ChatCompletionMessage{
		Role:    "user",
		Content: message,
	})
	resp, err := c.Client.CreateChatCompletion(context.Background(), gogpt.ChatCompletionRequest{
		Model:    gogpt.GPT3Dot5Turbo,
		Messages: c.Messages[user],
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
