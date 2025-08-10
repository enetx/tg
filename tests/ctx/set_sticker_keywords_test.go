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

func TestContext_SetStickerKeywords(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("sticker_id_123")

	result := ctx.SetStickerKeywords(file.Input(sticker).UnwrapOrDefault())

	if result == nil {
		t.Error("Expected SetStickerKeywords builder to be created")
	}

	// Test method chaining
	keywords := g.Slice[g.String]{}
	keywords.Push(g.String("happy"))
	withKeywords := result.Keywords(keywords)
	if withKeywords == nil {
		t.Error("Expected Keywords method to return builder")
	}
}

func TestSetStickerKeywords_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	if ctx.SetStickerKeywords(file.Input(sticker).UnwrapOrDefault()).Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSetStickerKeywords_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")
	if ctx.SetStickerKeywords(file.Input(sticker).UnwrapOrDefault()).APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSetStickerKeywords_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("sticker_id_123")

	keywords := g.Slice[g.String]{}
	keywords.Push(g.String("happy"))
	keywords.Push(g.String("smile"))

	sendResult := ctx.SetStickerKeywords(file.Input(sticker).UnwrapOrDefault()).Keywords(keywords).Send()

	if sendResult.IsErr() {
		t.Logf("SetStickerKeywords Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
