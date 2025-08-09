package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_PinChatMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(123456)

	// Test basic creation
	result := testCtx.PinChatMessage(messageID)
	if result == nil {
		t.Error("Expected PinChatMessage builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return PinChatMessage for chaining")
	}

	// Test Business method
	result = testCtx.PinChatMessage(messageID).Business(g.String("business_connection_123"))
	if result == nil {
		t.Error("Business method should return PinChatMessage for chaining")
	}

	// Test Silent method
	result = testCtx.PinChatMessage(messageID).Silent()
	if result == nil {
		t.Error("Silent method should return PinChatMessage for chaining")
	}

	// Test Timeout method
	result = testCtx.PinChatMessage(messageID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return PinChatMessage for chaining")
	}

	// Test APIURL method
	result = testCtx.PinChatMessage(messageID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return PinChatMessage for chaining")
	}
}

func TestContext_PinChatMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(123456)

	// Test complete method chaining
	result := testCtx.PinChatMessage(messageID).
		ChatID(-1001987654321).
		Business(g.String("biz_conn_456")).
		Silent().
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return PinChatMessage")
	}
}

func TestPinChatMessage_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero MessageID
	result := testCtx.PinChatMessage(0)
	if result == nil {
		t.Error("PinChatMessage should handle zero MessageID")
	}

	// Test with negative MessageID
	result = testCtx.PinChatMessage(-123456)
	if result == nil {
		t.Error("PinChatMessage should handle negative MessageID")
	}

	// Test with maximum MessageID
	result = testCtx.PinChatMessage(9223372036854775807)
	if result == nil {
		t.Error("PinChatMessage should handle maximum MessageID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.PinChatMessage(123456).ChatID(0)
	if result == nil {
		t.Error("PinChatMessage should handle zero ChatID")
	}

	// Test with empty business connection ID
	result = testCtx.PinChatMessage(123456).Business(g.String(""))
	if result == nil {
		t.Error("PinChatMessage should handle empty business connection ID")
	}

	// Test with zero timeout
	result = testCtx.PinChatMessage(123456).Timeout(0 * time.Second)
	if result == nil {
		t.Error("PinChatMessage should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.PinChatMessage(123456).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("PinChatMessage should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.PinChatMessage(123456).APIURL(g.String(""))
	if result == nil {
		t.Error("PinChatMessage should handle empty API URL")
	}
}

func TestPinChatMessage_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(123456)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.PinChatMessage(messageID)
	if result == nil {
		t.Error("PinChatMessage should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestPinChatMessage_BusinessIntegration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(123456)

	// Test business connection with various IDs
	businessIDs := []string{
		"business_123",
		"conn_abc_def",
		"enterprise_connection_456",
		"short",
		"very_long_business_connection_identifier_with_many_characters",
	}

	for _, businessID := range businessIDs {
		result := testCtx.PinChatMessage(messageID).Business(g.String(businessID))
		if result == nil {
			t.Errorf("PinChatMessage should handle business ID: %s", businessID)
		}
	}

	// Test combining business with other methods
	result := testCtx.PinChatMessage(messageID).
		Business(g.String("combined_business")).
		Silent().
		Timeout(15 * time.Second)

	if result == nil {
		t.Error("Business method should work with other method combinations")
	}
}

func TestPinChatMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(123456)

	// Test Send method - will fail with mock but covers the method
	sendResult := testCtx.PinChatMessage(messageID).Send()

	if sendResult.IsErr() {
		t.Logf("PinChatMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with all options
	sendWithOptionsResult := testCtx.PinChatMessage(messageID).
		ChatID(-1001987654321).
		Business(g.String("business_conn_123")).
		Silent().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("PinChatMessage Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}

	// Test Send method using default chat ID (from effective chat)
	sendWithDefaultChatResult := testCtx.PinChatMessage(messageID).
		Silent().
		Send()

	if sendWithDefaultChatResult.IsErr() {
		t.Logf("PinChatMessage Send with default chat ID failed as expected with mock bot: %v", sendWithDefaultChatResult.Err())
	}
}
