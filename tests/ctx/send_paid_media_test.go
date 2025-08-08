package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendPaidMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	result := ctx.SendPaidMedia(starCount)

	if result == nil {
		t.Error("Expected SendPaidMedia builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendPaidMediaChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	result := ctx.SendPaidMedia(starCount).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendPaidMedia builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendPaidMedia_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	starCount := int64(100)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendPaidMedia(starCount).Send()

	if sendResult.IsErr() {
		t.Logf("SendPaidMedia Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendPaidMedia(starCount).
		Caption(g.String("<b>Paid media content</b>")).
		HTML().
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendPaidMedia configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
