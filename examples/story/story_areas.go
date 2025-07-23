package main

import (
	"github.com/enetx/tg/areas"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	b.Command("flexsizing", func(ctx *ctx.Context) error {
		storyAreas := areas.New().
			// Simple position and size
			Position(10.0, 10.0).Size(25.0, 20.0).Location().

			// Custom position and size
			Position(40.0, 10.0).Size(30.0, 25.0).Reaction("❤️").

			// Different sizes
			Position(70.0, 10.0).Size(35.0, 15.0).Link("https://example.com").
			Position(10.0, 50.0).Size(20.0, 30.0).Weather().

			// Complex chain with all modifiers
			Position(50.0, 60.0). // Set position
			Size(25.0, 25.0).     // Set size
			Rotate(45.0).         // Add rotation
			Rounded(20.0).        // Add rounded corners
			UniqueGift()          // Set area type

		return ctx.PostPhotoStory("conn_id", "photo.jpg").
			Caption("Simplified sizing demo!").
			Areas(storyAreas).
			Send().Err()
	})

	b.Command("sizecomparison", func(ctx *ctx.Context) error {
		storyAreas := areas.New().
			// Small areas
			Position(20.0, 20.0).Size(10.0, 10.0).Location().
			Position(40.0, 20.0).Size(15.0, 10.0).Reaction("🔥").

			// Medium areas
			Position(20.0, 40.0).Size(20.0, 15.0).Link("https://example.com").
			Position(50.0, 40.0).Size(25.0, 20.0).Weather().

			// Large areas
			Position(20.0, 70.0).Size(30.0, 25.0).Rounded(15.0).UniqueGift().

			// Custom proportions - wide and short
			Position(60.0, 70.0).Size(35.0, 10.0).Rotate(30.0).Location()

		return ctx.PostVideoStory("conn_id", "video.mp4").
			Caption("Size comparison!").
			Areas(storyAreas).
			CoverFrame(1.5).
			Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
