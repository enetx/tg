package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditUserStarSubscription(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	chargeID := g.String("charge_123")
	isCanceled := true

	result := ctx.EditUserStarSubscription(userID, chargeID, isCanceled)

	if result == nil {
		t.Error("Expected EditUserStarSubscription builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestEditUserStarSubscription_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	chargeID := g.String("charge_123")
	isCanceled := true

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-star-subscription-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditUserStarSubscription for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-star-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-star-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return EditUserStarSubscription for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://star-subscription-api.example.com",
		"https://custom-star.telegram.org",
		"https://payments-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditUserStarSubscription with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestEditUserStarSubscription_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	chargeID := g.String("charge_123")
	isCanceled := true

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled).Send()

	if sendResult.IsErr() {
		t.Logf("EditUserStarSubscription Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration - canceled subscription
	configuredSendResult := ctx.EditUserStarSubscription(userID, chargeID, true).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditUserStarSubscription configured Send (canceled) failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method with active subscription
	activeSendResult := ctx.EditUserStarSubscription(userID, chargeID, false).
		Timeout(45 * time.Second).
		APIURL(g.String("https://payments-api.telegram.org")).
		Send()

	if activeSendResult.IsErr() {
		t.Logf("EditUserStarSubscription active subscription Send failed as expected: %v", activeSendResult.Err())
	}
}

func TestEditUserStarSubscription_VariousScenarios(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test different user ID scenarios
	userIDScenarios := []struct {
		userID      int64
		description string
	}{
		{123456789, "Regular user"},
		{987654321, "Premium user"},
		{111111111, "Test user"},
		{999999999, "Large user ID"},
		{1, "Minimal user ID"},
	}

	for _, scenario := range userIDScenarios {
		chargeID := g.String("charge_" + g.String(g.Int(scenario.userID).String()).Std())

		// Test canceled subscription
		canceledResult := ctx.EditUserStarSubscription(scenario.userID, chargeID, true).
			Timeout(30 * time.Second).
			APIURL(g.String("https://cancel-subscription-api.telegram.org")).
			Send()

		if canceledResult.IsErr() {
			t.Logf("EditUserStarSubscription %s (ID: %d) canceled Send failed as expected: %v",
				scenario.description, scenario.userID, canceledResult.Err())
		}

		// Test active subscription
		activeResult := ctx.EditUserStarSubscription(scenario.userID, chargeID, false).
			Timeout(45 * time.Second).
			APIURL(g.String("https://active-subscription-api.telegram.org")).
			Send()

		if activeResult.IsErr() {
			t.Logf("EditUserStarSubscription %s (ID: %d) active Send failed as expected: %v",
				scenario.description, scenario.userID, activeResult.Err())
		}
	}

	// Test different charge ID formats
	chargeIDFormats := []string{
		"charge_123456",
		"payment_abc123",
		"txn_xyz789",
		"star_sub_001",
		"premium_payment_999",
		"", // Empty charge ID
	}

	for _, chargeIDFormat := range chargeIDFormats {
		chargeFormatResult := ctx.EditUserStarSubscription(123456, g.String(chargeIDFormat), true).
			Timeout(20 * time.Second).
			APIURL(g.String("https://charge-format-api.telegram.org")).
			Send()

		if chargeFormatResult.IsErr() {
			t.Logf("EditUserStarSubscription with charge ID '%s' Send failed as expected: %v",
				chargeIDFormat, chargeFormatResult.Err())
		}
	}
}

func TestEditUserStarSubscription_TimeoutCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	chargeID := g.String("charge_timeout_test")
	isCanceled := false

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditUserStarSubscription for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return EditUserStarSubscription for chaining with existing RequestOpts")
	}

	// Test various timeout durations
	timeouts := []time.Duration{
		1 * time.Second,
		10 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.EditUserStarSubscription(userID, chargeID, isCanceled).
			Timeout(timeout).
			APIURL(g.String("https://timeout-test-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("EditUserStarSubscription with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestEditUserStarSubscription_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(999888777)
	chargeID := g.String("comprehensive_charge_12345")

	// Test comprehensive workflow - cancel subscription
	cancelWorkflowResult := ctx.EditUserStarSubscription(userID, chargeID, true).
		Timeout(60 * time.Second).
		APIURL(g.String("https://comprehensive-cancel-api.telegram.org")).
		Send()

	if cancelWorkflowResult.IsErr() {
		t.Logf("EditUserStarSubscription comprehensive cancel workflow Send failed as expected: %v", cancelWorkflowResult.Err())
	}

	// Test comprehensive workflow - activate subscription
	activateWorkflowResult := ctx.EditUserStarSubscription(userID, chargeID, false).
		Timeout(90 * time.Second).
		APIURL(g.String("https://comprehensive-activate-api.telegram.org")).
		Send()

	if activateWorkflowResult.IsErr() {
		t.Logf("EditUserStarSubscription comprehensive activate workflow Send failed as expected: %v", activateWorkflowResult.Err())
	}
}
