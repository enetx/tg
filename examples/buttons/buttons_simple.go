package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := g.NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
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

	b.On.Callback.Equal("cb_1", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("clicked the callback1 button").Send().Err()
	})

	b.On.Callback.Equal("cb_2", func(ctx *ctx.Context) error {
		return ctx.AnswerCallbackQuery("clicked the callback2 button").Alert().Send().Err()
	},
	)

	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.Reply("Echo: " + g.String(ctx.EffectiveMessage.Text)).Send().Err()
	})

	b.Polling().Start()
}
