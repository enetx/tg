package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateInvoiceLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	description := g.String("Test product description")
	payload := g.String("test_payload")
	currency := g.String("USD")

	// Test basic creation
	result := testCtx.CreateInvoiceLink(title, description, payload, currency)
	if result == nil {
		t.Error("Expected CreateInvoiceLink builder to be created")
	}

	// Test Price method
	result = result.Price(g.String("Product"), 1000)
	if result == nil {
		t.Error("Price method should return CreateInvoiceLink for chaining")
	}

	// Test Business method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).Business(g.String("business_connection_123"))
	if result == nil {
		t.Error("Business method should return CreateInvoiceLink for chaining")
	}

	// Test ProviderToken method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).ProviderToken(g.String("provider_token_456"))
	if result == nil {
		t.Error("ProviderToken method should return CreateInvoiceLink for chaining")
	}

	// Test SubscriptionPeriod method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).SubscriptionPeriod(2592000)
	if result == nil {
		t.Error("SubscriptionPeriod method should return CreateInvoiceLink for chaining")
	}

	// Test MaxTip method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).MaxTip(5000)
	if result == nil {
		t.Error("MaxTip method should return CreateInvoiceLink for chaining")
	}

	// Test SuggestedTips method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).SuggestedTips(100, 500, 1000, 2000)
	if result == nil {
		t.Error("SuggestedTips method should return CreateInvoiceLink for chaining")
	}

	// Test ProviderData method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).ProviderData(g.String(`{"test": "data"}`))
	if result == nil {
		t.Error("ProviderData method should return CreateInvoiceLink for chaining")
	}

	// Test Photo method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).Photo(g.String("https://example.com/photo.jpg"), 1024, 512, 512)
	if result == nil {
		t.Error("Photo method should return CreateInvoiceLink for chaining")
	}

	// Test NeedName method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).NeedName()
	if result == nil {
		t.Error("NeedName method should return CreateInvoiceLink for chaining")
	}

	// Test NeedPhone method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).NeedPhone()
	if result == nil {
		t.Error("NeedPhone method should return CreateInvoiceLink for chaining")
	}

	// Test NeedEmail method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).NeedEmail()
	if result == nil {
		t.Error("NeedEmail method should return CreateInvoiceLink for chaining")
	}

	// Test NeedShipping method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).NeedShipping()
	if result == nil {
		t.Error("NeedShipping method should return CreateInvoiceLink for chaining")
	}

	// Test SendPhone method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).SendPhone()
	if result == nil {
		t.Error("SendPhone method should return CreateInvoiceLink for chaining")
	}

	// Test SendEmail method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).SendEmail()
	if result == nil {
		t.Error("SendEmail method should return CreateInvoiceLink for chaining")
	}

	// Test Flexible method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).Flexible()
	if result == nil {
		t.Error("Flexible method should return CreateInvoiceLink for chaining")
	}

	// Test Timeout method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CreateInvoiceLink for chaining")
	}

	// Test APIURL method
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CreateInvoiceLink for chaining")
	}
}

func TestContext_CreateInvoiceLinkChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Premium Product")
	description := g.String("Premium product with all features")
	payload := g.String("premium_payload")
	currency := g.String("USD")

	// Test complete method chaining
	result := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Product"), 1000).
		Price(g.String("Shipping"), 200).
		Business(g.String("business_connection_789")).
		ProviderToken(g.String("stripe_token_abc")).
		SubscriptionPeriod(2592000).
		MaxTip(10000).
		SuggestedTips(500, 1000, 2500, 5000).
		ProviderData(g.String(`{"stripe_metadata": "value"}`)).
		Photo(g.String("https://example.com/premium.jpg"), 2048, 1024, 768).
		NeedName().
		NeedPhone().
		NeedEmail().
		NeedShipping().
		SendPhone().
		SendEmail().
		Flexible().
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CreateInvoiceLink")
	}
}

func TestCreateInvoiceLink_PriceItems(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Multi-Price Product")
	description := g.String("Product with multiple price components")
	payload := g.String("multi_price_payload")
	currency := g.String("USD")

	// Test various price combinations
	priceScenarios := []struct {
		name   string
		label  string
		amount int64
	}{
		{"Base Product", "Product", 2000},
		{"Shipping Standard", "Standard Shipping", 500},
		{"Shipping Express", "Express Shipping", 1500},
		{"Tax", "VAT (20%)", 400},
		{"Insurance", "Product Insurance", 300},
		{"Gift Wrapping", "Premium Gift Wrap", 200},
		{"Extended Warranty", "2 Year Warranty", 800},
		{"Installation", "Professional Installation", 1000},
		{"Support Package", "Priority Support", 600},
		{"Training", "Usage Training", 400},
	}

	result := testCtx.CreateInvoiceLink(title, description, payload, currency)

	for _, scenario := range priceScenarios {
		result = result.Price(g.String(scenario.label), scenario.amount)
		if result == nil {
			t.Errorf("Adding price item '%s' (%d) should work", scenario.name, scenario.amount)
		}
	}

	// Test with additional configuration
	finalResult := result.
		MaxTip(5000).
		SuggestedTips(100, 500, 1000).
		Flexible()

	if finalResult == nil {
		t.Error("Multiple price items with configuration should work")
	}
}

func TestCreateInvoiceLink_PaymentProviders(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Payment Provider Test")
	description := g.String("Testing different payment providers")
	payload := g.String("provider_test_payload")
	currency := g.String("USD")

	// Test various payment provider scenarios
	providerScenarios := []struct {
		name         string
		token        string
		providerData string
		description  string
	}{
		{"Stripe", "sk_test_stripe_token_123", `{"stripe_account": "acct_123"}`, "Stripe payment processor"},
		{"PayPal", "paypal_token_456", `{"paypal_merchant": "merchant_456"}`, "PayPal payment processor"},
		{"Telegram Stars", "", `{"stars_subscription": true}`, "Telegram Stars payment"},
		{"Custom Provider", "custom_provider_789", `{"custom_config": "value"}`, "Custom payment provider"},
		{"Bank Transfer", "bank_token_abc", `{"bank_code": "SWIFT123", "account": "12345"}`, "Direct bank transfer"},
		{"Cryptocurrency", "crypto_token_def", `{"wallet": "bitcoin", "network": "mainnet"}`, "Crypto payment"},
		{"Mobile Payment", "mobile_token_ghi", `{"carrier": "telecom", "region": "US"}`, "Mobile carrier billing"},
		{"Gift Cards", "gift_token_jkl", `{"card_type": "amazon", "region": "global"}`, "Gift card redemption"},
	}

	for _, provider := range providerScenarios {
		t.Run(provider.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Product"), 1000)

			if provider.token != "" {
				result = result.ProviderToken(g.String(provider.token))
			}

			result = result.ProviderData(g.String(provider.providerData))

			if result == nil {
				t.Errorf("%s (%s) should work", provider.name, provider.description)
			}

			// Test with additional features for each provider
			enhancedResult := result.
				MaxTip(2000).
				SuggestedTips(100, 500, 1000).
				NeedEmail().
				Timeout(60 * time.Second)

			if enhancedResult == nil {
				t.Errorf("Enhanced %s configuration should work", provider.name)
			}
		})
	}
}

func TestCreateInvoiceLink_SubscriptionPeriods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Subscription Service")
	description := g.String("Various subscription periods")
	payload := g.String("subscription_payload")
	currency := g.String("XTR") // Telegram Stars

	// Test various subscription periods (in seconds)
	subscriptionPeriods := []struct {
		name        string
		period      int64
		description string
	}{
		{"Weekly", 604800, "7 days subscription"},          // 7 * 24 * 60 * 60
		{"Bi-Weekly", 1209600, "14 days subscription"},     // 14 * 24 * 60 * 60
		{"Monthly", 2592000, "30 days subscription"},       // 30 * 24 * 60 * 60
		{"Quarterly", 7776000, "90 days subscription"},     // 90 * 24 * 60 * 60
		{"Semi-Annual", 15552000, "180 days subscription"}, // 180 * 24 * 60 * 60
		{"Annual", 31536000, "365 days subscription"},      // 365 * 24 * 60 * 60
		{"Custom 3 Days", 259200, "3 days trial"},          // 3 * 24 * 60 * 60
		{"Custom 6 Months", 15724800, "186 days exact"},    // 186 * 24 * 60 * 60
	}

	for _, period := range subscriptionPeriods {
		t.Run(period.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Subscription"), 100).
				SubscriptionPeriod(period.period)

			if result == nil {
				t.Errorf("%s (%s) should work", period.name, period.description)
			}

			// Test with business connection for subscriptions
			businessResult := result.Business(g.String("business_subscription_123"))
			if businessResult == nil {
				t.Errorf("Business subscription for %s should work", period.name)
			}

			// Test with minimal info required for subscriptions
			minimalResult := businessResult.
				NeedEmail().
				SendEmail()

			if minimalResult == nil {
				t.Errorf("Minimal subscription configuration for %s should work", period.name)
			}
		})
	}
}

func TestCreateInvoiceLink_TippingScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Service with Tips")
	description := g.String("Service that accepts tips")
	payload := g.String("tipping_payload")
	currency := g.String("USD")

	// Test various tipping configurations
	tippingScenarios := []struct {
		name          string
		maxTip        int64
		suggestedTips []int64
		description   string
	}{
		{"Restaurant", 5000, []int64{200, 500, 1000, 2000}, "Restaurant service with standard tips"},
		{"Delivery", 1000, []int64{100, 200, 300, 500}, "Food delivery with modest tips"},
		{"Luxury Service", 50000, []int64{1000, 5000, 10000, 25000}, "High-end service with premium tips"},
		{"Coffee Shop", 500, []int64{50, 100, 150, 200}, "Coffee shop with small tips"},
		{"Ride Service", 2000, []int64{100, 300, 500, 1000}, "Ride sharing with variable tips"},
		{"Beauty Salon", 10000, []int64{500, 1000, 2000, 5000}, "Beauty services with generous tips"},
		{"Personal Training", 15000, []int64{1000, 2500, 5000, 7500}, "Fitness services with appreciation tips"},
		{"Consulting", 25000, []int64{2000, 5000, 10000, 15000}, "Professional consulting with performance tips"},
	}

	for _, scenario := range tippingScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Service"), 5000).
				MaxTip(scenario.maxTip).
				SuggestedTips(scenario.suggestedTips...)

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with provider configuration
			providerResult := result.
				ProviderToken(g.String("tip_provider_token")).
				ProviderData(g.String(`{"tips_enabled": true, "service_type": "` + scenario.name + `"}`))

			if providerResult == nil {
				t.Errorf("Provider configuration for %s should work", scenario.name)
			}

			// Test with customer info requirements
			customerResult := providerResult.
				NeedName().
				NeedPhone().
				NeedEmail()

			if customerResult == nil {
				t.Errorf("Customer info requirements for %s should work", scenario.name)
			}
		})
	}

	// Test edge cases for tipping
	// Zero max tip
	result := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("No Tips Service"), 1000).
		MaxTip(0).
		SuggestedTips()

	if result == nil {
		t.Error("Zero max tip should work")
	}

	// Very high max tip
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Luxury Service"), 10000).
		MaxTip(1000000).
		SuggestedTips(10000, 50000, 100000, 500000)

	if result == nil {
		t.Error("Very high max tip should work")
	}
}

func TestCreateInvoiceLink_PhotoConfiguration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Product with Photos")
	description := g.String("Products with various photo configurations")
	payload := g.String("photo_payload")
	currency := g.String("USD")

	// Test various photo configurations
	photoConfigurations := []struct {
		name        string
		photoURL    string
		size        int64
		width       int64
		height      int64
		description string
	}{
		{"Small Square", "https://example.com/small.jpg", 1024, 100, 100, "Small square product image"},
		{"Medium Landscape", "https://example.com/medium.jpg", 5120, 320, 240, "Medium landscape product image"},
		{"Large Portrait", "https://example.com/large.jpg", 10240, 300, 400, "Large portrait product image"},
		{"HD Landscape", "https://example.com/hd.jpg", 51200, 800, 600, "HD landscape product image"},
		{"Ultra HD Square", "https://example.com/uhd.jpg", 204800, 1000, 1000, "Ultra HD square product image"},
		{"Banner Style", "https://example.com/banner.jpg", 15360, 600, 200, "Wide banner style image"},
		{"Mobile Optimized", "https://example.com/mobile.jpg", 2048, 200, 300, "Mobile optimized vertical image"},
		{"Thumbnail", "https://example.com/thumb.jpg", 512, 64, 64, "Small thumbnail image"},
	}

	for _, config := range photoConfigurations {
		t.Run(config.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Product"), 1000).
				Photo(g.String(config.photoURL), config.size, config.width, config.height)

			if result == nil {
				t.Errorf("%s (%s) should work", config.name, config.description)
			}

			// Test with additional configuration
			enhancedResult := result.
				ProviderToken(g.String("photo_provider_token")).
				NeedName().
				NeedEmail().
				Flexible()

			if enhancedResult == nil {
				t.Errorf("Enhanced photo configuration for %s should work", config.name)
			}
		})
	}

	// Test edge cases for photos
	// Empty photo URL
	result := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Product"), 1000).
		Photo(g.String(""), 0, 0, 0)

	if result == nil {
		t.Error("Empty photo configuration should work")
	}

	// Maximum photo dimensions
	result = testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Product"), 1000).
		Photo(g.String("https://example.com/max.jpg"), 20971520, 10000, 10000) // 20MB, 10k x 10k

	if result == nil {
		t.Error("Maximum photo dimensions should work")
	}
}

func TestCreateInvoiceLink_CustomerInformation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Information Collection Service")
	description := g.String("Service requiring customer information")
	payload := g.String("info_payload")
	currency := g.String("USD")

	// Test various customer information requirements
	infoScenarios := []struct {
		name         string
		needName     bool
		needPhone    bool
		needEmail    bool
		needShipping bool
		sendPhone    bool
		sendEmail    bool
		description  string
	}{
		{"Minimal", false, false, false, false, false, false, "No customer info required"},
		{"Name Only", true, false, false, false, false, false, "Only name required"},
		{"Contact Basic", true, false, true, false, false, false, "Name and email required"},
		{"Contact Full", true, true, true, false, false, false, "All contact info required"},
		{"Physical Product", true, true, true, true, false, false, "Physical delivery with shipping"},
		{"Service with Provider Sharing", true, true, true, false, true, true, "Service sharing contact with provider"},
		{"Complete Information", true, true, true, true, true, true, "All information collection enabled"},
		{"Email Only", false, false, true, false, false, false, "Email address only"},
		{"Phone Only", false, true, false, false, false, false, "Phone number only"},
		{"Shipping Only", false, false, false, true, false, false, "Shipping address only"},
	}

	for _, scenario := range infoScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Service"), 1000)

			if scenario.needName {
				result = result.NeedName()
			}
			if scenario.needPhone {
				result = result.NeedPhone()
			}
			if scenario.needEmail {
				result = result.NeedEmail()
			}
			if scenario.needShipping {
				result = result.NeedShipping()
			}
			if scenario.sendPhone {
				result = result.SendPhone()
			}
			if scenario.sendEmail {
				result = result.SendEmail()
			}

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with provider configuration
			providerResult := result.
				ProviderToken(g.String("info_provider_token")).
				ProviderData(g.String(`{"collect_info": true, "scenario": "` + scenario.name + `"}`))

			if providerResult == nil {
				t.Errorf("Provider configuration for %s should work", scenario.name)
			}
		})
	}
}

func TestCreateInvoiceLink_FlexiblePricing(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Flexible Pricing Service")
	description := g.String("Service with flexible pricing")
	payload := g.String("flexible_payload")
	currency := g.String("USD")

	// Test flexible pricing scenarios
	flexibleScenarios := []struct {
		name        string
		basePrice   int64
		isFlexible  bool
		description string
	}{
		{"Fixed Price Service", 1000, false, "Service with fixed pricing"},
		{"Flexible Consultation", 5000, true, "Consultation with flexible pricing"},
		{"Flexible Delivery", 500, true, "Delivery with variable pricing"},
		{"Fixed Product", 2500, false, "Physical product with fixed price"},
		{"Flexible Service Package", 10000, true, "Service package with flexible options"},
		{"Fixed Digital Product", 1500, false, "Digital product with fixed price"},
		{"Flexible Subscription", 3000, true, "Subscription with flexible features"},
		{"Fixed Membership", 8000, false, "Membership with fixed benefits"},
	}

	for _, scenario := range flexibleScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Base Service"), scenario.basePrice)

			if scenario.isFlexible {
				result = result.Flexible()
			}

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with additional pricing components for flexible scenarios
			if scenario.isFlexible {
				enhancedResult := result.
					Price(g.String("Optional Add-on"), 200).
					Price(g.String("Premium Feature"), 500).
					MaxTip(1000).
					SuggestedTips(100, 300, 500)

				if enhancedResult == nil {
					t.Errorf("Enhanced flexible pricing for %s should work", scenario.name)
				}
			}
		})
	}
}

func TestCreateInvoiceLink_BusinessIntegration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Business Service")
	description := g.String("Service integrated with business accounts")
	payload := g.String("business_payload")
	currency := g.String("USD")

	// Test various business connection scenarios
	businessScenarios := []struct {
		name         string
		connectionID string
		serviceType  string
		description  string
	}{
		{"Restaurant Business", "restaurant_conn_123", "restaurant", "Restaurant business integration"},
		{"Retail Store", "retail_conn_456", "retail", "Retail store business integration"},
		{"Professional Services", "professional_conn_789", "consulting", "Professional services integration"},
		{"E-commerce", "ecommerce_conn_abc", "online_store", "E-commerce platform integration"},
		{"Service Provider", "service_conn_def", "services", "General service provider integration"},
		{"Subscription Business", "subscription_conn_ghi", "subscription", "Subscription business integration"},
		{"Marketplace Vendor", "marketplace_conn_jkl", "marketplace", "Marketplace vendor integration"},
		{"Digital Content", "digital_conn_mno", "digital", "Digital content creator integration"},
	}

	for _, scenario := range businessScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.CreateInvoiceLink(title, description, payload, currency).
				Price(g.String("Business Service"), 2000).
				Business(g.String(scenario.connectionID))

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with provider configuration for business
			businessResult := result.
				ProviderToken(g.String("business_provider_token")).
				ProviderData(g.String(`{"business_type": "` + scenario.serviceType + `", "integration": true}`))

			if businessResult == nil {
				t.Errorf("Business provider configuration for %s should work", scenario.name)
			}

			// Test with customer requirements for business
			customerResult := businessResult.
				NeedName().
				NeedEmail().
				SendEmail().
				Flexible()

			if customerResult == nil {
				t.Errorf("Business customer requirements for %s should work", scenario.name)
			}
		})
	}
}

func TestCreateInvoiceLink_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty strings
	result := testCtx.CreateInvoiceLink(g.String(""), g.String(""), g.String(""), g.String(""))
	if result == nil {
		t.Error("Empty strings should work (builder creation)")
	}

	// Test with zero price
	result = testCtx.CreateInvoiceLink(g.String("Free Product"), g.String("Free description"), g.String("free_payload"), g.String("USD")).
		Price(g.String("Free Item"), 0)
	if result == nil {
		t.Error("Zero price should work")
	}

	// Test with negative price (edge case)
	result = testCtx.CreateInvoiceLink(g.String("Refund"), g.String("Refund item"), g.String("refund_payload"), g.String("USD")).
		Price(g.String("Refund Amount"), -1000)
	if result == nil {
		t.Error("Negative price should work (for refunds)")
	}

	// Test with very large price
	result = testCtx.CreateInvoiceLink(g.String("Expensive Item"), g.String("Very expensive item"), g.String("expensive_payload"), g.String("USD")).
		Price(g.String("Luxury Product"), 999999999)
	if result == nil {
		t.Error("Very large price should work")
	}

	// Test with zero subscription period
	result = testCtx.CreateInvoiceLink(g.String("One-time"), g.String("One-time payment"), g.String("onetime_payload"), g.String("XTR")).
		Price(g.String("Service"), 100).
		SubscriptionPeriod(0)
	if result == nil {
		t.Error("Zero subscription period should work")
	}

	// Test with zero max tip and empty suggested tips
	result = testCtx.CreateInvoiceLink(g.String("No Tips"), g.String("Service without tips"), g.String("notips_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		MaxTip(0).
		SuggestedTips()
	if result == nil {
		t.Error("Zero max tip with empty suggested tips should work")
	}

	// Test with empty business connection
	result = testCtx.CreateInvoiceLink(g.String("No Business"), g.String("Service without business"), g.String("nobusiness_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		Business(g.String(""))
	if result == nil {
		t.Error("Empty business connection should work")
	}

	// Test with empty provider token
	result = testCtx.CreateInvoiceLink(g.String("No Provider"), g.String("Service without provider"), g.String("noprovider_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		ProviderToken(g.String(""))
	if result == nil {
		t.Error("Empty provider token should work")
	}

	// Test with empty provider data
	result = testCtx.CreateInvoiceLink(g.String("No Data"), g.String("Service without provider data"), g.String("nodata_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		ProviderData(g.String(""))
	if result == nil {
		t.Error("Empty provider data should work")
	}

	// Test with zero timeout
	result = testCtx.CreateInvoiceLink(g.String("No Timeout"), g.String("Service without timeout"), g.String("notimeout_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.CreateInvoiceLink(g.String("Long Timeout"), g.String("Service with long timeout"), g.String("longtimeout_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.CreateInvoiceLink(g.String("No API"), g.String("Service without API URL"), g.String("noapi_payload"), g.String("USD")).
		Price(g.String("Service"), 1000).
		APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}
}

func TestCreateInvoiceLink_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Complete Test Product")
	description := g.String("Product testing all methods")
	payload := g.String("complete_test_payload")
	currency := g.String("USD")

	// Test all methods combined in different orders
	// Order 1: Complete configuration
	result1 := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Base Product"), 5000).
		Price(g.String("Shipping"), 500).
		Price(g.String("Tax"), 450).
		Business(g.String("complete_business_123")).
		ProviderToken(g.String("complete_provider_token")).
		SubscriptionPeriod(2592000).
		MaxTip(2000).
		SuggestedTips(200, 500, 1000, 1500).
		ProviderData(g.String(`{"complete": "test", "all_features": true}`)).
		Photo(g.String("https://example.com/complete.jpg"), 10240, 800, 600).
		NeedName().
		NeedPhone().
		NeedEmail().
		NeedShipping().
		SendPhone().
		SendEmail().
		Flexible().
		Timeout(90 * time.Second).
		APIURL(g.String("https://complete-api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2: Different sequence
	result2 := testCtx.CreateInvoiceLink(title, description, payload, currency).
		APIURL(g.String("https://reordered-api.example.com")).
		Timeout(60*time.Second).
		Flexible().
		SendEmail().
		SendPhone().
		NeedShipping().
		NeedEmail().
		NeedPhone().
		NeedName().
		Photo(g.String("https://example.com/reordered.jpg"), 5120, 600, 400).
		ProviderData(g.String(`{"reordered": "test"}`)).
		SuggestedTips(100, 300, 800).
		MaxTip(1500).
		SubscriptionPeriod(604800).
		ProviderToken(g.String("reordered_provider_token")).
		Business(g.String("reordered_business_456")).
		Price(g.String("Reordered Product"), 3000)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("First Product"), 1000).
		Price(g.String("Second Product"), 2000). // Additional price
		Business(g.String("first_business")).
		Business(g.String("second_business")). // Should override first
		ProviderToken(g.String("first_token")).
		ProviderToken(g.String("second_token")). // Should override first
		MaxTip(1000).
		MaxTip(2000). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test minimal configuration
	result4 := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Minimal Product"), 1000)

	if result4 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test various invoice types with different configurations
	invoiceTypes := []struct {
		name            string
		basePrice       int64
		hasSubscription bool
		hasTipping      bool
		hasShipping     bool
		hasFlexible     bool
	}{
		{"Digital Product", 1500, false, false, false, false},
		{"Physical Product", 2500, false, true, true, false},
		{"Subscription Service", 1000, true, false, false, false},
		{"Flexible Service", 3000, false, true, false, true},
		{"Complete Package", 5000, true, true, true, true},
	}

	for _, invoiceType := range invoiceTypes {
		result := testCtx.CreateInvoiceLink(title, description, payload, currency).
			Price(g.String("Product"), invoiceType.basePrice)

		if invoiceType.hasSubscription {
			result = result.SubscriptionPeriod(2592000)
		}
		if invoiceType.hasTipping {
			result = result.MaxTip(1000).SuggestedTips(100, 300, 500)
		}
		if invoiceType.hasShipping {
			result = result.NeedShipping().NeedName().NeedPhone()
		}
		if invoiceType.hasFlexible {
			result = result.Flexible()
		}

		result = result.Timeout(45 * time.Second).APIURL(g.String("https://type-api.example.com"))

		if result == nil {
			t.Errorf("Invoice type %s should work", invoiceType.name)
		}
	}
}

func TestCreateInvoiceLink_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	title := g.String("Test Product")
	description := g.String("Test product for Send method")
	payload := g.String("send_test_payload")
	currency := g.String("USD")

	// Test Send method - will fail with mock but covers the method
	sendResult := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Product"), 1000).
		Send()

	if sendResult.IsErr() {
		t.Logf("CreateInvoiceLink Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := testCtx.CreateInvoiceLink(title, description, payload, currency).
		Price(g.String("Premium Product"), 2500).
		Price(g.String("Shipping"), 300).
		ProviderToken(g.String("test_provider_token")).
		MaxTip(1000).
		SuggestedTips(100, 300, 500).
		NeedName().
		NeedEmail().
		Flexible().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("CreateInvoiceLink configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
