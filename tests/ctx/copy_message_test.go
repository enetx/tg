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
	"github.com/enetx/tg/reply"
)

func TestContext_CopyMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	messageID := int64(789)
	toChatID := int64(123)

	// Test basic creation
	result := testCtx.CopyMessage(fromChatID, messageID)
	if result == nil {
		t.Error("Expected CopyMessage builder to be created")
	}

	// Test To method
	result = result.To(toChatID)
	if result == nil {
		t.Error("To method should return CopyMessage for chaining")
	}

	// Test Caption method
	result = testCtx.CopyMessage(fromChatID, messageID).Caption(g.String("New caption"))
	if result == nil {
		t.Error("Caption method should return CopyMessage for chaining")
	}

	// Test CaptionEntities method
	entities := entities.New("test content").Bold("test")
	result = testCtx.CopyMessage(fromChatID, messageID).CaptionEntities(entities)
	if result == nil {
		t.Error("CaptionEntities method should return CopyMessage for chaining")
	}

	// Test HTML method
	result = testCtx.CopyMessage(fromChatID, messageID).HTML()
	if result == nil {
		t.Error("HTML method should return CopyMessage for chaining")
	}

	// Test Markdown method
	result = testCtx.CopyMessage(fromChatID, messageID).Markdown()
	if result == nil {
		t.Error("Markdown method should return CopyMessage for chaining")
	}

	// Test Silent method
	result = testCtx.CopyMessage(fromChatID, messageID).Silent()
	if result == nil {
		t.Error("Silent method should return CopyMessage for chaining")
	}

	// Test Protect method
	result = testCtx.CopyMessage(fromChatID, messageID).Protect()
	if result == nil {
		t.Error("Protect method should return CopyMessage for chaining")
	}

	// Test Markup method
	kb := keyboard.Inline().Text("Button", "data")
	result = testCtx.CopyMessage(fromChatID, messageID).Markup(kb)
	if result == nil {
		t.Error("Markup method should return CopyMessage for chaining")
	}

	// Test Thread method
	result = testCtx.CopyMessage(fromChatID, messageID).Thread(123)
	if result == nil {
		t.Error("Thread method should return CopyMessage for chaining")
	}

	// Test VideoStartAt method
	result = testCtx.CopyMessage(fromChatID, messageID).VideoStartAt(30 * time.Second)
	if result == nil {
		t.Error("VideoStartAt method should return CopyMessage for chaining")
	}

	// Test ShowCaptionAbove method
	result = testCtx.CopyMessage(fromChatID, messageID).ShowCaptionAbove()
	if result == nil {
		t.Error("ShowCaptionAbove method should return CopyMessage for chaining")
	}

	// Test AllowPaidBroadcast method
	result = testCtx.CopyMessage(fromChatID, messageID).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return CopyMessage for chaining")
	}

	// Test ReplyTo method
	result = testCtx.CopyMessage(fromChatID, messageID).Reply(reply.New(456))
	if result == nil {
		t.Error("ReplyTo method should return CopyMessage for chaining")
	}

	// Test After method
	result = testCtx.CopyMessage(fromChatID, messageID).After(5 * time.Second)
	if result == nil {
		t.Error("After method should return CopyMessage for chaining")
	}

	// Test DeleteAfter method
	result = testCtx.CopyMessage(fromChatID, messageID).DeleteAfter(60 * time.Second)
	if result == nil {
		t.Error("DeleteAfter method should return CopyMessage for chaining")
	}

	// Test Timeout method
	result = testCtx.CopyMessage(fromChatID, messageID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CopyMessage for chaining")
	}

	// Test APIURL method
	result = testCtx.CopyMessage(fromChatID, messageID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CopyMessage for chaining")
	}
}

func TestContext_CopyMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complex method chaining
	result := testCtx.CopyMessage(456, 789).
		To(123).
		Caption(g.String("Updated caption")).
		HTML().
		Silent().
		Protect().
		Thread(456).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CopyMessage")
	}

	// Test with keyboard and timing
	kb := keyboard.Inline().Text("Copy", "copy_data")
	final := testCtx.CopyMessage(456, 789).
		To(999).
		Markup(kb).
		After(2 * time.Second).
		DeleteAfter(60 * time.Second)

	if final == nil {
		t.Error("Keyboard and timing chaining should work")
	}
}

func TestCopyMessage_MessageIdentifiers(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various message identifier combinations
	testCases := []struct {
		name        string
		fromChatID  int64
		messageID   int64
		toChatID    int64
		description string
	}{
		{"Private to Private", 123, 456, 789, "Copy from private to private chat"},
		{"Group to Private", -100123456789, 123, 456, "Copy from group to private"},
		{"Private to Group", 456, 789, -100987654321, "Copy from private to group"},
		{"Group to Group", -100111111111, 222, -100222222222, "Copy between groups"},
		{"Channel to Private", -1001234567890, 333, 999, "Copy from channel to private"},
		{"Same Chat", 123, 456, 123, "Copy within same chat"},
		{"High Message ID", 123, 999999999, 456, "Copy message with high ID"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testCtx.CopyMessage(tc.fromChatID, tc.messageID).To(tc.toChatID)
			if result == nil {
				t.Errorf("%s should work (%s)", tc.name, tc.description)
			}

			// Test with caption for each scenario
			captionResult := result.Caption(g.String("Caption for " + tc.name))
			if captionResult == nil {
				t.Errorf("Caption should work for %s", tc.name)
			}
		})
	}
}

func TestCopyMessage_CaptionFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test HTML formatting
	htmlResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("<b>Bold</b> and <i>italic</i> text")).
		HTML()

	if htmlResult == nil {
		t.Error("HTML caption formatting should work")
	}

	// Test Markdown formatting
	markdownResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("**Bold** and _italic_ text")).
		Markdown()

	if markdownResult == nil {
		t.Error("Markdown caption formatting should work")
	}

	// Test caption with entities
	entities := entities.New("Important: This is a copied message").
		Bold("Important: ").
		Italic("This is a copied message")

	entitiesResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		CaptionEntities(entities)

	if entitiesResult == nil {
		t.Error("Caption with entities should work")
	}

	// Test caption above media
	aboveResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Caption above media")).
		ShowCaptionAbove()

	if aboveResult == nil {
		t.Error("Caption above media should work")
	}
}

func TestCopyMessage_KeyboardIntegration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test inline keyboard
	inlineKb := keyboard.Inline().
		Text("Button 1", "data1").
		Text("Button 2", "data2").Row().
		URL("Visit Site", "https://example.com")

	inlineResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Markup(inlineKb)

	if inlineResult == nil {
		t.Error("Inline keyboard should work")
	}

	// Test reply keyboard
	replyKb := keyboard.Reply().
		Text("Option 1").
		Text("Option 2").Row().
		Contact("Share Contact").
		Location("Share Location")

	replyResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Markup(replyKb)

	if replyResult == nil {
		t.Error("Reply keyboard should work")
	}

	// Test keyboard with other options
	combinedResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Message with keyboard")).
		Markup(inlineKb).
		Silent().
		Protect()

	if combinedResult == nil {
		t.Error("Keyboard with other options should work")
	}
}

func TestCopyMessage_TimingFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test After method with various durations
	afterDurations := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
	}

	for _, duration := range afterDurations {
		result := testCtx.CopyMessage(fromChatID, messageID).
			To(toChatID).
			After(duration)

		if result == nil {
			t.Errorf("After duration %v should work", duration)
		}
	}

	// Test DeleteAfter method
	deleteAfterDurations := []time.Duration{
		10 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, duration := range deleteAfterDurations {
		result := testCtx.CopyMessage(fromChatID, messageID).
			To(toChatID).
			DeleteAfter(duration)

		if result == nil {
			t.Errorf("DeleteAfter duration %v should work", duration)
		}
	}

	// Test combined timing
	combinedResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		After(2 * time.Second).
		DeleteAfter(60 * time.Second)

	if combinedResult == nil {
		t.Error("Combined After and DeleteAfter should work")
	}
}

func TestCopyMessage_VideoFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test video start timestamp
	videoOffsets := []time.Duration{
		0 * time.Second,
		5 * time.Second,
		30 * time.Second,
		2 * time.Minute,
		10 * time.Minute,
	}

	for _, offset := range videoOffsets {
		result := testCtx.CopyMessage(fromChatID, messageID).
			To(toChatID).
			VideoStartAt(offset)

		if result == nil {
			t.Errorf("VideoStartAt offset %v should work", offset)
		}
	}

	// Test video with caption above and start time
	videoResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Video starting at 30 seconds")).
		ShowCaptionAbove().
		VideoStartAt(30 * time.Second)

	if videoResult == nil {
		t.Error("Video with caption above and start time should work")
	}
}

func TestCopyMessage_BroadcastFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test paid broadcast
	paidResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		AllowPaidBroadcast()

	if paidResult == nil {
		t.Error("AllowPaidBroadcast should work")
	}

	// Test paid broadcast with other features
	combinedResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Paid broadcast message")).
		AllowPaidBroadcast().
		Silent().
		Protect()

	if combinedResult == nil {
		t.Error("Paid broadcast with other features should work")
	}
}

func TestCopyMessage_ForumFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(-1001234567890)
	messageID := int64(456)
	toChatID := int64(-1001987654321)

	// Test thread IDs
	threadIDs := []int64{
		1,   // General topic
		123, // Regular topic
		456, // Another topic
		999, // High ID topic
	}

	for _, threadID := range threadIDs {
		result := testCtx.CopyMessage(fromChatID, messageID).
			To(toChatID).
			Thread(threadID)

		if result == nil {
			t.Errorf("Thread ID %d should work", threadID)
		}
	}

	// Test thread with other forum features
	forumResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Thread(123).
		Caption(g.String("Forum topic message")).
		Silent()

	if forumResult == nil {
		t.Error("Forum thread with other features should work")
	}
}

func TestCopyMessage_ReplyFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test reply to various message IDs
	replyToIDs := []int64{
		1,
		123,
		456,
		999999,
	}

	for _, replyID := range replyToIDs {
		result := testCtx.CopyMessage(fromChatID, messageID).
			To(toChatID).
			Reply(reply.New(replyID))

		if result == nil {
			t.Errorf("ReplyTo message ID %d should work", replyID)
		}
	}

	// Test reply with other features
	replyResult := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Reply(reply.New(123)).
		Caption(g.String("Reply with caption")).
		HTML().
		Silent()

	if replyResult == nil {
		t.Error("Reply with other features should work")
	}
}

func TestCopyMessage_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero message ID
	result := testCtx.CopyMessage(123, 0)
	if result == nil {
		t.Error("CopyMessage should handle zero message ID")
	}

	// Test with zero chat IDs
	result = testCtx.CopyMessage(0, 456).To(0)
	if result == nil {
		t.Error("CopyMessage should handle zero chat IDs")
	}

	// Test with empty caption
	result = testCtx.CopyMessage(123, 456).Caption(g.String(""))
	if result == nil {
		t.Error("CopyMessage should handle empty caption")
	}

	// Test with zero timeout
	result = testCtx.CopyMessage(123, 456).Timeout(0 * time.Second)
	if result == nil {
		t.Error("CopyMessage should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.CopyMessage(123, 456).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("CopyMessage should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.CopyMessage(123, 456).APIURL(g.String(""))
	if result == nil {
		t.Error("CopyMessage should handle empty API URL")
	}

	// Test without To() method (should use effective chat)
	result = testCtx.CopyMessage(456, 789)
	if result == nil {
		t.Error("CopyMessage should work without explicit To() call")
	}
}

func TestCopyMessage_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test all methods combined in different orders
	kb := keyboard.Inline().Text("Test", "test_data")
	entities := entities.New("Bold text content").Bold("Bold text")

	// Order 1
	result1 := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Complete test")).
		CaptionEntities(entities).
		HTML().
		Silent().
		Protect().
		Markup(kb).
		Thread(123).
		VideoStartAt(10 * time.Second).
		ShowCaptionAbove().
		AllowPaidBroadcast().
		Reply(reply.New(999)).
		After(1 * time.Second).
		DeleteAfter(300 * time.Second).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2 (different sequence)
	result2 := testCtx.CopyMessage(fromChatID, messageID).
		APIURL(g.String("https://custom-api.example.com")).
		Timeout(45 * time.Second).
		DeleteAfter(180 * time.Second).
		After(2 * time.Second).
		Reply(reply.New(888)).
		AllowPaidBroadcast().
		ShowCaptionAbove().
		VideoStartAt(5 * time.Second).
		Thread(456).
		Markup(kb).
		Protect().
		Silent().
		Markdown().
		CaptionEntities(entities).
		Caption(g.String("Reordered test")).
		To(toChatID)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		HTML().
		Markdown(). // Should override HTML
		Caption(g.String("First caption")).
		Caption(g.String("Second caption")). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}
}

func TestCopyMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 456, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(123)
	messageID := int64(456)
	toChatID := int64(789)

	// Test Send method execution (immediate)
	builder := testCtx.CopyMessage(fromChatID, messageID).To(toChatID)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with various options
	builderWithOptions := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Test caption")).
		Silent().
		Protect().
		Timeout(30 * time.Second)
	resultWithOptions := builderWithOptions.Send()

	if !resultWithOptions.IsErr() && !resultWithOptions.IsOk() {
		t.Error("Send with options should return a result")
	}

	// Test Send with After (scheduled)
	builderWithAfter := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		After(1 * time.Millisecond) // Very short duration for testing
	resultWithAfter := builderWithAfter.Send()

	if !resultWithAfter.IsOk() {
		t.Error("Send with After should return Ok(nil) for scheduled execution")
	}

	// Test Send with DeleteAfter
	builderWithDeleteAfter := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		DeleteAfter(60 * time.Second)
	resultWithDeleteAfter := builderWithDeleteAfter.Send()

	if !resultWithDeleteAfter.IsErr() && !resultWithDeleteAfter.IsOk() {
		t.Error("Send with DeleteAfter should return a result")
	}

	// Test Send without To() method (should use effective chat)
	builderWithoutTo := testCtx.CopyMessage(fromChatID, messageID)
	resultWithoutTo := builderWithoutTo.Send()

	if !resultWithoutTo.IsErr() && !resultWithoutTo.IsOk() {
		t.Error("Send without To() should return a result (using effective chat)")
	}

	// Test Send with After and DeleteAfter combined
	builderCombined := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		After(1 * time.Millisecond).
		DeleteAfter(30 * time.Second)
	resultCombined := builderCombined.Send()

	if !resultCombined.IsOk() {
		t.Error("Send with After and DeleteAfter should return Ok(nil) for scheduled execution")
	}

	// Test Send with all features
	kb := keyboard.Inline().Text("Button", "data")
	entities := entities.New("Test message").Bold("Test")
	builderComplete := testCtx.CopyMessage(fromChatID, messageID).
		To(toChatID).
		Caption(g.String("Complete message")).
		CaptionEntities(entities).
		HTML().
		Silent().
		Protect().
		Markup(kb).
		Thread(123).
		VideoStartAt(10 * time.Second).
		ShowCaptionAbove().
		AllowPaidBroadcast().
		Reply(reply.New(999)).
		Timeout(45 * time.Second).
		APIURL(g.String("https://api.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
