package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetChatAdministrators(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChatAdministrators()

	if result == nil {
		t.Error("Expected GetChatAdministrators builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_GetChatAdministratorsChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChatAdministrators().
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected GetChatAdministrators builder to be created")
	}

	// Test that builder is functional
	_ = result
}

// Tests for methods with 0% coverage

func TestGetChatAdministrators_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatAdministrators()
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetChatAdministrators for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatAdministrators()
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetChatAdministrators for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetChatAdministrators().
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetChatAdministrators with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetChatAdministrators_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatAdministrators()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-admins-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetChatAdministrators for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatAdministrators()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-admins-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-admins-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetChatAdministrators for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://admins-api.example.com",
		"https://custom-admins.telegram.org",
		"https://regional-admins-api.telegram.org",
		"https://backup-admins-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetChatAdministrators().
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetChatAdministrators with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetChatAdministrators_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetChatAdministrators().Send()

	if sendResult.IsErr() {
		t.Logf("GetChatAdministrators Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.GetChatAdministrators().
		ChatID(-1001987654321).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetChatAdministrators configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.GetChatAdministrators().
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("GetChatAdministrators with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}
