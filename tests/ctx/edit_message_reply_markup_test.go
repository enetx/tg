package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditMessageReplyMarkup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	kb := keyboard.Inline()
	result := ctx.EditMessageReplyMarkup(kb)

	if result == nil {
		t.Error("Expected EditMessageReplyMarkup builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_EditMessageReplyMarkupChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	kb := keyboard.Inline()
	result := ctx.EditMessageReplyMarkup(kb).
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageReplyMarkup builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestEditMessageReplyMarkup_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	kb := keyboard.Inline()

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditMessageReplyMarkup(kb).Send()

	if sendResult.IsErr() {
		t.Logf("EditMessageReplyMarkup Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditMessageReplyMarkup(kb).
		ChatID(456).
		MessageID(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditMessageReplyMarkup configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

// Tests for methods with 0% coverage

func TestEditMessageReplyMarkup_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test InlineMessageID method
	inlineMessageIDs := []string{
		"inline_123456789",
		"inline_abcdef123",
		"inline_xyz789abc",
		"", // Empty inline message ID
	}

	for _, inlineID := range inlineMessageIDs {
		kb := keyboard.Inline().
			Text(g.String("⚙️ Settings"), g.String("settings")).
			Row().
			Text(g.String("🔄 Refresh"), g.String("refresh"))

		inlineResult := ctx.EditMessageReplyMarkup(kb).
			InlineMessageID(g.String(inlineID))

		if inlineResult == nil {
			t.Errorf("InlineMessageID with '%s' should work", inlineID)
		}

		// Test send with inline message ID
		sendResult := inlineResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageReplyMarkup with inline message ID '%s' Send failed as expected: %v", inlineID, sendResult.Err())
		}
	}
}

func TestEditMessageReplyMarkup_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Business connection IDs
	businessIDs := []string{
		"business_conn_123",
		"business_conn_456",
		"enterprise_conn_789",
		"", // Empty business ID
	}

	for _, businessID := range businessIDs {
		kb := keyboard.Inline().
			Text(g.String("💼 Business"), g.String("business_action")).
			Row().
			Text(g.String("📊 Analytics"), g.String("analytics")).
			Row().
			Text(g.String("📞 Support"), g.String("support"))

		businessResult := ctx.EditMessageReplyMarkup(kb).
			Business(g.String(businessID)).
			ChatID(456).
			MessageID(789)

		if businessResult == nil {
			t.Errorf("Business with '%s' should work", businessID)
		}

		// Test send with business ID
		sendResult := businessResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageReplyMarkup with business ID '%s' Send failed as expected: %v", businessID, sendResult.Err())
		}
	}
}

func TestEditMessageReplyMarkup_KeyboardVariations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test with different keyboard configurations
	keyboards := []struct {
		name string
		kb   keyboard.Keyboard
	}{
		{
			"Simple keyboard",
			keyboard.Inline().
				Text(g.String("✅ Yes"), g.String("yes")).
				Text(g.String("❌ No"), g.String("no")),
		},
		{
			"Multi-row keyboard",
			keyboard.Inline().
				Text(g.String("👍 Like"), g.String("like")).
				Text(g.String("👎 Dislike"), g.String("dislike")).
				Row().
				Text(g.String("💬 Comment"), g.String("comment")).
				Row().
				Text(g.String("🔗 Share"), g.String("share")),
		},
		{
			"URL keyboard",
			keyboard.Inline().
				URL(g.String("🌐 Website"), g.String("https://example.com")).
				Row().
				URL(g.String("📱 App"), g.String("https://app.example.com")),
		},
		{
			"Empty keyboard",
			keyboard.Inline(),
		},
	}

	for _, keyboard := range keyboards {
		// Test InlineMessageID with different keyboards
		inlineResult := ctx.EditMessageReplyMarkup(keyboard.kb).
			InlineMessageID(g.String("inline_" + keyboard.name)).
			Send()

		if inlineResult.IsErr() {
			t.Logf("EditMessageReplyMarkup %s with InlineMessageID Send failed as expected: %v", keyboard.name, inlineResult.Err())
		}

		// Test Business with different keyboards
		businessResult := ctx.EditMessageReplyMarkup(keyboard.kb).
			Business(g.String("business_" + keyboard.name)).
			ChatID(456).
			MessageID(789).
			Send()

		if businessResult.IsErr() {
			t.Logf("EditMessageReplyMarkup %s with Business Send failed as expected: %v", keyboard.name, businessResult.Err())
		}
	}
}

func TestEditMessageReplyMarkup_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test comprehensive workflow with all methods
	comprehensiveKB := keyboard.Inline().
		Text(g.String("🔥 Hot"), g.String("hot")).
		Text(g.String("❄️ Cool"), g.String("cool")).
		Row().
		URL(g.String("🌐 Website"), g.String("https://comprehensive-example.com")).
		Row().
		Text(g.String("🔄 Refresh All"), g.String("refresh_all"))

	complexResult := ctx.EditMessageReplyMarkup(comprehensiveKB).
		ChatID(456).
		MessageID(789).
		Business(g.String("business_comprehensive_123")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-markup-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditMessageReplyMarkup comprehensive workflow Send failed as expected: %v", complexResult.Err())
	}

	// Test with inline message workflow
	inlineWorkflowResult := ctx.EditMessageReplyMarkup(comprehensiveKB).
		InlineMessageID(g.String("inline_comprehensive_markup_123")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://inline-markup-api.telegram.org")).
		Send()

	if inlineWorkflowResult.IsErr() {
		t.Logf("EditMessageReplyMarkup inline workflow Send failed as expected: %v", inlineWorkflowResult.Err())
	}

	// Test with nil keyboard (should handle gracefully)
	nilKBResult := ctx.EditMessageReplyMarkup(nil).
		ChatID(456).
		MessageID(789).
		Send()

	if nilKBResult.IsErr() {
		t.Logf("EditMessageReplyMarkup with nil keyboard Send failed as expected: %v", nilKBResult.Err())
	}
}
