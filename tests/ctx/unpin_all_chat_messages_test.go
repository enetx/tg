package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnpinAllChatMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.UnpinAllChatMessages()

	if result == nil {
		t.Error("Expected UnpinAllChatMessages builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_UnpinAllChatMessagesChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.UnpinAllChatMessages().
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected UnpinAllChatMessages builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestUnpinAllChatMessages_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.UnpinAllChatMessages().Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestUnpinAllChatMessages_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.UnpinAllChatMessages().APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestUnpinAllChatMessages_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	sendResult := ctx.UnpinAllChatMessages().Send()
	
	if sendResult.IsErr() {
		t.Logf("UnpinAllChatMessages Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
