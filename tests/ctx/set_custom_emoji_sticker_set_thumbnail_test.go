package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetCustomEmojiStickerSetThumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")

	result := ctx.SetCustomEmojiStickerSetThumbnail(name)

	if result == nil {
		t.Error("Expected SetCustomEmojiStickerSetThumbnail builder to be created")
	}

	// Test method chaining
	withEmoji := result.CustomEmojiID(g.String("emoji_123"))
	if withEmoji == nil {
		t.Error("Expected CustomEmojiID method to return builder")
	}
}
