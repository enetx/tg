package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.Command("start", func(ctx *tg.Context) error {
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

	bot.On.Message.Contact(func(ctx *tg.Context) error {
		message := Format("Phone: {1.PhoneNumber}\nName: {1.FirstName}\nID: {1.UserId}", ctx.EffectiveMessage.Contact)
		return ctx.Message(message).Send().Err()
	})

	bot.Polling().Start()
}
