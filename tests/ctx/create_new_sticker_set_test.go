package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateNewStickerSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set_by_bot")
	title := g.String("Test Sticker Set")

	result := ctx.CreateNewStickerSet(userID, name, title)

	if result == nil {
		t.Error("Expected CreateNewStickerSet builder to be created")
	}

	// Test method chaining
	withType := result.StickerType(g.String("regular"))
	if withType == nil {
		t.Error("Expected StickerType method to return builder")
	}

	withRepainting := withType.NeedsRepainting()
	if withRepainting == nil {
		t.Error("Expected NeedsRepainting method to return builder")
	}

	// Test sticker builder
	emojiList := g.Slice[g.String]{}
	emojiList.Push(g.String("😀"))
	stickerBuilder := withRepainting.Sticker(
		g.String("sticker.png"),
		g.String("static"),
		emojiList,
	)
	if stickerBuilder == nil {
		t.Error("Expected Sticker method to return sticker builder")
	}

	keywordList := g.Slice[g.String]{}
	keywordList.Push(g.String("happy"))
	withKeywords := stickerBuilder.Keywords(keywordList)
	if withKeywords == nil {
		t.Error("Expected Keywords method to return sticker builder")
	}

	backToParent := withKeywords.Add()
	if backToParent == nil {
		t.Error("Expected Add method to return CreateNewStickerSet")
	}
}
