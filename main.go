package main

import (
	"log"
	"time"

	"github.com/h00s-go/h00s-bot/config"
	"github.com/h00s-go/h00s-bot/handlers"
	"github.com/h00s-go/h00s-bot/services"
	tele "gopkg.in/telebot.v3"
)

func main() {
	config := config.NewConfig()
	services := services.NewServices(config)
	handlers := handlers.NewHandlers(services)

	b, err := tele.NewBot(
		tele.Settings{
			Token:  config.Telegram.Token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		},
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("hello", handlers.Hello)
	b.Handle("new chat", handlers.NewChat)
	b.Handle(tele.OnText, handlers.All)

	b.Start()
}
