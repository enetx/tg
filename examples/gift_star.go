package main

import (
	. "github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
)

func main() {
	token := "YOUR_BOT_TOKEN"
	b := bot.New(token).Build().Unwrap()

	// Command to send a gift
	b.Command("gift", func(ctx *ctx.Context) error {
		giftID := "gift_premium"

		// Send gift to user with text
		result := ctx.SendGift(String(giftID)).
			To(ctx.EffectiveUser.Id).
			Text("Happy holidays! 🎁").
			HTML().
			PayForUpgrade().
			Send()

		if result.IsErr() {
			return ctx.Reply("Failed to send gift").Send().Err()
		}

		return ctx.Reply("Gift sent successfully! 🎉").Send().Err()
	})

	// Command to get available gifts
	b.Command("gifts", func(ctx *ctx.Context) error {
		result := ctx.GetAvailableGifts().Send()
		if result.IsErr() {
			return ctx.Reply("Failed to get gifts").Send().Err()
		}

		gifts := result.Ok()
		return ctx.Reply(Format("Available gifts: {}", len(gifts.Gifts))).Send().Err()
	})

	// Command to check star balance
	b.Command("balance", func(ctx *ctx.Context) error {
		result := ctx.GetMyStarBalance().Send()
		if result.IsErr() {
			return ctx.Reply("Failed to get balance").Send().Err()
		}

		balance := result.Ok()
		return ctx.Reply(Format("Bot star balance: {} ⭐", balance.Amount)).Send().Err()
	})

	// Command to get star transactions
	b.Command("transactions", func(ctx *ctx.Context) error {
		result := ctx.GetStarTransactions().
			Limit(10).
			Offset(0).
			Send()

		if result.IsErr() {
			return ctx.Reply("Failed to get transactions").Send().Err()
		}

		transactions := result.Ok()
		return ctx.Reply(Format("Recent transactions: {}", len(transactions.Transactions))).Send().Err()
	})

	// Business account gift management
	b.Command("business_gifts", func(ctx *ctx.Context) error {
		connectionID := String("your_business_connection_id")
		account := ctx.Business(connectionID)

		// Get business account gifts
		result := account.Balance().GetGifts().
			ExcludeUnsaved().
			SortByPrice().
			Limit(20).
			Send()

		if result.IsErr() {
			return ctx.Reply("Failed to get business gifts").Send().Err()
		}

		gifts := result.Ok()
		return ctx.Reply(Format("Business gifts: {}", gifts.TotalCount)).Send().Err()
	})

	// Business gift operations
	b.Command("gift_ops", func(ctx *ctx.Context) error {
		connectionID := String("your_business_connection_id")
		ownedGiftID := String("owned_gift_id")

		// Convert gift to stars
		if result := ctx.ConvertGiftToStars(connectionID, ownedGiftID).Send(); result.IsOk() {
			return ctx.Reply("Gift converted to stars! ⭐").Send().Err()
		}

		// Transfer gift to another user
		newOwnerChatID := int64(123456789)
		if result := ctx.TransferGift(connectionID, ownedGiftID, newOwnerChatID).
			StarCount(100).Send(); result.IsOk() {
			return ctx.Reply("Gift transferred! 🎁➡️").Send().Err()
		}

		// Upgrade gift
		if result := ctx.UpgradeGift(connectionID, ownedGiftID).
			KeepOriginalDetails().
			StarCount(500).Send(); result.IsOk() {
			return ctx.Reply("Gift upgraded! ✨").Send().Err()
		}

		return ctx.Reply("Gift operation failed").Send().Err()
	})

	// Star subscription management
	b.Command("subscription", func(ctx *ctx.Context) error {
		userID := ctx.EffectiveUser.Id
		chargeID := String("telegram_payment_charge_id")

		// Cancel user's star subscription
		result := ctx.EditUserStarSubscription(userID, chargeID, true).Send()
		if result.IsErr() {
			return ctx.Reply("Failed to cancel subscription").Send().Err()
		}

		return ctx.Reply("Subscription cancelled").Send().Err()
	})

	b.Polling().Start()
}
