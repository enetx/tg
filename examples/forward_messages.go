package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Forward multiple messages from another chat
	b.Command("forwardmessages", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /forwardmessages <from_chat_id> <message_id1> <message_id2> ...").Send().Err()
		}

		fromChatID := args[0].ToInt().Unwrap().Int64()

		var messageIDs Slice[int64]

		args.Iter().Skip(1).ForEach(func(arg String) {
			messageIDs.Push(arg.ToInt().Unwrap().Int64())
		})

		result := ctx.ForwardMessages().
			From(fromChatID).
			MessageIDs(messageIDs).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Successfully forwarded " + result.Ok().Len().String() + " messages").
			Send().
			Err()
	})

	b.Polling().AllowedUpdates().Start()
}
