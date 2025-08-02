package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerSetThumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")
	userID := int64(456)

	result := ctx.SetStickerSetThumbnail(name, userID)

	if result == nil {
		t.Error("Expected SetStickerSetThumbnail builder to be created")
	}

	// Test method chaining
	withFormat := result.Format(g.String("static"))
	if withFormat == nil {
		t.Error("Expected Format method to return builder")
	}
}
