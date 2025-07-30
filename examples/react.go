package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Handle /react command to demonstrate message reactions
	b.Command("react", func(ctx *ctx.Context) error {
		if ctx.EffectiveMessage.ReplyToMessage == nil {
			return ctx.Reply("Reply to a message to react to it").Send().Err()
		}

		result := ctx.SetMessageReaction(ctx.EffectiveMessage.ReplyToMessage.MessageId).
			Reaction("❤️").
			Big().
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to set reaction: {}", result.Err())).Send().Err()
		}

		return ctx.Reply("Reactions set successfully!").Send().Err()
	})

	b.Polling().Start()
}
