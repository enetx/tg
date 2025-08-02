package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendChatAction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction()

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test method chaining
	chained := result.Typing()
	if chained == nil {
		t.Error("Expected Typing method to return builder")
	}
}

func TestContext_SendChatActionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().
		Typing().
		To(123).
		Thread(456)

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test that builder is functional
	_ = result
}
