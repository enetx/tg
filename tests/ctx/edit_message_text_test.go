package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditMessageText(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	result := ctx.EditMessageText(text)

	if result == nil {
		t.Error("Expected EditMessageText builder to be created")
	}

	// Test method chaining
	chained := result.HTML()
	if chained == nil {
		t.Error("Expected HTML method to return builder")
	}
}

func TestContext_EditMessageTextChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Edited text")

	result := ctx.EditMessageText(text).
		HTML().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageText builder to be created")
	}

	// Test that builder is functional
	_ = result
}
