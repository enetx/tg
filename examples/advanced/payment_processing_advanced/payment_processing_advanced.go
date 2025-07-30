// Package main demonstrates advanced payment processing features in TG Framework.
// This example showcases Telegram Stars, invoices, pre-checkout validation,
// refunds, subscription management, and payment webhooks.
package main

import (
	"log"
	"os"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

// Global bot instance for payment operations
var botInstance *bot.Bot

func main() {
	// Get bot token from environment
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create bot instance
	botInstance = bot.New(token).Build().Unwrap()

	// Start command handler
	botInstance.Command("payments", handlePaymentPanel).Register()

	// Payment menu handlers
	botInstance.On.Callback.Equal("star_payments", handleStarPayments)
	botInstance.On.Callback.Equal("invoice_management", handleInvoiceManagement)
	botInstance.On.Callback.Equal("subscription_management", handleSubscriptionManagement)
	botInstance.On.Callback.Equal("payment_history", handlePaymentHistory)

	// Star payment handlers
	botInstance.On.Callback.Equal("buy_stars", handleBuyStars)
	botInstance.On.Callback.Equal("check_star_balance", handleCheckStarBalance)
	botInstance.On.Callback.Equal("star_transactions", handleStarTransactionHistory)

	// Invoice handlers
	botInstance.On.Callback.Equal("create_invoice", handleCreateInvoice)
	botInstance.On.Callback.Equal("create_invoice_link", handleCreateInvoiceLink)
	botInstance.On.Callback.Equal("refund_payment", handleRefundPayment)

	// Subscription handlers
	botInstance.On.Callback.Equal("premium_subscription", handlePremiumSubscription)
	botInstance.On.Callback.Equal("manage_subscription", handleManageSubscription)
	botInstance.On.Callback.Equal("cancel_subscription", handleCancelSubscription)

	// Payment processing handlers
	botInstance.On.PreCheckout.Any(handlePreCheckoutQuery)
	botInstance.On.Message.SuccessfulPayment(handleSuccessfulPayment)

	// Shipping query handler for physical goods
	botInstance.On.Shipping.Any(handleShippingQuery)

	// Back navigation
	botInstance.On.Callback.Equal("back_payments", handlePaymentPanel)

	// Start the bot
	log.Println("🚀 Advanced Payment Processing Example started...")
	botInstance.Polling().AllowedUpdates().Start()
}

// handlePaymentPanel provides main payment processing menu
func handlePaymentPanel(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("⭐ Star Payments", "star_payments").
		Text("🧾 Invoice Management", "invoice_management").
		Row().
		Text("🔄 Subscriptions", "subscription_management").
		Text("📊 Payment History", "payment_history")

	return ctx.Reply("💳 <b>Advanced Payment Processing</b>\n\n" +
		"Complete payment solution for Telegram bots:\n\n" +
		"⭐ <b>Star Payments</b> - Telegram Stars transactions\n" +
		"🧾 <b>Invoice Management</b> - Create and manage invoices\n" +
		"🔄 <b>Subscriptions</b> - Recurring payment management\n" +
		"📊 <b>Payment History</b> - Transaction tracking and analytics\n\n" +
		"<i>Supports both Telegram Stars and traditional payment providers.</i>").
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ STAR PAYMENTS ================

func handleStarPayments(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("💰 Buy Stars", "buy_stars").
		Text("💎 Check Balance", "check_star_balance").
		Row().
		Text("📋 Star Transactions", "star_transactions").
		Row().
		Text("🔙 Back", "back_payments")

	return ctx.EditMessageText("⭐ <b>Telegram Stars Payment System</b>\n\n" +
		"Manage Telegram Stars - the native payment currency:\n\n" +
		"💰 <b>Buy Stars</b> - Purchase stars with various amounts\n" +
		"💎 <b>Check Balance</b> - View current star balance\n" +
		"📋 <b>Star Transactions</b> - Transaction history\n\n" +
		"<b>Star Payment Benefits:</b>\n" +
		"• Instant transactions\n" +
		"• Low processing fees\n" +
		"• Integrated with Telegram\n" +
		"• Global availability\n" +
		"• No external payment processor needed").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleBuyStars(ctx *ctx.Context) error {
	// Create invoice for star purchase
	result := ctx.SendInvoice(
		"⭐ Telegram Stars Purchase",
		"Purchase Telegram Stars to unlock premium features and support the bot",
		"star_purchase_payload",
		"XTR", // XTR is Telegram Stars currency
	).
		Price("Stars", 100). // 100 stars
		Photo("https://telegram.org/img/t_logo.png", 512, 512, 512).
		Send()

	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to create star purchase invoice: {}", result.Err())).Send().Err()
	}

	return ctx.AnswerCallbackQuery("💰 Star purchase invoice created!").Send().Err()
}

func handleCheckStarBalance(ctx *ctx.Context) error {
	// Get current star balance
	result := ctx.GetMyStarBalance().Send()
	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to get star balance: " + result.Err().Error())).Send().Err()
	}

	balance := result.Ok()
	return ctx.Reply(g.String("💎 <b>Current Star Balance</b>\n\n" +
		"<b>Available Stars:</b> ⭐ " + g.Int(balance.Amount).String().Std() + "\n" +
		"<b>Last Updated:</b> " + time.Now().Format("2006-01-02 15:04:05") + "\n\n" +
		"<b>Star Usage:</b>\n" +
		"• Premium features access\n" +
		"• Bot service payments\n" +
		"• Gift purchases\n" +
		"• Subscription renewals\n\n" +
		"<i>Stars are Telegram's universal payment currency.</i>")).
		HTML().Send().Err()
}

func handleStarTransactionHistory(ctx *ctx.Context) error {
	// Get star transaction history
	result := ctx.GetStarTransactions().Send()
	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to get star transactions: {}", result.Err())).Send().Err()
	}

	transactions := result.Ok()
	transactionText := g.String("📋 <b>Star Transaction History</b>\n\n")

	if len(transactions.Transactions) == 0 {
		transactionText += "<i>No transactions found.</i>"
	} else {
		transactionText += "<b>Recent Transactions:</b>\n\n"
		for i, tx := range transactions.Transactions {
			if i >= 5 { // Show only last 5 transactions
				break
			}

			transactionText += "<b>Transaction " + g.Int(i+1).String() + ":</b>\n"
			transactionText += "• <b>Amount:</b> ⭐ " + g.Int(tx.Amount).String() + "\n"
			transactionText += "• <b>Date:</b> " + g.String(time.Unix(int64(tx.Date), 0).Format("2006-01-02 15:04")) + "\n"
			transactionText += "• <b>Type:</b> " + getTransactionType(tx) + "\n\n"
		}
	}

	transactionText += "<b>Balance Operations:</b>\n"
	transactionText += "• <code>ctx.GetMyStarBalance()</code> - Check balance\n"
	transactionText += "• <code>ctx.GetStarTransactions()</code> - View history\n"
	transactionText += "• <code>ctx.RefundStarPayment()</code> - Process refunds"

	return ctx.Reply(transactionText).HTML().Send().Err()
}

// Helper function to determine transaction type
func getTransactionType(tx gotgbot.StarTransaction) g.String {
	if tx.Source != nil {
		return "Incoming"
	}

	if tx.Receiver != nil {
		return "Outgoing"
	}

	return "Unknown"
}

// ================ INVOICE MANAGEMENT ================

func handleInvoiceManagement(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("📄 Create Invoice", "create_invoice").
		Text("🔗 Invoice Link", "create_invoice_link").
		Row().
		Text("💸 Process Refund", "refund_payment").
		Row().
		Text("🔙 Back", "back_payments")

	return ctx.EditMessageText("🧾 <b>Invoice Management System</b>\n\n" +
		"Create and manage payment invoices:\n\n" +
		"📄 <b>Create Invoice</b> - Generate payment invoice\n" +
		"🔗 <b>Invoice Link</b> - Create shareable payment link\n" +
		"💸 <b>Process Refund</b> - Handle payment refunds\n\n" +
		"<b>Invoice Features:</b>\n" +
		"• Multiple payment providers\n" +
		"• Custom pricing and tips\n" +
		"• Customer information collection\n" +
		"• Shipping address support\n" +
		"• Flexible payment options").
		HTML().
		Markup(kb).
		Send().Err()
}

func handleCreateInvoice(ctx *ctx.Context) error {
	// Create comprehensive invoice with all features
	result := ctx.SendInvoice(
		"🎯 Premium Bot Features",
		"Unlock premium features including advanced analytics, priority support, and exclusive content access",
		"premium_features_v1",
		"USD",
	).
		Price("Premium Features", 999). // $9.99
		MaxTip(500).                    // Max $5.00 tip
		SuggestedTips(100, 200, 300).   // $1, $2, $3 tip suggestions
		ProviderData(`{
			"product_id": "premium_features",
			"version": "1.0",
			"features": ["analytics", "priority_support", "exclusive_content"],
			"duration_months": 1
		}`).
		Photo("https://via.placeholder.com/512x512/4CAF50/white?text=Premium", 512, 512, 512).
		NeedName().
		NeedPhone().
		NeedEmail().
		NeedShipping().
		SendEmail().
		SendPhone().
		Flexible().
		Silent().
		Send()

	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to create invoice: {}", result.Err())).Send().Err()
	}

	return ctx.AnswerCallbackQuery("📄 Premium features invoice created!").Send().Err()
}

func handleCreateInvoiceLink(ctx *ctx.Context) error {
	// Create shareable invoice link
	result := ctx.CreateInvoiceLink(
		"🚀 Bot Premium Subscription",
		"Monthly premium subscription with advanced features, priority support, and exclusive content",
		"premium_monthly_sub",
		"USD",
	).
		Price("Premium Subscription", 1999). // $19.99
		MaxTip(1000).                        // Max $10 tip
		SuggestedTips(200, 500, 1000).       // $2, $5, $10 tips
		ProviderData(`{
			"subscription_type": "monthly",
			"tier": "premium",
			"auto_renew": true,
			"features": ["unlimited_usage", "priority_support", "advanced_analytics", "custom_integrations"]
		}`).
		Photo("https://via.placeholder.com/600x400/2196F3/white?text=Premium+Subscription", 600, 400, 600).
		NeedName().
		NeedEmail().
		NeedPhone().
		SendEmail().
		SendPhone().
		Flexible().
		Send()

	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to create invoice link: {}", result.Err())).Send().Err()
	}

	invoiceLink := g.String(result.Ok())

	return ctx.Reply("🔗 <b>Invoice Link Created</b>\n\n" +
		"<b>Product:</b> Bot Premium Subscription\n" +
		"<b>Price:</b> $19.99/month\n" +
		"<b>Link:</b> <code>" + invoiceLink + "</code>\n\n" +
		"<b>Link Features:</b>\n" +
		"• Share anywhere on Telegram\n" +
		"• No bot interaction required\n" +
		"• Direct payment processing\n" +
		"• Automatic fulfillment\n\n" +
		"<i>Users can pay directly through this link!</i>").
		HTML().Send().Err()
}

func handleRefundPayment(ctx *ctx.Context) error {
	// Example refund processing (in real implementation, you'd track actual transactions)
	telegramPaymentChargeID := g.String("tpc_demo_12345")

	result := ctx.RefundStarPayment(telegramPaymentChargeID).Send()
	if result.IsErr() {
		return ctx.Reply(g.Format("❌ Failed to process refund: {}", result.Err())).Send().Err()
	}

	return ctx.Reply(g.String("💸 <b>Refund Processed Successfully</b>\n\n" +
		"<b>Transaction ID:</b> <code>" + telegramPaymentChargeID + "</code>\n" +
		"<b>Status:</b> ✅ Refunded\n" +
		"<b>Processing Time:</b> Instant\n\n" +
		"<b>Refund Details:</b>\n" +
		"• Full amount refunded to original payment method\n" +
		"• User will receive confirmation notification\n" +
		"• Transaction record updated in system\n" +
		"• Customer service notified\n\n" +
		"<i>Refund has been processed and will appear in user's account shortly.</i>")).
		HTML().Send().Err()
}

// ================ SUBSCRIPTION MANAGEMENT ================

func handleSubscriptionManagement(ctx *ctx.Context) error {
	kb := keyboard.Inline().
		Row().
		Text("⭐ Premium Subscription", "premium_subscription").
		Row().
		Text("⚙️ Manage Subscription", "manage_subscription").
		Text("❌ Cancel Subscription", "cancel_subscription").
		Row().
		Text("🔙 Back", "back_payments")

	return ctx.EditMessageText("🔄 <b>Subscription Management</b>\n\n" +
		"Manage recurring payments and subscriptions:\n\n" +
		"⭐ <b>Premium Subscription</b> - Subscribe to premium features\n" +
		"⚙️ <b>Manage Subscription</b> - Update subscription settings\n" +
		"❌ <b>Cancel Subscription</b> - End recurring payments\n\n" +
		"<b>Subscription Benefits:</b>\n" +
		"• Recurring revenue model\n" +
		"• Automatic payment processing\n" +
		"• Flexible billing cycles\n" +
		"• Easy cancellation handling\n" +
		"• Customer retention tools").
		HTML().
		Markup(kb).
		Send().Err()
}

func handlePremiumSubscription(ctx *ctx.Context) error {
	// Create premium subscription invoice
	result := ctx.SendInvoice(
		"⭐ Premium Bot Subscription",
		"Monthly premium subscription with advanced features, priority support, unlimited usage, and exclusive content",
		"premium_subscription_monthly",
		"USD",
	).
		Price("Premium Subscription", 1999). // $19.99
		ProviderData(`{
			"subscription_type": "premium",
			"billing_cycle": "monthly",
			"tier": "pro",
			"features": {
				"unlimited_requests": true,
				"priority_support": true,
				"advanced_analytics": true,
				"custom_integrations": true,
				"api_access": true,
				"white_label": false
			},
			"trial_period_days": 7,
			"auto_renew": true
		}`).
		Photo("https://via.placeholder.com/512x512/FF9800/white?text=Premium", 512, 512, 512).
		NeedEmail().
		SendEmail().
		Send()

	if result.IsErr() {
		return ctx.Reply(g.String("❌ Failed to create subscription invoice: " + result.Err().Error())).Send().Err()
	}

	return ctx.AnswerCallbackQuery("⭐ Premium subscription invoice created!").Send().Err()
}

func handleManageSubscription(ctx *ctx.Context) error {
	// In a real implementation, you would fetch actual subscription data
	return ctx.Reply(g.String("⚙️ <b>Subscription Management</b>\n\n" +
		"<b>Current Subscription:</b>\n" +
		"• <b>Plan:</b> Pro Monthly\n" +
		"• <b>Status:</b> ✅ Active\n" +
		"• <b>Price:</b> $19.99/month\n" +
		"• <b>Next Billing:</b> " + time.Now().AddDate(0, 1, 0).Format("2006-01-02") + "\n" +
		"• <b>Auto Renew:</b> Enabled\n\n" +
		"<b>Subscription Features:</b>\n" +
		"• ✅ Unlimited requests\n" +
		"• ✅ Priority support\n" +
		"• ✅ Advanced analytics\n" +
		"• ✅ Custom integrations\n" +
		"• ✅ API access\n\n" +
		"<b>Management Options:</b>\n" +
		"• Upgrade/downgrade plan\n" +
		"• Update payment method\n" +
		"• Change billing cycle\n" +
		"• Enable/disable auto-renewal\n" +
		"• View billing history\n\n" +
		"<i>Contact support for subscription changes.</i>")).
		HTML().Send().Err()
}

func handleCancelSubscription(ctx *ctx.Context) error {
	// Create confirmation keyboard
	kb := keyboard.Inline().
		Row().
		Text("✅ Confirm Cancellation", "confirm_cancel").
		Text("❌ Keep Subscription", "subscription_management")

	return ctx.Reply(g.String("❌ <b>Cancel Subscription</b>\n\n" +
		"<b>⚠️ Are you sure you want to cancel your subscription?</b>\n\n" +
		"<b>Current Plan:</b> Pro Monthly ($19.99/month)\n" +
		"<b>Active Until:</b> " + time.Now().AddDate(0, 1, 0).Format("2006-01-02") + "\n\n" +
		"<b>What happens when you cancel:</b>\n" +
		"• ❌ No more automatic charges\n" +
		"• ✅ Keep access until end of billing period\n" +
		"• ✅ Can resubscribe anytime\n" +
		"• ✅ Data and settings preserved\n\n" +
		"<b>Alternative Options:</b>\n" +
		"• Pause subscription temporarily\n" +
		"• Downgrade to basic plan\n" +
		"• Switch to annual billing (save 20%)\n\n" +
		"<i>You can always reactivate your subscription later.</i>")).
		HTML().
		Markup(kb).
		Send().Err()
}

// ================ PAYMENT HISTORY ================

func handlePaymentHistory(ctx *ctx.Context) error {
	// In a real implementation, you would fetch actual payment history from database
	return ctx.Reply(g.String("📊 <b>Payment History & Analytics</b>\n\n" +
		"<b>Recent Transactions:</b>\n\n" +
		"<b>1. Premium Subscription</b>\n" +
		"• <b>Date:</b> " + time.Now().AddDate(0, 0, -7).Format("2006-01-02") + "\n" +
		"• <b>Amount:</b> $19.99\n" +
		"• <b>Status:</b> ✅ Completed\n" +
		"• <b>Method:</b> Telegram Stars\n\n" +
		"<b>2. Star Package Purchase</b>\n" +
		"• <b>Date:</b> " + time.Now().AddDate(0, 0, -14).Format("2006-01-02") + "\n" +
		"• <b>Amount:</b> ⭐ 500 stars\n" +
		"• <b>Status:</b> ✅ Completed\n" +
		"• <b>Method:</b> Credit Card\n\n" +
		"<b>3. Premium Features</b>\n" +
		"• <b>Date:</b> " + time.Now().AddDate(0, 0, -21).Format("2006-01-02") + "\n" +
		"• <b>Amount:</b> $9.99\n" +
		"• <b>Status:</b> ✅ Completed\n" +
		"• <b>Method:</b> Telegram Stars\n\n" +
		"<b>Monthly Summary:</b>\n" +
		"• <b>Total Spent:</b> $49.97\n" +
		"• <b>Transactions:</b> 3\n" +
		"• <b>Average:</b> $16.66\n" +
		"• <b>Savings:</b> $5.00 (via subscription)\n\n" +
		"<b>Payment Methods:</b>\n" +
		"• ⭐ Telegram Stars (preferred)\n" +
		"• 💳 Credit/Debit Cards\n" +
		"• 🏦 Bank Transfers\n" +
		"• 📱 Mobile Payments\n\n" +
		"<i>All transactions are secure and encrypted.</i>")).
		HTML().Send().Err()
}

// ================ PAYMENT PROCESSING HANDLERS ================

func handlePreCheckoutQuery(ctx *ctx.Context) error {
	query := ctx.Update.PreCheckoutQuery

	// Validate the payment before processing
	log.Printf("Pre-checkout query: %+v", query)

	// Perform validation logic here
	// - Check inventory
	// - Validate pricing
	// - Verify user eligibility
	// - Check for fraud

	// Example validation
	isValid := true
	var errorMessage g.String

	// Simple payload validation
	if query.InvoicePayload == "" {
		isValid = false
		errorMessage = "Invalid payment payload"
	}

	// Currency validation
	if query.Currency != "USD" && query.Currency != "XTR" {
		isValid = false
		errorMessage = "Unsupported currency"
	}

	if isValid {
		// Answer pre-checkout query positively
		return ctx.AnswerPreCheckoutQuery().Ok().Send().Err()
	}

	// Answer pre-checkout query with error
	return ctx.AnswerPreCheckoutQuery().
		Error(errorMessage).
		Send().Err()
}

func handleSuccessfulPayment(ctx *ctx.Context) error {
	payment := ctx.EffectiveMessage.SuccessfulPayment

	// Log successful payment
	log.Printf("Successful payment: %+v", payment)

	// Process the successful payment
	// - Grant access to paid features
	// - Update user subscription status
	// - Send confirmation email
	// - Update analytics
	// - Trigger fulfillment

	// Send confirmation message
	confirmationText := "🎉 <b>Payment Successful!</b>\n\n" +
		"<b>Transaction Details:</b>\n" +
		"• <b>Amount:</b> " + g.Int(payment.TotalAmount).String().Std() + " " + payment.Currency + "\n" +
		"• <b>Invoice Payload:</b> <code>" + payment.InvoicePayload + "</code>\n" +
		"• <b>Telegram Payment ID:</b> <code>" + payment.TelegramPaymentChargeId + "</code>\n"

	if payment.ProviderPaymentChargeId != "" {
		confirmationText += "• <b>Provider Payment ID:</b> <code>" + payment.ProviderPaymentChargeId + "</code>\n"
	}

	if payment.ShippingOptionId != "" {
		confirmationText += "• <b>Shipping Option:</b> " + payment.ShippingOptionId + "\n"
	}

	confirmationText += "\n<b>✅ Your purchase has been processed successfully!</b>\n\n"

	// Add specific fulfillment message based on payload
	switch payment.InvoicePayload {
	case "premium_subscription_monthly":
		confirmationText += "🔓 <b>Premium features have been activated!</b>\n" +
			"• Unlimited bot usage\n" +
			"• Priority support access\n" +
			"• Advanced analytics dashboard\n" +
			"• Custom integration options\n\n" +
			"Your subscription will auto-renew monthly."
	case "premium_features_v1":
		confirmationText += "🎯 <b>Premium features unlocked!</b>\n" +
			"• Access granted immediately\n" +
			"• Features valid for 1 month\n" +
			"• Check /premium for feature list"
	case "star_purchase_payload":
		confirmationText += "⭐ <b>Stars added to your account!</b>\n" +
			"• Check your balance with /balance\n" +
			"• Use stars for premium features\n" +
			"• Stars never expire"
	default:
		confirmationText += "🛍️ <b>Purchase completed successfully!</b>"
	}

	confirmationText += "\n\n<i>Thank you for your purchase! If you have any questions, contact our support team.</i>"

	return ctx.Reply(g.String(confirmationText)).HTML().Send().Err()
}

func handleShippingQuery(ctx *ctx.Context) error {
	query := ctx.Update.ShippingQuery

	// Log shipping query
	log.Printf("Shipping query: %+v", query)

	// Validate shipping address and provide options
	// In a real implementation, you would validate the address
	// and calculate shipping costs based on location

	return ctx.AnswerShippingQuery().
		Ok().
		Option("standard", "Standard Shipping").
		Price("Shipping", 500).
		Done().
		Option("express", "Express Shipping").
		Price("Express Shipping", 1000).
		Price("Express Handling", 200).
		Done().
		Send().Err()
}
