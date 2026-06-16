package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/types/updates"
)

// Guest messages (Bot API 10.0): a guest bot summoned into a chat receives a message
// carrying a guest_query_id; reply to it with answerGuestQuery.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.On.Message.Any(func(ctx *ctx.Context) error {
		// Only messages summoned from a guest bot carry a guest query id.
		if ctx.EffectiveMessage.GuestQueryId == "" {
			return nil
		}

		queryID := g.String(ctx.EffectiveMessage.GuestQueryId)

		result := inline.NewArticle(
			"answer-1",
			"Hello from the guest bot",
			input.Text(g.String("Thanks for summoning me!")),
		)

		return ctx.AnswerGuestQuery(queryID, result).Send().Err()
	})

	b.Polling().AllowedUpdates(updates.All...).Start()
}
