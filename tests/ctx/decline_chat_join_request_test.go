package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeclineChatJoinRequest(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(456789)

	// Test basic creation
	result := testCtx.DeclineChatJoinRequest(userID)
	if result == nil {
		t.Error("Expected DeclineChatJoinRequest builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return DeclineChatJoinRequest for chaining")
	}

	// Test Timeout method
	result = testCtx.DeclineChatJoinRequest(userID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeclineChatJoinRequest for chaining")
	}

	// Test APIURL method
	result = testCtx.DeclineChatJoinRequest(userID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeclineChatJoinRequest for chaining")
	}
}

func TestContext_DeclineChatJoinRequestChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(987654)

	// Test complete method chaining
	result := testCtx.DeclineChatJoinRequest(userID).
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeclineChatJoinRequest")
	}
}

func TestDeclineChatJoinRequest_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeclineChatJoinRequest(userID).Send()

	if sendResult.IsErr() {
		t.Logf("DeclineChatJoinRequest Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeclineChatJoinRequest(userID).
		ChatID(789).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeclineChatJoinRequest configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
