package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	// /start command with inline button
	bot.Command("start", func(ctx *tg.Context) error {
		ctx.Reply(Format("Hello, I'm @{.Raw.Username}. I <b>repeat</b> all your messages.", bot)).
			HTML().
			Markup(keyboard.Inline().Text("Press me", "start_callback")).
			Send()

		// Delete original /start message
		return ctx.Delete().Send().Err()
	})

	// Callback query handler
	bot.On.Callback.Equal("start_callback", func(ctx *tg.Context) error {
		ctx.Answer("You pressed a button!").Alert().Send()
		return ctx.EditText("You edited the start message.").Send().Err()
	})

	// Start polling
	bot.Polling().DropPendingUpdates().Start()
}
