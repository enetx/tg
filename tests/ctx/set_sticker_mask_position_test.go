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

func TestContext_SetStickerMaskPosition(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")

	result := ctx.SetStickerMaskPosition(file.Input(sticker).UnwrapOrDefault())

	if result == nil {
		t.Error("Expected SetStickerMaskPosition builder to be created")
	}

	// Test method chaining
	withMask := result.MaskPosition(g.String("forehead"), 0.5, 0.5, 1.0)
	if withMask == nil {
		t.Error("Expected MaskPosition method to return builder")
	}
}

func TestSetStickerMaskPosition_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	if ctx.SetStickerMaskPosition(file.Input(sticker).UnwrapOrDefault()).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSetStickerMaskPosition_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	if ctx.SetStickerMaskPosition(file.Input(sticker).UnwrapOrDefault()).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetStickerMaskPosition_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")

	sendResult := ctx.SetStickerMaskPosition(file.Input(sticker).UnwrapOrDefault()).MaskPosition(g.String("forehead"), 0.5, 0.5, 1.0).Send()

	if sendResult.IsErr() {
		t.Logf("SetStickerMaskPosition Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
