package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CopyMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test basic creation
	result := testCtx.CopyMessages()
	if result == nil {
		t.Error("Expected CopyMessages builder to be created")
	}

	// Test To method
	result = result.To(123)
	if result == nil {
		t.Error("To method should return CopyMessages for chaining")
	}

	// Test From method
	result = testCtx.CopyMessages().From(456)
	if result == nil {
		t.Error("From method should return CopyMessages for chaining")
	}

	// Test MessageIDs method
	messageIDs := []int64{1, 2, 3, 4, 5}
	result = testCtx.CopyMessages().MessageIDs(messageIDs)
	if result == nil {
		t.Error("MessageIDs method should return CopyMessages for chaining")
	}

	// Test AddMessages method
	result = testCtx.CopyMessages().AddMessages(10, 20, 30)
	if result == nil {
		t.Error("AddMessages method should return CopyMessages for chaining")
	}

	// Test Thread method
	result = testCtx.CopyMessages().Thread(123)
	if result == nil {
		t.Error("Thread method should return CopyMessages for chaining")
	}

	// Test Silent method
	result = testCtx.CopyMessages().Silent()
	if result == nil {
		t.Error("Silent method should return CopyMessages for chaining")
	}

	// Test Protect method
	result = testCtx.CopyMessages().Protect()
	if result == nil {
		t.Error("Protect method should return CopyMessages for chaining")
	}

	// Test RemoveCaption method
	result = testCtx.CopyMessages().RemoveCaption()
	if result == nil {
		t.Error("RemoveCaption method should return CopyMessages for chaining")
	}

	// Test Timeout method
	result = testCtx.CopyMessages().Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return CopyMessages for chaining")
	}

	// Test APIURL method
	result = testCtx.CopyMessages().APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return CopyMessages for chaining")
	}
}

func TestContext_CopyMessagesChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test complex method chaining
	result := testCtx.CopyMessages().
		To(123).
		From(456).
		MessageIDs([]int64{1, 2, 3}).
		Thread(789).
		Silent().
		Protect().
		RemoveCaption().
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return CopyMessages")
	}

	// Test with AddMessages chaining
	final := testCtx.CopyMessages().
		To(999).
		From(888).
		AddMessages(10, 20, 30).
		AddMessages(40, 50).
		Silent().
		Protect()

	if final == nil {
		t.Error("AddMessages chaining should work")
	}
}

func TestCopyMessages_MessageIDManagement(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test single message ID
	result := testCtx.CopyMessages().MessageIDs([]int64{123})
	if result == nil {
		t.Error("Single message ID should work")
	}

	// Test multiple message IDs
	multipleIDs := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result = testCtx.CopyMessages().MessageIDs(multipleIDs)
	if result == nil {
		t.Error("Multiple message IDs should work")
	}

	// Test maximum message IDs (100)
	maxIDs := make([]int64, 100)
	for i := 0; i < 100; i++ {
		maxIDs[i] = int64(i + 1)
	}
	result = testCtx.CopyMessages().MessageIDs(maxIDs)
	if result == nil {
		t.Error("Maximum 100 message IDs should work")
	}

	// Test AddMessages with various counts
	addCounts := [][]int64{
		{1},
		{1, 2},
		{1, 2, 3, 4, 5},
		{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
	}

	for _, ids := range addCounts {
		result = testCtx.CopyMessages().AddMessages(ids...)
		if result == nil {
			t.Errorf("AddMessages with %d IDs should work", len(ids))
		}
	}

	// Test combining MessageIDs and AddMessages
	combinedResult := testCtx.CopyMessages().
		MessageIDs([]int64{1, 2, 3}).
		AddMessages(4, 5, 6).
		AddMessages(7, 8, 9, 10)

	if combinedResult == nil {
		t.Error("Combining MessageIDs and AddMessages should work")
	}
}

func TestCopyMessages_ChatIdentifiers(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various chat ID combinations
	testCases := []struct {
		name       string
		fromChatID int64
		toChatID   int64
		description string
	}{
		{"Private to Private", 123, 456, "Copy from private to private chat"},
		{"Group to Private", -100123456789, 456, "Copy from group to private"},
		{"Private to Group", 456, -100987654321, "Copy from private to group"},
		{"Group to Group", -100111111111, -100222222222, "Copy between groups"},
		{"Channel to Private", -1001234567890, 999, "Copy from channel to private"},
		{"Supergroup to Channel", -1001111111111, -1002222222222, "Copy from supergroup to channel"},
		{"Same Chat", 123, 123, "Copy within same chat"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testCtx.CopyMessages().
				From(tc.fromChatID).
				To(tc.toChatID).
				MessageIDs([]int64{1, 2, 3})

			if result == nil {
				t.Errorf("%s should work (%s)", tc.name, tc.description)
			}

			// Test with additional features
			enhancedResult := result.Silent().Protect()
			if enhancedResult == nil {
				t.Errorf("Enhanced %s should work", tc.name)
			}
		})
	}
}

func TestCopyMessages_ForumFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(-1001234567890)
	toChatID := int64(-1001987654321)
	messageIDs := []int64{123, 456, 789}

	// Test various thread IDs
	threadIDs := []int64{
		1,   // General topic
		123, // Regular topic
		456, // Another topic
		999, // High ID topic
		123456789, // Very high ID topic
	}

	for _, threadID := range threadIDs {
		result := testCtx.CopyMessages().
			From(fromChatID).
			To(toChatID).
			MessageIDs(messageIDs).
			Thread(threadID)

		if result == nil {
			t.Errorf("Thread ID %d should work", threadID)
		}
	}

	// Test thread with other forum features
	forumResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		Thread(123).
		Silent().
		Protect()

	if forumResult == nil {
		t.Error("Forum thread with other features should work")
	}
}

func TestCopyMessages_ContentOptions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	toChatID := int64(789)
	messageIDs := []int64{1, 2, 3, 4, 5}

	// Test Silent option
	silentResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		Silent()

	if silentResult == nil {
		t.Error("Silent option should work")
	}

	// Test Protect option
	protectResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		Protect()

	if protectResult == nil {
		t.Error("Protect option should work")
	}

	// Test RemoveCaption option
	removeCaptionResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		RemoveCaption()

	if removeCaptionResult == nil {
		t.Error("RemoveCaption option should work")
	}

	// Test all content options combined
	combinedResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		Silent().
		Protect().
		RemoveCaption()

	if combinedResult == nil {
		t.Error("All content options combined should work")
	}
}

func TestCopyMessages_TimeoutAndAPI(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	toChatID := int64(789)
	messageIDs := []int64{1, 2, 3}

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		30 * time.Second,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, timeout := range timeouts {
		result := testCtx.CopyMessages().
			From(fromChatID).
			To(toChatID).
			MessageIDs(messageIDs).
			Timeout(timeout)

		if result == nil {
			t.Errorf("Timeout %v should work", timeout)
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom-api.example.com",
		"https://bulk-copy-api.example.com",
		"https://localhost:8080",
		"https://api-staging.telegram.org",
		"https://enterprise-copy-api.com",
	}

	for _, apiURL := range apiURLs {
		result := testCtx.CopyMessages().
			From(fromChatID).
			To(toChatID).
			MessageIDs(messageIDs).
			APIURL(g.String(apiURL))

		if result == nil {
			t.Errorf("API URL %s should work", apiURL)
		}
	}

	// Test timeout and API URL combined
	combinedResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		MessageIDs(messageIDs).
		Timeout(60 * time.Second).
		APIURL(g.String("https://combined-api.example.com"))

	if combinedResult == nil {
		t.Error("Timeout and API URL combined should work")
	}
}

func TestCopyMessages_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty message IDs slice
	result := testCtx.CopyMessages().MessageIDs([]int64{})
	if result == nil {
		t.Error("Empty message IDs slice should work (builder creation)")
	}

	// Test with zero chat IDs
	result = testCtx.CopyMessages().From(0).To(0)
	if result == nil {
		t.Error("Zero chat IDs should work (builder creation)")
	}

	// Test with zero thread ID
	result = testCtx.CopyMessages().Thread(0)
	if result == nil {
		t.Error("Zero thread ID should work")
	}

	// Test with zero timeout
	result = testCtx.CopyMessages().Timeout(0 * time.Second)
	if result == nil {
		t.Error("Zero timeout should work")
	}

	// Test with very long timeout
	result = testCtx.CopyMessages().Timeout(24 * time.Hour)
	if result == nil {
		t.Error("Very long timeout should work")
	}

	// Test with empty API URL
	result = testCtx.CopyMessages().APIURL(g.String(""))
	if result == nil {
		t.Error("Empty API URL should work")
	}

	// Test without To() method (should use effective chat)
	result = testCtx.CopyMessages().
		From(456).
		MessageIDs([]int64{1, 2, 3})

	if result == nil {
		t.Error("CopyMessages should work without explicit To() call")
	}
}

func TestCopyMessages_BulkOperations(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	toChatID := int64(789)

	// Test bulk copy scenarios
	bulkScenarios := []struct {
		name        string
		messageCount int
		description string
	}{
		{"Small Batch", 5, "Small batch of messages"},
		{"Medium Batch", 25, "Medium batch of messages"},
		{"Large Batch", 50, "Large batch of messages"},
		{"Maximum Batch", 100, "Maximum allowed batch size"},
	}

	for _, scenario := range bulkScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Generate sequential message IDs
			messageIDs := make([]int64, scenario.messageCount)
			for i := 0; i < scenario.messageCount; i++ {
				messageIDs[i] = int64(i + 1)
			}

			result := testCtx.CopyMessages().
				From(fromChatID).
				To(toChatID).
				MessageIDs(messageIDs).
				Silent().
				Protect()

			if result == nil {
				t.Errorf("%s (%s) should work", scenario.name, scenario.description)
			}
		})
	}

	// Test incremental adding
	incrementalResult := testCtx.CopyMessages().
		From(fromChatID).
		To(toChatID).
		AddMessages(1, 2, 3).
		AddMessages(4, 5, 6).
		AddMessages(7, 8, 9, 10)

	if incrementalResult == nil {
		t.Error("Incremental message adding should work")
	}
}

func TestCopyMessages_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	fromChatID := int64(456)
	toChatID := int64(789)
	messageIDs := []int64{1, 2, 3, 4, 5}

	// Test all methods combined in different orders
	// Order 1
	result1 := testCtx.CopyMessages().
		To(toChatID).
		From(fromChatID).
		MessageIDs(messageIDs).
		Thread(123).
		Silent().
		Protect().
		RemoveCaption().
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result1 == nil {
		t.Error("All methods combined (order 1) should work")
	}

	// Order 2 (different sequence)
	result2 := testCtx.CopyMessages().
		APIURL(g.String("https://custom-api.example.com")).
		Timeout(45 * time.Second).
		RemoveCaption().
		Protect().
		Silent().
		Thread(456).
		AddMessages(messageIDs...).
		From(fromChatID).
		To(toChatID)

	if result2 == nil {
		t.Error("All methods combined (order 2) should work")
	}

	// Test overriding methods
	result3 := testCtx.CopyMessages().
		To(toChatID).
		To(999). // Should override first
		From(fromChatID).
		From(888). // Should override first
		MessageIDs([]int64{1, 2}).
		MessageIDs([]int64{3, 4, 5}). // Should override first
		Timeout(30 * time.Second).
		Timeout(60 * time.Second) // Should override first

	if result3 == nil {
		t.Error("Method overriding should work")
	}

	// Test combining MessageIDs and AddMessages
	result4 := testCtx.CopyMessages().
		To(toChatID).
		From(fromChatID).
		MessageIDs([]int64{1, 2, 3}).
		AddMessages(4, 5, 6).
		AddMessages(7, 8, 9, 10).
		Silent().
		Protect()

	if result4 == nil {
		t.Error("Combining MessageIDs and AddMessages should work")
	}
}
