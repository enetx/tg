package main

import (
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Handle shipping queries during payment process
	b.On.Shipping.Any(func(ctx *ctx.Context) error {
		// Check shipping address (example validation)
		shipping := ctx.Update.ShippingQuery
		address := shipping.ShippingAddress

		// Example: Don't deliver to certain countries
		if address.CountryCode == "XX" {
			return ctx.AnswerShippingQuery().
				Error("We don't deliver to your country").
				Send().Err()
		}

		// Example: Different shipping for different regions
		if address.CountryCode == "US" {
			return ctx.AnswerShippingQuery().
				Ok().
				Option("express", "Express Shipping").
				Price("Express Delivery", 500). // $5.00
				Done().
				Option("standard", "Standard Shipping").
				Price("Standard Delivery", 200). // $2.00
				Done().
				Send().Err()
		}

		// International shipping
		return ctx.AnswerShippingQuery().
			Ok().
			Option("international", "International Shipping").
			Price("International Delivery", 1000). // $10.00
			Price("Handling Fee", 100).            // $1.00
			Done().
			Send().Err()
	})

	// Handle /start command to initiate payment
	b.Command("start", func(ctx *ctx.Context) error {
		return ctx.SendInvoice(
			"Test Product",
			"A test product for shipping demo",
			"test-payload",
			"USD",
		).
			Price("Test Product", 1500). // $15.00
			ProviderToken("284685532:TEST:your_test_token").
			StartParameter("test_shipping").
			NeedShipping().
			Send().Err()
	})

	// More advanced shipping example
	b.Command("premium", func(ctx *ctx.Context) error {
		return ctx.SendInvoice(
			"Premium Package",
			"Premium package with multiple shipping options",
			"premium-payload",
			"USD",
		).
			Price("Premium Package", 5000). // $50.00
			Price("Gift Wrapping", 300).    // $3.00
			ProviderToken("284685532:TEST:your_test_token").
			StartParameter("premium_shipping").
			NeedShipping().
			NeedName().
			NeedPhone().
			Send().Err()
	})

	// Advanced shipping handler with weight-based pricing
	b.On.Shipping.Any(func(ctx *ctx.Context) error {
		shipping := ctx.Update.ShippingQuery
		address := shipping.ShippingAddress

		// Block certain regions
		blockedCountries := g.SliceOf("XX", "YY", "ZZ")
		if blockedCountries.Contains(address.CountryCode) {
			return ctx.AnswerShippingQuery().
				Error("Sorry, we don't ship to your region").
				Send().Err()
		}

		// Different shipping for different payloads
		switch shipping.InvoicePayload {
		case "test-payload":
			return ctx.AnswerShippingQuery().
				Ok().
				Option("fast", "Fast Delivery (1-2 days)").
				Price("Fast Shipping", 800).
				Price("Insurance", 200).
				Done().
				Option("economy", "Economy Delivery (5-7 days)").
				Price("Economy Shipping", 300).
				Done().
				Send().Err()

		case "premium-payload":
			answer := ctx.AnswerShippingQuery().Ok()

			// Premium white-glove service
			answer.Option("white_glove", "White Glove Service").
				Price("Premium Delivery", 2000).
				Price("Setup Service", 1000).
				Price("Insurance", 500).
				Done()

			// Express service
			answer.Option("express", "Express Service").
				Price("Express Delivery", 1200).
				Price("Signature Required", 300).
				Done()

			// International handling for premium
			if address.CountryCode != "US" {
				answer.Option("international", "International Premium").
					Price("International Shipping", 3000).
					Price("Customs Handling", 800).
					Price("Tracking", 200).
					Done()
			}

			return answer.Send().Err()

		default:
			// Standard shipping for unknown payloads
			return ctx.AnswerShippingQuery().
				Ok().
				Option("standard", "Standard Shipping").
				Price("Standard Delivery", 500).
				Done().
				Send().Err()
		}
	})

	b.Polling().Start()
}
