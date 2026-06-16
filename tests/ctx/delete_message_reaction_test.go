package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteMessageReaction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	result := testCtx.DeleteMessageReaction(messageID)
	if result == nil {
		t.Error("Expected DeleteMessageReaction builder to be created")
	}

	result = result.ChatID(456)
	if result == nil {
		t.Error("ChatID method should return DeleteMessageReaction for chaining")
	}

	result = testCtx.DeleteMessageReaction(messageID).UserID(111222)
	if result == nil {
		t.Error("UserID method should return DeleteMessageReaction for chaining")
	}

	result = testCtx.DeleteMessageReaction(messageID).ActorChatID(-1001234567890)
	if result == nil {
		t.Error("ActorChatID method should return DeleteMessageReaction for chaining")
	}

	result = testCtx.DeleteMessageReaction(messageID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeleteMessageReaction for chaining")
	}

	result = testCtx.DeleteMessageReaction(messageID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeleteMessageReaction for chaining")
	}
}

func TestContext_DeleteMessageReactionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	messageID := int64(789)

	result := testCtx.DeleteMessageReaction(messageID).
		ChatID(-1001987654321).
		UserID(111222).
		ActorChatID(-1009999999999).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteMessageReaction")
	}
}

func TestDeleteMessageReaction_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.DeleteMessageReaction(789).Send()
	if sendResult.IsErr() {
		t.Logf("DeleteMessageReaction Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configuredSendResult := testCtx.DeleteMessageReaction(789).
		ChatID(456).
		UserID(111).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteMessageReaction configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// ActorChatID branch.
	actorChatSendResult := testCtx.DeleteMessageReaction(789).
		ActorChatID(-1001111111111).
		Send()
	if actorChatSendResult.IsErr() {
		t.Logf("DeleteMessageReaction ActorChatID Send failed as expected: %v", actorChatSendResult.Err())
	}
}
