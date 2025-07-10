package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	quiz := func(ctx *tg.Context) *tg.Poll {
		return ctx.Poll("🧠 Choose the correct option:").
			Option("Option A").
			Option("Option B").
			Option("Option C").
			Quiz(1). // Correct option is index 1 (Option B)
			Explanation("<i>Correct answer is B because it's awesome.</i>").
			ExplanationHTML().         // Explanation in HTML
			Anonymous().               // Anonymous poll
			MultipleAnswers().         // Allows multiple answers (ignored in quiz mode)
			CloseIn(60).               // Auto-close in 60 seconds
			Protect().                 // Protects message from being forwarded
			Silent().                  // Sends message silently
			Paid().                    // Telegram Stars premium delivery (if enabled)
			Effect(effects.Fireworks). // Adds fireworks effect (if supported)
			Markup(                    // Adds inline keyboard
				keyboard.Inline().
					URL("🌐 Learn more", "https://example.com").
					Row().
					Text("👍 Vote again", "vote_again"))
	}

	bot.Command("start", func(ctx *tg.Context) error { return quiz(ctx).Send().Err() })

	// Handle the "vote_again" callback
	bot.On.Callback.Equal("vote_again", func(ctx *tg.Context) error {
		ctx.Delete()
		return quiz(ctx).Send().Err()
	})

	bot.Polling().DropPendingUpdates().Start()
}
