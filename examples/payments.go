package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	b.Command("start", func(ctx *ctx.Context) error {
		if ctx.EffectiveChat.Type != "private" {
			return nil
		}

		ctx.SendMessage("Welcome! Here's your invoice.").Send()

		return ctx.Invoice("Digital Product", "A cool digital item.", "invoice_payload_123", "XTR").
			Price("Cool item", 1).
			Protect().
			// Markup(keyboard.Inline().Pay("Buy with 100 ⭐")).
			// Markup(keyboard.Inline().Button(keyboard.NewButton().Text("Buy with 100 ⭐").Pay())).
			Send().
			Err()
	})

	b.Command("refund", func(ctx *ctx.Context) error {
		if ctx.EffectiveChat.Type != "private" {
			return nil
		}

		chargeID := ctx.Args().Get(0).Some()

		if result := ctx.RefundStarPayment(chargeID).Send(); result.IsErr() {
			err := String(result.Err().Error())
			switch {
			case err.Contains("CHARGE_ALREADY_REFUNDED"):
				return ctx.Reply("This charge has already been refunded.").Send().Err()
			case err.Contains("CHARGE_NOT_FOUND"):
				return ctx.Reply("No such payment was found. Please check the charge ID.").Send().Err()
			default:
				return ctx.Reply("An unknown error occurred while processing the refund.").Send().Err()
			}
		}

		return nil
	})

	b.On.PreCheckout.Any(func(ctx *ctx.Context) error {
		// you can validate payload/user here if needed

		// return ctx.PreCheckout().Error("Payment declined").Send().Err()
		return ctx.PreCheckout().Ok().Send().Err()
	})

	b.On.Message.SuccessfulPayment(func(ctx *ctx.Context) error {
		user := ctx.EffectiveUser
		payment := ctx.EffectiveMessage.SuccessfulPayment
		chargeID := ctx.EffectiveMessage.SuccessfulPayment.TelegramPaymentChargeId

		Println("User {1.FirstName} ({1.Id}) paid {2.TotalAmount} {2.Currency} with payload {2.InvoicePayload}",
			user, payment)

		return ctx.Message(Format("Payment complete! Thank you, {}!\nChargeID:\n{}", user.FirstName, chargeID)).
			Send().
			Err()
	})

	b.On.Message.RefundedPayment(func(ctx *ctx.Context) error {
		return ctx.SendMessage("The refund was successful.").Send().Err()
	})

	b.Polling().Start()
}
