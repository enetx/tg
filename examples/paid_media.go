package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Send paid photo
	b.Command("paidphoto", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /paidphoto <star_count> <photo_path>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()
		photo := file.Input(args[1]).Unwrap()

		defer photo.File.Close()

		result := ctx.SendPaidMedia(stars).
			Photo(input.PaidPhoto(photo)).
			Caption("📷 Exclusive photo content for " + args[0] + " stars!").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("✅ Paid photo sent successfully!").Send().Err()
	})

	// Send paid video
	b.Command("paidvideo", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /paidvideo <star_count> <video_path>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()
		video := file.Input(args[1]).Unwrap()

		defer video.File.Close()

		result := ctx.SendPaidMedia(stars).
			Video(input.PaidVideo(video).SupportsStreaming()).
			Caption("🎬 Premium video content").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("✅ Paid video sent successfully!").Send().Err()
	})

	// Send multiple paid media
	b.Command("paidalbum", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 4 {
			return ctx.Reply("Usage: /paidalbum <star_count> <photo1> <photo2> <video1>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()

		photo1 := file.Input(args[1]).Unwrap()
		defer photo1.File.Close()

		photo2 := file.Input(args[2]).Unwrap()
		defer photo2.File.Close()

		video := file.Input(args[3]).Unwrap()
		defer video.File.Close()

		result := ctx.SendPaidMedia(stars).
			Photo(input.PaidPhoto(photo1)).
			Photo(input.PaidPhoto(photo2)).
			Video(input.PaidVideo(video).SupportsStreaming()).
			Caption("🖼 Premium media album with photos and video").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("✅ Paid media album sent successfully!").Send().Err()
	})

	// Send with custom payload and protection
	b.Command("paidprotected", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /paidprotected <star_count> <photo_path>").Send().Err()
		}

		stars := args[0].ToInt().Unwrap().Int64()

		photo := file.Input(args[1]).Unwrap()
		defer photo.File.Close()

		result := ctx.SendPaidMedia(stars).
			Photo(input.PaidPhoto(photo)).
			Caption("🔒 Protected premium content").
			Payload("premium_content_v1").
			Protect().
			Silent().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("✅ Protected paid media sent successfully!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
