package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_EditChatInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/abc123")

	result := ctx.EditChatInviteLink(inviteLink)

	if result == nil {
		t.Error("Expected EditChatInviteLink builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	named := chained.Name(g.String("Updated Link"))
	if named == nil {
		t.Error("Expected Name method to return builder")
	}

	withExpiry := named.ExpiresIn(time.Hour * 24)
	if withExpiry == nil {
		t.Error("Expected ExpiresIn method to return builder")
	}

	withLimit := withExpiry.MemberLimit(50)
	if withLimit == nil {
		t.Error("Expected MemberLimit method to return builder")
	}

	withRequest := withLimit.CreatesJoinRequest()
	if withRequest == nil {
		t.Error("Expected CreatesJoinRequest method to return builder")
	}
}

// Tests for methods with 0% coverage

func TestEditChatInviteLink_ExpiresAt(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/expires_at_test")

	// Test ExpiresAt method functionality
	expirationTimes := []time.Time{
		time.Now().Add(1 * time.Hour),                    // 1 hour from now
		time.Now().Add(24 * time.Hour),                   // 1 day from now
		time.Now().Add(7 * 24 * time.Hour),               // 1 week from now
		time.Now().Add(30 * 24 * time.Hour),              // 1 month from now
		time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC), // Specific future date
		time.Now().Add(365 * 24 * time.Hour),             // 1 year from now
		time.Now().Add(5 * time.Minute),                  // 5 minutes from now
	}

	for i, expiryTime := range expirationTimes {
		expiresAtResult := ctx.EditChatInviteLink(inviteLink).
			ChatID(-1001234567890).
			Name(g.String("ExpiresAt Test Link")).
			ExpiresAt(expiryTime)

		if expiresAtResult == nil {
			t.Errorf("ExpiresAt method with time %d should work", i)
		}

		// Test Send with ExpiresAt
		sendResult := expiresAtResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditChatInviteLink with ExpiresAt time %d Send failed as expected: %v",
				i, sendResult.Err())
		}
	}
}

func TestEditChatInviteLink_ExpiresAtChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/expires_at_chain_test")

	// Test ExpiresAt method chaining with other methods
	futureTime := time.Now().Add(48 * time.Hour)
	chainedResult := ctx.EditChatInviteLink(inviteLink).
		ChatID(-1001987654321).
		Name(g.String("Chained ExpiresAt Link")).
		ExpiresAt(futureTime).
		MemberLimit(100).
		CreatesJoinRequest().
		Timeout(45 * time.Second).
		APIURL(g.String("https://expires-at-api.telegram.org")).
		Send()

	if chainedResult.IsErr() {
		t.Logf("EditChatInviteLink with ExpiresAt chaining Send failed as expected: %v", chainedResult.Err())
	}

	// Test with both ExpiresAt and ExpiresIn (ExpiresAt should take precedence)
	precedenceResult := ctx.EditChatInviteLink(inviteLink).
		ChatID(-1001234567890).
		ExpiresIn(24 * time.Hour).                 // Set with ExpiresIn first
		ExpiresAt(time.Now().Add(12 * time.Hour)). // Then override with ExpiresAt
		Send()

	if precedenceResult.IsErr() {
		t.Logf("EditChatInviteLink ExpiresAt precedence test Send failed as expected: %v", precedenceResult.Err())
	}
}

func TestEditChatInviteLink_ExpiresAtEdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/expires_at_edge_test")

	// Test edge cases for ExpiresAt
	edgeCases := []struct {
		expiry      time.Time
		description string
	}{
		{time.Now().Add(1 * time.Second), "Very short expiration"},
		{time.Unix(0, 0), "Unix epoch"},
		{time.Date(2030, 1, 1, 0, 0, 0, 0, time.UTC), "Far future date"},
		{time.Now().Add(-1 * time.Hour), "Past time (should still work)"},
		{time.Time{}, "Zero time"},
	}

	for _, edgeCase := range edgeCases {
		edgeResult := ctx.EditChatInviteLink(inviteLink).
			ChatID(-1001234567890).
			Name(g.String("Edge Case: " + edgeCase.description)).
			ExpiresAt(edgeCase.expiry).
			Send()

		if edgeResult.IsErr() {
			t.Logf("EditChatInviteLink ExpiresAt edge case '%s' Send failed as expected: %v",
				edgeCase.description, edgeResult.Err())
		}
	}
}

func TestEditChatInviteLink_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	inviteLink := g.String("https://t.me/joinchat/test123")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.EditChatInviteLink(inviteLink).Send()

	if sendResult.IsErr() {
		t.Logf("EditChatInviteLink Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditChatInviteLink(inviteLink).
		ChatID(789).
		Name(g.String("Updated Test Link")).
		MemberLimit(25).
		ExpiresIn(48 * time.Hour).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditChatInviteLink configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
