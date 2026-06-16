package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteAllMessageReactions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.DeleteAllMessageReactions()
	if result == nil {
		t.Error("Expected DeleteAllMessageReactions builder to be created")
	}

	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return DeleteAllMessageReactions for chaining")
	}

	result = testCtx.DeleteAllMessageReactions().UserID(111222)
	if result == nil {
		t.Error("UserID method should return DeleteAllMessageReactions for chaining")
	}

	result = testCtx.DeleteAllMessageReactions().ActorChatID(-1001234567890)
	if result == nil {
		t.Error("ActorChatID method should return DeleteAllMessageReactions for chaining")
	}

	result = testCtx.DeleteAllMessageReactions().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeleteAllMessageReactions for chaining")
	}

	result = testCtx.DeleteAllMessageReactions().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeleteAllMessageReactions for chaining")
	}
}

func TestContext_DeleteAllMessageReactionsChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.DeleteAllMessageReactions().
		ChatID(-1001987654321).
		UserID(111222).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteAllMessageReactions")
	}
}

func TestDeleteAllMessageReactions_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.DeleteAllMessageReactions().Send()
	if sendResult.IsErr() {
		t.Logf("DeleteAllMessageReactions Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configuredSendResult := testCtx.DeleteAllMessageReactions().
		ChatID(456).
		UserID(111).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteAllMessageReactions configured Send failed as expected: %v", configuredSendResult.Err())
	}

	actorChatSendResult := testCtx.DeleteAllMessageReactions().
		ActorChatID(-1001111111111).
		Send()
	if actorChatSendResult.IsErr() {
		t.Logf("DeleteAllMessageReactions ActorChatID Send failed as expected: %v", actorChatSendResult.Err())
	}
}
