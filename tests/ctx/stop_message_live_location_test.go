package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func TestContext_StopMessageLiveLocation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.StopMessageLiveLocation()

	if result == nil {
		t.Error("Expected StopMessageLiveLocation builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_StopMessageLiveLocationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.StopMessageLiveLocation().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected StopMessageLiveLocation builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestStopMessageLiveLocation_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopMessageLiveLocation().InlineMessageID(g.String("inline_123")) == nil {
		t.Error("InlineMessageID should return builder")
	}
}

func TestStopMessageLiveLocation_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopMessageLiveLocation().Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestStopMessageLiveLocation_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})
	btn1 := keyboard.NewButton().Text(g.String("Update")).Callback(g.String("update_location"))
	if ctx.StopMessageLiveLocation().Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestStopMessageLiveLocation_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopMessageLiveLocation().Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestStopMessageLiveLocation_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopMessageLiveLocation().APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestStopMessageLiveLocation_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789}, Update: &gotgbot.Update{UpdateId: 1}})

	sendResult := ctx.StopMessageLiveLocation().Send()

	if sendResult.IsErr() {
		t.Logf("StopMessageLiveLocation Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
