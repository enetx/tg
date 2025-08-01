package main

import (
	"time"

	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
)

func main() {
	// Create a bot instance
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Handle business connections
	b.On.BusinessConnection.Enabled(func(ctx *ctx.Context) error {
		return ctx.Reply("Bot connected! You can now manage stories.").Send().Err()
	})

	// Photo story command
	b.Command("photostory", func(ctx *ctx.Context) error {
		// Simple usage - just pass filename/URL as string
		photo := input.StoryPhoto("photo.jpg")

		return ctx.PostStory("your_business_connection_id", photo).
			Caption("Amazing photo story!").
			HTML().
			ActiveFor(24 * time.Hour). // 24 hours
			PostToChatPage().
			Send().Err()
	})

	// Video story command
	b.Command("videostory", func(ctx *ctx.Context) error {
		video := input.StoryVideo("video.mp4").
			CoverFrameTimestamp(2 * time.Second) // Cover at 2 seconds

		return ctx.PostStory("your_business_connection_id", video).
			Caption("Epic video!").
			HTML().
			Send().Err()
	})

	// Story deletion
	b.Command("deletestory", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /deletestory <story_id>").Send().Err()
		}

		storyID := args[0].ToInt().Unwrap().Int64()
		return ctx.DeleteStory("your_business_connection_id", storyID).Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
