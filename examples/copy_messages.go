package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Copy multiple messages from another chat
	b.Command("copymessages", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /copymessages <from_chat_id> <message_id1> <message_id2> ...").Send().Err()
		}

		fromChatID := args[0].ToInt().Unwrap().Int64()

		var messageIDs g.Slice[int64]

		args.Iter().Skip(1).ForEach(func(arg g.String) {
			messageIDs.Push(arg.ToInt().Unwrap().Int64())
		})

		result := ctx.CopyMessages().
			From(fromChatID).
			MessageIDs(messageIDs).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Successfully copied " + result.Ok().Len().String() + " messages").
			Send().
			Err()
	})

	b.Polling().AllowedUpdates().Start()
}
