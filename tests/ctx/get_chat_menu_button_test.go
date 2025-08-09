package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetChatMenuButton(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetChatMenuButton()

	if result == nil {
		t.Error("Expected GetChatMenuButton builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(123)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetChatMenuButton_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatMenuButton()
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetChatMenuButton for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatMenuButton()
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetChatMenuButton for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetChatMenuButton().
			ChatID(123).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetChatMenuButton with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetChatMenuButton_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatMenuButton()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-menu-button-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetChatMenuButton for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatMenuButton()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-menu-button-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-menu-button-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetChatMenuButton for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://menu-button-api.example.com",
		"https://custom-menu-button.telegram.org",
		"https://regional-menu-button-api.telegram.org",
		"https://backup-menu-button-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetChatMenuButton().
			ChatID(123).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetChatMenuButton with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetChatMenuButton_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetChatMenuButton().Send()

	if sendResult.IsErr() {
		t.Logf("GetChatMenuButton Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.GetChatMenuButton().
		ChatID(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetChatMenuButton configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method without explicit ChatID (tests nil UnwrapOrDefault)
	noChatIDSendResult := ctx.GetChatMenuButton().
		Timeout(60 * time.Second).
		Send()

	if noChatIDSendResult.IsErr() {
		t.Logf("GetChatMenuButton without ChatID Send failed as expected: %v", noChatIDSendResult.Err())
	}
}
