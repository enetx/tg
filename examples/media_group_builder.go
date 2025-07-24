package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Basic video with auto metadata
	b.Command("videometa", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /videometa <video_path>").Send().Err()
		}

		video := args[0]

		result := ctx.MediaGroup().
			Video(video).
			ApplyMetadata().
			Caption("Video with auto-extracted metadata").
			HTML().
			Add().
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

		video := args[0]
		thumb := args[1]

		result := ctx.MediaGroup().
			Video(video).
			Width(1920).Height(1080).Duration(300).
			Thumbnail(thumb).
			Caption("Video with custom thumbnail").
			Markdown().
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sent with custom thumbnail!").Send().Err()
	})

	// Video with auto-generated thumbnail
	b.Command("videogen", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /videogen <video_path> [seek_time]").Send().Err()
		}

		video := args[0]

		videoBuilder := ctx.MediaGroup().
			Video(video).
			ApplyMetadata().
			Caption("Video with auto-generated thumbnail")

		// Optional custom seek time for thumbnail
		if args.Len() > 1 {
			seekTime := args[1]
			videoBuilder = videoBuilder.GenerateThumbnail(seekTime)
		} else {
			videoBuilder = videoBuilder.GenerateThumbnail() // Middle of video
		}

		result := videoBuilder.Add().Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sent with generated thumbnail!").Send().Err()
	})

	// Video with streaming support
	b.Command("videostream", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /videostream <video_path>").Send().Err()
		}

		video := args[0]

		result := ctx.MediaGroup().
			Video(video).
			ApplyMetadata().
			Streamable().
			Caption("Streamable video content").
			ShowCaptionAbove().
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Streamable video sent!").Send().Err()
	})

	// Mixed media album with advanced video
	b.Command("mixedalbum", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 3 {
			return ctx.Reply("Usage: /mixedalbum <photo1> <video1> <photo2>").Send().Err()
		}

		photo1 := args[0]
		video1 := args[1]
		photo2 := args[2]

		result := ctx.MediaGroup().
			Photo(photo1, "First photo").
			Video(video1).
			ApplyMetadata().
			GenerateThumbnail().
			Caption("Advanced video in album").
			HTML().
			Streamable().
			Add().
			Photo(photo2, "Second photo").
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Mixed media album sent!").Send().Err()
	})

	// Video with spoiler effect
	b.Command("videospoiler", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 1 {
			return ctx.Reply("Usage: /videospoiler <video_path>").Send().Err()
		}

		video := args[0]

		result := ctx.MediaGroup().
			Video(video).
			ApplyMetadata().
			HasSpoiler().
			Caption("Spoiler video content").
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Spoiler video sent!").Send().Err()
	})

	// Video with custom dimensions
	b.Command("videohd", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 4 {
			return ctx.Reply("Usage: /videohd <video_path> <width> <height> <duration>").Send().Err()
		}

		video := args[0]
		width := args[1].ToInt().Unwrap().Int64()
		height := args[2].ToInt().Unwrap().Int64()
		duration := args[3].ToInt().Unwrap().Int64()

		result := ctx.MediaGroup().
			Video(video).
			Width(width).
			Height(height).
			Duration(duration).
			Streamable().
			Caption("HD Video " + args[1] + "x" + args[2]).
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("HD video sent!").Send().Err()
	})

	// Multiple videos with different configurations
	b.Command("multiVideo", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /multivideo <video1> <video2>").Send().Err()
		}

		video1 := args.Get(0).Unwrap()
		video2 := args.Get(1).Unwrap()

		result := ctx.MediaGroup().
			Video(video1).
			ApplyMetadata().
			GenerateThumbnail("00:00:05"). // Thumbnail at 5 seconds
			Caption("First video with early thumbnail").
			Add().
			Video(video2).
			ApplyMetadata().
			GenerateThumbnail(). // Thumbnail at middle
			Caption("Second video with middle thumbnail").
			Streamable().
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Multiple videos sent!").Send().Err()
	})

	// Video to specific chat
	b.Command("videoto", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if args.Len() < 2 {
			return ctx.Reply("Usage: /videoto <chat_id> <video_path>").Send().Err()
		}

		chatID := args[0].ToInt().Unwrap().Int64()
		video := args[1]

		result := ctx.MediaGroup().
			To(chatID).
			Video(video).
			ApplyMetadata().
			GenerateThumbnail().
			Caption("🚀 Cross-chat video delivery").
			Add().
			Send()

		if result.IsErr() {
			return result.Err()
		}

		return ctx.Reply("Video sent to chat " + args.Get(0).Unwrap()).Send().Err()
	})

	b.Polling().AllowedUpdates().Start()
}
