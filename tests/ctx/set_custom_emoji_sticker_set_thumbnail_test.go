package ctx_test

import (
	"testing"
	"time"

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

func TestSetCustomEmojiStickerSetThumbnail_DropThumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	if ctx.SetCustomEmojiStickerSetThumbnail(name).DropThumbnail() == nil { t.Error("DropThumbnail should return builder") }
}

func TestSetCustomEmojiStickerSetThumbnail_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	if ctx.SetCustomEmojiStickerSetThumbnail(name).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetCustomEmojiStickerSetThumbnail_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	if ctx.SetCustomEmojiStickerSetThumbnail(name).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetCustomEmojiStickerSetThumbnail_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	
	sendResult := ctx.SetCustomEmojiStickerSetThumbnail(name).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetCustomEmojiStickerSetThumbnail Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
