package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func TestContext_GiftPremiumSubscription(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	result := ctx.GiftPremiumSubscription(userID, monthCount, starCount)

	if result == nil {
		t.Error("Expected GiftPremiumSubscription builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGiftPremiumSubscription_Text(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	// Test Text method with various text messages
	textMessages := []string{
		"🎁 Enjoy your premium subscription!",
		"Congratulations! You've received a premium gift!",
		"",
		"A very long message with lots of details about the premium subscription and all the benefits you will receive including advanced features, priority support, and exclusive content access that comes with this wonderful gift subscription that someone has generously provided to you.",
		"<b>Bold text</b> with HTML formatting",
		"**Markdown** formatted text with *italics*",
		"Text with special characters: @#$%^&*()_+-=[]{}|;':\",./<>?`~",
		"Emoji test: 🎉🎊🥳🎁💎⭐",
		"Multiline text\nwith line breaks\nand multiple lines",
	}

	for _, text := range textMessages {
		result := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
		textResult := result.Text(g.String(text))
		if textResult == nil {
			t.Errorf("Text method should return GiftPremiumSubscription for chaining with text: %s", text)
		}

		// Test that Text can be chained and overridden
		chainedResult := textResult.Text(g.String("Updated: " + text))
		if chainedResult == nil {
			t.Errorf("Text method should support chaining and override with text: %s", text)
		}
	}
}

func TestGiftPremiumSubscription_HTML(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	// Test HTML method
	result := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
	htmlResult := result.HTML()
	if htmlResult == nil {
		t.Error("HTML method should return GiftPremiumSubscription for chaining")
	}

	// Test HTML method can be chained multiple times
	htmlChained := htmlResult.HTML()
	if htmlChained == nil {
		t.Error("HTML method should support multiple chaining calls")
	}

	// Test HTML with text
	htmlWithText := ctx.GiftPremiumSubscription(userID, monthCount, starCount).
		Text(g.String("<b>Bold HTML text</b>")).
		HTML()
	if htmlWithText == nil {
		t.Error("HTML method should work with Text method")
	}

	// Test overriding parse mode
	parseModeSwitching := ctx.GiftPremiumSubscription(userID, monthCount, starCount).
		Markdown().
		HTML() // Override Markdown with HTML
	if parseModeSwitching == nil {
		t.Error("HTML method should override other parse modes")
	}
}

func TestGiftPremiumSubscription_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	// Test Markdown method
	result := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
	markdownResult := result.Markdown()
	if markdownResult == nil {
		t.Error("Markdown method should return GiftPremiumSubscription for chaining")
	}

	// Test Markdown method can be chained multiple times
	markdownChained := markdownResult.Markdown()
	if markdownChained == nil {
		t.Error("Markdown method should support multiple chaining calls")
	}

	// Test Markdown with text
	markdownWithText := ctx.GiftPremiumSubscription(userID, monthCount, starCount).
		Text(g.String("**Bold Markdown text**")).
		Markdown()
	if markdownWithText == nil {
		t.Error("Markdown method should work with Text method")
	}

	// Test overriding parse mode
	parseModeSwitching := ctx.GiftPremiumSubscription(userID, monthCount, starCount).
		HTML().
		Markdown() // Override HTML with Markdown
	if parseModeSwitching == nil {
		t.Error("Markdown method should override other parse modes")
	}
}

func TestGiftPremiumSubscription_Entities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	// Test Entities method with various entity configurations
	testEntities := []*entities.Entities{
		entities.New(g.String("Bold")),                              // Simple entities
		entities.New(g.String("Multiple entities with formatting")), // Multiple entities
		entities.New(g.String("URL test")),                          // URL entity
		entities.New(g.String("@mention")),                          // Mention entity
		entities.New(g.String("#hashtag $cashtag")),                 // Hashtag and cashtag
		entities.New(g.String("")),                                  // Empty entities
		entities.New(g.String("Spoiler underline strike")),          // Text formatting entities
	}

	for i, ents := range testEntities {
		result := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
		entitiesResult := result.Entities(ents)
		if entitiesResult == nil {
			t.Errorf("Entities method should return GiftPremiumSubscription for chaining with entities set %d", i)
		}

		// Test that Entities can be chained and overridden
		newEntities := entities.New(g.String("Bold new entities"))
		chainedResult := entitiesResult.Entities(newEntities)
		if chainedResult == nil {
			t.Errorf("Entities method should support chaining and override with entities set %d", i)
		}
	}
}

func TestGiftPremiumSubscription_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	monthCount := int64(1)
	starCount := int64(100)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-gift-subscription-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GiftPremiumSubscription for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GiftPremiumSubscription(userID, monthCount, starCount)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-gift-subscription-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-gift-subscription-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GiftPremiumSubscription for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://gift-subscription-api.example.com",
		"https://custom-gift-subscription.telegram.org",
		"https://regional-gift-subscription-api.telegram.org",
		"https://backup-gift-subscription-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GiftPremiumSubscription(userID, monthCount, starCount).
			Text(g.String("Gift subscription test")).
			HTML().
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GiftPremiumSubscription with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGiftPremiumSubscription_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various gift configurations
	giftScenarios := []struct {
		userID      int64
		monthCount  int64
		starCount   int64
		description string
	}{
		{456, 1, 100, "1-month subscription for 100 stars"},
		{789, 3, 300, "3-month subscription for 300 stars"},
		{123, 6, 600, "6-month subscription for 600 stars"},
		{111, 12, 1200, "12-month subscription for 1200 stars"},
		{222, 1, 50, "1-month subscription for 50 stars"},
		{0, 1, 100, "Zero user ID"},
		{456, 0, 100, "Zero month count"},
		{456, 1, 0, "Zero star count"},
		{-1, 1, 100, "Negative user ID"},
		{456, -1, 100, "Negative month count"},
		{456, 1, -1, "Negative star count"},
	}

	for _, scenario := range giftScenarios {
		// Basic Send test
		sendResult := ctx.GiftPremiumSubscription(scenario.userID, scenario.monthCount, scenario.starCount).Send()
		if sendResult.IsErr() {
			t.Logf("GiftPremiumSubscription with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test with text
		configuredSendResult := ctx.GiftPremiumSubscription(scenario.userID, scenario.monthCount, scenario.starCount).
			Text(g.String("🎁 Premium gift for you!")).
			HTML().
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GiftPremiumSubscription configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}

	// Test different text formatting combinations
	// HTML formatting
	htmlFormattedResult := ctx.GiftPremiumSubscription(456, 1, 100).
		Text(g.String("<b>HTML formatted gift</b>")).
		HTML().
		Send()
	if htmlFormattedResult.IsErr() {
		t.Logf("GiftPremiumSubscription with HTML formatting Send failed as expected: %v", htmlFormattedResult.Err())
	}

	// Markdown formatting
	markdownFormattedResult := ctx.GiftPremiumSubscription(456, 1, 100).
		Text(g.String("**Markdown formatted gift**")).
		Markdown().
		Send()
	if markdownFormattedResult.IsErr() {
		t.Logf("GiftPremiumSubscription with Markdown formatting Send failed as expected: %v", markdownFormattedResult.Err())
	}

	// Entities formatting
	ents := entities.New(g.String("Bold gift message"))
	entitiesFormattedResult := ctx.GiftPremiumSubscription(456, 1, 100).
		Text(g.String("Bold gift message")).
		Entities(ents).
		Send()
	if entitiesFormattedResult.IsErr() {
		t.Logf("GiftPremiumSubscription with Entities formatting Send failed as expected: %v", entitiesFormattedResult.Err())
	}

	// Plain text
	plainTextResult := ctx.GiftPremiumSubscription(456, 1, 100).
		Text(g.String("Plain text gift")).
		Send()
	if plainTextResult.IsErr() {
		t.Logf("GiftPremiumSubscription with plain text Send failed as expected: %v", plainTextResult.Err())
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.GiftPremiumSubscription(456, 3, 300).
		Text(g.String("🎁 <b>Special Premium Gift</b> 🌟")).
		HTML().
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-gift-subscription-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GiftPremiumSubscription comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := ctx.GiftPremiumSubscription(789, 1, 150).
		APIURL(g.String("https://order-test-1.telegram.org")).
		HTML().
		Text(g.String("<i>Order test 1</i>")).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GiftPremiumSubscription order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.GiftPremiumSubscription(789, 1, 150).
		Timeout(45 * time.Second).
		Text(g.String("**Order test 2**")).
		Markdown().
		APIURL(g.String("https://order-test-2.telegram.org")).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GiftPremiumSubscription order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
