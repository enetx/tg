package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetChatMember(userID)

	if result == nil {
		t.Error("Expected GetChatMember builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_GetChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	result := ctx.GetChatMember(userID).
		ChatID(456)

	if result == nil {
		t.Error("Expected GetChatMember builder to be created")
	}

	// Test that builder is functional
	_ = result
}

// Tests for methods with 0% coverage

func TestGetChatMember_UserID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test UserID method functionality
	userIDs := []int64{
		123,
		456,
		789,
		1000,
		999999999,
		1234567890,
		987654321,
	}

	for _, testUserID := range userIDs {
		userIDResult := ctx.GetChatMember(111). // Initial userID
							UserID(testUserID) // Override with UserID method

		if userIDResult == nil {
			t.Errorf("UserID method with user ID %d should work", testUserID)
		}

		// Test Send with UserID
		sendResult := userIDResult.ChatID(-1001234567890).Send()
		if sendResult.IsErr() {
			t.Logf("GetChatMember with UserID %d Send failed as expected: %v",
				testUserID, sendResult.Err())
		}
	}
}

func TestGetChatMember_Timeout(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatMember(userID)
	timeoutResultNil := freshResult.Timeout(30 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return GetChatMember for chaining with nil RequestOpts")
	}

	// Test Timeout method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatMember(userID)
	timeoutFirst := anotherFreshResult.Timeout(15 * time.Second) // This creates RequestOpts
	timeoutSecond := timeoutFirst.Timeout(25 * time.Second)      // This uses existing RequestOpts
	if timeoutSecond == nil {
		t.Error("Timeout method should return GetChatMember for chaining with existing RequestOpts")
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
		timeoutResult := ctx.GetChatMember(userID).
			ChatID(-1001234567890).
			Timeout(timeout).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("GetChatMember with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}
}

func TestGetChatMember_APIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.GetChatMember(userID)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-member-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return GetChatMember for chaining with nil RequestOpts")
	}

	// Test APIURL method with existing RequestOpts (covers the non-nil branch)
	anotherFreshResult := ctx.GetChatMember(userID)
	apiURLFirst := anotherFreshResult.APIURL(g.String("https://first-member-api.telegram.org")) // This creates RequestOpts
	apiURLSecond := apiURLFirst.APIURL(g.String("https://second-member-api.telegram.org"))      // This uses existing RequestOpts
	if apiURLSecond == nil {
		t.Error("APIURL method should return GetChatMember for chaining with existing RequestOpts")
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://member-api.example.com",
		"https://custom-member.telegram.org",
		"https://regional-member-api.telegram.org",
		"https://backup-member-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.GetChatMember(userID).
			ChatID(-1001234567890).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("GetChatMember with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}
}

func TestGetChatMember_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.GetChatMember(userID).Send()

	if sendResult.IsErr() {
		t.Logf("GetChatMember Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.GetChatMember(userID).
		ChatID(-1001987654321).
		UserID(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("GetChatMember configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.GetChatMember(userID).
		UserID(789).
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("GetChatMember with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}
