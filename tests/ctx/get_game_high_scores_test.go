package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetGameHighScores(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	result := ctx.GetGameHighScores(userID)

	if result == nil {
		t.Error("Expected GetGameHighScores builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(123)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}
