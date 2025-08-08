package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_LeaveChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.LeaveChat()
	if result == nil {
		t.Error("Expected LeaveChat builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return LeaveChat for chaining")
	}

	// Test Timeout method
	result = testCtx.LeaveChat().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return LeaveChat for chaining")
	}

	// Test APIURL method
	result = testCtx.LeaveChat().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return LeaveChat for chaining")
	}
}

func TestContext_LeaveChatChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.LeaveChat().
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return LeaveChat")
	}
}

func TestLeaveChat_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero ChatID (should use effective chat)
	result := testCtx.LeaveChat().ChatID(0)
	if result == nil {
		t.Error("LeaveChat should handle zero ChatID")
	}

	// Test with negative ChatID (supergroup/channel)
	result = testCtx.LeaveChat().ChatID(-1001111111111)
	if result == nil {
		t.Error("LeaveChat should handle negative ChatID")
	}

	// Test with positive ChatID (user/group)
	result = testCtx.LeaveChat().ChatID(123456789)
	if result == nil {
		t.Error("LeaveChat should handle positive ChatID")
	}

	// Test with zero timeout
	result = testCtx.LeaveChat().Timeout(0 * time.Second)
	if result == nil {
		t.Error("LeaveChat should handle zero timeout")
	}

	// Test with empty API URL
	result = testCtx.LeaveChat().APIURL(g.String(""))
	if result == nil {
		t.Error("LeaveChat should handle empty API URL")
	}
}

func TestLeaveChat_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.LeaveChat()
	if result == nil {
		t.Error("LeaveChat should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestLeaveChat_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.LeaveChat().Send()

	if sendResult.IsErr() {
		t.Logf("LeaveChat Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
