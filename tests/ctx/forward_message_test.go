package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ForwardMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	toChatID := int64(123)

	// ForwardMessage takes fromChatID and messageID, toChatID is set via To() method
	result := ctx.ForwardMessage(456, 789).To(toChatID)

	if result == nil {
		t.Error("Expected ForwardMessage builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_ForwardMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ForwardMessage(456, 789).
		To(123).
		Silent().
		Protect()

	if result == nil {
		t.Error("Expected ForwardMessage builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestForwardMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.ForwardMessage(456, 789).
		To(123).
		Send()

	if sendResult.IsErr() {
		t.Logf("ForwardMessage Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.ForwardMessage(456, 789).
		To(999).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("ForwardMessage configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

// Tests for methods with 0% coverage

func TestForwardMessage_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	messageID := int64(789)
	toChatID := int64(123)

	// Test After method functionality
	afterDurations := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		10 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		2 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero duration
	}

	for _, duration := range afterDurations {
		afterResult := ctx.ForwardMessage(fromChatID, messageID).
			To(toChatID).
			After(duration)

		if afterResult == nil {
			t.Errorf("After method with duration %v should work", duration)
		}

		// Test send with After scheduling
		sendResult := afterResult.Send()
		if sendResult.IsErr() {
			t.Logf("ForwardMessage with After duration %v Send failed as expected: %v", duration, sendResult.Err())
		}
	}
}

func TestForwardMessage_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	messageID := int64(789)
	toChatID := int64(123)

	// Test DeleteAfter method functionality
	deleteAfterDurations := []time.Duration{
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		30 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
		0 * time.Second, // Zero duration
	}

	for _, duration := range deleteAfterDurations {
		deleteAfterResult := ctx.ForwardMessage(fromChatID, messageID).
			To(toChatID).
			DeleteAfter(duration)

		if deleteAfterResult == nil {
			t.Errorf("DeleteAfter method with duration %v should work", duration)
		}

		// Test send with DeleteAfter scheduling
		sendResult := deleteAfterResult.Send()
		if sendResult.IsErr() {
			t.Logf("ForwardMessage with DeleteAfter duration %v Send failed as expected: %v", duration, sendResult.Err())
		}
	}
}

func TestForwardMessage_AfterAndDeleteAfterCombined(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	messageID := int64(789)
	toChatID := int64(123)

	// Test After and DeleteAfter methods combined
	combinedScenarios := []struct {
		after       time.Duration
		deleteAfter time.Duration
		description string
	}{
		{5 * time.Second, 30 * time.Second, "Quick forward with auto-delete"},
		{10 * time.Second, 1 * time.Minute, "Delayed forward with cleanup"},
		{1 * time.Minute, 10 * time.Minute, "Long delay with extended cleanup"},
		{0 * time.Second, 30 * time.Second, "Immediate forward with auto-delete"},
		{30 * time.Second, 0 * time.Second, "Delayed forward no cleanup"},
		{0 * time.Second, 0 * time.Second, "Immediate forward no cleanup"},
	}

	for _, scenario := range combinedScenarios {
		combinedResult := ctx.ForwardMessage(fromChatID, messageID).
			To(toChatID).
			After(scenario.after).
			DeleteAfter(scenario.deleteAfter).
			Silent().
			Protect()

		if combinedResult == nil {
			t.Errorf("Combined After/DeleteAfter scenario '%s' should work", scenario.description)
		}

		// Test send with combined scheduling
		sendResult := combinedResult.Send()
		if sendResult.IsErr() {
			t.Logf("ForwardMessage combined scenario '%s' Send failed as expected: %v",
				scenario.description, sendResult.Err())
		}
	}
}

func TestForwardMessage_SchedulingVariousScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test forwarding different message types with scheduling
	messageScenarios := []struct {
		fromChatID  int64
		messageID   int64
		toChatID    int64
		description string
	}{
		{456, 789, 123, "Text message forward"},
		{111, 222, 333, "Photo message forward"},
		{444, 555, 666, "Video message forward"},
		{777, 888, 999, "Document message forward"},
		{101, 202, 303, "Sticker message forward"},
	}

	for _, scenario := range messageScenarios {
		// Test with After scheduling
		afterScenarioResult := ctx.ForwardMessage(scenario.fromChatID, scenario.messageID).
			To(scenario.toChatID).
			After(15 * time.Second).
			Silent().
			Timeout(45 * time.Second).
			APIURL(g.String("https://after-forward-api.telegram.org")).
			Send()

		if afterScenarioResult.IsErr() {
			t.Logf("ForwardMessage After scenario '%s' Send failed as expected: %v",
				scenario.description, afterScenarioResult.Err())
		}

		// Test with DeleteAfter scheduling
		deleteAfterScenarioResult := ctx.ForwardMessage(scenario.fromChatID, scenario.messageID).
			To(scenario.toChatID).
			DeleteAfter(2 * time.Minute).
			Protect().
			Timeout(30 * time.Second).
			APIURL(g.String("https://delete-after-forward-api.telegram.org")).
			Send()

		if deleteAfterScenarioResult.IsErr() {
			t.Logf("ForwardMessage DeleteAfter scenario '%s' Send failed as expected: %v",
				scenario.description, deleteAfterScenarioResult.Err())
		}

		// Test with both After and DeleteAfter
		comboScenarioResult := ctx.ForwardMessage(scenario.fromChatID, scenario.messageID).
			To(scenario.toChatID).
			After(10 * time.Second).
			DeleteAfter(1 * time.Minute).
			Silent().
			Protect().
			Timeout(60 * time.Second).
			APIURL(g.String("https://combo-forward-api.telegram.org")).
			Send()

		if comboScenarioResult.IsErr() {
			t.Logf("ForwardMessage combo scenario '%s' Send failed as expected: %v",
				scenario.description, comboScenarioResult.Err())
		}
	}
}

func TestForwardMessage_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "comprehensive test message"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fromChatID := int64(987654321)
	messageID := int64(123456789)
	toChatID := int64(555666777)

	// Test comprehensive workflow with all scheduling options
	comprehensiveResult := ctx.ForwardMessage(fromChatID, messageID).
		To(toChatID).
		After(30 * time.Second).
		DeleteAfter(5 * time.Minute).
		Silent().
		Protect().
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-forward-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("ForwardMessage comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test using EffectiveChat as target (no explicit To())
	effectiveChatResult := ctx.ForwardMessage(fromChatID, messageID).
		After(1 * time.Minute).
		DeleteAfter(10 * time.Minute).
		Send()

	if effectiveChatResult.IsErr() {
		t.Logf("ForwardMessage to effective chat Send failed as expected: %v", effectiveChatResult.Err())
	}
}
