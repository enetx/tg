package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	bot.Command("start", func(ctx *tg.Context) error {
		markup := keyboard.Inline().
			Row().
			Text("Callback1", "cb_1").
			Text("Callback2", "cb_2").
			Row().
			URL("URL", "https://example.com").
			WebApp("WebApp", "https://web.example").
			// LoginURL("Login", "https://login.example").
			Row().
			CopyText("Copy", "Copied!").
			// Pay("Buy").
			Row().
			// Game("Game").
			SwitchInlineQuery("Inline", "query").
			SwitchInlineQueryCurrentChat("Inline Here", "query")

		return ctx.Reply("Choose a button:").Markup(markup).Send().Err()
	})

	bot.On.Callback.Equal("cb_1", func(ctx *tg.Context) error {
		return ctx.Answer("clicked the callback1 button").Send().Err()
	})

	bot.On.Callback.Equal("cb_2", func(ctx *tg.Context) error {
		return ctx.Answer("clicked the callback2 button").Alert().Send().Err()
	},
	)

	bot.On.Message.Text(func(ctx *tg.Context) error {
		return ctx.Reply("Echo: " + String(ctx.EffectiveMessage.Text)).Send().Err()
	})

	bot.Polling().Start()
}
