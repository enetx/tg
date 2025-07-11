package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.On.Message.Dice(func(ctx *tg.Context) error {
		return ctx.Dice().Send().Err()
	})

	bot.On.Message.Text(func(ctx *tg.Context) error {
		return ctx.Dice().Slot().Send().Err()
	})

	bot.Polling().AllowedUpdates().DropPendingUpdates().Start()
}
