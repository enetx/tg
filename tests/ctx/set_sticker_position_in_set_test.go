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

func TestContext_SetStickerPositionInSet(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")
	position := int64(2)

	result := ctx.SetStickerPositionInSet(file.Input(sticker).UnwrapOrDefault(), position)

	if result == nil {
		t.Error("Expected SetStickerPositionInSet builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetStickerPositionInSet_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	position := int64(2)
	if ctx.SetStickerPositionInSet(file.Input(sticker).UnwrapOrDefault(), position).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetStickerPositionInSet_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	position := int64(2)
	
	sendResult := ctx.SetStickerPositionInSet(file.Input(sticker).UnwrapOrDefault(), position).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetStickerPositionInSet Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
