package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessage()

	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}

	// Test method chaining
	chained := result.MessageID(123)
	if chained == nil {
		t.Error("Expected MessageID method to return builder")
	}
}

func TestContext_DeleteMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.DeleteMessage().
		MessageID(123).
		ChatID(456)

	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}

	// Test that builder is functional
	_ = result
}
