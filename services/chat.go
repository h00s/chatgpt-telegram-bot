package services

import (
	"time"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type Chat struct {
	Messages      []gogpt.ChatCompletionMessage
	LastMessageAt time.Time
}

func (c Chat) AddMessage(message string) Chat {
	c.LastMessageAt = time.Now()
	c.Messages = append(c.Messages, gogpt.ChatCompletionMessage{
		Role:    "user",
		Content: message,
	})
	return c
}

type Chats struct {
	Chats map[int64]Chat
}

func NewChats() *Chats {
	return &Chats{
		Chats: map[int64]Chat{},
	}
}

func (c *Chats) ResetChat(user int64) {
	c.Chats[user] = Chat{
		Messages:      []gogpt.ChatCompletionMessage{},
		LastMessageAt: time.Now(),
	}
}

func (c *Chats) AddMessage(user int64, message string) {
	if _, ok := c.Chats[user]; !ok {
		c.ResetChat(user)
	}
	if (time.Since(c.Chats[user].LastMessageAt) / time.Minute) > 5 {
		c.ResetChat(user)
	}
	c.Chats[user] = c.Chats[user].AddMessage(message)
}
