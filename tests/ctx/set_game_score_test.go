package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetGameScore(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	score := int64(1000)

	result := ctx.SetGameScore(userID, score)

	if result == nil {
		t.Error("Expected SetGameScore builder to be created")
	}

	// Test method chaining
	withForce := result.Force()
	if withForce == nil {
		t.Error("Expected Force method to return builder")
	}
}
