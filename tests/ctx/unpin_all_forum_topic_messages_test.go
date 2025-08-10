package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnpinAllForumTopicMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	messageThreadID := int64(123)

	result := ctx.UnpinAllForumTopicMessages(messageThreadID)

	if result == nil {
		t.Error("Expected UnpinAllForumTopicMessages builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestUnpinAllForumTopicMessages_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	messageThreadID := int64(123)
	if ctx.UnpinAllForumTopicMessages(messageThreadID).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestUnpinAllForumTopicMessages_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	messageThreadID := int64(123)
	if ctx.UnpinAllForumTopicMessages(messageThreadID).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestUnpinAllForumTopicMessages_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	messageThreadID := int64(123)

	sendResult := ctx.UnpinAllForumTopicMessages(messageThreadID).Send()

	if sendResult.IsErr() {
		t.Logf("UnpinAllForumTopicMessages Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
