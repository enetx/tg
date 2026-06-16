package input_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestStickerMedia(t *testing.T) {
	mediaFile := file.Input(testFileID).Ok()
	sticker := input.StickerMedia(mediaFile)
	if sticker == nil {
		t.Error("Expected StickerMedia to be created")
	}
	if !assertPollOptionMedia(sticker) {
		t.Error("StickerMedia should implement PollOptionMedia correctly")
	}
}

func TestStickerMedia_Emoji(t *testing.T) {
	mediaFile := file.Input(testFileID).Ok()
	sticker := input.StickerMedia(mediaFile)
	result := sticker.Emoji(g.String("🔥"))
	if result != sticker {
		t.Error("Expected Emoji to return same MediaSticker instance")
	}

	built := result.BuildPollOptionMedia()
	if v, ok := built.(gotgbot.InputMediaSticker); ok {
		if v.Emoji != "🔥" {
			t.Error("Expected Emoji to be set correctly")
		}
	} else {
		t.Error("Expected result to be InputMediaSticker")
	}
}
