package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Basic video with auto metadata
	b.Command("videometa", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /videometa <video_path>").Send().Err()
		}

		video := file.Input(args[0]).Unwrap()
		defer video.File.Close()

		result := ctx.MediaGroup().
			Video(input.Video(video).Caption("Video with auto-extracted metadata").HTML()).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sent with metadata!").Send().Err()
	})

	// Video with custom thumbnail
	b.Command("videothumb", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /videothumb <video_path> <thumbnail_path>").Send().Err()
		}

		video := file.Input(args[0]).Unwrap()
		defer video.File.Close()

		thumb := file.Input(args[1]).Unwrap()
		defer thumb.File.Close()

		result := ctx.MediaGroup().
			Video(input.Video(video).
				Size(1920, 1080).
				Duration(300).
				Caption("Video with custom thumbnail")).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sent with custom thumbnail!").Send().Err()
	})

	// Mixed media album with advanced video
	b.Command("mixedalbum", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /mixedalbum <photo1> <video1> <photo2>").Send().Err()
		}

		photo1 := file.Input(args[0]).Unwrap()
		defer photo1.File.Close()

		video := file.Input(args[1]).Unwrap()
		defer video.File.Close()

		photo2 := file.Input(args[2]).Unwrap()
		defer photo2.File.Close()

		result := ctx.MediaGroup().
			Photo(input.Photo(photo1).Caption("First photo")).
			Video(input.Video(video).Caption("Advanced video in album").SupportsStreaming()).
			Photo(input.Photo(photo2).Caption("Second photo")).
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Mixed media album sent!").Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
