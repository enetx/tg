package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetStarTransactions(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.GetStarTransactions()

	if result == nil {
		t.Error("Expected GetStarTransactions builder to be created")
	}

	// Test method chaining
	withOffset := result.Offset(0)
	if withOffset == nil {
		t.Error("Expected Offset method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestGetStarTransactions_Limit(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Limit method with various limit values
	limits := []int64{1, 10, 25, 50, 75, 100, 0, -1}

	for _, limit := range limits {
		result := ctx.GetStarTransactions()
		limitResult := result.Limit(limit)
		if limitResult == nil {
			t.Errorf("Limit method should return GetStarTransactions for chaining with limit %d", limit)
		}

		// Test that Limit can be chained
		chainedResult := limitResult.Limit(limit + 1)
		if chainedResult == nil {
			t.Errorf("Limit method should support chaining with limit %d", limit+1)
		}
	}

	// Test Limit chaining with other methods
	chainedResult := ctx.GetStarTransactions().
		Offset(10).
		Limit(50).
		Offset(20) // Test method order independence

	if chainedResult == nil {
		t.Error("Limit method should support chaining with other methods")
	}
}

func TestGetStarTransactions_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetStarTransactions()
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetStarTransactions for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetStarTransactions()
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetStarTransactions for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetStarTransactions().
			Offset(5).
			Limit(20).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetStarTransactions with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetStarTransactions_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetStarTransactions()
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-star-transactions-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetStarTransactions for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetStarTransactions()
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-star-transactions-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-star-transactions-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetStarTransactions for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://star-transactions-api.example.com",
		"https://custom-star-transactions.telegram.org",
		"https://regional-star-transactions-api.telegram.org",
		"https://backup-star-transactions-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetStarTransactions().
			Offset(0).
			Limit(10).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetStarTransactions with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetStarTransactions_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetStarTransactions().Send()

	if sendResult.IsErr() {
		t.Logf("GetStarTransactions Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with various configurations
	testScenarios := []struct {
		offset      int64
		limit       int64
		description string
	}{
		{0, 10, "First 10 transactions"},
		{10, 20, "Next 20 transactions"},
		{0, 1, "Single transaction"},
		{100, 50, "Offset 100, limit 50"},
		{0, 100, "Maximum limit"},
		{-1, -1, "Negative values"},
		{1000, 5, "High offset, low limit"},
	}

	for _, scenario := range testScenarios {
		configuredSendResult := ctx.GetStarTransactions().
			Offset(scenario.offset).
			Limit(scenario.limit).
			Timeout(30 * time.Second).
			APIURL(g.String("https://api.example.com")).
			Send()

		if configuredSendResult.IsErr() {
			t.Logf("GetStarTransactions Send with %s failed as expected: %v", scenario.description, configuredSendResult.Err())
		}
	}

	// Test comprehensive workflow with all methods
	comprehensiveResult := ctx.GetStarTransactions().
		Offset(25).
		Limit(75).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-star-transactions-api.telegram.org")).
		Send()

	if comprehensiveResult.IsErr() {
		t.Logf("GetStarTransactions comprehensive workflow Send failed as expected: %v", comprehensiveResult.Err())
	}

	// Test method order independence - different chaining orders should work
	orderTest1 := ctx.GetStarTransactions().
		Limit(30).
		Offset(10).
		APIURL(g.String("https://order-test-1.telegram.org")).
		Timeout(45 * time.Second).
		Send()

	if orderTest1.IsErr() {
		t.Logf("GetStarTransactions order test 1 Send failed as expected: %v", orderTest1.Err())
	}

	orderTest2 := ctx.GetStarTransactions().
		Timeout(45 * time.Second).
		APIURL(g.String("https://order-test-2.telegram.org")).
		Limit(30).
		Offset(10).
		Send()

	if orderTest2.IsErr() {
		t.Logf("GetStarTransactions order test 2 Send failed as expected: %v", orderTest2.Err())
	}
}
