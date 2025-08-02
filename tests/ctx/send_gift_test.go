package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendGift(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	giftID := g.String("gift_123")

	result := ctx.SendGift(giftID)

	if result == nil {
		t.Error("Expected SendGift builder to be created")
	}

	// Test method chaining
	chained := result.Text(g.String("Happy Birthday!"))
	if chained == nil {
		t.Error("Expected Text method to return builder")
	}
}

func TestContext_SendGiftChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	giftID := g.String("gift_123")

	result := ctx.SendGift(giftID).
		Text(g.String("Happy Birthday!"))

	if result == nil {
		t.Error("Expected SendGift builder to be created")
	}

	// Test that builder is functional
	_ = result
}
