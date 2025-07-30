package main

import (
	"fmt"

	"github.com/enetx/g"
	"github.com/enetx/tg/bot"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

func main() {
	token := g.NewFile("../.env").Read().Ok().Trim().Split("=").Collect().Last().Some()
	b := bot.New(token).Build().Unwrap()

	// Command to create a simple product invoice link
	b.Command("product", func(ctx *ctx.Context) error {
		result := ctx.CreateInvoiceLink(
			"🚀 Premium Bot Features",
			"Unlock advanced features including priority support, custom commands, and enhanced analytics for your bot experience.",
			"premium_access_v1",
			"USD",
		).
			Price("Premium Feature Access", 1500). // $15.00 (in cents)
			Price("Processing Fee", 100).          // $1.00
			Photo("https://picsum.photos/400/300", 0, 400, 300).
			NeedName().
			NeedEmail().
			MaxTip(500).                  // $5.00 max tip
			SuggestedTips(100, 200, 300). // $1, $2, $3
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to create invoice link: {}", result.Err())).Send().Err()
		}

		link := result.Ok()

		return ctx.Reply(g.Format(`
💳 <b>Premium Features Invoice</b>

Your payment link has been created!

🔗 <b>Payment Link:</b>
<a href="{}">Click here to pay</a>

💰 <b>Total:</b> $16.00 (+ optional tip)
📦 <b>Product:</b> Premium Bot Features
🎯 <b>Features:</b>
• Priority support
• Custom commands
• Enhanced analytics
• Advanced automation

<i>Complete your payment to unlock premium features instantly!</i>
		`, link)).HTML().Preview(preview.New().Disable()).Send().Err()
	})

	// Command to create Telegram Stars invoice link
	b.Command("stars", func(ctx *ctx.Context) error {
		result := ctx.CreateInvoiceLink(
			"⭐ Telegram Stars Package",
			"Purchase Telegram Stars to unlock premium bot features and support the development.",
			"stars_package_100",
			"XTR", // Telegram Stars currency
		).
			Price("Premium Stars Package", 100). // 100 Telegram Stars
			Photo("https://picsum.photos/400/300?random=stars", 0, 400, 300).
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to create Stars invoice: {}", result.Err())).Send().Err()
		}

		link := result.Ok()

		return ctx.Reply(g.Format(`
⭐ <b>Telegram Stars Package</b>

<a href="{}">💳 Pay with Telegram Stars</a>

🌟 <b>Amount:</b> 100 Stars
🎁 <b>Package includes:</b>
• Premium bot access
• Ad-free experience
• Priority support
• Exclusive features

<i>Pay instantly with your Telegram Stars balance!</i>
		`, link)).HTML().Send().Err()
	})

	// Command to create subscription invoice link
	b.Command("subscription", func(ctx *ctx.Context) error {
		result := ctx.CreateInvoiceLink(
			"🔄 Monthly Premium Subscription",
			"Subscribe to premium features with automatic monthly billing. Cancel anytime.",
			"monthly_sub_v1",
			"XTR", // Must be XTR for subscriptions
		).
			Price("Monthly Subscription", 250). // 250 Telegram Stars per month
			SubscriptionPeriod(2592000).        // 30 days in seconds
			Photo("https://picsum.photos/400/300?random=subscription", 0, 400, 300).
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to create subscription link: {}", result.Err())).Send().Err()
		}

		link := result.Ok()

		return ctx.Reply(g.Format(`
🔄 <b>Monthly Premium Subscription</b>

<a href="{}">📅 Subscribe Now</a>

💎 <b>Monthly Cost:</b> 250 Stars
⏰ <b>Billing:</b> Every 30 days
🎯 <b>Benefits:</b>
• All premium features
• Unlimited usage
• Priority support
• New features first

<i>Subscription auto-renews monthly. Cancel anytime in settings.</i>
		`, link)).HTML().Send().Err()
	})

	// Command to create flexible pricing invoice
	b.Command("flexible", func(ctx *ctx.Context) error {
		result := ctx.CreateInvoiceLink(
			"🛠 Custom Service Package",
			"Professional bot customization service with flexible pricing based on requirements.",
			"custom_service_v1",
			"USD",
		).
			Price("Base Service", 2000). // $20.00
			Flexible().                  // Enable flexible pricing
			NeedName().
			NeedEmail().
			NeedPhone().
			NeedShipping().
			MaxTip(1000). // $10.00 max tip
			ProviderData(`{"service_type": "custom", "complexity": "standard"}`).
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to create flexible invoice: {}", result.Err())).Send().Err()
		}

		link := result.Ok()

		return ctx.Reply(g.Format(`
🛠 <b>Custom Service Package</b>

<a href="{}">💼 Get Quote & Pay</a>

💰 <b>Starting at:</b> $20.00
⚡ <b>Flexible Pricing:</b> Final price based on requirements
📝 <b>Service includes:</b>
• Custom bot development
• Feature implementation
• Testing & deployment
• 30-day support

<i>Price will be calculated based on your specific needs during checkout.</i>
		`, link)).HTML().Send().Err()
	})

	// Command to create business invoice link
	b.Command("business", func(ctx *ctx.Context) error {
		args := ctx.Args()
		if len(args) == 0 {
			return ctx.Reply("Usage: /business <business_connection_id>").Send().Err()
		}

		businessID := args[0]

		result := ctx.CreateInvoiceLink(
			"💼 Business Account Service",
			"Premium service for business accounts with enhanced features and priority support.",
			"business_service_v1",
			"XTR",
		).
			Price("Business Service", 50). // 50 Telegram Stars
			Business(businessID).
			Send()

		if result.IsErr() {
			return ctx.Reply(g.Format("Failed to create business invoice: {}", result.Err())).Send().Err()
		}

		link := result.Ok()

		return ctx.Reply(g.Format(`
💼 <b>Business Account Service</b>

<a href="{}">🏢 Pay via Business Account</a>

⭐ <b>Cost:</b> 50 Stars
🎯 <b>Business Features:</b>
• Enhanced API access
• Priority processing
• Business analytics
• Dedicated support

<i>Payment processed through your business account.</i>
		`, link)).HTML().Send().Err()
	})

	// Command to demonstrate invoice link management
	b.Command("invoice_demo", func(ctx *ctx.Context) error {
		return ctx.Reply(`
💳 <b>Invoice Link Demo Commands</b>

Try these commands to create different types of invoice links:

🛒 <code>/product</code> - Standard product with USD payment
⭐ <code>/stars</code> - Telegram Stars one-time payment
🔄 <code>/subscription</code> - Monthly Stars subscription
🛠 <code>/flexible</code> - Flexible pricing with shipping
💼 <code>/business &lt;business_id&gt;</code> - Business account payment

<b>Features Demonstrated:</b>
• Multiple currency support (USD, XTR)
• Flexible and fixed pricing
• Subscription billing
• Business account integration
• Custom photos and descriptions
• Tip suggestions
• User info requirements

<i>Note: Use test payment providers for development</i>
		`).HTML().Markup(
			keyboard.Inline().
				Text("💳 Create Product Link", "demo_product").
				Text("⭐ Create Stars Link", "demo_stars").
				Row().
				Text("🔄 Subscription Link", "demo_subscription").
				Text("🛠 Flexible Pricing", "demo_flexible")).
			Send().Err()
	})

	// Handle demo callbacks
	b.On.Callback.Equal("demo_product", func(ctx *ctx.Context) error {
		result := ctx.CreateInvoiceLink(
			"🎮 Demo Product",
			"This is a demonstration product for testing invoice links.",
			"demo_product_123",
			"USD",
		).
			Price("Demo Product", 999). // $9.99
			Send()

		if result.IsErr() {
			fmt.Println(result.Err())
			return ctx.AnswerCallbackQuery("Failed to create demo link").Send().Err()
		}

		link := result.Ok()

		ctx.EditMessageText(g.Format("🎮 Demo product link created!\n\n<a href=\"{}\">Click to view invoice</a>", link)).
			HTML().
			Send()

		return ctx.AnswerCallbackQuery("Demo product link created! 🎮").Send().Err()
	})

	b.Polling().Start()
}
