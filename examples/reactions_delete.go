package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

// Removing reactions (Bot API 10.0): admins with the can_delete_messages right can
// remove a single reaction from a message or wipe all recent reactions of a user/chat.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /unreact — reply to a message to remove the reaction added by the replied user.
	b.Command("unreact", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to a message to remove its reaction.").Send().Err()
		}

		reply := ctx.EffectiveMessage.ReplyToMessage

		return ctx.DeleteMessageReaction(reply.MessageId).
			UserID(reply.From.Id).
			Send().Err()
	})

	// /clearreacts — remove up to 10000 recent reactions added by the replied user.
	b.Command("clearreacts", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil || ctx.EffectiveMessage.ReplyToMessage.From == nil {
			return ctx.Reply("Reply to a user's message to clear their reactions.").Send().Err()
		}

		userID := ctx.EffectiveMessage.ReplyToMessage.From.Id

		result := ctx.DeleteAllMessageReactions().
			UserID(userID).
			Send()
		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to clear reactions: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("All recent reactions from this user were removed.").Send().Err()
	})

	b.Polling().Start()
}
