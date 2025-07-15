package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		button := keyboard.Reply().WebApp("peet", "https://tls.peet.ws/api/all")
		return ctx.Message("оk").Markup(button).Send().Err()
	})

	b.Polling().Start()
}
