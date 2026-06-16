package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserPersonalChatMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.GetUserPersonalChatMessages(123456, 10)
	if result == nil {
		t.Error("Expected GetUserPersonalChatMessages builder to be created")
	}

	result = result.Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return GetUserPersonalChatMessages for chaining")
	}

	result = testCtx.GetUserPersonalChatMessages(123456, 5).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return GetUserPersonalChatMessages for chaining")
	}
}

func TestContext_GetUserPersonalChatMessages_LimitBoundaries(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Per Bot API 10.0 the limit must be 1-20; the builder accepts any value and
	// defers validation to the server.
	for _, limit := range []int64{1, 5, 10, 15, 20} {
		result := testCtx.GetUserPersonalChatMessages(123456, limit)
		if result == nil {
			t.Errorf("GetUserPersonalChatMessages should accept limit %d", limit)
		}
	}

	zero := testCtx.GetUserPersonalChatMessages(123456, 0)
	if zero == nil {
		t.Error("GetUserPersonalChatMessages should still accept zero limit; the API rejects it at send time")
	}
}

func TestGetUserPersonalChatMessages_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.GetUserPersonalChatMessages(456, 10).Send()
	if sendResult.IsErr() {
		t.Logf("GetUserPersonalChatMessages Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configuredSendResult := testCtx.GetUserPersonalChatMessages(456, 5).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetUserPersonalChatMessages configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
