package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("ban", func(ctx *ctx.Context) error {
		if admin := ctx.IsAdmin(); admin.IsErr() || !admin.Ok() {
			return nil
		}

		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the user you want to ban.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		ctx.BanChatMember(userID).For(time.Hour).Send()

		return ctx.Reply("User has been banned for 1 hour.").Send().Err()
	})

	b.Polling().Start()
}
