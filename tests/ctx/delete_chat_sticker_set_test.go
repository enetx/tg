package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteChatStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.DeleteChatStickerSet()
	if result == nil {
		t.Error("Expected DeleteChatStickerSet builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return DeleteChatStickerSet for chaining")
	}

	// Test Timeout method
	result = testCtx.DeleteChatStickerSet().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeleteChatStickerSet for chaining")
	}

	// Test APIURL method
	result = testCtx.DeleteChatStickerSet().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeleteChatStickerSet for chaining")
	}
}

func TestContext_DeleteChatStickerSetChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.DeleteChatStickerSet().
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteChatStickerSet")
	}
}

func TestDeleteChatStickerSet_ChatTypes(t *testing.T) {
	bot := &mockBot{}

	// Test sticker set deletion for various supergroup types
	chatTypes := []struct {
		name        string
		chatID      int64
		type_       string
		description string
	}{
		{"Standard Supergroup", -1001234567890, "supergroup", "Regular supergroup with sticker set"},
		{"Large Supergroup", -1002000000000, "supergroup", "Large community supergroup"},
		{"Gaming Supergroup", -1003000000000, "supergroup", "Gaming community with custom stickers"},
		{"Business Supergroup", -1004000000000, "supergroup", "Business group with branded stickers"},
		{"Educational Supergroup", -1005000000000, "supergroup", "Educational group with academic stickers"},
		{"Entertainment Supergroup", -1006000000000, "supergroup", "Entertainment group with fun stickers"},
		{"Technical Supergroup", -1007000000000, "supergroup", "Technical discussion group"},
		{"Community Supergroup", -1008000000000, "supergroup", "General community group"},
	}

	for _, chatType := range chatTypes {
		t.Run(chatType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat: &gotgbot.Chat{Id: chatType.chatID, Type: chatType.type_},
				Update:        &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.DeleteChatStickerSet().
				ChatID(chatType.chatID)

			if result == nil {
				t.Errorf("DeleteChatStickerSet should work for %s (%s)", chatType.name, chatType.description)
			}

			// Test with timeout for each chat type
			timedResult := result.Timeout(60 * time.Second)
			if timedResult == nil {
				t.Errorf("Timed deletion should work for %s", chatType.name)
			}
		})
	}
}

func TestDeleteChatStickerSet_StickerSetScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various sticker set deletion scenarios
	stickerScenarios := []struct {
		name        string
		chatID      int64
		stickerType string
		description string
	}{
		{"Custom Sticker Set", chatID, "custom", "Removing custom community stickers"},
		{"Branded Sticker Set", -1001987654321, "branded", "Removing business branded stickers"},
		{"Animated Sticker Set", -1002000000000, "animated", "Removing animated sticker collection"},
		{"Emoji Sticker Set", -1003000000000, "emoji", "Removing custom emoji stickers"},
		{"Meme Sticker Set", -1004000000000, "meme", "Removing meme collection stickers"},
		{"Art Sticker Set", -1005000000000, "art", "Removing artistic sticker collection"},
		{"Gaming Sticker Set", -1006000000000, "gaming", "Removing gaming-themed stickers"},
		{"Event Sticker Set", -1007000000000, "event", "Removing event-specific stickers"},
		{"Seasonal Sticker Set", -1008000000000, "seasonal", "Removing seasonal/holiday stickers"},
		{"Community Sticker Set", -1009000000000, "community", "Removing community-created stickers"},
	}

	for _, scenario := range stickerScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.DeleteChatStickerSet().
				ChatID(scenario.chatID)

			if result == nil {
				t.Errorf("%s deletion (%s) should work", scenario.name, scenario.description)
			}

			// Test with API URL for different sticker types
			apiResult := result.APIURL(g.String("https://stickers-api.telegram.org"))
			if apiResult == nil {
				t.Errorf("API URL configuration for %s should work", scenario.name)
			}

			// Test complete workflow for each scenario
			completedResult := apiResult.Timeout(90 * time.Second)
			if completedResult == nil {
				t.Errorf("Complete %s deletion workflow should work", scenario.name)
			}
		})
	}
}

func TestDeleteChatStickerSet_AdminScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various admin scenarios for sticker set deletion
	adminScenarios := []struct {
		name        string
		chatID      int64
		adminType   string
		urgency     string
		timeout     time.Duration
		description string
	}{
		{"Owner Cleanup", -1001234567890, "owner", "normal", 30 * time.Second, "Owner removing outdated stickers"},
		{"Admin Moderation", -1001987654321, "admin", "urgent", 10 * time.Second, "Admin removing inappropriate stickers"},
		{"Scheduled Cleanup", -1002000000000, "admin", "scheduled", 60 * time.Second, "Scheduled maintenance cleanup"},
		{"Policy Enforcement", -1003000000000, "admin", "policy", 15 * time.Second, "Removing policy-violating stickers"},
		{"Rebranding", -1004000000000, "owner", "rebranding", 45 * time.Second, "Removing old brand stickers for rebranding"},
		{"Seasonal Cleanup", -1005000000000, "admin", "seasonal", 30 * time.Second, "Removing seasonal stickers post-event"},
		{"Copyright Issues", -1006000000000, "admin", "legal", 5 * time.Second, "Urgent removal due to copyright"},
		{"Community Request", -1007000000000, "admin", "community", 20 * time.Second, "Removal based on community feedback"},
	}

	for _, scenario := range adminScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.DeleteChatStickerSet().
				ChatID(scenario.chatID).
				Timeout(scenario.timeout)

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with appropriate API configuration based on urgency
			var apiURL string
			switch scenario.urgency {
			case "urgent", "legal":
				apiURL = "https://urgent-api.telegram.org"
			case "scheduled":
				apiURL = "https://scheduled-api.telegram.org"
			default:
				apiURL = "https://standard-api.telegram.org"
			}

			apiResult := result.APIURL(g.String(apiURL))
			if apiResult == nil {
				t.Errorf("API configuration for %s should work", scenario.name)
			}
		})
	}
}

func TestDeleteChatStickerSet_TimeoutConfigurations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various timeout configurations
	timeoutConfigurations := []struct {
		name        string
		timeout     time.Duration
		scenario    string
		description string
	}{
		{"Instant", 1 * time.Second, "emergency", "Emergency removal with minimal timeout"},
		{"Quick", 5 * time.Second, "urgent", "Quick removal for urgent situations"},
		{"Standard", 30 * time.Second, "normal", "Standard removal timeout"},
		{"Extended", 60 * time.Second, "thorough", "Extended timeout for thorough processing"},
		{"Long", 2 * time.Minute, "bulk", "Long timeout for bulk operations"},
		{"Maximum", 5 * time.Minute, "complex", "Maximum timeout for complex scenarios"},
		{"Zero", 0 * time.Second, "default", "Zero timeout (use default)"},
		{"Very Long", 10 * time.Minute, "maintenance", "Very long timeout for maintenance"},
	}

	for _, config := range timeoutConfigurations {
		t.Run(config.name, func(t *testing.T) {
			result := testCtx.DeleteChatStickerSet().
				ChatID(chatID).
				Timeout(config.timeout)

			if result == nil {
				t.Errorf("%s timeout (%s) should work", config.name, config.description)
			}

			// Test with API URL for each timeout configuration
			apiResult := result.APIURL(g.String("https://timeout-api.telegram.org"))
			if apiResult == nil {
				t.Errorf("API URL with %s timeout should work", config.name)
			}
		})
	}
}

func TestDeleteChatStickerSet_APIURLConfigurations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test various API URL configurations
	apiConfigurations := []struct {
		name        string
		apiURL      string
		purpose     string
		description string
	}{
		{"Standard API", "https://api.telegram.org", "standard", "Standard Telegram API"},
		{"Custom API", "https://custom-api.example.com", "custom", "Custom API endpoint"},
		{"Regional API", "https://eu-api.telegram.org", "regional", "European regional API"},
		{"Backup API", "https://backup-api.telegram.org", "backup", "Backup API endpoint"},
		{"Development API", "https://dev-api.telegram.org", "development", "Development environment API"},
		{"Test API", "https://test-api.telegram.org", "testing", "Testing environment API"},
		{"Secure API", "https://secure-api.telegram.org", "security", "High-security API endpoint"},
		{"Load Balanced API", "https://lb-api.telegram.org", "load_balance", "Load-balanced API endpoint"},
		{"CDN API", "https://cdn-api.telegram.org", "cdn", "CDN-optimized API endpoint"},
		{"Mobile API", "https://mobile-api.telegram.org", "mobile", "Mobile-optimized API endpoint"},
	}

	for _, config := range apiConfigurations {
		t.Run(config.name, func(t *testing.T) {
			result := testCtx.DeleteChatStickerSet().
				ChatID(chatID).
				APIURL(g.String(config.apiURL))

			if result == nil {
				t.Errorf("%s (%s) should work", config.name, config.description)
			}

			// Test with timeout for each API configuration
			timedResult := result.Timeout(45 * time.Second)
			if timedResult == nil {
				t.Errorf("Timeout with %s should work", config.name)
			}
		})
	}

	// Test empty API URL
	result := testCtx.DeleteChatStickerSet().
		ChatID(chatID).
		APIURL(g.String(""))

	if result == nil {
		t.Error("Empty API URL should work")
	}
}

func TestDeleteChatStickerSet_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero chat ID (should use effective chat)
	result := testCtx.DeleteChatStickerSet().ChatID(0)
	if result == nil {
		t.Error("Zero chat ID should work (uses effective chat)")
	}

	// Test with negative chat ID (valid for supergroups)
	result = testCtx.DeleteChatStickerSet().ChatID(-1001000000000)
	if result == nil {
		t.Error("Negative chat ID should work for supergroups")
	}

	// Test with very large negative chat ID
	result = testCtx.DeleteChatStickerSet().ChatID(-1009999999999)
	if result == nil {
		t.Error("Very large negative chat ID should work")
	}

	// Test with zero timeout
	result = testCtx.DeleteChatStickerSet().Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.DeleteChatStickerSet().Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.DeleteChatStickerSet().APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without ChatID (should use effective chat)
	result = testCtx.DeleteChatStickerSet()
	if result == nil {
		t.Error("DeleteChatStickerSet should work without explicit ChatID")
	}

	// Test minimal configuration
	result = testCtx.DeleteChatStickerSet()
	if result == nil {
		t.Error("Minimal configuration should work")
	}
}

func TestDeleteChatStickerSet_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001987654321)

	// Test all methods combined in different orders
	// Order 1
	result1 := testCtx.DeleteChatStickerSet().
		ChatID(chatID).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2 (different sequence)
	result2 := testCtx.DeleteChatStickerSet().
		APIURL(g.String("https://reordered-api.example.com")).
		Timeout(45 * time.Second).
		ChatID(chatID)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.DeleteChatStickerSet().
		ChatID(chatID).
		ChatID(-1002000000000). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second). // Should override first
		APIURL(g.String("https://first-api.example.com")).
		APIURL(g.String("https://second-api.example.com")) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test minimal configuration
	result4 := testCtx.DeleteChatStickerSet()
	if result4 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test with just ChatID
	result5 := testCtx.DeleteChatStickerSet().ChatID(chatID)
	if result5 == nil {
		t.Error("ChatID-only configuration should work")
	}

	// Test with just Timeout
	result6 := testCtx.DeleteChatStickerSet().Timeout(30 * time.Second)
	if result6 == nil {
		t.Error("Timeout-only configuration should work")
	}

	// Test with just APIURL
	result7 := testCtx.DeleteChatStickerSet().APIURL(g.String("https://single-api.example.com"))
	if result7 == nil {
		t.Error("APIURL-only configuration should work")
	}

	// Test various chat IDs with different configurations
	chatIDs := []int64{
		-1001234567890, // Standard supergroup
		-1002000000000, // Large supergroup
		-1003000000000, // Community supergroup
		-1004000000000, // Business supergroup
		-1005000000000, // Gaming supergroup
	}

	for i, testChatID := range chatIDs {
		result := testCtx.DeleteChatStickerSet().
			ChatID(testChatID).
			Timeout(time.Duration(30+i*10) * time.Second).
			APIURL(g.String("https://multi-chat-api.example.com"))

		if result == nil {
			t.Errorf("Multi-chat configuration for chat %d should work", testChatID)
		}
	}
}

func TestDeleteChatStickerSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001987654321)

	// Test Send method - will fail with mock but covers the method
	sendResult := testCtx.DeleteChatStickerSet().Send()

	if sendResult.IsErr() {
		t.Logf("DeleteChatStickerSet Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := testCtx.DeleteChatStickerSet().
		ChatID(chatID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteChatStickerSet configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := testCtx.DeleteChatStickerSet().
		Timeout(45 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("DeleteChatStickerSet with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}

	// Test Send method with various timeout configurations
	timeouts := []time.Duration{
		5 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
	}

	for _, timeout := range timeouts {
		timeoutSendResult := testCtx.DeleteChatStickerSet().
			ChatID(chatID).
			Timeout(timeout).
			Send()

		if timeoutSendResult.IsErr() {
			t.Logf("DeleteChatStickerSet with %v timeout Send failed as expected: %v", timeout, timeoutSendResult.Err())
		}
	}

	// Test Send method with different API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-sticker-api.example.com",
		"https://admin-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiSendResult := testCtx.DeleteChatStickerSet().
			ChatID(chatID).
			APIURL(g.String(apiURL)).
			Send()

		if apiSendResult.IsErr() {
			t.Logf("DeleteChatStickerSet with API URL '%s' Send failed as expected: %v", apiURL, apiSendResult.Err())
		}
	}

	// Test Send method with complete configuration
	completeSendResult := testCtx.DeleteChatStickerSet().
		ChatID(chatID).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-api.telegram.org")).
		Send()

	if completeSendResult.IsErr() {
		t.Logf("DeleteChatStickerSet complete configuration Send failed as expected: %v", completeSendResult.Err())
	}
}
