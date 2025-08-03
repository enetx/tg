package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
)

func TestContext_SetStickerMaskPosition(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")

	result := ctx.SetStickerMaskPosition(file.Input(sticker).UnwrapOrDefault())

	if result == nil {
		t.Error("Expected SetStickerMaskPosition builder to be created")
	}

	// Test method chaining
	withMask := result.MaskPosition(g.String("forehead"), 0.5, 0.5, 1.0)
	if withMask == nil {
		t.Error("Expected MaskPosition method to return builder")
	}
}
