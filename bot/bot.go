package bot

import (
	"time"

	"github.com/h00s/chatgpt-telegram-bot/config"
	"github.com/h00s/chatgpt-telegram-bot/services"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Client      *tele.Bot
	Services    *services.Services
	sendOptions *tele.SendOptions
}

func NewBot(c *config.Telegram, s *services.Services) (*Bot, error) {
	b, err := tele.NewBot(
		tele.Settings{
			Token:  c.Token,
			Poller: &tele.LongPoller{Timeout: 10 * time.Second},
		},
	)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Client:   b,
		Services: s,
		sendOptions: &tele.SendOptions{
			ParseMode: tele.ModeMarkdown,
		},
	}, nil
}

func (b *Bot) Start() {
	b.SetHandlers()
	b.Client.Start()
}
