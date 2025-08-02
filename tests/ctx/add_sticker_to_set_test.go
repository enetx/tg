package ctx_test

import (
	"testing"
	"time"

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

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")

	// Test basic creation
	result := testCtx.AddStickerToSet(userID, name)
	if result == nil {
		t.Error("Expected AddStickerToSet builder to be created")
	}

	// Test File method
	result = result.File(g.String("test_sticker.png"))
	if result == nil {
		t.Error("File method should return AddStickerToSet for chaining")
	}

	// Test Format method
	result = result.Format(g.String("static"))
	if result == nil {
		t.Error("Format method should return AddStickerToSet for chaining")
	}

	// Test EmojiList method
	emojis := g.SliceOf[g.String](g.String("😀"), g.String("😁"))
	result = result.EmojiList(emojis)
	if result == nil {
		t.Error("EmojiList method should return AddStickerToSet for chaining")
	}

	// Test Keywords method
	keywords := g.SliceOf[g.String](g.String("happy"), g.String("smile"))
	result = result.Keywords(keywords)
	if result == nil {
		t.Error("Keywords method should return AddStickerToSet for chaining")
	}

	// Test MaskPosition method
	result = result.MaskPosition(g.String("forehead"), 0.5, 0.5, 1.0)
	if result == nil {
		t.Error("MaskPosition method should return AddStickerToSet for chaining")
	}

	// Test Timeout method
	result = result.Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return AddStickerToSet for chaining")
	}

	// Test APIURL method
	result = result.APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return AddStickerToSet for chaining")
	}
}

func TestAddStickerToSet_CompleteChain(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(789)
	name := g.String("complete_test_set")

	// Test complete method chaining
	result := testCtx.AddStickerToSet(userID, name).
		File(g.String("sticker.webp")).
		Format(g.String("animated")).
		EmojiList(g.SliceOf[g.String](g.String("🎉"), g.String("🎊"))).
		Keywords(g.SliceOf[g.String](g.String("party"), g.String("celebration"))).
		MaskPosition(g.String("eyes"), 0.0, 0.0, 2.0).
		Timeout(60 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return AddStickerToSet")
	}
}

func TestAddStickerToSet_EmptyValues(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(0)   // Test with zero userID
	name := g.String("") // Test with empty name

	result := testCtx.AddStickerToSet(userID, name)
	if result == nil {
		t.Error("AddStickerToSet should handle empty/zero values")
	}

	// Test with empty values in methods
	result = result.
		File(g.String("")).
		Format(g.String("")).
		EmojiList(g.Slice[g.String]{}).
		Keywords(g.Slice[g.String]{})

	if result == nil {
		t.Error("Methods should handle empty values gracefully")
	}
}
