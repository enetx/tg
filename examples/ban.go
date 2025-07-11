package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token).Build().Unwrap()

	bot.Command("ban", func(ctx *tg.Context) error {
		if admin := ctx.IsAdmin(); admin.IsErr() || !admin.Ok() {
			return nil
		}

		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to the user you want to ban.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id
		ctx.Ban(userID, tg.NewBanChatMemberOpts().For(1*time.Hour))

		return ctx.Reply("User has been banned for 1 hour.").Send().Err()
	})

	bot.Polling().Start()
}
