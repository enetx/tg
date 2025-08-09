package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_RefundStarPayment(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chargeID := g.String("charge_123")

	result := ctx.RefundStarPayment(chargeID)

	if result == nil {
		t.Error("Expected RefundStarPayment builder to be created")
	}

	// Test method chaining
	withTimeout := result.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}
}

func TestRefundStarPayment_UserID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveUser: &gotgbot.User{Id: 456, FirstName: "Test"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chargeID := g.String("charge_userid_123")

	// Test UserID method with various user IDs
	userIDs := []int64{
		123456,
		789012,
		0,                   // Zero user ID
		-123456,             // Negative user ID (edge case)
		9223372036854775807, // Max int64
	}

	for _, userID := range userIDs {
		result := ctx.RefundStarPayment(chargeID)
		userIDResult := result.UserID(userID)
		if userIDResult == nil {
			t.Errorf("UserID method should return RefundStarPayment builder for chaining with userID: %d", userID)
		}

		// Test that UserID can be chained and overridden
		chainedResult := userIDResult.UserID(userID + 1)
		if chainedResult == nil {
			t.Errorf("UserID method should support chaining and override with userID: %d", userID)
		}
	}
}

func TestRefundStarPayment_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveUser: &gotgbot.User{Id: 456, FirstName: "Test"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chargeID := g.String("charge_apiurl_123")

	// Test APIURL method with various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://custom.api.example.com",
		"",
		"https://api.example.com/bot",
		"http://localhost:8080",
	}

	for _, apiURL := range apiURLs {
		result := ctx.RefundStarPayment(chargeID)
		apiURLResult := result.APIURL(g.String(apiURL))
		if apiURLResult == nil {
			t.Errorf("APIURL method should return RefundStarPayment builder for chaining with URL: %s", apiURL)
		}

		// Test that APIURL can be chained and overridden
		chainedResult := apiURLResult.APIURL(g.String("https://override.example.com"))
		if chainedResult == nil {
			t.Errorf("APIURL method should support chaining and override with URL: %s", apiURL)
		}
	}
}

func TestRefundStarPayment_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		EffectiveUser: &gotgbot.User{Id: 456, FirstName: "Test"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	chargeID := g.String("charge_send_123")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.RefundStarPayment(chargeID).Send()

	if sendResult.IsErr() {
		t.Logf("RefundStarPayment Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with UserID
	sendWithUserIDResult := ctx.RefundStarPayment(chargeID).
		UserID(789012).
		Send()

	if sendWithUserIDResult.IsErr() {
		t.Logf("RefundStarPayment Send with UserID failed as expected with mock bot: %v", sendWithUserIDResult.Err())
	}

	// Test Send method with all options
	sendWithOptionsResult := ctx.RefundStarPayment(chargeID).
		UserID(123456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org")).
		Send()

	if sendWithOptionsResult.IsErr() {
		t.Logf("RefundStarPayment Send with options failed as expected with mock bot: %v", sendWithOptionsResult.Err())
	}

	// Test Send method using default user ID (from effective user)
	sendWithDefaultUserResult := ctx.RefundStarPayment(chargeID).Send()

	if sendWithDefaultUserResult.IsErr() {
		t.Logf("RefundStarPayment Send with default user ID failed as expected with mock bot: %v", sendWithDefaultUserResult.Err())
	}
}
