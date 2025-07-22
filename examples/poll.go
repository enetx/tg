package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	quiz := func(ctx *ctx.Context) *ctx.SendPoll {
		return ctx.SendPoll("🧠 Choose the correct option:").
			Option("Option A").
			Option("Option B").
			Option("Option C").
			Quiz(1). // Correct option is index 1 (Option B)
			Explanation("<i>Correct answer is B because it's awesome.</i>").
			ExplanationHTML().    // Explanation in HTML
			Anonymous().          // Anonymous poll
			MultipleAnswers().    // Allows multiple answers (ignored in quiz mode)
			CloseIn(60).          // Auto-close in 60 seconds
			Protect().            // Protects message from being forwarded
			Silent().             // Sends message silently
			Effect(effects.Fire). // Adds fireworks effect (if supported)
			Markup(               // Adds inline keyboard
				keyboard.Inline().
					URL("🌐 Learn more", "https://example.com").
					Row().
					Text("👍 Vote again", "vote_again"))
	}

	b.Command("start", func(ctx *ctx.Context) error { return quiz(ctx).Send().Err() })

	// Handle the "vote_again" callback
	b.On.Callback.Equal("vote_again", func(ctx *ctx.Context) error {
		ctx.Delete().Send()
		return quiz(ctx).Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
