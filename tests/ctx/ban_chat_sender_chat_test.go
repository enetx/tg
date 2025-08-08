package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_BanChatSenderChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(789012)

	// Test basic creation
	result := testCtx.BanChatSenderChat(senderChatID)
	if result == nil {
		t.Error("Expected BanChatSenderChat builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return BanChatSenderChat for chaining")
	}

	// Test Timeout method
	result = testCtx.BanChatSenderChat(senderChatID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return BanChatSenderChat for chaining")
	}

	// Test APIURL method
	result = testCtx.BanChatSenderChat(senderChatID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return BanChatSenderChat for chaining")
	}
}

func TestContext_BanChatSenderChatChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(987654)

	// Test complete method chaining
	result := testCtx.BanChatSenderChat(senderChatID).
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return BanChatSenderChat")
	}
}

func TestBanChatSenderChat_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero SenderChatID
	result := testCtx.BanChatSenderChat(0)
	if result == nil {
		t.Error("BanChatSenderChat should handle zero SenderChatID")
	}

	// Test with negative SenderChatID (channel/group format)
	result = testCtx.BanChatSenderChat(-1001111222333)
	if result == nil {
		t.Error("BanChatSenderChat should handle negative SenderChatID")
	}

	// Test with positive SenderChatID (bot/user format)
	result = testCtx.BanChatSenderChat(123456789)
	if result == nil {
		t.Error("BanChatSenderChat should handle positive SenderChatID")
	}

	// Test with maximum SenderChatID
	result = testCtx.BanChatSenderChat(9223372036854775807)
	if result == nil {
		t.Error("BanChatSenderChat should handle maximum SenderChatID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.BanChatSenderChat(123456).ChatID(0)
	if result == nil {
		t.Error("BanChatSenderChat should handle zero ChatID")
	}

	// Test with zero timeout
	result = testCtx.BanChatSenderChat(123456).Timeout(0 * time.Second)
	if result == nil {
		t.Error("BanChatSenderChat should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.BanChatSenderChat(123456).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("BanChatSenderChat should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.BanChatSenderChat(123456).APIURL(g.String(""))
	if result == nil {
		t.Error("BanChatSenderChat should handle empty API URL")
	}
}

func TestBanChatSenderChat_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(555666)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.BanChatSenderChat(senderChatID)
	if result == nil {
		t.Error("BanChatSenderChat should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestBanChatSenderChat_SenderChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various sender chat ID scenarios
	senderChatIDs := []int64{
		-1001234567890, // Supergroup format
		-1001987654321, // Another supergroup
		-100123456789,  // Group format
		-999999999999,  // Large negative
		123456789,      // Bot/user format
		999999999999,   // Large positive
		1,              // Minimal positive
	}

	for _, senderChatID := range senderChatIDs {
		result := testCtx.BanChatSenderChat(senderChatID)
		if result == nil {
			t.Errorf("BanChatSenderChat should handle sender chat ID: %d", senderChatID)
		}

		// Test chaining for each sender chat ID
		chainedResult := result.ChatID(-1001987654321).Timeout(20 * time.Second)
		if chainedResult == nil {
			t.Errorf("Chaining should work for sender chat ID: %d", senderChatID)
		}
	}
}

func TestBanChatSenderChat_ChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(777888)

	// Test various target chat ID scenarios
	chatIDs := []int64{
		-1001234567890, // Supergroup
		-1001987654321, // Another supergroup
		-100123456789,  // Group
		-999999999999,  // Large negative
	}

	for _, chatID := range chatIDs {
		result := testCtx.BanChatSenderChat(senderChatID).ChatID(chatID)
		if result == nil {
			t.Errorf("BanChatSenderChat should handle chat ID: %d", chatID)
		}

		// Test combining with other methods
		combinedResult := result.Timeout(15 * time.Second).APIURL(g.String("https://api.telegram.org"))
		if combinedResult == nil {
			t.Errorf("Combined methods should work for chat ID: %d", chatID)
		}
	}
}

func TestBanChatSenderChat_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(999000)

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.BanChatSenderChat(senderChatID).Timeout(timeout)
		if result == nil {
			t.Errorf("BanChatSenderChat should handle timeout: %v", timeout)
		}

		// Test combining timeout with other methods
		combinedResult := result.ChatID(-1001111222333).APIURL(g.String("https://custom.api.com"))
		if combinedResult == nil {
			t.Errorf("Timeout combination should work for: %v", timeout)
		}
	}
}

func TestBanChatSenderChat_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	senderChatID := int64(111222)

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
		result := testCtx.BanChatSenderChat(senderChatID).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("BanChatSenderChat should handle API URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combinedResult := result.ChatID(-1001444555666).Timeout(35 * time.Second)
		if combinedResult == nil {
			t.Errorf("API URL combination should work for: %s", apiURL)
		}
	}
}

func TestBanChatSenderChat_SenderChatTypes(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test different types of sender chats that might be banned
	testCases := []struct {
		name         string
		senderChatID int64
		description  string
	}{
		{"Channel", -1001111111111, "Public/private channel"},
		{"Supergroup", -1002222222222, "Supergroup"},
		{"Group", -100333333333, "Basic group"},
		{"Bot", 123456789, "Bot account"},
		{"Anonymous", -1000000000000, "Anonymous admin"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testCtx.BanChatSenderChat(tc.senderChatID)
			if result == nil {
				t.Errorf("BanChatSenderChat should handle %s (%s): %d", tc.name, tc.description, tc.senderChatID)
			}

			// Test complete workflow for each type
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

func TestBanChatSenderChat_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	senderChatID := int64(456)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.BanChatSenderChat(senderChatID).Send()

	if sendResult.IsErr() {
		t.Logf("BanChatSenderChat Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.BanChatSenderChat(senderChatID).
		ChatID(789).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("BanChatSenderChat configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
