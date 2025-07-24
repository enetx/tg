package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Delete multiple messages by IDs
	b.Command("deletemessages", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletemessages <message_id1> <message_id2> ...").Send().Err()
		}

		var messageIDs Slice[int64]

		args.Iter().ForEach(func(arg String) {
			messageIDs.Push(arg.ToInt().Unwrap().Int64())
		})

		return ctx.DeleteMessages().MessageIDs(messageIDs).Send().Err()
	})

	// Delete recent messages (last N messages)
	b.Command("deleterecent", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deleterecent <count>").Send().Err()
		}

		count := args[0].ToInt().UnwrapOr(100).Std()

		// Build list of recent message IDs (this is a simplified example)
		currentMsgID := ctx.EffectiveMessage.MessageId
		deleteBuilder := ctx.DeleteMessages()

		for i := 1; i <= count; i++ {
			deleteBuilder.AddMessages(currentMsgID - int64(i))
		}

		return deleteBuilder.Send().Err()
	})

	// Scheduled deletion
	b.Command("deletescheduled", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /deletescheduled <seconds> <message_id1> <message_id2> ...").Send().Err()
		}

		seconds := args[0].ToInt().Unwrap()

		var messageIDs Slice[int64]

		args.Iter().Skip(1).ForEach(func(arg String) {
			messageIDs.Push(arg.ToInt().Unwrap().Int64())
		})

		ctx.DeleteMessages().
			MessageIDs(messageIDs).
			After(time.Duration(seconds) * time.Second).
			Send()

		return ctx.Reply(String("Messages scheduled for deletion in " + seconds.String() + " seconds")).
			Send().
			Err()
	})

	b.Polling().AllowedUpdates().Start()
}
