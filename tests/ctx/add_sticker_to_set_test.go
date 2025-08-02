package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AddStickerToSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")

	result := ctx.AddStickerToSet(userID, name)

	if result == nil {
		t.Error("Expected AddStickerToSet builder to be created")
	}

	// Test method chaining
	chained := result.File(g.String("new_file.png"))
	if chained == nil {
		t.Error("Expected File method to return builder")
	}

	formatted := chained.Format(g.String("static"))
	if formatted == nil {
		t.Error("Expected Format method to return builder")
	}
}
