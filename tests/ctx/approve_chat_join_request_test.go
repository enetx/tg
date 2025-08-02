package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ApproveChatJoinRequest(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(456789)

	// Test basic creation
	result := testCtx.ApproveChatJoinRequest(userID)
	if result == nil {
		t.Error("Expected ApproveChatJoinRequest builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return ApproveChatJoinRequest for chaining")
	}

	// Test Timeout method
	result = testCtx.ApproveChatJoinRequest(userID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return ApproveChatJoinRequest for chaining")
	}

	// Test APIURL method
	result = testCtx.ApproveChatJoinRequest(userID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return ApproveChatJoinRequest for chaining")
	}
}

func TestContext_ApproveChatJoinRequestChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(987654)

	// Test complete method chaining
	result := testCtx.ApproveChatJoinRequest(userID).
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return ApproveChatJoinRequest")
	}
}

func TestApproveChatJoinRequest_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero UserID
	result := testCtx.ApproveChatJoinRequest(0)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle zero UserID")
	}

	// Test with negative UserID
	result = testCtx.ApproveChatJoinRequest(-123456)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle negative UserID")
	}

	// Test with maximum UserID
	result = testCtx.ApproveChatJoinRequest(9223372036854775807)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle maximum UserID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.ApproveChatJoinRequest(123456).ChatID(0)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle zero ChatID")
	}

	// Test with zero timeout
	result = testCtx.ApproveChatJoinRequest(123456).Timeout(0 * time.Second)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.ApproveChatJoinRequest(123456).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.ApproveChatJoinRequest(123456).APIURL(g.String(""))
	if result == nil {
		t.Error("ApproveChatJoinRequest should handle empty API URL")
	}
}

func TestApproveChatJoinRequest_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(555666)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.ApproveChatJoinRequest(userID)
	if result == nil {
		t.Error("ApproveChatJoinRequest should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestApproveChatJoinRequest_UserIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various user ID scenarios
	userIDs := []int64{
		1,
		100,
		123456789,
		999999999999,
		9223372036854775807, // max int64
	}

	for _, userID := range userIDs {
		result := testCtx.ApproveChatJoinRequest(userID)
		if result == nil {
			t.Errorf("ApproveChatJoinRequest should handle user ID: %d", userID)
		}

		// Test chaining for each user ID
		chainedResult := result.ChatID(-1001987654321).Timeout(20 * time.Second)
		if chainedResult == nil {
			t.Errorf("Chaining should work for user ID: %d", userID)
		}
	}
}

func TestApproveChatJoinRequest_ChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(777888)

	// Test various chat ID scenarios
	chatIDs := []int64{
		-1001234567890, // Supergroup
		-1001987654321, // Another supergroup
		-100123456789,  // Group
		-999999999999,  // Large negative
		123456789,      // Private chat (unlikely for join requests but should work)
	}

	for _, chatID := range chatIDs {
		result := testCtx.ApproveChatJoinRequest(userID).ChatID(chatID)
		if result == nil {
			t.Errorf("ApproveChatJoinRequest should handle chat ID: %d", chatID)
		}

		// Test combining with other methods
		combinedResult := result.Timeout(15 * time.Second).APIURL(g.String("https://api.telegram.org"))
		if combinedResult == nil {
			t.Errorf("Combined methods should work for chat ID: %d", chatID)
		}
	}
}

func TestApproveChatJoinRequest_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(999000)

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.ApproveChatJoinRequest(userID).Timeout(timeout)
		if result == nil {
			t.Errorf("ApproveChatJoinRequest should handle timeout: %v", timeout)
		}

		// Test combining timeout with other methods
		combinedResult := result.ChatID(-1001111222333).APIURL(g.String("https://custom.api.com"))
		if combinedResult == nil {
			t.Errorf("Timeout combination should work for: %v", timeout)
		}
	}
}

func TestApproveChatJoinRequest_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(111222)

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
		result := testCtx.ApproveChatJoinRequest(userID).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("ApproveChatJoinRequest should handle API URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combinedResult := result.ChatID(-1001444555666).Timeout(35 * time.Second)
		if combinedResult == nil {
			t.Errorf("API URL combination should work for: %s", apiURL)
		}
	}
}
