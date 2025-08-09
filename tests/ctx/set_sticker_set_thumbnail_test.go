package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerSetThumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")
	userID := int64(456)

	result := ctx.SetStickerSetThumbnail(name, userID)

	if result == nil {
		t.Error("Expected SetStickerSetThumbnail builder to be created")
	}

	// Test method chaining
	withFormat := result.Format(g.String("static"))
	if withFormat == nil {
		t.Error("Expected Format method to return builder")
	}
}

func TestSetStickerSetThumbnail_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).Thumbnail(g.String("thumbnail.png")) == nil { t.Error("Thumbnail should return builder") }
}

func TestSetStickerSetThumbnail_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetStickerSetThumbnail_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	if ctx.SetStickerSetThumbnail(name, userID).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetStickerSetThumbnail_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	userID := int64(456)
	
	sendResult := ctx.SetStickerSetThumbnail(name, userID).Format(g.String("static")).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetStickerSetThumbnail Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
