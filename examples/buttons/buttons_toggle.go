package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func main() {
	token := NewFile("../../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Create toggleable buttons for fruits
	apple := keyboard.NewButton().Callback("fruit:apple").On("🍏 Apple").Off("Apple")
	banan := keyboard.NewButton().Callback("fruit:banan").On("🍌 Banan").Off("Banan")

	// Create the initial inline keyboard with two rows:
	// - First row: a URL button linking to Google
	// - Second row: toggle buttons for Apple and Banan
	markup := keyboard.Inline().
		Row().URL("google", "google.com").
		Row().Button(apple).Button(banan)

	// Handle /start command: send initial keyboard
	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.Reply("Choose your fruits:").Markup(markup).Send().Err()
	})

	// Handle toggle for Apple button
	b.On.Callback.Equal("fruit:apple", func(ctx *ctx.Context) error {
		apple.Flip()

		markup.Edit(func(btn *keyboard.Button) {
			if btn.Get.URL() == "google.com" || btn.Get.URL() == "www.banan.com" {
				btn.Text("apple.com")
				btn.URL("www.apple.com")
			}
		})

		return ctx.EditMessageReplyMarkup(markup).Send().Err()
	})

	// Handle toggle for Banan button
	b.On.Callback.Equal("fruit:banan", func(ctx *ctx.Context) error {
		banan.Flip()

		markup.Edit(func(btn *keyboard.Button) {
			if btn.Get.URL() == "google.com" || btn.Get.URL() == "www.apple.com" {
				btn.Text("banan.com")
				btn.URL("www.banan.com")
			}
		})

		return ctx.EditMessageReplyMarkup(markup).Send().Err()
	})

	b.Polling().DropPendingUpdates().Start()
}
