package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func Hello(c tele.Context) error {
	return c.Send("Hello!")
}
