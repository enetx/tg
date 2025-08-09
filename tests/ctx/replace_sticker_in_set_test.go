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

func TestContext_ReplaceStickerInSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")
	oldSticker := g.String("old_sticker_id")
	newSticker := gotgbot.InputSticker{
		Sticker:   file.Input("new_sticker.png").UnwrapOrDefault().Doc,
		Format:    "static",
		EmojiList: []string{"😀"},
	}

	result := ctx.ReplaceStickerInSet(userID, name, oldSticker, newSticker)

	if result == nil {
		t.Error("Expected ReplaceStickerInSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestReplaceStickerInSet_Sticker(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")
	oldSticker := g.String("old_sticker_id")
	initialSticker := gotgbot.InputSticker{
		Sticker:   file.Input("initial_sticker.png").UnwrapOrDefault().Doc,
		Format:    "static",
		EmojiList: []string{"😀"},
	}

	stickers := []gotgbot.InputSticker{
		{
			Sticker:   file.Input("sticker1.png").UnwrapOrDefault().Doc,
			Format:    "static",
			EmojiList: []string{"😀"},
		},
		{
			Sticker:   file.Input("sticker2.webm").UnwrapOrDefault().Doc,
			Format:    "video",
			EmojiList: []string{"😁", "😂"},
		},
		{
			Sticker:   file.Input("sticker3.tgs").UnwrapOrDefault().Doc,
			Format:    "animated",
			EmojiList: []string{"🎉"},
		},
	}

	for i, sticker := range stickers {
		result := ctx.ReplaceStickerInSet(userID, name, oldSticker, initialSticker)
		stickerResult := result.Sticker(sticker)
		if stickerResult == nil {
			t.Errorf("Sticker method should return ReplaceStickerInSet builder for chaining with sticker %d", i)
		}

		chainedSticker := gotgbot.InputSticker{
			Sticker:   file.Input("chained.png").UnwrapOrDefault().Doc,
			Format:    "static",
			EmojiList: []string{"🔥"},
		}
		chainedResult := stickerResult.Sticker(chainedSticker)
		if chainedResult == nil {
			t.Errorf("Sticker method should support chaining and override with sticker %d", i)
		}
	}
}

func TestReplaceStickerInSet_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")
	oldSticker := g.String("old_sticker_id")
	newSticker := gotgbot.InputSticker{
		Sticker:   file.Input("new_sticker.png").UnwrapOrDefault().Doc,
		Format:    "static",
		EmojiList: []string{"😀"},
	}

	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.ReplaceStickerInSet(userID, name, oldSticker, newSticker)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return ReplaceStickerInSet builder for chaining with URL: %s", apiURL)
		}

		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestReplaceStickerInSet_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	name := g.String("test_sticker_set")
	oldSticker := g.String("old_sticker_id")
	newSticker := gotgbot.InputSticker{
		Sticker:   file.Input("new_sticker.png").UnwrapOrDefault().Doc,
		Format:    "static",
		EmojiList: []string{"😀"},
	}

	sendResult := ctx.ReplaceStickerInSet(userID, name, oldSticker, newSticker).Send()

	if sendResult.IsErr() {
		t.Logf("ReplaceStickerInSet Send failed as expected with mock bot: %v", sendResult.Err())
	}

	replacementSticker := gotgbot.InputSticker{
		Sticker:   file.Input("replacement.webm").UnwrapOrDefault().Doc,
		Format:    "video",
		EmojiList: []string{"🎥"},
	}

	sendWithOptionsResult := ctx.ReplaceStickerInSet(userID, name, oldSticker, newSticker).
		Sticker(replacementSticker).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("ReplaceStickerInSet Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}
}
