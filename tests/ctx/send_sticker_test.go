package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendSticker(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	result := ctx.SendSticker(sticker)

	if result == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendStickerChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	result := ctx.SendSticker(sticker).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test continued chaining
	final := result.Thread(456)
	if final == nil {
		t.Error("Expected Thread method to return builder")
	}
}
