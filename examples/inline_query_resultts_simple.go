package main

import (
	"fmt"

	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/types/updates"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.On.Inline.Any(func(ctx *ctx.Context) error {
		fmt.Println("inline received:", ctx.Update.InlineQuery.Query)

		query := ctx.Update.InlineQuery.Query
		queryID := ctx.Update.InlineQuery.Id

		result := inline.Article("id1", "Echo", input.Text(String(query)))

		return ctx.AnswerInlineQuery(String(queryID)).
			AddResult(result).
			CacheFor(0).
			Personal().
			Send().Err()
	})

	b.Polling().AllowedUpdates(updates.All...).Start()
}
