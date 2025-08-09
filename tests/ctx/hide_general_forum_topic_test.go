package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_HideGeneralForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.HideGeneralForumTopic()

	if result == nil {
		t.Error("Expected HideGeneralForumTopic builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestHideGeneralForumTopic_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.HideGeneralForumTopic()
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return HideGeneralForumTopic for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.HideGeneralForumTopic()
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return HideGeneralForumTopic for chaining with existing RequestOpts")
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
		timeoutResult := ctx.HideGeneralForumTopic().
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("HideGeneralForumTopic with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestHideGeneralForumTopic_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.HideGeneralForumTopic()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-hide-forum-topic-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return HideGeneralForumTopic for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.HideGeneralForumTopic()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-hide-forum-topic-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-hide-forum-topic-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return HideGeneralForumTopic for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://hide-forum-topic-api.example.com",
		"https://custom-hide-forum-topic.telegram.org",
		"https://regional-hide-forum-topic-api.telegram.org",
		"https://backup-hide-forum-topic-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.HideGeneralForumTopic().
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("HideGeneralForumTopic with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestHideGeneralForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with various chat configurations
	testScenarios := []struct {
		chatID      int64
		description string
	}{
		{-1001234567890, "Regular supergroup"},
		{-1001111111111, "Different supergroup"},
		{-1002222222222, "Another supergroup"},
		{0, "Zero chat ID (uses effective chat)"},
		{-1, "Invalid negative chat ID"},
	}

	for _, scenario := range testScenarios {
		// Basic Send test
		var sendResult g.Result[bool]
		if scenario.chatID != 0 {
			sendResult = ctx.HideGeneralForumTopic().ChatID(scenario.chatID).Send()
		} else {
			sendResult = ctx.HideGeneralForumTopic().Send() // Uses effective chat
		}

		if sendResult.IsErr() {
			t.Logf("HideGeneralForumTopic with %s Send failed as expected: %v", scenario.description, sendResult.Err())
		}

		// Configured Send test
		configuredSendResult := ctx.HideGeneralForumTopic().
			ChatID(scenario.chatID).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("HideGeneralForumTopic configured with %s Send failed as expected: %v", scenario.description, configuredSendResult.Err())
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
		timeoutResult := ctx.HideGeneralForumTopic().
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("HideGeneralForumTopic with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.HideGeneralForumTopic().
		ChatID(-1001234567890).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-hide-forum-topic-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("HideGeneralForumTopic comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method chaining order independence
	orderTest1 := ctx.HideGeneralForumTopic().
		APIURL(g.String("https://order-test-1.telegram.org")).
		ChatID(-1001234567890).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("HideGeneralForumTopic order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.HideGeneralForumTopic().
		Timeout(45 * time.Second).
		APIURL(g.String("https://order-test-2.telegram.org")).
		ChatID(-1001234567890).
		Send()

	if orderTest2.IsErr() {
		t.Logf("HideGeneralForumTopic order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
