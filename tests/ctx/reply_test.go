package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
	"github.com/enetx/tg/types/effects"
)

func TestContext_Reply(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Reply text")

	// Test basic creation
	result := testCtx.Reply(text)
	if result == nil {
		t.Error("Expected Reply builder to be created")
	}

	// Test HTML method
	result = result.HTML()
	if result == nil {
		t.Error("HTML method should return Reply for chaining")
	}

	// Test Markdown method
	result = testCtx.Reply(text).Markdown()
	if result == nil {
		t.Error("Markdown method should return Reply for chaining")
	}

	// Test Silent method
	result = testCtx.Reply(text).Silent()
	if result == nil {
		t.Error("Silent method should return Reply for chaining")
	}

	// Test Effect method
	result = testCtx.Reply(text).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return Reply for chaining")
	}

	// Test ReplyTo method
	result = testCtx.Reply(text).ReplyTo(123)
	if result == nil {
		t.Error("ReplyTo method should return Reply for chaining")
	}

	// Test AllowPaidBroadcast method
	result = testCtx.Reply(text).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return Reply for chaining")
	}

	// Test Thread method
	result = testCtx.Reply(text).Thread(456)
	if result == nil {
		t.Error("Thread method should return Reply for chaining")
	}

	// Test ForceReply method
	result = testCtx.Reply(text).ForceReply()
	if result == nil {
		t.Error("ForceReply method should return Reply for chaining")
	}

	// Test RemoveKeyboard method
	result = testCtx.Reply(text).RemoveKeyboard()
	if result == nil {
		t.Error("RemoveKeyboard method should return Reply for chaining")
	}

	// Test Business method
	result = testCtx.Reply(text).Business(g.String("business_123"))
	if result == nil {
		t.Error("Business method should return Reply for chaining")
	}

	// Test Protect method
	result = testCtx.Reply(text).Protect()
	if result == nil {
		t.Error("Protect method should return Reply for chaining")
	}

	// Test After method
	result = testCtx.Reply(text).After(5 * time.Second)
	if result == nil {
		t.Error("After method should return Reply for chaining")
	}

	// Test DeleteAfter method
	result = testCtx.Reply(text).DeleteAfter(10 * time.Second)
	if result == nil {
		t.Error("DeleteAfter method should return Reply for chaining")
	}

	// Test Timeout method
	result = testCtx.Reply(text).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return Reply for chaining")
	}

	// Test APIURL method
	result = testCtx.Reply(text).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return Reply for chaining")
	}
}

func TestContext_ReplyChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Complete reply with all options")

	// Test complete method chaining
	result := testCtx.Reply(text).
		HTML().
		Silent().
		Protect().
		AllowPaidBroadcast().
		Thread(123).
		ReplyTo(456).
		Business(g.String("biz_conn_789")).
		Effect(effects.Celebration).
		After(2 * time.Second).
		DeleteAfter(30 * time.Second).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return Reply")
	}
}

func TestReply_ComplexFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Reply with complex features")

	result := testCtx.Reply(text)

	// Test Entities method
	ent := entities.New("Reply with bold text").Bold(g.String("bold"))
	result = result.Entities(ent)
	if result == nil {
		t.Error("Entities method should return Reply for chaining")
	}

	// Test Preview method
	prev := preview.New().Disable()
	result = testCtx.Reply(text).Preview(prev)
	if result == nil {
		t.Error("Preview method should return Reply for chaining")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Test", "test_data")
	result = testCtx.Reply(text).Markup(kb)
	if result == nil {
		t.Error("Markup method should return Reply for chaining")
	}
}

func TestReply_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "original"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty text
	result := testCtx.Reply(g.String(""))
	if result == nil {
		t.Error("Reply should handle empty text")
	}

	// Test with zero values
	result = testCtx.Reply(g.String("test")).
		ReplyTo(0).
		Thread(0).
		After(0 * time.Second).
		DeleteAfter(0 * time.Second)

	if result == nil {
		t.Error("Reply should handle zero values gracefully")
	}

	// Test switching between parse modes
	result = testCtx.Reply(g.String("switching modes")).
		HTML().
		Markdown() // Should override HTML

	if result == nil {
		t.Error("Parse mode switching should work")
	}
}
