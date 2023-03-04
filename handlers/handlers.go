package handlers

import (
	"fmt"
	"log"

	"github.com/h00s-go/h00s-bot/services"
	tele "gopkg.in/telebot.v3"
)

type Handlers struct {
	Services    *services.Services
	SendOptions *tele.SendOptions
}

func NewHandlers(s *services.Services) *Handlers {
	return &Handlers{
		Services: s,
		SendOptions: &tele.SendOptions{
			ParseMode: tele.ModeMarkdown,
		},
	}
}

func (h *Handlers) LogMessage(c tele.Context) {
	log.Printf("Message from %s (%v): %s", c.Message().Sender.Username, c.Message().Sender.ID, c.Message().Text)
}

func (h *Handlers) Hello(c tele.Context) error {
	h.LogMessage(c)
	return c.Send("Hello!")
}

func (h *Handlers) NewChat(c tele.Context) error {
	h.LogMessage(c)
	h.Services.ChatGPT.Reset(c.Message().Sender.ID)
	return c.Send("Ok, new chat started!")
}

func (h *Handlers) All(c tele.Context) error {
	c.Notify(tele.Typing)
	h.LogMessage(c)
	response, err := h.Services.ChatGPT.Chat(c.Message().Sender.ID, c.Message().Text)
	if err != nil {
		fmt.Println(err)
		c.Send("Something went wrong...")
	}

	return c.Send(response, h.SendOptions)
}
