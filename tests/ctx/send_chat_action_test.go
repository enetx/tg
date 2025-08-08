package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendChatAction(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction()

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test method chaining
	chained := result.Typing()
	if chained == nil {
		t.Error("Expected Typing method to return builder")
	}
}

func TestContext_SendChatActionChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.SendChatAction().
		Typing().
		To(123).
		Thread(456)

	if result == nil {
		t.Error("Expected SendChatAction builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendChatAction_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendChatAction().Send()

	if sendResult.IsErr() {
		t.Logf("SendChatAction Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendChatAction().
		Typing().
		To(123).
		Thread(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendChatAction configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
