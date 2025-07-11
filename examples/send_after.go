package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	// Read the bot token from the .env file
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	// Register a command handler for /start
	bot.Command("start", func(ctx *tg.Context) error {
		// Send an immediate message so Telegram considers the update as "handled"
		ctx.Message("Preparing self-destruct...").Send()

		// Schedule a second message to be sent after 3 seconds,
		// and automatically delete it 5 seconds after it is sent
		ctx.Message("This message will self-destruct in 5 seconds.").
			After(time.Second * 3).       // Delay sending by 3 seconds
			DeleteAfter(time.Second * 5). // Delete 5 seconds after it is sent
			Send()

		// Delete the original /start message (from the user)
		// This should be done after responding to avoid Telegram resending the update
		return ctx.Delete().Send().Err()
	})

	// Start polling for updates and drop any pending ones from before startup
	bot.Polling().DropPendingUpdates().Start()
}
