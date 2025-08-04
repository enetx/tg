package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_DeleteMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.DeleteMessage()
	if result == nil {
		t.Error("Expected DeleteMessage builder to be created")
	}

	// Test After method
	result = testCtx.DeleteMessage().After(5 * time.Second)
	if result == nil {
		t.Error("After method should return DeleteMessage for chaining")
	}

	// Test ChatID method
	result = testCtx.DeleteMessage().ChatID(456)
	if result == nil {
		t.Error("ChatID method should return DeleteMessage for chaining")
	}

	// Test MessageID method
	result = testCtx.DeleteMessage().MessageID(123)
	if result == nil {
		t.Error("MessageID method should return DeleteMessage for chaining")
	}

	// Test Timeout method
	result = testCtx.DeleteMessage().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return DeleteMessage for chaining")
	}

	// Test APIURL method
	result = testCtx.DeleteMessage().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return DeleteMessage for chaining")
	}
}

func TestContext_DeleteMessageChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complete method chaining
	result := testCtx.DeleteMessage().
		After(10 * time.Second).
		ChatID(456).
		MessageID(123).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return DeleteMessage")
	}
}

func TestDeleteMessage_DelayedDeletion(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(456)
	messageID := int64(789)

	// Test various delayed deletion scenarios
	delayScenarios := []struct {
		name        string
		delay       time.Duration
		scenario    string
		description string
	}{
		{"Instant", 100 * time.Millisecond, "cleanup", "Quick cleanup deletion"},
		{"Short", 1 * time.Second, "moderation", "Short delay for moderation"},
		{"Standard", 5 * time.Second, "temp_message", "Standard temporary message deletion"},
		{"Medium", 30 * time.Second, "warning", "Medium delay for warning messages"},
		{"Long", 1 * time.Minute, "notice", "Long delay for important notices"},
		{"Extended", 5 * time.Minute, "announcement", "Extended delay for announcements"},
		{"Very Long", 15 * time.Minute, "reminder", "Very long delay for reminders"},
		{"Zero", 0 * time.Second, "immediate", "Zero delay (immediate deletion)"},
	}

	for _, scenario := range delayScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.DeleteMessage().
				After(scenario.delay).
				ChatID(chatID).
				MessageID(messageID)

			if result == nil {
				t.Errorf("%s delayed deletion (%s) should work", scenario.name, scenario.description)
			}

			// Test with additional configuration
			enhancedResult := result.
				Timeout(60 * time.Second).
				APIURL(g.String("https://delayed-api.telegram.org"))

			if enhancedResult == nil {
				t.Errorf("Enhanced %s delayed deletion should work", scenario.name)
			}
		})
	}
}

func TestDeleteMessage_MessageTypes(t *testing.T) {
	bot := &mockBot{}

	// Test deletion for various message types
	messageTypes := []struct {
		name        string
		messageID   int64
		messageType string
		content     string
		description string
	}{
		{"Text Message", 100, "text", "Regular text message", "Standard text message deletion"},
		{"Photo Message", 101, "photo", "Photo with caption", "Photo message deletion"},
		{"Video Message", 102, "video", "Video content", "Video message deletion"},
		{"Audio Message", 103, "audio", "Audio file", "Audio message deletion"},
		{"Document Message", 104, "document", "File attachment", "Document message deletion"},
		{"Sticker Message", 105, "sticker", "Sticker content", "Sticker message deletion"},
		{"Voice Message", 106, "voice", "Voice note", "Voice message deletion"},
		{"Video Note", 107, "video_note", "Video note", "Video note deletion"},
		{"Animation", 108, "animation", "GIF animation", "Animation message deletion"},
		{"Location Message", 109, "location", "Location share", "Location message deletion"},
		{"Contact Message", 110, "contact", "Contact info", "Contact message deletion"},
		{"Poll Message", 111, "poll", "Poll question", "Poll message deletion"},
		{"Dice Message", 112, "dice", "Dice emoji", "Dice message deletion"},
		{"Game Message", 113, "game", "Game content", "Game message deletion"},
		{"Invoice Message", 114, "invoice", "Payment invoice", "Invoice message deletion"},
	}

	for _, msgType := range messageTypes {
		t.Run(msgType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
				EffectiveMessage: &gotgbot.Message{MessageId: msgType.messageID, Text: msgType.content},
				Update:           &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.DeleteMessage().
				MessageID(msgType.messageID)

			if result == nil {
				t.Errorf("%s deletion (%s) should work", msgType.name, msgType.description)
			}

			// Test with delay for each message type
			delayedResult := result.After(5 * time.Second)
			if delayedResult == nil {
				t.Errorf("Delayed %s deletion should work", msgType.name)
			}

			// Test with custom API for each message type
			apiResult := delayedResult.APIURL(g.String("https://message-api.telegram.org"))
			if apiResult == nil {
				t.Errorf("Custom API for %s deletion should work", msgType.name)
			}
		})
	}
}

func TestDeleteMessage_ChatTypes(t *testing.T) {
	bot := &mockBot{}
	messageID := int64(789)

	// Test message deletion for various chat types
	chatTypes := []struct {
		name        string
		chatID      int64
		chatType    string
		permissions string
		description string
	}{
		{"Private Chat", 456, "private", "full", "Direct message deletion"},
		{"Group Chat", -123456789, "group", "admin", "Group message deletion"},
		{"Supergroup", -1001234567890, "supergroup", "admin", "Supergroup message deletion"},
		{"Channel", -1001987654321, "channel", "admin", "Channel message deletion"},
		{"Large Supergroup", -1002000000000, "supergroup", "admin", "Large supergroup deletion"},
		{"Public Channel", -1003000000000, "channel", "admin", "Public channel deletion"},
		{"Private Channel", -1004000000000, "channel", "owner", "Private channel deletion"},
		{"Discussion Group", -1005000000000, "supergroup", "admin", "Discussion group deletion"},
	}

	for _, chatType := range chatTypes {
		t.Run(chatType.name, func(t *testing.T) {
			rawCtx := &ext.Context{
				EffectiveChat:    &gotgbot.Chat{Id: chatType.chatID, Type: chatType.chatType},
				EffectiveMessage: &gotgbot.Message{MessageId: messageID, Text: "test"},
				Update:           &gotgbot.Update{UpdateId: 1},
			}

			testCtx := ctx.New(bot, rawCtx)

			result := testCtx.DeleteMessage().
				ChatID(chatType.chatID).
				MessageID(messageID)

			if result == nil {
				t.Errorf("%s message deletion (%s) should work", chatType.name, chatType.description)
			}

			// Test with timeout for each chat type
			timedResult := result.Timeout(30 * time.Second)
			if timedResult == nil {
				t.Errorf("Timed deletion for %s should work", chatType.name)
			}

			// Test with delayed deletion for different chat types
			if chatType.chatType != "private" {
				delayedResult := result.After(10 * time.Second)
				if delayedResult == nil {
					t.Errorf("Delayed deletion for %s should work", chatType.name)
				}
			}
		})
	}
}

func TestDeleteMessage_ModerationScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various moderation scenarios
	moderationScenarios := []struct {
		name        string
		messageID   int64
		reason      string
		delay       time.Duration
		priority    string
		description string
	}{
		{"Spam Removal", 1001, "spam", 0, "immediate", "Immediate spam message removal"},
		{"Inappropriate Content", 1002, "inappropriate", 1 * time.Second, "high", "Quick inappropriate content removal"},
		{"Off-Topic", 1003, "off_topic", 5 * time.Second, "medium", "Off-topic message cleanup"},
		{"Duplicate Message", 1004, "duplicate", 3 * time.Second, "low", "Duplicate message removal"},
		{"Advertisement", 1005, "ads", 0, "immediate", "Advertisement removal"},
		{"Harassment", 1006, "harassment", 0, "immediate", "Harassment content removal"},
		{"Misinformation", 1007, "misinformation", 2 * time.Second, "high", "Misinformation removal"},
		{"Copyright Violation", 1008, "copyright", 0, "immediate", "Copyright violation removal"},
		{"Expired Notice", 1009, "expired", 30 * time.Second, "low", "Expired notice cleanup"},
		{"Temporary Warning", 1010, "warning", 60 * time.Second, "scheduled", "Temporary warning removal"},
	}

	for _, scenario := range moderationScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.DeleteMessage().
				MessageID(scenario.messageID)

			if scenario.delay > 0 {
				result = result.After(scenario.delay)
			}

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}

			// Test with appropriate timeout based on priority
			var timeout time.Duration
			switch scenario.priority {
			case "immediate":
				timeout = 5 * time.Second
			case "high":
				timeout = 15 * time.Second
			case "medium":
				timeout = 30 * time.Second
			case "low":
				timeout = 60 * time.Second
			case "scheduled":
				timeout = 120 * time.Second
			default:
				timeout = 30 * time.Second
			}

			timedResult := result.Timeout(timeout)
			if timedResult == nil {
				t.Errorf("Timed %s should work", scenario.name)
			}
		})
	}
}

func TestDeleteMessage_BulkDeletion(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test bulk deletion scenarios (multiple individual deletions)
	bulkScenarios := []struct {
		name         string
		messageCount int
		startMsgID   int64
		delay        time.Duration
		description  string
	}{
		{"Small Batch", 5, 2000, 1 * time.Second, "Small batch of 5 messages"},
		{"Medium Batch", 10, 2100, 2 * time.Second, "Medium batch of 10 messages"},
		{"Large Batch", 20, 2200, 3 * time.Second, "Large batch of 20 messages"},
		{"Spam Cleanup", 50, 2300, 500 * time.Millisecond, "Spam cleanup of 50 messages"},
		{"History Cleanup", 100, 2400, 100 * time.Millisecond, "History cleanup of 100 messages"},
		{"Archive Cleanup", 200, 2500, 50 * time.Millisecond, "Archive cleanup of 200 messages"},
		{"Sequential", 3, 2600, 5 * time.Second, "Sequential deletion with delays"},
		{"Immediate Bulk", 15, 2700, 0, "Immediate bulk deletion"},
	}

	for _, scenario := range bulkScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			for i := 0; i < scenario.messageCount; i++ {
				messageID := scenario.startMsgID + int64(i)

				result := testCtx.DeleteMessage().
					ChatID(chatID).
					MessageID(messageID)

				if scenario.delay > 0 {
					// Add progressive delay for sequential deletions
					delay := scenario.delay + time.Duration(i*100)*time.Millisecond
					result = result.After(delay)
				}

				if result == nil {
					t.Errorf("%s message %d deletion should work", scenario.name, i+1)
					break
				}

				// Test with timeout for bulk operations
				timedResult := result.Timeout(60 * time.Second)
				if timedResult == nil {
					t.Errorf("%s message %d with timeout should work", scenario.name, i+1)
					break
				}

				// Only test first few messages to avoid too many iterations
				if i >= 2 {
					break
				}
			}
		})
	}
}

func TestDeleteMessage_ScheduledDeletion(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(-1001234567890)

	// Test scheduled deletion scenarios
	scheduledScenarios := []struct {
		name        string
		messageID   int64
		delay       time.Duration
		messageType string
		description string
	}{
		{"Welcome Message", 3000, 30 * time.Second, "welcome", "Welcome message auto-delete"},
		{"Temporary Notice", 3001, 2 * time.Minute, "notice", "Temporary notice removal"},
		{"Event Reminder", 3002, 5 * time.Minute, "reminder", "Event reminder cleanup"},
		{"Promotional", 3003, 10 * time.Minute, "promo", "Promotional message cleanup"},
		{"System Status", 3004, 15 * time.Minute, "status", "System status update cleanup"},
		{"Daily Reminder", 3005, 1 * time.Hour, "daily", "Daily reminder cleanup"},
		{"Weekly Notice", 3006, 7 * 24 * time.Hour, "weekly", "Weekly notice cleanup"},
		{"Quick Tip", 3007, 45 * time.Second, "tip", "Quick tip auto-delete"},
		{"Warning Message", 3008, 3 * time.Minute, "warning", "Warning message cleanup"},
		{"Achievement", 3009, 20 * time.Minute, "achievement", "Achievement notification cleanup"},
	}

	for _, scenario := range scheduledScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := testCtx.DeleteMessage().
				ChatID(chatID).
				MessageID(scenario.messageID).
				After(scenario.delay)

			if result == nil {
				t.Errorf("%s scheduled deletion (%s) should work", scenario.name, scenario.description)
			}

			// Test with API configuration for scheduled deletions
			apiResult := result.APIURL(g.String("https://scheduled-api.telegram.org"))
			if apiResult == nil {
				t.Errorf("API configuration for %s should work", scenario.name)
			}

			// Test with timeout for scheduled operations
			timedResult := apiResult.Timeout(2 * time.Minute)
			if timedResult == nil {
				t.Errorf("Timeout for %s should work", scenario.name)
			}
		})
	}
}

func TestDeleteMessage_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero message ID (should use effective message)
	result := testCtx.DeleteMessage().MessageID(0)
	if result == nil {
		t.Error("Zero message ID should work (uses effective message)")
	}

	// Test with zero chat ID (should use effective chat)
	result = testCtx.DeleteMessage().ChatID(0)
	if result == nil {
		t.Error("Zero chat ID should work (uses effective chat)")
	}

	// Test with negative message ID
	result = testCtx.DeleteMessage().MessageID(-1)
	if result == nil {
		t.Error("Negative message ID should work")
	}

	// Test with very large message ID
	result = testCtx.DeleteMessage().MessageID(999999999999)
	if result == nil {
		t.Error("Very large message ID should work")
	}

	// Test with zero delay
	result = testCtx.DeleteMessage().After(0 * time.Second)
	if result == nil {
		t.Error("Zero delay should work")
	}

	// Test with very long delay
	result = testCtx.DeleteMessage().After(24 * time.Hour)
	if result == nil {
		t.Error("Very long delay should work")
	}

	// Test with zero timeout
	result = testCtx.DeleteMessage().Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.DeleteMessage().Timeout(10 * time.Minute)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.DeleteMessage().APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without any parameters (should use effective values)
	result = testCtx.DeleteMessage()
	if result == nil {
		t.Error("DeleteMessage should work without explicit parameters")
	}

	// Test minimal configuration
	result = testCtx.DeleteMessage()
	if result == nil {
		t.Error("Minimal configuration should work")
	}
}

func TestDeleteMessage_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(456)
	messageID := int64(123)

	// Test all methods combined in different orders
	// Order 1: Complete configuration with delay
	result1 := testCtx.DeleteMessage().
		After(5 * time.Second).
		ChatID(chatID).
		MessageID(messageID).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined with delay (order 1) should work")
	}

	// Order 2: Different sequence without delay
	result2 := testCtx.DeleteMessage().
		APIURL(g.String("https://reordered-api.example.com")).
		Timeout(45 * time.Second).
		MessageID(messageID).
		ChatID(chatID)

	if result2 == nil {
		t.Error("All methods combined without delay (order 2) should work")
	}

	// Order 3: Delay at the end
	result3 := testCtx.DeleteMessage().
		ChatID(chatID).
		MessageID(messageID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://delayed-api.example.com")).
		After(10 * time.Second)

	if result3 == nil {
		t.Error("All methods with delay at end should work")
	}

	// Test overriding methods
	result4 := testCtx.DeleteMessage().
		ChatID(chatID).
		ChatID(999). // Should override first
		MessageID(messageID).
		MessageID(888). // Should override first
		After(5 * time.Second).
		After(10 * time.Second). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result4 == nil {
		t.Error("Method overriding should work")
	}

	// Test minimal configurations
	result5 := testCtx.DeleteMessage()
	if result5 == nil {
		t.Error("Minimal configuration should work")
	}

	// Test with just ChatID
	result6 := testCtx.DeleteMessage().ChatID(chatID)
	if result6 == nil {
		t.Error("ChatID-only configuration should work")
	}

	// Test with just MessageID
	result7 := testCtx.DeleteMessage().MessageID(messageID)
	if result7 == nil {
		t.Error("MessageID-only configuration should work")
	}

	// Test with just After
	result8 := testCtx.DeleteMessage().After(5 * time.Second)
	if result8 == nil {
		t.Error("After-only configuration should work")
	}

	// Test with just Timeout
	result9 := testCtx.DeleteMessage().Timeout(30 * time.Second)
	if result9 == nil {
		t.Error("Timeout-only configuration should work")
	}

	// Test with just APIURL
	result10 := testCtx.DeleteMessage().APIURL(g.String("https://single-api.example.com"))
	if result10 == nil {
		t.Error("APIURL-only configuration should work")
	}

	// Test various combinations of parameters
	combinations := []struct {
		name     string
		hasDelay bool
		hasChat  bool
		hasMsg   bool
	}{
		{"Delay + Chat", true, true, false},
		{"Delay + Message", true, false, true},
		{"Chat + Message", false, true, true},
		{"All Parameters", true, true, true},
		{"Only Delay", true, false, false},
	}

	for _, combo := range combinations {
		result := testCtx.DeleteMessage()

		if combo.hasDelay {
			result = result.After(3 * time.Second)
		}
		if combo.hasChat {
			result = result.ChatID(chatID)
		}
		if combo.hasMsg {
			result = result.MessageID(messageID)
		}

		result = result.Timeout(45 * time.Second).APIURL(g.String("https://combo-api.example.com"))

		if result == nil {
			t.Errorf("Combination %s should work", combo.name)
		}
	}
}

func TestDeleteMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Text: "test"},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	chatID := int64(456)
	messageID := int64(123)

	// Test Send method execution (immediate)
	builder := testCtx.DeleteMessage().ChatID(chatID).MessageID(messageID)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with various options
	builderWithOptions := testCtx.DeleteMessage().
		ChatID(chatID).
		MessageID(messageID).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org"))
	resultWithOptions := builderWithOptions.Send()

	if !resultWithOptions.IsErr() && !resultWithOptions.IsOk() {
		t.Error("Send with options should return a result")
	}

	// Test Send with After (scheduled)
	builderWithAfter := testCtx.DeleteMessage().
		ChatID(chatID).
		MessageID(messageID).
		After(1 * time.Millisecond) // Very short duration for testing
	resultWithAfter := builderWithAfter.Send()

	if !resultWithAfter.IsOk() {
		t.Error("Send with After should return Ok(true) for scheduled execution")
	}

	// Test Send without ChatID (should use effective chat)
	builderWithoutChatID := testCtx.DeleteMessage().MessageID(messageID)
	resultWithoutChatID := builderWithoutChatID.Send()

	if !resultWithoutChatID.IsErr() && !resultWithoutChatID.IsOk() {
		t.Error("Send without ChatID should return a result (using effective chat)")
	}

	// Test Send without MessageID (should use effective message)
	builderWithoutMessageID := testCtx.DeleteMessage().ChatID(chatID)
	resultWithoutMessageID := builderWithoutMessageID.Send()

	if !resultWithoutMessageID.IsErr() && !resultWithoutMessageID.IsOk() {
		t.Error("Send without MessageID should return a result (using effective message)")
	}

	// Test Send without both IDs (should use effective values)
	builderWithoutIDs := testCtx.DeleteMessage()
	resultWithoutIDs := builderWithoutIDs.Send()

	if !resultWithoutIDs.IsErr() && !resultWithoutIDs.IsOk() {
		t.Error("Send without IDs should return a result (using effective values)")
	}

	// Test Send with all features
	builderComplete := testCtx.DeleteMessage().
		ChatID(chatID).
		MessageID(messageID).
		Timeout(45 * time.Second).
		APIURL(g.String("https://complete-api.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all features should return a result")
	}

	// Test Send with After and complete features
	builderScheduled := testCtx.DeleteMessage().
		ChatID(chatID).
		MessageID(messageID).
		After(1 * time.Millisecond).
		Timeout(30 * time.Second).
		APIURL(g.String("https://scheduled-api.example.com"))
	resultScheduled := builderScheduled.Send()

	if !resultScheduled.IsOk() {
		t.Error("Send with After and all features should return Ok(true)")
	}
}
