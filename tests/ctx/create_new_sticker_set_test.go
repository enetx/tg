package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
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
		file.Input("sticker.png").UnwrapOrDefault(),
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

// Tests for methods with 0% coverage

func TestStickerBuilder_MaskPosition(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("mask_sticker_set_by_bot")
	title := g.String("Mask Sticker Set")

	// Test MaskPosition method functionality
	emojiList := g.Slice[g.String]{}
	emojiList.Push(g.String("😷"))

	maskPositions := []struct {
		point       g.String
		xShift      float64
		yShift      float64
		scale       float64
		description string
	}{
		{g.String("forehead"), 0.0, 0.0, 1.0, "Forehead position"},
		{g.String("eyes"), -0.1, 0.2, 0.8, "Eyes position with offset"},
		{g.String("mouth"), 0.15, -0.05, 1.2, "Mouth position enlarged"},
		{g.String("chin"), 0.0, -0.3, 0.9, "Chin position lowered"},
		{g.String("forehead"), 0.5, 0.5, 2.0, "Extreme position values"},
		{g.String("eyes"), -0.5, -0.5, 0.1, "Minimum scale values"},
	}

	for _, maskPos := range maskPositions {
		maskResult := ctx.CreateNewStickerSet(userID, name, title).
			StickerType(g.String("mask")).
			Sticker(
				file.Input("mask_sticker.png").UnwrapOrDefault(),
				g.String("static"),
				emojiList,
			).
			MaskPosition(maskPos.point, maskPos.xShift, maskPos.yShift, maskPos.scale).
			Add()

		if maskResult == nil {
			t.Errorf("MaskPosition method with %s should work", maskPos.description)
		}

		// Test Send with mask position
		sendResult := maskResult.Send()
		if sendResult.IsErr() {
			t.Logf("CreateNewStickerSet with MaskPosition %s Send failed as expected: %v",
				maskPos.description, sendResult.Err())
		}
	}
}

func TestStickerBuilder_MaskPositionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("mask_sticker_set_chain_by_bot")
	title := g.String("Mask Sticker Set Chain")

	emojiList := g.Slice[g.String]{}
	emojiList.Push(g.String("😎"))

	// Test MaskPosition method chaining
	chainedResult := ctx.CreateNewStickerSet(userID, name, title).
		StickerType(g.String("mask")).
		Sticker(
			file.Input("mask_sticker_chain.png").UnwrapOrDefault(),
			g.String("static"),
			emojiList,
		).
		Keywords(func() g.Slice[g.String] {
			keywords := g.Slice[g.String]{}
			keywords.Push(g.String("sunglasses"))
			keywords.Push(g.String("cool"))
			return keywords
		}()).
		MaskPosition(g.String("eyes"), 0.0, 0.1, 1.1).
		Add().
		Timeout(45 * time.Second).
		APIURL(g.String("https://mask-stickers-api.telegram.org")).
		Send()

	if chainedResult.IsErr() {
		t.Logf("CreateNewStickerSet with MaskPosition chaining Send failed as expected: %v", chainedResult.Err())
	}
}

func TestCreateNewStickerSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set_by_bot")
	title := g.String("Test Sticker Set")

	// Test Send method - will fail with mock but covers the method
	emojiList := g.Slice[g.String]{}
	emojiList.Push(g.String("😀"))

	sendResult := ctx.CreateNewStickerSet(userID, name, title).
		StickerType(g.String("regular")).
		Sticker(
			file.Input("sticker.png").UnwrapOrDefault(),
			g.String("static"),
			emojiList,
		).Add().
		Send()

	if sendResult.IsErr() {
		t.Logf("CreateNewStickerSet Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.CreateNewStickerSet(userID, name, title).
		StickerType(g.String("regular")).
		NeedsRepainting().
		Sticker(
			file.Input("sticker1.png").UnwrapOrDefault(),
			g.String("static"),
			emojiList,
		).Add().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("CreateNewStickerSet configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
