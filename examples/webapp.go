package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	bot.Command("start", func(ctx *tg.Context) error {
		button := keyboard.Reply().WebApp("peet", "https://tls.peet.ws/api/all")
		return ctx.Message("оk").Markup(button).Send().Err()
	})

	bot.Polling().Start()
}
