package main

import (
	"time"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.String("YOUR_BOT_TOKEN")
	b := bot.New(token).Build().Unwrap()

	// Demo of Timeout and APIURL methods across different ctx types
	b.Command("demo", func(ctx *ctx.Context) error {
		// Message with custom timeout and API URL
		result := ctx.SendMessage("Hello with custom settings! 👋").
			Timeout(30 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Silent().
			Send()

		if result.IsErr() {
			return ctx.Reply("Message failed").Send().Err()
		}

		// Invoice with timeout
		ctx.SendInvoice("Test Product", "Test Description", "test_payload", "USD").
			Price("Product", 100).
			Timeout(15 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Send()

		// Gift with custom settings
		ctx.SendGift("gift_premium").
			To(ctx.EffectiveUser.Id).
			Text("Custom timeout gift! 🎁").
			Timeout(20 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Send()

		// Star balance with timeout
		ctx.GetMyStarBalance().
			Timeout(10 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Send()

		// Business account operations
		account := ctx.Business("connection_id")

		account.SetName("Business Name").
			LastName("Inc").
			Timeout(25 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Send()

		account.Balance().GetStarBalance().
			Timeout(15 * time.Second).
			APIURL("https://custom.api.telegram.org/bot" + token).
			Send()

		return ctx.Reply("Demo completed! All methods support Timeout and APIURL").Send().Err()
	})

	// Demo showing method chaining flexibility
	b.Command("chain", func(ctx *ctx.Context) error {
		// All these are equivalent and demonstrate method chaining flexibility

		// Style 1: Timeout first
		ctx.SendMessage("Style 1").
			Timeout(5 * time.Second).
			APIURL("https://api1.example.com/bot" + token).
			HTML().
			Silent().
			Send()

		// Style 2: APIURL first
		ctx.SendMessage("Style 2").
			APIURL("https://api2.example.com/bot" + token).
			Timeout(5 * time.Second).
			Markdown().
			Protect().
			Send()

		// Style 3: Mixed with other methods
		ctx.SendPhoto("https://example.com/photo.jpg").
			Silent().
			Timeout(10 * time.Second).
			Protect().
			APIURL("https://api3.example.com/bot" + token).
			Caption("Photo with custom settings").
			Send()

		return ctx.Reply("Method chaining demo completed! ⛓️").Send().Err()
	})

	b.Polling().Start()
}
