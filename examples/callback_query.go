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

	// /start command with inline button
	b.Command("start", func(ctx *ctx.Context) error {
		ctx.Reply(Format("Hello, I'm @{.Raw.Username}. I <b>repeat</b> all your messages.", b)).
			HTML().
			Markup(keyboard.Inline().Text("Press me", "start_callback")).
			Send()

		// Delete original /start message
		return ctx.DeleteMessage().Send().Err()
	})

	// Callback query handler
	b.On.Callback.Equal("start_callback", func(ctx *ctx.Context) error {
		ctx.AnswerCallbackQuery("You pressed a button!").Alert().Send()
		return ctx.EditMessageText("You edited the start message.").Send().Err()
	})

	// Start polling
	b.Polling().DropPendingUpdates().Start()
}
