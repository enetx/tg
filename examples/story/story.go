package main

import (
	"time"

	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
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
		return ctx.PostPhotoStory("your_business_connection_id", "photo.jpg").
			Caption("Amazing photo story!").
			HTML().
			ActiveFor(24 * time.Hour). // 24 hours
			PostToChatPage().
			Send().Err()
	})

	// Video story command
	b.Command("videostory", func(ctx *ctx.Context) error {
		return ctx.PostVideoStory("your_business_connection_id", "video.mp4").
			Caption("Epic video!").
			HTML().
			CoverFrame(2.5). // Cover at 2.5 seconds
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
