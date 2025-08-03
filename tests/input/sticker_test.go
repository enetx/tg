package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestNewSticker(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf(g.String("😀"), g.String("😃"))

	sticker := input.Sticker(stickerID, format, emojiList)
	if sticker == nil {
		t.Error("Expected Sticker to be created")
	}
}

func TestSticker_MaskPosition(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf(g.String("😀"))

	sticker := input.Sticker(stickerID, format, emojiList)
	maskPosition := &gotgbot.MaskPosition{
		Point:  "forehead",
		XShift: 0.1,
		YShift: 0.2,
		Scale:  1.5,
	}
	result := sticker.MaskPosition(maskPosition)
	if result == nil {
		t.Error("Expected MaskPosition method to return Sticker")
	}
	if result != sticker {
		t.Error("Expected MaskPosition to return same Sticker instance")
	}

	built := result.Build()
	if built.MaskPosition == nil {
		t.Error("Expected MaskPosition to be set")
	}
	if built.MaskPosition.Point != "forehead" {
		t.Error("Expected MaskPosition point to be set correctly")
	}
}

func TestSticker_Keywords(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf[g.String](g.String("😀"))

	sticker := input.Sticker(stickerID, format, emojiList)
	keywords := g.SliceOf(g.String("happy"), g.String("smile"))
	result := sticker.Keywords(keywords)
	if result == nil {
		t.Error("Expected Keywords method to return Sticker")
	}
	if result != sticker {
		t.Error("Expected Keywords to return same Sticker instance")
	}

	built := result.Build()
	if len(built.Keywords) != 2 {
		t.Errorf("Expected 2 keywords, got %d", len(built.Keywords))
	}
	if built.Keywords[0] != "happy" || built.Keywords[1] != "smile" {
		t.Error("Expected Keywords to be set correctly")
	}
}

func TestSticker_Build(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf(g.String("😀"), g.String("😃"))

	sticker := input.Sticker(stickerID, format, emojiList)
	built := sticker.Build()

	if built.Format != format.Std() {
		t.Errorf("Expected Format to be %s, got %s", format.Std(), built.Format)
	}
	if len(built.EmojiList) != 2 {
		t.Errorf("Expected 2 emojis, got %d", len(built.EmojiList))
	}
	if built.EmojiList[0] != "😀" || built.EmojiList[1] != "😃" {
		t.Error("Expected EmojiList to be set correctly")
	}
}

func TestSticker_BuildReturnsCorrectType(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf(g.String("😀"))

	sticker := input.Sticker(stickerID, format, emojiList)
	built := sticker.Build()

	// Verify that Build() returns the correct type
	if _, ok := any(built).(gotgbot.InputSticker); !ok {
		t.Error("Expected Build() to return gotgbot.InputSticker")
	}
}

func TestSticker_MethodChaining(t *testing.T) {
	stickerID := file.Input(testURL).Ok()
	format := g.String("static")
	emojiList := g.SliceOf(g.String("😀"))
	keywords := g.SliceOf(g.String("happy"), g.String("smile"))
	maskPosition := &gotgbot.MaskPosition{
		Point:  "forehead",
		XShift: 0.1,
		YShift: 0.2,
		Scale:  1.5,
	}

	result := input.Sticker(stickerID, format, emojiList).
		Keywords(keywords).
		MaskPosition(maskPosition)

	if result == nil {
		t.Error("Expected method chaining to work")
	}

	built := result.Build()
	// if built.Sticker == "" {
	// 	t.Error("Expected chained Sticker to build correctly")
	// }

	if _, ok := any(built).(gotgbot.InputSticker); !ok {
		t.Error("Expected result to be InputSticker")
	}
}
