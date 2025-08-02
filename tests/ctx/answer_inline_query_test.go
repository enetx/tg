package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerInlineQuery(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "inline123"}
	rawCtx := &ext.Context{Update: &gotgbot.Update{InlineQuery: inlineQuery}}
	ctx := ctx.New(bot, rawCtx)

	// Test AnswerInlineQuery
	result := ctx.AnswerInlineQuery(g.String("inline123"))
	if result == nil {
		t.Error("Expected AnswerInlineQuery builder to be created")
	}

	// Test method chaining
	chained := result.CacheFor(300 * time.Second).Personal()
	if chained == nil {
		t.Error("Expected chained methods to return builder")
	}
}

func TestContext_AnswerInlineQueryChaining(t *testing.T) {
	bot := &mockBot{}
	inlineQuery := &gotgbot.InlineQuery{Id: "inline123"}
	rawCtx := &ext.Context{Update: &gotgbot.Update{InlineQuery: inlineQuery}}
	ctx := ctx.New(bot, rawCtx)

	result := ctx.AnswerInlineQuery(g.String("inline123")).
		CacheFor(300 * time.Second).
		Personal()

	if result == nil {
		t.Error("Expected AnswerInlineQuery builder to be created")
	}

	// Test continued chaining
	final := result.NextOffset(g.String("next_offset"))
	if final == nil {
		t.Error("Expected NextOffset method to return builder")
	}
}
