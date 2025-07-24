package main

import (
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Create a new forum topic
	b.Command("createtopic", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /createtopic <topic_name>").Send().Err()
		}

		topicName := args[0]

		return ctx.CreateForumTopic(topicName).
			IconColor(0x6FB9F0). // Blue color
			Send().Err()
	})

	// Edit forum topic
	b.Command("edittopic", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /edittopic <thread_id> <new_name>").Send().Err()
		}

		threadID := args[0].ToInt().Unwrap().Int64()
		newName := args[1]

		return ctx.EditForumTopic(threadID).
			Name(newName).
			Send().Err()
	})

	// Close forum topic
	b.Command("closetopic", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /closetopic <thread_id>").Send().Err()
		}

		threadID := args[0].ToInt().Unwrap().Int64()

		return ctx.CloseForumTopic(threadID).Send().Err()
	})

	// Reopen forum topic
	b.Command("reopentopic", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /reopentopic <thread_id>").Send().Err()
		}

		threadID := args[0].ToInt().Unwrap().Int64()

		return ctx.ReopenForumTopic(threadID).Send().Err()
	})

	// Delete forum topic
	b.Command("deletetopic", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletetopic <thread_id>").Send().Err()
		}

		threadID := args[0].ToInt().Unwrap().Int64()

		return ctx.DeleteForumTopic(threadID).Send().Err()
	})

	// Edit general forum topic
	b.Command("editgeneral", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /editgeneral <new_name>").Send().Err()
		}

		newName := args[0]

		return ctx.EditGeneralForumTopic(newName).Send().Err()
	})

	// Close general forum topic
	b.Command("closegeneral", func(ctx *ctx.Context) error {
		return ctx.CloseGeneralForumTopic().Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
