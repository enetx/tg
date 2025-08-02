package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CloseForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123456)

	// Test basic creation
	result := testCtx.CloseForumTopic(messageThreadID)
	if result == nil {
		t.Error("Expected CloseForumTopic builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return CloseForumTopic for chaining")
	}

	// Test Timeout method
	result = testCtx.CloseForumTopic(messageThreadID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CloseForumTopic for chaining")
	}

	// Test APIURL method
	result = testCtx.CloseForumTopic(messageThreadID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CloseForumTopic for chaining")
	}
}

func TestContext_CloseForumTopicChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(987654)

	// Test complete method chaining
	result := testCtx.CloseForumTopic(messageThreadID).
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CloseForumTopic")
	}
}

func TestCloseForumTopic_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero MessageThreadID
	result := testCtx.CloseForumTopic(0)
	if result == nil {
		t.Error("CloseForumTopic should handle zero MessageThreadID")
	}

	// Test with negative MessageThreadID
	result = testCtx.CloseForumTopic(-123456)
	if result == nil {
		t.Error("CloseForumTopic should handle negative MessageThreadID")
	}

	// Test with maximum MessageThreadID
	result = testCtx.CloseForumTopic(9223372036854775807)
	if result == nil {
		t.Error("CloseForumTopic should handle maximum MessageThreadID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.CloseForumTopic(123456).ChatID(0)
	if result == nil {
		t.Error("CloseForumTopic should handle zero ChatID")
	}

	// Test with zero timeout
	result = testCtx.CloseForumTopic(123456).Timeout(0 * time.Second)
	if result == nil {
		t.Error("CloseForumTopic should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.CloseForumTopic(123456).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("CloseForumTopic should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.CloseForumTopic(123456).APIURL(g.String(""))
	if result == nil {
		t.Error("CloseForumTopic should handle empty API URL")
	}
}

func TestCloseForumTopic_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(555666)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.CloseForumTopic(messageThreadID)
	if result == nil {
		t.Error("CloseForumTopic should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestCloseForumTopic_MessageThreadIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various message thread ID scenarios
	messageThreadIDs := []int64{
		1,                   // Minimal topic ID
		2,                   // General topic (often used as default)
		100,                 // Small topic ID
		123456789,           // Medium topic ID
		999999999999,        // Large topic ID
		9223372036854775807, // Maximum int64
	}

	for _, threadID := range messageThreadIDs {
		result := testCtx.CloseForumTopic(threadID)
		if result == nil {
			t.Errorf("CloseForumTopic should handle message thread ID: %d", threadID)
		}

		// Test chaining for each thread ID
		chainedResult := result.ChatID(-1001987654321).Timeout(20 * time.Second)
		if chainedResult == nil {
			t.Errorf("Chaining should work for message thread ID: %d", threadID)
		}
	}
}

func TestCloseForumTopic_ChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(777888)

	// Test various chat ID scenarios (forum groups must be supergroups)
	chatIDs := []int64{
		-1001234567890, // Supergroup with forum enabled
		-1001987654321, // Another supergroup
		-1002000000000, // Large supergroup ID
		-1009999999999, // Very large supergroup ID
	}

	for _, chatID := range chatIDs {
		result := testCtx.CloseForumTopic(messageThreadID).ChatID(chatID)
		if result == nil {
			t.Errorf("CloseForumTopic should handle chat ID: %d", chatID)
		}

		// Test combining with other methods
		combinedResult := result.Timeout(15 * time.Second).APIURL(g.String("https://api.telegram.org"))
		if combinedResult == nil {
			t.Errorf("Combined methods should work for chat ID: %d", chatID)
		}
	}
}

func TestCloseForumTopic_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(999000)

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.CloseForumTopic(messageThreadID).Timeout(timeout)
		if result == nil {
			t.Errorf("CloseForumTopic should handle timeout: %v", timeout)
		}

		// Test combining timeout with other methods
		combinedResult := result.ChatID(-1001111222333).APIURL(g.String("https://custom.api.com"))
		if combinedResult == nil {
			t.Errorf("Timeout combination should work for: %v", timeout)
		}
	}
}

func TestCloseForumTopic_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageThreadID := int64(111222)

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-api.example.com",
		"https://bot-api.mycompany.com",
		"https://localhost:8080",
		"https://api-staging.telegram.org",
		"https://proxy.telegram-api.com",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.CloseForumTopic(messageThreadID).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("CloseForumTopic should handle API URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combinedResult := result.ChatID(-1001444555666).Timeout(35 * time.Second)
		if combinedResult == nil {
			t.Errorf("API URL combination should work for: %s", apiURL)
		}
	}
}

func TestCloseForumTopic_ForumScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test different forum topic scenarios
	testCases := []struct {
		name        string
		threadID    int64
		description string
	}{
		{"General Topic", 1, "General forum topic (special case)"},
		{"Regular Topic", 123, "Regular forum topic"},
		{"High Activity Topic", 456789, "Topic with high message count"},
		{"Recent Topic", 999888777, "Recently created topic"},
		{"System Topic", 2, "System/admin topic"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testCtx.CloseForumTopic(tc.threadID)
			if result == nil {
				t.Errorf("CloseForumTopic should handle %s (%s): %d", tc.name, tc.description, tc.threadID)
			}

			// Test complete workflow for each scenario
			completedResult := result.
				ChatID(-1001234567890).
				Timeout(30 * time.Second).
				APIURL(g.String("https://api.telegram.org"))

			if completedResult == nil {
				t.Errorf("Complete workflow should work for %s (%s)", tc.name, tc.description)
			}
		})
	}
}
