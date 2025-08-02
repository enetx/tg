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

func TestContextSendMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Hello World")

	// Test basic creation
	result := testCtx.SendMessage(text)
	if result == nil {
		t.Error("Expected SendMessage builder to be created")
	}

	// Test To method
	result = result.To(123456789)
	if result == nil {
		t.Error("To method should return SendMessage for chaining")
	}

	// Test HTML method
	result = testCtx.SendMessage(text).HTML()
	if result == nil {
		t.Error("HTML method should return SendMessage for chaining")
	}

	// Test Markdown method
	result = testCtx.SendMessage(text).Markdown()
	if result == nil {
		t.Error("Markdown method should return SendMessage for chaining")
	}

	// Test Silent method
	result = testCtx.SendMessage(text).Silent()
	if result == nil {
		t.Error("Silent method should return SendMessage for chaining")
	}

	// Test Effect method
	result = testCtx.SendMessage(text).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendMessage for chaining")
	}

	// Test ReplyTo method
	result = testCtx.SendMessage(text).ReplyTo(789)
	if result == nil {
		t.Error("ReplyTo method should return SendMessage for chaining")
	}

	// Test AllowPaidBroadcast method
	result = testCtx.SendMessage(text).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendMessage for chaining")
	}

	// Test Thread method
	result = testCtx.SendMessage(text).Thread(456)
	if result == nil {
		t.Error("Thread method should return SendMessage for chaining")
	}

	// Test ForceReply method
	result = testCtx.SendMessage(text).ForceReply()
	if result == nil {
		t.Error("ForceReply method should return SendMessage for chaining")
	}

	// Test RemoveKeyboard method
	result = testCtx.SendMessage(text).RemoveKeyboard()
	if result == nil {
		t.Error("RemoveKeyboard method should return SendMessage for chaining")
	}

	// Test Business method
	result = testCtx.SendMessage(text).Business(g.String("business_123"))
	if result == nil {
		t.Error("Business method should return SendMessage for chaining")
	}

	// Test Protect method
	result = testCtx.SendMessage(text).Protect()
	if result == nil {
		t.Error("Protect method should return SendMessage for chaining")
	}

	// Test After method
	result = testCtx.SendMessage(text).After(5 * time.Second)
	if result == nil {
		t.Error("After method should return SendMessage for chaining")
	}

	// Test DeleteAfter method
	result = testCtx.SendMessage(text).DeleteAfter(10 * time.Second)
	if result == nil {
		t.Error("DeleteAfter method should return SendMessage for chaining")
	}

	// Test Timeout method
	result = testCtx.SendMessage(text).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return SendMessage for chaining")
	}

	// Test APIURL method
	result = testCtx.SendMessage(text).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return SendMessage for chaining")
	}
}

func TestContextSendMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Complete message with all options")

	// Test complete method chaining
	result := testCtx.SendMessage(text).
		To(123456789).
		HTML().
		Silent().
		Effect(effects.Celebration).
		ReplyTo(789).
		AllowPaidBroadcast().
		Thread(456).
		Business(g.String("biz_conn_123")).
		Protect().
		After(2 * time.Second).
		DeleteAfter(30 * time.Second).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return SendMessage")
	}
}

func TestSendMessage_ComplexFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Message with complex features")

	result := testCtx.SendMessage(text)

	// Test Entities method
	ent := entities.New("Message with bold text").Bold(g.String("bold"))
	result = result.Entities(ent)
	if result == nil {
		t.Error("Entities method should return SendMessage for chaining")
	}

	// Test Preview method
	prev := preview.New().Disable()
	result = testCtx.SendMessage(text).Preview(prev)
	if result == nil {
		t.Error("Preview method should return SendMessage for chaining")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback_data")
	result = testCtx.SendMessage(text).Markup(kb)
	if result == nil {
		t.Error("Markup method should return SendMessage for chaining")
	}
}

func TestSendMessage_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty text
	result := testCtx.SendMessage(g.String(""))
	if result == nil {
		t.Error("SendMessage should handle empty text")
	}

	// Test with zero values
	result = testCtx.SendMessage(g.String("test")).
		To(0).
		ReplyTo(0).
		Thread(0).
		After(0 * time.Second).
		DeleteAfter(0 * time.Second)

	if result == nil {
		t.Error("SendMessage should handle zero values gracefully")
	}

	// Test switching between parse modes
	result = testCtx.SendMessage(g.String("switching modes")).
		HTML().
		Markdown() // Should override HTML

	if result == nil {
		t.Error("Parse mode switching should work")
	}

	// Test with maximum integer values
	result = testCtx.SendMessage(g.String("max values")).
		To(9223372036854775807).
		ReplyTo(9223372036854775807).
		Thread(9223372036854775807)

	if result == nil {
		t.Error("SendMessage should handle large integer values")
	}
}

func TestSendMessage_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test without setting To (should use EffectiveChat.Id)
	result := testCtx.SendMessage(g.String("default chat"))
	if result == nil {
		t.Error("SendMessage should work without explicit To")
	}

	// The Send() method should use EffectiveChat.Id when no To is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}
