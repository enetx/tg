package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.On.Message.Dice(func(ctx *ctx.Context) error {
		return ctx.Dice().Send().Err()
	})

	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.Dice().Slot().Send().Err()
	})

	b.Polling().AllowedUpdates().DropPendingUpdates().Start()
}
