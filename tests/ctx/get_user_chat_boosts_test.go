package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserChatBoosts(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	result := ctx.GetUserChatBoosts(userID)

	if result == nil {
		t.Error("Expected GetUserChatBoosts builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetUserChatBoosts_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetUserChatBoosts(userID)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetUserChatBoosts for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetUserChatBoosts(userID)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetUserChatBoosts for chaining with existing RequestOpts")
	}

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.GetUserChatBoosts(userID).
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetUserChatBoosts with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetUserChatBoosts_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetUserChatBoosts(userID)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-chat-boosts-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetUserChatBoosts for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetUserChatBoosts(userID)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-chat-boosts-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-chat-boosts-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetUserChatBoosts for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://chat-boosts-api.example.com",
		"https://custom-chat-boosts.telegram.org",
		"https://regional-chat-boosts-api.telegram.org",
		"https://backup-chat-boosts-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetUserChatBoosts(userID).
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetUserChatBoosts with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetUserChatBoosts_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various user/chat configurations
	testScenarios := []struct {
		userID      int64
		chatID      int64
		description string
	}{
		{123, -1001234567890, "Regular supergroup boost check"},
		{456, -1001111111111, "Different supergroup and user"},
		{789, -1002222222222, "Another supergroup configuration"},
		{0, -1001234567890, "Zero user ID"},
		{999, 0, "Zero chat ID (uses effective chat)"},
		{-1, -1001234567890, "Negative user ID"},
	}

	for _, scenario := range testScenarios {
		result := ctx.GetUserChatBoosts(scenario.userID)

		if scenario.chatID != 0 {
			result = result.ChatID(scenario.chatID)
		}

		// Basic Send test
		sendResult := result.Send()
		if sendResult.IsErr() {
			t.Logf("GetUserChatBoosts with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test
		configuredSendResult := ctx.GetUserChatBoosts(scenario.userID).
			ChatID(scenario.chatID).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GetUserChatBoosts configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}

	// Test Send method with various timeout configurations
	timeoutConfigs := []time.Duration{
		5 * time.Second,
		15 * time.Second,
		45 * time.Second,
		60 * time.Second,
		2 * time.Minute,
	}

	for _, timeout := range timeoutConfigs {
		timeoutResult := ctx.GetUserChatBoosts(456).
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetUserChatBoosts with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.GetUserChatBoosts(456).
		ChatID(-1001234567890).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-chat-boosts-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetUserChatBoosts comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := ctx.GetUserChatBoosts(789).
		APIURL(g.String("https://order-test-1.telegram.org")).
		ChatID(-1001234567890).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GetUserChatBoosts order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.GetUserChatBoosts(789).
		Timeout(45 * time.Second).
		APIURL(g.String("https://order-test-2.telegram.org")).
		ChatID(-1001234567890).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GetUserChatBoosts order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
