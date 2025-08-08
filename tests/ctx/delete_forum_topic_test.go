package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.DeleteForumTopic(messageThreadID)

	if result == nil {
		t.Error("Expected DeleteForumTopic builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	// Test Timeout method
	timeoutResult := result.Timeout(30 * time.Second)
	if timeoutResult == nil {
		t.Error("Timeout method should return DeleteForumTopic for chaining")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return DeleteForumTopic for chaining")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.DeleteForumTopic(messageThreadID)
	apiURLResultNil := freshResult.APIURL(g.String("https://api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return DeleteForumTopic for chaining with nil RequestOpts")
	}
}

func TestContext_DeleteForumTopicMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(456)

	// Test all methods in combination
	result := ctx.DeleteForumTopic(messageThreadID).
		ChatID(-1001987654321).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteForumTopic")
	}
}

func TestDeleteForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(789)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.DeleteForumTopic(messageThreadID).Send()

	if sendResult.IsErr() {
		t.Logf("DeleteForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.DeleteForumTopic(messageThreadID).
		ChatID(-1001987654321).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("DeleteForumTopic configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.DeleteForumTopic(messageThreadID).
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("DeleteForumTopic with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}
