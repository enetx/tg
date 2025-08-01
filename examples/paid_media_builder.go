package main

import (
	"time"

	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
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

		video := file.Input(args[1]).Unwrap()
		defer video.File.Close()

		result := ctx.SendPaidMedia(stars).
			Video(input.PaidVideo(video).Streamable().StartAt(10 * time.Second)).
			Caption("HD Premium Video with builder pattern and auto metadata!").
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

		photo1 := file.Input(args[1]).Unwrap()
		defer photo1.File.Close()

		video := file.Input(args[2]).Unwrap()
		defer video.File.Close()

		photo2 := file.Input(args[3]).Unwrap()
		defer photo2.File.Close()

		result := ctx.SendPaidMedia(stars).
			Photo(input.PaidPhoto(photo1)).
			Video(input.PaidVideo(video).Streamable()).
			Photo(input.PaidPhoto(photo2)).
			Caption("Mixed premium content with video builder!").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Mixed paid media sent with builder pattern!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
