package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// /start command — sends inline keyboard with 3 options
	bot.Command("start", func(ctx *tg.Context) error {
		markup := keyboard.Inline().
			Row().
			Text("Option 1", "1").
			Text("Option 2", "2").
			Row().
			Text("Option 3", "3")

		return ctx.Reply("Please choose:").Markup(markup).Send().Err()
	})

	// /help command
	bot.Command("help", func(ctx *tg.Context) error {
		return ctx.Reply("Use /start to test this bot.").Send().Err()
	})

	// Handles any callback data
	bot.On.Callback.Any(func(ctx *tg.Context) error {
		data := ctx.Update.CallbackQuery.Data
		return ctx.EditText(Format("Selected option: <b>{}</b>", data)).
			HTML().
			Preview(preview.New().URL("https://www.enetx.surf").Above()).
			Send().Err()
	})

	bot.Polling().Start()
}
