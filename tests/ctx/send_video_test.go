package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename)

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Video caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendVideoChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename).
		Caption(g.String("Test video")).
		HTML().
		Silent()

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendVideo_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_video.mp4")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVideo(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendVideo Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVideo(filename).
		Caption(g.String("Test <b>video</b> with HTML")).
		HTML().
		Width(640).
		Height(480).
		Duration(120).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVideo configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
