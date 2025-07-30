package main

import (
	"fmt"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/types/updates"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.On.Inline.Any(func(ctx *ctx.Context) error {
		fmt.Println("inline received:", ctx.Update.InlineQuery.Query)

		query := g.String(ctx.Update.InlineQuery.Query)
		queryID := g.String(ctx.Update.InlineQuery.Id)

		result := inline.Article("id1", "Echo", input.Text(query))

		return ctx.AnswerInlineQuery(queryID).
			AddResult(result).
			CacheFor(0).
			Personal().
			Send().Err()
	})

	b.Polling().AllowedUpdates(updates.All...).Start()
}
