package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerCallbackQuery(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{
		Id:   "callback123",
		Data: "test_data",
	}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	ctx := ctx.New(bot, rawCtx)

	if ctx.Callback != callback {
		t.Error("Expected callback query to be set")
	}

	// Test callback query answer
	result := ctx.AnswerCallbackQuery(g.String("Test answer"))

	if result == nil {
		t.Error("Expected AnswerCallbackQuery builder to be created")
	}

	// Test method chaining
	chained := result.Alert()
	if chained == nil {
		t.Error("Expected Alert method to return builder")
	}
}

func TestContext_AnswerCallbackQueryChaining(t *testing.T) {
	bot := &mockBot{}
	callback := &gotgbot.CallbackQuery{
		Id:   "callback123",
		Data: "test_data",
	}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1, CallbackQuery: callback},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.AnswerCallbackQuery(g.String("Test answer")).
		Alert().
		URL(g.String("https://example.com"))

	if result == nil {
		t.Error("Expected AnswerCallbackQuery builder to be created")
	}

	// Test continued chaining
	final := result.CacheFor(60)
	if final == nil {
		t.Error("Expected CacheFor method to return builder")
	}
}
