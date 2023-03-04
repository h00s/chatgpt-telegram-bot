package handlers

import (
	"fmt"
	"log"

	"github.com/h00s-go/h00s-bot/services"
	tele "gopkg.in/telebot.v3"
)

type Handlers struct {
	Services *services.Services
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		Services: s,
	}
}

func (h *Handlers) LogMessage(c tele.Context) {
	log.Println("Message from", c.Message().Sender.Username, ":", c.Message().Text)
}

func (h *Handlers) Hello(c tele.Context) error {
	h.LogMessage(c)
	return c.Send("Hello!")
}

func (h *Handlers) NewChat(c tele.Context) error {
	h.LogMessage(c)
	h.Services.ChatGPT.Reset(c.Message().Sender.Username)
	return c.Send("Ok, new chat started!")
}

func (h *Handlers) All(c tele.Context) error {
	h.LogMessage(c)
	response, err := h.Services.ChatGPT.Chat(c.Message().Sender.Username, c.Message().Text)
	if err != nil {
		fmt.Println(err)
		c.Send("Something went wrong...")
	}

	return c.Send(response)
}
