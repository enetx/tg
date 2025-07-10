package main

import (
	"github.com/enetx/tg"
	"github.com/enetx/tg/keyboard"

	. "github.com/enetx/g"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	bot := tg.NewBot(token)

	// Create the initial inline keyboard with two rows:
	// - First row: a URL button linking to Google
	// - Second row: (to be filled with toggle buttons)
	markup := keyboard.Inline().
		Row().URL("google", "google.com").
		Row()

	// Create toggleable buttons for fruits
	apple := keyboard.NewButton().On("🍏 Apple").Off("Apple")
	banan := keyboard.NewButton().On("🍌 Banan").Off("Banan")

	// Map of callback data to buttons
	buttons := Map[String, *keyboard.Button]{
		"fruit:apple": apple,
		"fruit:banan": banan,
	}

	// Add fruit buttons dynamically from the map
	buttons.Iter().ForEach(func(cb String, btn *keyboard.Button) {
		markup.Button(btn.Callback(cb))
	})

	// Handle /start command: send initial keyboard
	bot.Command("start", func(ctx *tg.Context) error {
		return ctx.Reply("Choose your fruits:").Markup(markup).Send().Err()
	})

	// Handle all callback queries starting with "fruit"
	bot.On.Callback.Prefix("fruit", func(ctx *tg.Context) error {
		cb := String(ctx.Callback.Data)

		if btn := buttons.Get(cb); btn.IsSome() {
			// Toggle the button state
			btn.Some().Flip()

			// Update URL button based on which fruit was toggled
			markup.Edit(func(b *keyboard.Button) {
				if b.Get.URL() == "google.com" || b.Get.URL() == "www.apple.com" || b.Get.URL() == "www.banan.com" {
					switch cb {
					case "fruit:apple":
						b.Text("apple.com").URL("www.apple.com")
					case "fruit:banan":
						b.Text("banan.com").URL("www.banan.com")
					}
				}
			})

			return ctx.EditMarkup(markup).Send().Err()
		}

		// Unknown button
		return ctx.Answer("Unknown fruit").Send().Err()
	})

	bot.Polling().Start()
}
