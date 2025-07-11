package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	// Read the bot token from the .env file
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.Command("start", func(ctx *tg.Context) error {
		// Self-destruct message
		ctx.Message("This message will self-destruct in 5 seconds.").
			DeleteAfter(5 * time.Second).
			Send()

		// Delete original /start message
		return ctx.Delete().Send().Err()
	})

	bot.Polling().DropPendingUpdates().Start()
}
