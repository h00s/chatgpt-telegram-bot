package main

import (
	"log"
	"time"

	"github.com/h00s-go/h00s-bot/config"
	"github.com/h00s-go/h00s-bot/handlers"
	tele "gopkg.in/telebot.v3"
)

func main() {
	config := config.NewConfig()

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

	b.Handle("cujes", handlers.Hello)

	b.Start()
}
