package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

// Date-time entities (Bot API 9.5): the "date_time" entity renders a Unix timestamp as a
// localized, auto-updating date/time. An optional format string (e.g. "t", "r", "wDT")
// controls how it is displayed.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /when — show a message where the word "soon" renders as a live date/time one hour from now.
	b.Command("when", func(ctx *ctx.Context) error {
		text := g.String("The event starts soon — don't be late!")
		startsAt := time.Now().Add(time.Hour).Unix()

		e := entities.New(text).
			Bold("event").
			DateTime("soon", startsAt, g.String("wDT")) // weekday, date and time

		return ctx.Reply(text).Entities(e).Send().Err()
	})

	b.Polling().Start()
}
