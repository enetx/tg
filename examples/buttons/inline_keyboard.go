package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /start command — sends inline keyboard with 3 options
	b.Command("start", func(ctx *ctx.Context) error {
		markup := keyboard.Inline().
			Row().
			Text("Option 1", "1").
			Text("Option 2", "2").
			Row().
			Text("Option 3", "3")

		return ctx.Reply("Please choose:").Markup(markup).Send().Err()
	})

	// /help command
	b.Command("help", func(ctx *ctx.Context) error {
		return ctx.Reply("Use /start to test this bot.").Send().Err()
	})

	// Handles any callback data
	b.On.Callback.Any(func(ctx *ctx.Context) error {
		data := ctx.Update.CallbackQuery.Data
		return ctx.EditText(Format("Selected option: <b>{}</b>", data)).
			HTML().
			Preview(preview.New().URL("https://enetx.surf").Above()).
			Send().Err()
	})

	b.Polling().Start()
}
