package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateChatSubscriptionInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	subscriptionPeriod := int64(2592000) // 30 days in seconds
	subscriptionPrice := int64(100)      // 100 stars

	// Test basic creation
	result := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice)
	if result == nil {
		t.Error("Expected CreateChatSubscriptionInviteLink builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return CreateChatSubscriptionInviteLink for chaining")
	}

	// Test Name method
	result = testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).Name(g.String("Premium Subscription"))
	if result == nil {
		t.Error("Name method should return CreateChatSubscriptionInviteLink for chaining")
	}

	// Test Timeout method
	result = testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CreateChatSubscriptionInviteLink for chaining")
	}

	// Test APIURL method
	result = testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CreateChatSubscriptionInviteLink for chaining")
	}
}

func TestContext_CreateChatSubscriptionInviteLinkChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	subscriptionPeriod := int64(2592000) // 30 days
	subscriptionPrice := int64(150)      // 150 stars

	// Test complete method chaining
	result := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
		ChatID(-1001987654321).
		Name(g.String("VIP Subscription Access")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CreateChatSubscriptionInviteLink")
	}
}

func TestCreateChatSubscriptionInviteLink_SubscriptionPeriods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)
	subscriptionPrice := int64(100)

	// Test various subscription periods (in seconds)
	subscriptionPeriods := []struct {
		name        string
		period      int64
		description string
	}{
		{"Weekly", 604800, "7 days subscription"},     // 7 * 24 * 60 * 60
		{"Monthly", 2592000, "30 days subscription"},   // 30 * 24 * 60 * 60
		{"Quarterly", 7776000, "90 days subscription"}, // 90 * 24 * 60 * 60
		{"Yearly", 31536000, "365 days subscription"},  // 365 * 24 * 60 * 60
		{"Custom 14 Days", 1209600, "14 days subscription"},   // 14 * 24 * 60 * 60
		{"Custom 6 Months", 15552000, "180 days subscription"}, // 180 * 24 * 60 * 60
	}

	for _, period := range subscriptionPeriods {
		t.Run(period.name, func(t *testing.T) {
			result := testCtx.CreateChatSubscriptionInviteLink(period.period, subscriptionPrice).
				ChatID(chatID).
				Name(g.String(period.description))

			if result == nil {
				t.Errorf("%s (%s) should work", period.name, period.description)
			}

			// Test with additional options
			enhancedResult := result.
				Timeout(60 * time.Second).
				APIURL(g.String("https://subscription-api.example.com"))

			if enhancedResult == nil {
				t.Errorf("Enhanced %s should work", period.name)
			}
		})
	}
}

func TestCreateChatSubscriptionInviteLink_SubscriptionPrices(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)
	subscriptionPeriod := int64(2592000) // 30 days

	// Test various subscription prices (in stars)
	subscriptionPrices := []struct {
		name        string
		price       int64
		description string
	}{
		{"Basic", 50, "Basic subscription - 50 stars"},
		{"Standard", 100, "Standard subscription - 100 stars"},
		{"Premium", 200, "Premium subscription - 200 stars"},
		{"VIP", 500, "VIP subscription - 500 stars"},
		{"Enterprise", 1000, "Enterprise subscription - 1000 stars"},
		{"Ultra", 2500, "Ultra subscription - 2500 stars"},
		{"Minimum", 1, "Minimum subscription - 1 star"},
		{"High Value", 10000, "High value subscription - 10000 stars"},
	}

	for _, price := range subscriptionPrices {
		t.Run(price.name, func(t *testing.T) {
			result := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, price.price).
				ChatID(chatID).
				Name(g.String(price.description))

			if result == nil {
				t.Errorf("%s (%s) should work", price.name, price.description)
			}

			// Test with timeout for each price tier
			timedResult := result.Timeout(30 * time.Second)
			if timedResult == nil {
				t.Errorf("Timed %s should work", price.name)
			}
		})
	}
}

func TestCreateChatSubscriptionInviteLink_SubscriptionNames(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)
	subscriptionPeriod := int64(2592000) // 30 days
	subscriptionPrice := int64(100)

	// Test various subscription link names (0-32 characters)
	subscriptionNames := []string{
		"Premium",
		"VIP Access",
		"Monthly Subscription",
		"Elite Members Only",
		"Pro Subscription 2024",
		"Advanced Features Access",
		"Executive Premium Plan",
		"🌟 Premium Star Membership",
		"Special Offer - Limited Time",
		"Enterprise Business Package",
		"Ultra VIP Exclusive Access",
		"Developer Pro Subscription",
		"", // Empty name (should work)
		"A", // Single character
		"This is exactly 32 characters!!", // Max length (32 chars)
	}

	for _, name := range subscriptionNames {
		displayName := name
		if name == "" {
			displayName = "[empty]"
		}

		t.Run(displayName, func(t *testing.T) {
			result := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
				ChatID(chatID).
				Name(g.String(name))

			if result == nil {
				t.Errorf("Subscription name '%s' should work", displayName)
			}

			// Test with additional options
			combinedResult := result.
				Timeout(45 * time.Second).
				APIURL(g.String("https://api.telegram.org"))

			if combinedResult == nil {
				t.Errorf("Combined options with name '%s' should work", displayName)
			}
		})
	}
}

func TestCreateChatSubscriptionInviteLink_SubscriptionCombinations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test realistic subscription combinations
	subscriptionCombos := []struct {
		name   string
		period int64
		price  int64
		linkName string
		description string
	}{
		{"Weekly Basic", 604800, 25, "Weekly Basic", "Weekly basic subscription"},
		{"Monthly Standard", 2592000, 100, "Monthly Standard", "Monthly standard plan"},
		{"Quarterly Premium", 7776000, 250, "Quarterly Premium", "Quarterly premium access"},
		{"Yearly VIP", 31536000, 800, "Yearly VIP", "Annual VIP membership"},
		{"Trial 3 Days", 259200, 10, "Trial Access", "3-day trial subscription"},
		{"Enterprise Annual", 31536000, 5000, "Enterprise Plan", "Enterprise annual plan"},
		{"Student Monthly", 2592000, 50, "Student Discount", "Student monthly discount"},
		{"Professional Bi-Annual", 15552000, 400, "Pro Bi-Annual", "Professional 6-month plan"},
	}

	for _, combo := range subscriptionCombos {
		t.Run(combo.name, func(t *testing.T) {
			result := testCtx.CreateChatSubscriptionInviteLink(combo.period, combo.price).
				ChatID(chatID).
				Name(g.String(combo.linkName))

			if result == nil {
				t.Errorf("%s (%s) should work", combo.name, combo.description)
			}

			// Test complete workflow for each combination
			completedResult := result.
				Timeout(60 * time.Second).
				APIURL(g.String("https://subscription-api.telegram.org"))

			if completedResult == nil {
				t.Errorf("Complete %s workflow should work", combo.name)
			}
		})
	}
}

func TestCreateChatSubscriptionInviteLink_ChatTypes(t *testing.T) {
	bot := &mockBot{}
	subscriptionPeriod := int64(2592000) // 30 days
	subscriptionPrice := int64(100)

	// Test subscription links for various chat types
	chatTypes := []struct {
		name   string
		chatID int64
		type_  string
	}{
		{"Supergroup", -1001234567890, "supergroup"},
		{"Large Supergroup", -1002000000000, "supergroup"},
		{"Channel", -1001987654321, "channel"},
		{"Enterprise Channel", -1003000000000, "channel"},
	}

	for _, chatType := range chatTypes {
		t.Run(chatType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat: &gotgbot.Chat{Id: chatType.chatID, Type: chatType.type_},
				Update:        &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
				ChatID(chatType.chatID).
				Name(g.String("Subscription for " + chatType.name))

			if result == nil {
				t.Errorf("CreateChatSubscriptionInviteLink should work for %s", chatType.name)
			}

			// Test with full features for each chat type
			fullResult := result.
				Timeout(90 * time.Second).
				APIURL(g.String("https://subscription-api.example.com"))

			if fullResult == nil {
				t.Errorf("Full features should work for %s", chatType.name)
			}
		})
	}
}

func TestCreateChatSubscriptionInviteLink_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero values
	result := testCtx.CreateChatSubscriptionInviteLink(0, 0)
	if result == nil {
		t.Error("Zero period and price should work (builder creation)")
	}

	// Test with very small values
	result = testCtx.CreateChatSubscriptionInviteLink(1, 1)
	if result == nil {
		t.Error("Minimum period and price should work")
	}

	// Test with very large values
	result = testCtx.CreateChatSubscriptionInviteLink(31536000000, 1000000) // ~1000 years, 1M stars
	if result == nil {
		t.Error("Very large period and price should work")
	}

	// Test with zero chat ID (should use effective chat)
	result = testCtx.CreateChatSubscriptionInviteLink(2592000, 100).ChatID(0)
	if result == nil {
		t.Error("Zero chat ID should work")
	}

	// Test with zero timeout
	result = testCtx.CreateChatSubscriptionInviteLink(2592000, 100).Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.CreateChatSubscriptionInviteLink(2592000, 100).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.CreateChatSubscriptionInviteLink(2592000, 100).APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without ChatID (should use effective chat)
	result = testCtx.CreateChatSubscriptionInviteLink(2592000, 100).Name(g.String("Default Chat Subscription"))
	if result == nil {
		t.Error("CreateChatSubscriptionInviteLink should work without explicit ChatID")
	}
}

func TestCreateChatSubscriptionInviteLink_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001987654321)
	subscriptionPeriod := int64(2592000) // 30 days
	subscriptionPrice := int64(250)      // 250 stars

	// Test all methods combined in different orders
	// Order 1
	result1 := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
		ChatID(chatID).
		Name(g.String("Complete Test Subscription")).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2 (different sequence)
	result2 := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
		APIURL(g.String("https://custom-subscription-api.example.com")).
		Timeout(45 * time.Second).
		Name(g.String("Reordered Test Subscription")).
		ChatID(chatID)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
		ChatID(chatID).
		ChatID(-1002000000000). // Should override first
		Name(g.String("First Name")).
		Name(g.String("Second Name")). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test minimal configuration
	result4 := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice)
	if result4 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test with just name
	result5 := testCtx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice).
		Name(g.String("Simple Subscription"))
	if result5 == nil {
		t.Error("Simple name configuration should work")
	}

	// Test various period/price combinations with different method chains
	combinations := []struct {
		period int64
		price  int64
		name   string
	}{
		{604800, 50, "Weekly Basic"},     // 7 days, 50 stars
		{2592000, 100, "Monthly Standard"}, // 30 days, 100 stars
		{7776000, 300, "Quarterly Premium"}, // 90 days, 300 stars
		{31536000, 1000, "Annual VIP"},      // 365 days, 1000 stars
	}

	for _, combo := range combinations {
		result := testCtx.CreateChatSubscriptionInviteLink(combo.period, combo.price).
			ChatID(chatID).
			Name(g.String(combo.name)).
			Timeout(30 * time.Second).
			APIURL(g.String("https://combo-api.example.com"))

		if result == nil {
			t.Errorf("Combination %s should work", combo.name)
		}
	}
}
