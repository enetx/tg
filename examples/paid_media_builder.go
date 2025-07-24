package main

import (
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	b := bot.New("YOUR_BOT_TOKEN").Build().Unwrap()

	// Example with video builder pattern
	b.Command("paidvideobuilder", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /paidvideobuilder <star_count> <video_path>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()
		if stars < 1 || stars > 10000 {
			return ctx.Reply("Star count must be between 1-10000").Send().Err()
		}

		video := args[1]

		result := ctx.SendPaidMedia(stars).
			Video(video).
			ApplyMetadata().
			Streamable().
			StartTimestamp(10).
			GenerateThumbnail().
			Add().
			Caption("HD Premium Video with builder pattern and auto metadata!").
			HTML().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Paid video sent with builder pattern!").Send().Err()
	})

	// Example with mixed media using builder
	b.Command("paidmixedbuilder", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 4 {
			return ctx.Reply("Usage: /paidmixedbuilder <star_count> <photo1> <video1> <photo2>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()
		if stars < 1 || stars > 10000 {
			return ctx.Reply("Star count must be between 1-10000").Send().Err()
		}

		result := ctx.SendPaidMedia(stars).
			Photo(args[1]).
			Video(args[2]).
			ApplyMetadata().
			Streamable().
			Thumbnail("custom_thumb.jpg").
			Add().
			Photo(args[3]).
			Caption("Mixed premium content with video builder!").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Mixed paid media sent with builder pattern!").Send().Err()
	})

	// Example with multiple videos using builder
	b.Command("paidmultiplevideos", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /paidmultiplevideos <star_count> <video1> <video2>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()
		if stars < 1 || stars > 10000 {
			return ctx.Reply("Star count must be between 1-10000").Send().Err()
		}

		result := ctx.SendPaidMedia(stars).
			Video(args[1]).
			ApplyMetadata().
			Streamable().
			GenerateThumbnail("5.5").
			Add().
			Video(args[2]).
			ApplyMetadata().
			Cover("cover2.jpg").
			Add().
			Caption("Double premium videos with auto metadata and thumbnails!").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Multiple paid videos sent with builder pattern!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
