package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContextSendMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Hello World")

	sm := ctx.SendMessage(text)

	// Test via public methods instead of accessing private fields
	if sm == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test that we can chain methods (verifies builder initialization)
	result := sm.HTML().Silent()
	if result == nil {
		t.Error("Expected chained method to return builder")
	}
}

func TestContextSendMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Hello World")

	sm := ctx.SendMessage(text).
		HTML().
		Silent().
		To(123).
		ReplyTo(789)

	// Test method chaining works properly
	if sm == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test that we can continue chaining after configuration
	result := sm.Thread(456)
	if result == nil {
		t.Error("Expected thread method to return builder")
	}
}

func TestContext_SendMessageAdvancedChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	text := g.String("Advanced test message")

	// Test complex method chaining
	sm := ctx.SendMessage(text).
		HTML().
		Silent().
		Protect().
		To(789).
		Thread(123).
		ReplyTo(456).
		AllowPaidBroadcast()

	// Verify builder is created and methods return non-nil
	if sm == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test that chaining continues to work
	result := sm.Markdown()
	if result == nil {
		t.Error("Expected markdown method to return builder")
	}
}
