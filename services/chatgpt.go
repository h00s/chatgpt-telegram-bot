package services

import (
	"context"
	"fmt"

	"github.com/h00s-go/h00s-bot/config"
	gogpt "github.com/sashabaranov/go-openai"
)

type ChatGPT struct {
	Client *gogpt.Client
	Chats  *Chats
}

func NewChatGPT(c *config.OpenAI) *ChatGPT {
	return &ChatGPT{
		Client: gogpt.NewClient(c.APIKey),
		Chats:  NewChats(),
	}
}

func (c *ChatGPT) Chat(user int64, message string) (string, error) {
	c.Chats.AddMessage(user, message)
	resp, err := c.Client.CreateChatCompletion(context.Background(), gogpt.ChatCompletionRequest{
		Model:    gogpt.GPT3Dot5Turbo0301,
		Messages: c.Chats.Chats[user].Messages,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
