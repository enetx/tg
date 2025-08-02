package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerEmojiList(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")

	result := ctx.SetStickerEmojiList(sticker)

	if result == nil {
		t.Error("Expected SetStickerEmojiList builder to be created")
	}

	// Test method chaining
	emojiList := g.Slice[g.String]{}
	emojiList.Push(g.String("😀"))
	withEmojis := result.EmojiList(emojiList)
	if withEmojis == nil {
		t.Error("Expected EmojiList method to return builder")
	}

	withTimeout := withEmojis.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}
