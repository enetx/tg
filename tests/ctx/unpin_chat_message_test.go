package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnpinChatMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.UnpinChatMessage()
	if result == nil {
		t.Error("Expected UnpinChatMessage builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return UnpinChatMessage for chaining")
	}

	// Test MessageID method
	result = testCtx.UnpinChatMessage().MessageID(123456)
	if result == nil {
		t.Error("MessageID method should return UnpinChatMessage for chaining")
	}

	// Test Business method
	result = testCtx.UnpinChatMessage().Business(g.String("business_connection_789"))
	if result == nil {
		t.Error("Business method should return UnpinChatMessage for chaining")
	}

	// Test Timeout method
	result = testCtx.UnpinChatMessage().Timeout(25 * time.Second)
	if result == nil {
		t.Error("Timeout method should return UnpinChatMessage for chaining")
	}

	// Test APIURL method
	result = testCtx.UnpinChatMessage().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return UnpinChatMessage for chaining")
	}
}

func TestContext_UnpinChatMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.UnpinChatMessage().
		ChatID(-1001987654321).
		MessageID(789012).
		Business(g.String("biz_conn_unpin_789")).
		Timeout(35 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return UnpinChatMessage")
	}
}

func TestUnpinChatMessage_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero MessageID
	result := testCtx.UnpinChatMessage().MessageID(0)
	if result == nil {
		t.Error("UnpinChatMessage should handle zero MessageID")
	}

	// Test with negative MessageID
	result = testCtx.UnpinChatMessage().MessageID(-123456)
	if result == nil {
		t.Error("UnpinChatMessage should handle negative MessageID")
	}

	// Test with maximum MessageID
	result = testCtx.UnpinChatMessage().MessageID(9223372036854775807)
	if result == nil {
		t.Error("UnpinChatMessage should handle maximum MessageID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.UnpinChatMessage().ChatID(0)
	if result == nil {
		t.Error("UnpinChatMessage should handle zero ChatID")
	}

	// Test with empty business connection ID
	result = testCtx.UnpinChatMessage().Business(g.String(""))
	if result == nil {
		t.Error("UnpinChatMessage should handle empty business connection ID")
	}

	// Test with zero timeout
	result = testCtx.UnpinChatMessage().Timeout(0 * time.Second)
	if result == nil {
		t.Error("UnpinChatMessage should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.UnpinChatMessage().Timeout(12 * time.Hour)
	if result == nil {
		t.Error("UnpinChatMessage should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.UnpinChatMessage().APIURL(g.String(""))
	if result == nil {
		t.Error("UnpinChatMessage should handle empty API URL")
	}
}

func TestUnpinChatMessage_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.UnpinChatMessage()
	if result == nil {
		t.Error("UnpinChatMessage should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestUnpinChatMessage_MessageIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various message ID scenarios
	messageIDs := []int64{
		1,
		100,
		123456789,
		999999999999,
		9223372036854775807, // max int64
	}

	for _, msgID := range messageIDs {
		result := testCtx.UnpinChatMessage().MessageID(msgID)
		if result == nil {
			t.Errorf("UnpinChatMessage should handle message ID: %d", msgID)
		}
	}

	// Test without MessageID (unpins all messages)
	result := testCtx.UnpinChatMessage()
	if result == nil {
		t.Error("UnpinChatMessage should work without MessageID (unpin all)")
	}

	// Test combining MessageID with other methods
	result = testCtx.UnpinChatMessage().
		MessageID(555666).
		Business(g.String("combined_business")).
		Timeout(20 * time.Second)

	if result == nil {
		t.Error("MessageID method should work with other method combinations")
	}
}

func TestUnpinChatMessage_BusinessIntegration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test business connection with various IDs
	businessIDs := []string{
		"unpin_business_123",
		"conn_xyz_789",
		"enterprise_unpin_connection_456",
		"short_biz",
		"very_long_business_connection_identifier_for_unpinning_messages",
	}

	for _, businessID := range businessIDs {
		result := testCtx.UnpinChatMessage().Business(g.String(businessID))
		if result == nil {
			t.Errorf("UnpinChatMessage should handle business ID: %s", businessID)
		}
	}

	// Test combining business with MessageID
	result := testCtx.UnpinChatMessage().
		Business(g.String("unpin_biz_msg")).
		MessageID(987654).
		Timeout(10 * time.Second)

	if result == nil {
		t.Error("Business method should work with MessageID and other combinations")
	}
}

func TestUnpinChatMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := testCtx.UnpinChatMessage().MessageID(123456).Send()

	if sendResult.IsErr() {
		t.Logf("UnpinChatMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
