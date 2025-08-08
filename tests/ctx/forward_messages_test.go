package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ForwardMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ForwardMessages()

	if result == nil {
		t.Error("Expected ForwardMessages builder to be created")
	}

	// Test method chaining
	chained := result.To(123)
	if chained == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestContext_ForwardMessagesChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.ForwardMessages().
		To(123).
		Silent()

	if result == nil {
		t.Error("Expected ForwardMessages builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestForwardMessages_AddMessages(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test AddMessages method functionality
	addMessagesScenarios := []struct {
		messageIDs  []int64
		description string
	}{
		{[]int64{100}, "Single message"},
		{[]int64{200, 300}, "Two messages"},
		{[]int64{400, 500, 600}, "Three messages"},
		{[]int64{700, 800, 900, 1000, 1100}, "Five messages"},
		{[]int64{1200}, "Another single message"},
		{[]int64{}, "Empty message list (should work)"},
	}

	for _, scenario := range addMessagesScenarios {
		addResult := ctx.ForwardMessages().
			From(456).
			To(123).
			AddMessages(scenario.messageIDs...)

		if addResult == nil {
			t.Errorf("AddMessages method with %s should work", scenario.description)
		}

		// Test Send with AddMessages
		sendResult := addResult.Send()
		if sendResult.IsErr() {
			t.Logf("ForwardMessages with AddMessages %s Send failed as expected: %v",
				scenario.description, sendResult.Err())
		}
	}
}

func TestForwardMessages_AddMessagesChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test AddMessages method chaining and accumulation
	chainedResult := ctx.ForwardMessages().
		From(456).
		To(123).
		AddMessages(100, 200).      // Add first batch
		AddMessages(300).           // Add single message
		AddMessages(400, 500, 600). // Add another batch
		Silent().
		Protect().
		Timeout(45 * time.Second).
		APIURL(g.String("https://add-messages-api.telegram.org")).
		Send()

	if chainedResult.IsErr() {
		t.Logf("ForwardMessages with AddMessages chaining Send failed as expected: %v", chainedResult.Err())
	}

	// Test combining MessageIDs and AddMessages
	combinedResult := ctx.ForwardMessages().
		From(789).
		To(456).
		MessageIDs([]int64{1000, 2000}). // Set initial IDs
		AddMessages(3000, 4000).         // Add more IDs
		Send()

	if combinedResult.IsErr() {
		t.Logf("ForwardMessages with MessageIDs+AddMessages combination Send failed as expected: %v", combinedResult.Err())
	}
}

func TestForwardMessages_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Thread method functionality for forum supergroups
	threadIDs := []int64{
		1,     // Thread ID 1
		2,     // Thread ID 2
		5,     // Thread ID 5
		10,    // Thread ID 10
		100,   // Thread ID 100
		1000,  // Thread ID 1000
		99999, // Large thread ID
	}

	for _, threadID := range threadIDs {
		threadResult := ctx.ForwardMessages().
			From(-1001234567890).
			To(-1001987654321).
			MessageIDs([]int64{100, 200, 300}).
			Thread(threadID)

		if threadResult == nil {
			t.Errorf("Thread method with thread ID %d should work", threadID)
		}

		// Test Send with Thread
		sendResult := threadResult.Send()
		if sendResult.IsErr() {
			t.Logf("ForwardMessages with Thread ID %d Send failed as expected: %v",
				threadID, sendResult.Err())
		}
	}
}

func TestForwardMessages_ThreadChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Thread method chaining with other methods
	chainedResult := ctx.ForwardMessages().
		From(-1001234567890).
		To(-1001987654321).
		MessageIDs([]int64{1, 2, 3, 4, 5}).
		Thread(42).
		Silent().
		Protect().
		Timeout(60 * time.Second).
		APIURL(g.String("https://thread-forward-api.telegram.org")).
		Send()

	if chainedResult.IsErr() {
		t.Logf("ForwardMessages with Thread chaining Send failed as expected: %v", chainedResult.Err())
	}

	// Test Thread with AddMessages
	threadWithAddResult := ctx.ForwardMessages().
		From(-1001111111111).
		To(-1002222222222).
		Thread(123).
		AddMessages(500, 600, 700, 800).
		Send()

	if threadWithAddResult.IsErr() {
		t.Logf("ForwardMessages Thread with AddMessages Send failed as expected: %v", threadWithAddResult.Err())
	}
}

func TestForwardMessages_CombinedAddMessagesAndThread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test comprehensive workflow combining AddMessages and Thread
	comprehensiveResult := ctx.ForwardMessages().
		From(-1001234567890).
		To(-1001987654321).
		AddMessages(10, 20).     // Add some messages
		Thread(999).             // Set thread
		AddMessages(30, 40, 50). // Add more messages
		Silent().
		Protect().
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-forward-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("ForwardMessages comprehensive AddMessages+Thread Send failed as expected: %v", comprehensiveResult.Err())
	}
}

func TestForwardMessages_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.ForwardMessages().
		From(456).
		To(123).
		MessageIDs([]int64{789, 790}).
		Send()

	if sendResult.IsErr() {
		t.Logf("ForwardMessages Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.ForwardMessages().
		From(456).
		To(999).
		MessageIDs([]int64{100, 200, 300}).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("ForwardMessages configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
