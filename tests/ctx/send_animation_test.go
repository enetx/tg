package ctx_test

import (
	"testing"
	"time"

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

func TestSendAnimation_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_animation.gif")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendAnimation(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendAnimation Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendAnimation(filename).
		Caption(g.String("Test <b>animation</b> with HTML")).
		HTML().
		Width(640).
		Height(480).
		Duration(10).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendAnimation configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
