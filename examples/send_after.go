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

	// Register a command handler for /start
	b.Command("start", func(ctx *ctx.Context) error {
		// Send an immediate message so Telegram considers the update as "handled"
		ctx.SendMessage("Preparing self-destruct...").Send()

		// Schedule a second message to be sent after 3 seconds,
		// and automatically delete it 5 seconds after it is sent
		ctx.SendMessage("This message will self-destruct in 5 seconds.").
			After(3 * time.Second).       // Delay sending by 3 seconds
			DeleteAfter(5 * time.Second). // Delete 5 seconds after it is sent
			Send()

		// Delete the original /start message (from the user)
		// This should be done after responding to avoid Telegram resending the update
		return ctx.DeleteMessage().Send().Err()
	})

	// Start polling for updates and drop any pending ones from before startup
	b.Polling().DropPendingUpdates().Start()
}
