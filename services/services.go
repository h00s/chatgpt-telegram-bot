package services

import "github.com/h00s-go/h00s-bot/config"

type Services struct {
	ChatGPT *ChatGPT
}

func NewServices(c *config.Config) *Services {
	return &Services{
		ChatGPT: NewChatGPT(&c.ChatGPT),
	}
}
