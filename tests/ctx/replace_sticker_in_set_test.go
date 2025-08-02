package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ReplaceStickerInSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")
	oldSticker := g.String("old_sticker_id")
	newSticker := gotgbot.InputSticker{
		Sticker:   "new_sticker.png",
		Format:    "static",
		EmojiList: []string{"😀"},
	}

	result := ctx.ReplaceStickerInSet(userID, name, oldSticker, newSticker)

	if result == nil {
		t.Error("Expected ReplaceStickerInSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}
