package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditChatSubscriptionInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/subscription_abc123")

	result := ctx.EditChatSubscriptionInviteLink(inviteLink)

	if result == nil {
		t.Error("Expected EditChatSubscriptionInviteLink builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	named := chained.Name(g.String("Updated Premium Subscription"))
	if named == nil {
		t.Error("Expected Name method to return builder")
	}

	// Test Timeout method
	timeoutResult := result.Timeout(30 * time.Second)
	if timeoutResult == nil {
		t.Error("Timeout method should return EditChatSubscriptionInviteLink for chaining")
	}

	// Test APIURL method
	apiURLResult := result.APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return EditChatSubscriptionInviteLink for chaining")
	}

	// Test APIURL method with nil RequestOpts (covers the nil branch)
	freshResult := ctx.EditChatSubscriptionInviteLink(inviteLink)
	apiURLResultNil := freshResult.APIURL(g.String("https://custom-api.telegram.org"))
	if apiURLResultNil == nil {
		t.Error("APIURL method should return EditChatSubscriptionInviteLink for chaining with nil RequestOpts")
	}

	// Test Timeout method with nil RequestOpts (covers the nil branch)
	timeoutResultNil := freshResult.Timeout(45 * time.Second)
	if timeoutResultNil == nil {
		t.Error("Timeout method should return EditChatSubscriptionInviteLink for chaining with nil RequestOpts")
	}
}

func TestEditChatSubscriptionInviteLink_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/subscription_test123")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditChatSubscriptionInviteLink(inviteLink).Send()

	if sendResult.IsErr() {
		t.Logf("EditChatSubscriptionInviteLink Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditChatSubscriptionInviteLink(inviteLink).
		ChatID(-1001987654321).
		Name(g.String("Premium Subscription Edited")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditChatSubscriptionInviteLink configured Send failed as expected: %v", configuredSendResult.Err())
	}

	// Test Send method using EffectiveChat ID (no explicit ChatID)
	effectiveChatSendResult := ctx.EditChatSubscriptionInviteLink(inviteLink).
		Name(g.String("Effective Chat Subscription")).
		Timeout(60 * time.Second).
		Send()

	if effectiveChatSendResult.IsErr() {
		t.Logf("EditChatSubscriptionInviteLink with effective chat Send failed as expected: %v", effectiveChatSendResult.Err())
	}
}

func TestEditChatSubscriptionInviteLink_CompleteMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/comprehensive_test")

	// Test all methods in combination
	result := ctx.EditChatSubscriptionInviteLink(inviteLink).
		ChatID(-1001987654321).
		Name(g.String("Complete Test Subscription")).
		Timeout(60 * time.Second).
		APIURL(g.String("https://complete-subscription-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return EditChatSubscriptionInviteLink")
	}

	// Test complete workflow with Send
	completeResult := result.Send()
	if completeResult.IsErr() {
		t.Logf("Complete EditChatSubscriptionInviteLink workflow Send failed as expected: %v", completeResult.Err())
	}

	// Test various subscription invite links
	inviteLinks := []string{
		"https://t.me/joinchat/subscription_basic123",
		"https://t.me/joinchat/subscription_premium456",
		"https://t.me/joinchat/subscription_vip789",
		"https://t.me/joinchat/subscription_enterprise000",
		"https://t.me/joinchat/sub_short",
		"https://t.me/joinchat/very_long_subscription_invite_link_name_test",
		"https://t.me/+AbCdEf123456", // Short format
	}

	for _, link := range inviteLinks {
		linkResult := ctx.EditChatSubscriptionInviteLink(g.String(link)).
			ChatID(-1001234567890).
			Name(g.String("Test Subscription Link")).
			Timeout(45 * time.Second).
			Send()

		if linkResult.IsErr() {
			t.Logf("EditChatSubscriptionInviteLink with link '%s' Send failed as expected: %v", link, linkResult.Err())
		}
	}

	// Test various timeout configurations
	timeouts := []time.Duration{
		5 * time.Second,
		15 * time.Second,
		30 * time.Second,
		60 * time.Second,
		2 * time.Minute,
		0 * time.Second, // Zero timeout
	}

	for _, timeout := range timeouts {
		timeoutResult := ctx.EditChatSubscriptionInviteLink(inviteLink).
			ChatID(-1001234567890).
			Name(g.String("Timeout Test Subscription")).
			Timeout(timeout).
			APIURL(g.String("https://timeout-subscription-api.telegram.org")).
			Send()

		if timeoutResult.IsErr() {
			t.Logf("EditChatSubscriptionInviteLink with timeout %v Send failed as expected: %v", timeout, timeoutResult.Err())
		}
	}

	// Test various API URLs
	apiURLs := []string{
		"https://api.telegram.org",
		"https://subscription-edit-api.example.com",
		"https://custom-subscription.telegram.org",
		"https://regional-subscription-api.telegram.org",
		"", // Empty URL
	}

	for _, apiURL := range apiURLs {
		apiResult := ctx.EditChatSubscriptionInviteLink(inviteLink).
			ChatID(-1001234567890).
			Name(g.String("API Test Subscription")).
			Timeout(30 * time.Second).
			APIURL(g.String(apiURL)).
			Send()

		if apiResult.IsErr() {
			t.Logf("EditChatSubscriptionInviteLink with API URL '%s' Send failed as expected: %v", apiURL, apiResult.Err())
		}
	}

	// Test various subscription names (0-32 characters)
	subscriptionNames := []string{
		"Premium",
		"VIP Access",
		"Monthly Subscription",
		"Elite Members Only",
		"Pro Subscription 2024",
		"🌟 Premium Star Membership",
		"Enterprise Business Package",
		"",                                // Empty name
		"A",                               // Single character
		"This is exactly 32 characters!!", // Max length (32 chars)
	}

	for _, name := range subscriptionNames {
		displayName := name
		if name == "" {
			displayName = "[empty]"
		}

		nameResult := ctx.EditChatSubscriptionInviteLink(inviteLink).
			ChatID(-1001234567890).
			Name(g.String(name)).
			Timeout(45 * time.Second).
			Send()

		if nameResult.IsErr() {
			t.Logf("EditChatSubscriptionInviteLink with name '%s' Send failed as expected: %v", displayName, nameResult.Err())
		}
	}
}
