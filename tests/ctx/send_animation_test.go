package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendAnimation(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename)

	if result == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Animation caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendAnimationChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("animation.gif")

	result := ctx.SendAnimation(filename).
		Caption(g.String("Test animation")).
		HTML().
		Width(400).
		Height(300)

	if result == nil {
		t.Error("Expected SendAnimation builder to be created")
	}

	// Test continued chaining
	final := result.Silent()
	if final == nil {
		t.Error("Expected Silent method to return builder")
	}
}
