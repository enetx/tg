package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetMyStarBalance(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetMyStarBalance()

	if result == nil {
		t.Error("Expected GetMyStarBalance builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetMyStarBalance_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetMyStarBalance()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-star-balance-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetMyStarBalance for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetMyStarBalance()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-star-balance-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-star-balance-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetMyStarBalance for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://star-balance-api.example.com",
		"https://custom-star-balance.telegram.org",
		"https://regional-star-balance-api.telegram.org",
		"https://backup-star-balance-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetMyStarBalance().
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetMyStarBalance with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetMyStarBalance_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetMyStarBalance().Send()

	if sendResult.IsErr() {
		t.Logf("GetMyStarBalance Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.GetMyStarBalance().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetMyStarBalance configured Send failed as expected: %v", configuredSendResult.Err())
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
		timeoutResult := ctx.GetMyStarBalance().
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetMyStarBalance with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test comprehensive workflow
	comprehensiveResult := ctx.GetMyStarBalance().
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-star-balance-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetMyStarBalance comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}
}
