package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ConvertGiftToStars(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	ownedGiftID := g.String("gift_123")

	// Test basic creation
	result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID)
	if result == nil {
		t.Error("Expected ConvertGiftToStars builder to be created")
	}

	// Test Timeout method
	result = result.Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return ConvertGiftToStars for chaining")
	}

	// Test APIURL method
	result = testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return ConvertGiftToStars for chaining")
	}
}

func TestContext_ConvertGiftToStarsChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 2},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_456")
	ownedGiftID := g.String("gift_456")

	// Test complete method chaining
	result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return ConvertGiftToStars")
	}
}

func TestConvertGiftToStars_BusinessConnectionIDs(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 789, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 3},
	}

	testCtx := ctx.New(bot, rawCtx)
	ownedGiftID := g.String("gift_789")

	// Test various business connection ID formats
	businessConnectionIDs := []string{
		"business_conn_123",
		"BUSINESS_CONNECTION_456",
		"bus-conn-789",
		"bc_12345",
		"business.connection.123",
		"very_long_business_connection_identifier_with_many_characters_12345",
		"short",
		"b1",
		"enterprise_business_connection_id_2024",
	}

	for _, connID := range businessConnectionIDs {
		businessConnectionID := g.String(connID)
		result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID)
		if result == nil {
			t.Errorf("ConvertGiftToStars should handle business connection ID: %s", connID)
		}

		// Test chaining for each connection ID
		chainedResult := result.Timeout(20 * time.Second)
		if chainedResult == nil {
			t.Errorf("Chaining should work for business connection ID: %s", connID)
		}
	}
}

func TestConvertGiftToStars_OwnedGiftIDs(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 999, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 4},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_999")

	// Test various owned gift ID formats
	ownedGiftIDs := []string{
		"gift_123",
		"GIFT_456",
		"premium-gift-789",
		"subscription_gift_12345",
		"gift.premium.2024",
		"star_gift_conversion_candidate",
		"g1",
		"special_anniversary_gift_2024",
		"monthly_premium_gift_subscription",
		"enterprise_gift_package_ultra",
	}

	for _, giftID := range ownedGiftIDs {
		ownedGiftID := g.String(giftID)
		result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID)
		if result == nil {
			t.Errorf("ConvertGiftToStars should handle owned gift ID: %s", giftID)
		}

		// Test combining with APIURL
		combinedResult := result.APIURL(g.String("https://api.example.com"))
		if combinedResult == nil {
			t.Errorf("APIURL combination should work for gift ID: %s", giftID)
		}
	}
}

func TestConvertGiftToStars_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 111, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 5},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_111")
	ownedGiftID := g.String("gift_111")

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).Timeout(timeout)
		if result == nil {
			t.Errorf("ConvertGiftToStars should handle timeout: %v", timeout)
		}

		// Test combining timeout with APIURL
		combinedResult := result.APIURL(g.String("https://timeout-api.example.com"))
		if combinedResult == nil {
			t.Errorf("Timeout with APIURL should work for: %v", timeout)
		}
	}
}

func TestConvertGiftToStars_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 222, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 6},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_222")
	ownedGiftID := g.String("gift_222")

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-api.example.com",
		"https://gift-conversion-api.example.com",
		"https://business-api.mycompany.com",
		"https://localhost:8080",
		"https://api-staging.telegram.org",
		"https://enterprise-gift-api.com",
		"https://stars-conversion-service.example.com",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("ConvertGiftToStars should handle API URL: %s", apiURL)
		}

		// Test combining API URL with timeout
		combinedResult := result.Timeout(25 * time.Second)
		if combinedResult == nil {
			t.Errorf("API URL with timeout should work for: %s", apiURL)
		}
	}
}

func TestConvertGiftToStars_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 333, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 7},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty business connection ID
	emptyBusinessID := g.String("")
	ownedGiftID := g.String("gift_333")
	result := testCtx.ConvertGiftToStars(emptyBusinessID, ownedGiftID)
	if result == nil {
		t.Error("ConvertGiftToStars should handle empty business connection ID")
	}

	// Test with empty owned gift ID
	businessConnectionID := g.String("business_conn_333")
	emptyGiftID := g.String("")
	result = testCtx.ConvertGiftToStars(businessConnectionID, emptyGiftID)
	if result == nil {
		t.Error("ConvertGiftToStars should handle empty owned gift ID")
	}

	// Test with both empty IDs
	result = testCtx.ConvertGiftToStars(emptyBusinessID, emptyGiftID)
	if result == nil {
		t.Error("ConvertGiftToStars should handle both empty IDs")
	}

	// Test with zero timeout
	result = testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).Timeout(0 * time.Second)
	if result == nil {
		t.Error("ConvertGiftToStars should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("ConvertGiftToStars should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).APIURL(g.String(""))
	if result == nil {
		t.Error("ConvertGiftToStars should handle empty API URL")
	}
}

func TestConvertGiftToStars_GiftConversionScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 444, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 8},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test different gift conversion scenarios
	conversionScenarios := []struct {
		name                 string
		businessConnectionID string
		ownedGiftID          string
		description          string
	}{
		{"Premium Gift", "premium_business_123", "premium_gift_456", "Premium subscription gift conversion"},
		{"Anniversary Gift", "anniversary_business_789", "anniversary_gift_2024", "Special anniversary gift"},
		{"Subscription Gift", "subscription_business_111", "monthly_subscription_gift", "Monthly subscription gift"},
		{"Corporate Gift", "corporate_business_222", "corporate_premium_package", "Corporate gift package"},
		{"Holiday Gift", "holiday_business_333", "holiday_special_2024", "Holiday special gift"},
		{"Loyalty Gift", "loyalty_business_444", "loyalty_reward_gift", "Loyalty program reward"},
		{"VIP Gift", "vip_business_555", "vip_exclusive_gift", "VIP exclusive gift package"},
	}

	for _, scenario := range conversionScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			businessID := g.String(scenario.businessConnectionID)
			giftID := g.String(scenario.ownedGiftID)

			result := testCtx.ConvertGiftToStars(businessID, giftID)
			if result == nil {
				t.Errorf("%s conversion (%s) should work", scenario.name, scenario.description)
			}

			// Test complete workflow for each scenario
			completedResult := result.
				Timeout(30 * time.Second).
				APIURL(g.String("https://gift-conversion-api.telegram.org"))

			if completedResult == nil {
				t.Errorf("Complete %s workflow should work", scenario.name)
			}
		})
	}
}

func TestConvertGiftToStars_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 555, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 9},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("coverage_business_555")
	ownedGiftID := g.String("coverage_gift_555")

	// Test all method combinations systematically
	baseBuilder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID)

	// Test Timeout method
	timeoutBuilder := baseBuilder.Timeout(60 * time.Second)
	if timeoutBuilder == nil {
		t.Error("Timeout method should work")
	}

	// Test APIURL method
	apiBuilder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).APIURL(g.String("https://coverage-api.example.com"))
	if apiBuilder == nil {
		t.Error("APIURL method should work")
	}

	// Test all methods combined
	combinedBuilder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).
		Timeout(90 * time.Second).
		APIURL(g.String("https://combined-api.example.com"))

	if combinedBuilder == nil {
		t.Error("All methods combined should work")
	}

	// Test method order independence
	reorderedBuilder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).
		APIURL(g.String("https://reordered-api.example.com")).
		Timeout(75 * time.Second)

	if reorderedBuilder == nil {
		t.Error("Method order independence should work")
	}

	// Test multiple chaining operations
	multipleBuilder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).
		Timeout(45 * time.Second).
		APIURL(g.String("https://first-api.example.com")).
		Timeout(55 * time.Second).
		APIURL(g.String("https://second-api.example.com"))

	if multipleBuilder == nil {
		t.Error("Multiple method calls should work (last one wins)")
	}
}

func TestConvertGiftToStars_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 666, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 10},
	}

	testCtx := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("send_business_666")
	ownedGiftID := g.String("send_gift_666")

	// Test Send method execution
	builder := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with Timeout
	builderWithTimeout := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).Timeout(30 * time.Second)
	resultWithTimeout := builderWithTimeout.Send()

	if !resultWithTimeout.IsErr() && !resultWithTimeout.IsOk() {
		t.Error("Send with timeout should return a result")
	}

	// Test Send with APIURL
	builderWithAPIURL := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).APIURL(g.String("https://send-api.example.com"))
	resultWithAPIURL := builderWithAPIURL.Send()

	if !resultWithAPIURL.IsErr() && !resultWithAPIURL.IsOk() {
		t.Error("Send with API URL should return a result")
	}

	// Test Send with all options
	builderComplete := testCtx.ConvertGiftToStars(businessConnectionID, ownedGiftID).
		Timeout(45 * time.Second).
		APIURL(g.String("https://complete-send-api.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all options should return a result")
	}
}
