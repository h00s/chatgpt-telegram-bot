package bot

import (
	"fmt"
	"log"

	tele "gopkg.in/telebot.v3"
)

func (b *Bot) SetHandlers() {
	b.Client.Handle("/hello", b.Hello)
	b.Client.Handle("/new", b.NewChat)
	b.Client.Handle(tele.OnText, b.All)
}

func (b *Bot) LogMessage(c tele.Context) {
	log.Printf("Message from %s (%v): %s", c.Message().Sender.Username, c.Message().Sender.ID, c.Message().Text)
}

func (b *Bot) Hello(c tele.Context) error {
	b.LogMessage(c)
	return c.Send("Hello!")
}

func (b *Bot) NewChat(c tele.Context) error {
	b.LogMessage(c)
	b.Services.ChatGPT.Chats.ResetChat(c.Message().Sender.ID)
	return c.Send("Ok, new chat started!")
}

func (b *Bot) All(c tele.Context) error {
	c.Notify(tele.Typing)
	b.LogMessage(c)
	response, err := b.Services.ChatGPT.Chat(c.Message().Sender.ID, c.Message().Text)
	if err != nil {
		fmt.Println(err)
		c.Send("Something went wrong...")
	}

	return c.Send(response, b.sendOptions)
}
