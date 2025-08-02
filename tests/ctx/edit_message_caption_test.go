package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditMessageCaption(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("New caption")

	result := ctx.EditMessageCaption(caption)

	if result == nil {
		t.Error("Expected EditMessageCaption builder to be created")
	}

	// Test method chaining
	chained := result.HTML()
	if chained == nil {
		t.Error("Expected HTML method to return builder")
	}
}

func TestContext_EditMessageCaptionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Caption: "original caption"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	caption := g.String("New caption")

	result := ctx.EditMessageCaption(caption).
		HTML().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageCaption builder to be created")
	}

	// Test that builder is functional
	_ = result
}
