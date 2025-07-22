package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		markup := keyboard.Reply().
			Row().
			Text("Hello").
			Contact("Share Phone").
			Row().
			Location("Send Location").
			WebApp("Open WebApp", "https://example.com").
			Row().
			Poll("Create Poll")

		return ctx.Reply("Choose a button:").Markup(markup).Send().Err()
	})

	b.On.Message.Contact(func(ctx *ctx.Context) error {
		message := Format("Phone: {1.PhoneNumber}\nName: {1.FirstName}\nID: {1.UserId}", ctx.EffectiveMessage.Contact)
		return ctx.SendMessage(message).Send().Err()
	})

	b.Polling().Start()
}
