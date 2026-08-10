package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/g/fs"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	// Read the bot token from the .env file
	token := g.Slice[g.String](fs.NewFile("../.env").Read().Ok().Trim().Split("=")).Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		// Self-destruct message
		ctx.SendMessage("This message will self-destruct in 5 seconds.").
			DeleteAfter(5 * time.Second).
			Send()

		// Delete original /start message
		return ctx.DeleteMessage().Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
