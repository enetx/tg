package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

// Message drafts (Bot API 9.4+): stream a temporary "draft" preview to a chat while a
// reply is being generated. The draft is ephemeral (~30s); send the final message with
// SendMessage to persist it. Pass empty text to show a "Thinking..." placeholder.
func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// /draft — show a thinking placeholder, then a partial draft, then the final message.
	b.Command("draft", func(ctx *ctx.Context) error {
		draftID := int64(1)

		// Empty text renders a "Thinking..." placeholder.
		if err := ctx.SendMessageDraft(draftID).Text(g.String("")).Send().Err(); err != nil {
			return err
		}

		// Stream a partial draft.
		if err := ctx.SendMessageDraft(draftID).
			Text(g.String("Generating your answer")).
			HTML().
			Send().Err(); err != nil {
			return err
		}

		// Persist the final message.
		return ctx.SendMessage(g.String("Here is your complete answer!")).Send().Err()
	})

	b.Polling().Start()
}
