package main

import (
	"time"

	. "github.com/enetx/g"
	"github.com/enetx/tg"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	bot.Command("start", func(ctx *tg.Context) error {
		return ctx.Reply(Format("Welcome to <b>{}</b>", ctx.Bot.Std().Username)).HTML().Send().Err()
	})

	bot.Command("doc", func(ctx *tg.Context) error {
		return ctx.Document("doc.pdf").Caption("pdf doc").Send().Err()
	})

	bot.Command("audio", func(ctx *tg.Context) error {
		return ctx.Audio("audio.mp3").
			Caption("some audio").
			ReplyTo(ctx.EffectiveMessage.MessageId).
			Timeout(time.Second * 30).
			Send().
			Err()
	})

	bot.Command("photo", func(ctx *tg.Context) error {
		return ctx.Photo("photo.png").Send().Err()
	})

	bot.Command("video", func(ctx *tg.Context) error {
		return ctx.Video("video.mp4").
			Caption("Look at this cat").
			Spoiler().
			Timeout(time.Minute * 3).
			ApplyMetadata().     // ffprobe
			GenerateThumbnail(). // ffmpeg
			Send().
			Err()
	})

	bot.Polling().Start()
}
