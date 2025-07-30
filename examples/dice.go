package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.On.Message.Dice(func(ctx *ctx.Context) error {
		return ctx.SendDice().Send().Err()
	})

	b.On.Message.Text(func(ctx *ctx.Context) error {
		return ctx.SendDice().Slot().Send().Err()
	})

	b.Polling().AllowedUpdates().DropPendingUpdates().Start()
}
