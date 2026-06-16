package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

// Live photos (Bot API 10.0): a short video (up to 10 seconds, 10 MB) paired with a
// static cover photo. Sending live photos by URL is not supported — upload local files.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /live — send a live photo (the short video plus its cover photo).
	b.Command("live", func(ctx *ctx.Context) error {
		return ctx.SendLivePhoto(g.String("live.mp4"), g.String("cover.jpg")).
			Caption(g.String("A moment caught in motion")).
			HTML().
			ShowCaptionAboveMedia().
			Spoiler().
			Send().Err()
	})

	b.Polling().Start()
}
