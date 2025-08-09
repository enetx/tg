package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetStickerSetTitle(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	name := g.String("test_sticker_set")
	title := g.String("Updated Sticker Set Title")

	result := ctx.SetStickerSetTitle(name, title)

	if result == nil {
		t.Error("Expected SetStickerSetTitle builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(time.Second * 30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestSetStickerSetTitle_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	title := g.String("Updated Title")
	if ctx.SetStickerSetTitle(name, title).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetStickerSetTitle_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	name := g.String("test_sticker_set")
	title := g.String("Updated Title")
	
	sendResult := ctx.SetStickerSetTitle(name, title).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetStickerSetTitle Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
