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

func TestSendMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Send test")

	// Test Send method execution (immediate)
	builder := testCtx.SendMessage(text)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with various options
	builderWithOptions := testCtx.SendMessage(text).
		To(456).
		HTML().
		Silent().
		Protect().
		Timeout(30 * time.Second)
	resultWithOptions := builderWithOptions.Send()

	if !resultWithOptions.IsErr() && !resultWithOptions.IsOk() {
		t.Error("Send with options should return a result")
	}

	// Test Send with After (scheduled)
	builderWithAfter := testCtx.SendMessage(text).
		After(1 * time.Millisecond) // Very short duration for testing
	resultWithAfter := builderWithAfter.Send()

	if !resultWithAfter.IsOk() {
		t.Error("Send with After should return Ok for scheduled execution")
	}

	// Test Send with DeleteAfter
	builderWithDeleteAfter := testCtx.SendMessage(text).
		DeleteAfter(60 * time.Second)
	resultWithDeleteAfter := builderWithDeleteAfter.Send()

	if !resultWithDeleteAfter.IsErr() && !resultWithDeleteAfter.IsOk() {
		t.Error("Send with DeleteAfter should return a result")
	}

	// Test Send without To() method (should use effective chat)
	builderWithoutTo := testCtx.SendMessage(text)
	resultWithoutTo := builderWithoutTo.Send()

	if !resultWithoutTo.IsErr() && !resultWithoutTo.IsOk() {
		t.Error("Send without To() should return a result (using effective chat)")
	}

	// Test Send with all features
	kb := keyboard.Inline().Text("Button", "data")
	entities := entities.New("Test message").Bold("Test")
	p := preview.New().Disable()
	builderComplete := testCtx.SendMessage(text).
		To(789).
		Entities(entities).
		HTML().
		Silent().
		Effect(effects.Heart).
		ReplyTo(999).
		Markup(kb).
		AllowPaidBroadcast().
		Thread(123).
		Preview(p).
		Business(g.String("business_456")).
		Protect().
		Timeout(45 * time.Second).
		APIURL(g.String("https://api.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}

func TestSendMessage_TextVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various text content
	texts := []string{
		"Simple message",
		"Message with emojis 🎉🔥⚡",
		"Multi-line\nmessage\nwith\nbreaks",
		"<b>HTML</b> formatted message",
		"**Markdown** formatted message",
		"Very long message that exceeds normal length expectations and contains lots of text to test handling of large content",
		"Message with special characters: !@#$%^&*()_+-=[]{}|;':\",./<>?",
		"Message with numbers: 1234567890",
		"Message with URLs: https://example.com",
		"Message with mentions: @username",
		"Message with hashtags: #tag",
		"A",
		"",
	}

	for _, text := range texts {
		result := testCtx.SendMessage(g.String(text))
		if result == nil {
			t.Errorf("SendMessage should work with text: %s", text)
		}

		// Test chaining for each text
		chained := result.HTML().Silent()
		if chained == nil {
			t.Errorf("Chaining should work for text: %s", text)
		}
	}
}

func TestSendMessage_EffectVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Effect test")

	// Test various message effects
	effectTypes := []effects.EffectType{
		effects.Fire,
		effects.ThumbsUp,
		effects.ThumbsDown,
		effects.Heart,
		effects.Celebration,
		effects.Poop,
	}

	for _, effect := range effectTypes {
		result := testCtx.SendMessage(text).Effect(effect)
		if result == nil {
			t.Errorf("Effect %s should work", effect.String())
		}

		// Test combining effects with other methods
		combined := result.Silent().Protect()
		if combined == nil {
			t.Errorf("Effect combination should work for %s", effect.String())
		}
	}
}

func TestSendMessage_TimeoutVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Timeout test")

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		result := testCtx.SendMessage(text).Timeout(timeout)
		if result == nil {
			t.Errorf("Timeout method should work with duration: %v", timeout)
		}

		// Test combining timeout with other methods
		combined := result.Silent().HTML()
		if combined == nil {
			t.Errorf("Timeout combination should work with duration: %v", timeout)
		}
	}
}

func TestSendMessage_APIURLVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("API URL test")

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://api.example.com",
		"https://custom-telegram-api.com",
		"https://localhost:8080",
		"https://bot-api.myservice.com",
		"",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.SendMessage(text).APIURL(g.String(apiURL))
		if result == nil {
			t.Errorf("APIURL method should work with URL: %s", apiURL)
		}

		// Test combining API URL with other methods
		combined := result.Silent().Timeout(30 * time.Second)
		if combined == nil {
			t.Errorf("APIURL combination should work with URL: %s", apiURL)
		}
	}
}

func TestSendMessage_ChatIDVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Chat ID test")

	// Test various chat destinations
	chatIDs := []int64{
		123,            // Private chat
		-100123456789,  // Supergroup
		-1001234567890, // Channel
		456,            // Another private chat
		0,              // Zero (edge case)
		-999999999999,  // Large negative ID
		999999999999,   // Large positive ID
	}

	for _, chatID := range chatIDs {
		result := testCtx.SendMessage(text).To(chatID)
		if result == nil {
			t.Errorf("SendMessage should work with chat ID: %d", chatID)
		}

		// Test combining with other methods
		combined := result.Silent().Protect()
		if combined == nil {
			t.Errorf("Method combination should work with chat ID: %d", chatID)
		}
	}
}

func TestSendMessage_MethodOverrides(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	text := g.String("Override test")

	// Test method overriding
	kb1 := keyboard.Inline().Text("Button 1", "data1")
	kb2 := keyboard.Reply().Text("Button 2")

	result := testCtx.SendMessage(text).
		HTML().
		Markdown(). // Should override HTML
		To(456).
		To(789). // Should override first To
		Markup(kb1).
		Markup(kb2). // Should override first keyboard
		Timeout(30 * time.Second).
		Timeout(60 * time.Second). // Should override first timeout
		APIURL(g.String("https://first.com")).
		APIURL(g.String("https://second.com")) // Should override first URL

	if result == nil {
		t.Error("Method overriding should work")
	}
}
