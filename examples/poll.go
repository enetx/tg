package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	quiz := func(ctx *ctx.Context) *ctx.SendPoll {
		return ctx.SendPoll("🧠 Choose the correct option:").
			Option(input.Choice("Option A")).
			Option(input.Choice("Option B")).
			Option(input.Choice("Option C")).
			Quiz(1). // Correct option is index 1 (Option B)
			Explanation("<i>Correct answer is B because it's awesome.</i>").
			ExplanationHTML().          // Explanation in HTML
			Anonymous().                // Anonymous poll
			MultipleAnswers().          // Allows multiple answers (ignored in quiz mode)
			ClosesIn(60 * time.Second). // Auto-close in 60 seconds
			Protect().                  // Protects message from being forwarded
			Silent().                   // Sends message silently
			Effect(effects.Fire).       // Adds fireworks effect (if supported)
			Markup(                     // Adds inline keyboard
				keyboard.Inline().
					URL("🌐 Learn more", "https://example.com").
					Row().
					Text("👍 Vote again", "vote_again"))
	}

	b.Command("start", func(ctx *ctx.Context) error { return quiz(ctx).Send().Err() })

	// Handle the "vote_again" callback
	b.On.Callback.Equal("vote_again", func(ctx *ctx.Context) error {
		ctx.DeleteMessage().Send()
		return quiz(ctx).Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
