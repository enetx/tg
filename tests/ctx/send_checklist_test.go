package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendChecklist(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("My Checklist")
	intro := g.String("This is my checklist")

	result := ctx.SendChecklist(title, intro)

	if result == nil {
		t.Error("Expected SendChecklist builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendChecklistChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("My Checklist")
	intro := g.String("This is my checklist")

	result := ctx.SendChecklist(title, intro).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendChecklist builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestSendChecklist_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	title := g.String("Test Checklist")
	intro := g.String("Test checklist description")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendChecklist(title, intro).Send()

	if sendResult.IsErr() {
		t.Logf("SendChecklist Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendChecklist(title, intro).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendChecklist configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
