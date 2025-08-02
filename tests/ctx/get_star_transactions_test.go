package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetStarTransactions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetStarTransactions()

	if result == nil {
		t.Error("Expected GetStarTransactions builder to be created")
	}

	// Test method chaining
	withOffset := result.Offset(0)
	if withOffset == nil {
		t.Error("Expected Offset method to return builder")
	}
}
