package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	// Read the bot token from the .env file
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		// Self-destruct message
		ctx.Message("This message will self-destruct in 5 seconds.").
			DeleteAfter(5 * time.Second).
			Send()

		// Delete original /start message
		return ctx.Delete().Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
