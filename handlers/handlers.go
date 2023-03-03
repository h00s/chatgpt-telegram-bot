package handlers

import (
	tele "gopkg.in/telebot.v3"
)

func Hello(c tele.Context) error {
	return c.Send("Cujem!")
}

func All(c tele.Context) error {
	m := c.Message()

	return c.Send("You said: " + m.Text)
}
