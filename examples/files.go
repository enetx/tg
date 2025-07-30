package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.Reply(g.Format("Welcome to <b>{}</b>", ctx.Bot.Raw().Username)).HTML().Send().Err()
	})

	b.Command("doc", func(ctx *ctx.Context) error {
		return ctx.SendDocument("doc.pdf").Caption("pdf doc").Send().Err()
	})

	b.Command("audio", func(ctx *ctx.Context) error {
		return ctx.SendAudio("audio.mp3").
			Caption("some audio").
			ReplyTo(ctx.EffectiveMessage.MessageId).
			Timeout(time.Second * 30).
			Send().
			Err()
	})

	b.Command("photo", func(ctx *ctx.Context) error {
		return ctx.SendPhoto("photo.png").Send().Err()
	})

	b.Command("video", func(ctx *ctx.Context) error {
		return ctx.SendVideo("video.mp4").
			Caption("Look at this cat").
			Spoiler().
			Timeout(time.Minute * 3). // Custom timeout
			ApplyMetadata().          // Extract video info (ffprobe)
			GenerateThumbnail().      // Auto-generate thumbnail (ffmpeg)
			Send().
			Err()
	})

	b.Polling().Start()
}
