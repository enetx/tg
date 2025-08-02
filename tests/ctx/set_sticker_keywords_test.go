package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerKeywords(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")

	result := ctx.SetStickerKeywords(sticker)

	if result == nil {
		t.Error("Expected SetStickerKeywords builder to be created")
	}

	// Test method chaining
	keywords := g.Slice[g.String]{}
	keywords.Push(g.String("happy"))
	withKeywords := result.Keywords(keywords)
	if withKeywords == nil {
		t.Error("Expected Keywords method to return builder")
	}
}
