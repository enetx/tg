package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_BanChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test basic creation
	result := testCtx.BanChatMember(userID)
	if result == nil {
		t.Error("Expected BanChatMember builder to be created")
	}

	// Test ChatID method
	result = result.ChatID(-1001234567890)
	if result == nil {
		t.Error("ChatID method should return BanChatMember for chaining")
	}

	// Test RevokeMessages method
	result = testCtx.BanChatMember(userID).RevokeMessages()
	if result == nil {
		t.Error("RevokeMessages method should return BanChatMember for chaining")
	}

	// Test Until method
	futureTime := time.Now().Add(24 * time.Hour)
	result = testCtx.BanChatMember(userID).Until(futureTime)
	if result == nil {
		t.Error("Until method should return BanChatMember for chaining")
	}

	// Test For method
	result = testCtx.BanChatMember(userID).For(48 * time.Hour)
	if result == nil {
		t.Error("For method should return BanChatMember for chaining")
	}

	// Test Timeout method
	result = testCtx.BanChatMember(userID).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return BanChatMember for chaining")
	}

	// Test APIURL method
	result = testCtx.BanChatMember(userID).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return BanChatMember for chaining")
	}
}

func TestContext_BanChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test complete method chaining
	result := testCtx.BanChatMember(userID).
		ChatID(-1001987654321).
		RevokeMessages().
		For(7 * 24 * time.Hour). // 7 days
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return BanChatMember")
	}
}

func TestBanChatMember_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with zero UserID
	result := testCtx.BanChatMember(0)
	if result == nil {
		t.Error("BanChatMember should handle zero UserID")
	}

	// Test with negative UserID
	result = testCtx.BanChatMember(-123456)
	if result == nil {
		t.Error("BanChatMember should handle negative UserID")
	}

	// Test with maximum UserID
	result = testCtx.BanChatMember(9223372036854775807)
	if result == nil {
		t.Error("BanChatMember should handle maximum UserID")
	}

	// Test with zero ChatID (should use effective chat)
	result = testCtx.BanChatMember(123456).ChatID(0)
	if result == nil {
		t.Error("BanChatMember should handle zero ChatID")
	}

	// Test with past time for Until
	pastTime := time.Now().Add(-24 * time.Hour)
	result = testCtx.BanChatMember(123456).Until(pastTime)
	if result == nil {
		t.Error("BanChatMember should handle past time for Until")
	}

	// Test with zero duration for For
	result = testCtx.BanChatMember(123456).For(0 * time.Second)
	if result == nil {
		t.Error("BanChatMember should handle zero duration for For")
	}

	// Test with very long duration
	result = testCtx.BanChatMember(123456).For(365 * 24 * time.Hour)
	if result == nil {
		t.Error("BanChatMember should handle very long duration")
	}
}

func TestBanChatMember_TimeComparison(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test that For() and Until() both work with same duration
	duration := 7 * 24 * time.Hour
	expectedTime := time.Now().Add(duration)

	// For method adds duration from now
	result1 := testCtx.BanChatMember(userID).For(duration)
	if result1 == nil {
		t.Error("For method should work")
	}

	// Until method sets specific time
	result2 := testCtx.BanChatMember(userID).Until(expectedTime)
	if result2 == nil {
		t.Error("Until method should work")
	}

	// Both should create valid builders
	if result1 == nil || result2 == nil {
		t.Error("Both For and Until methods should create valid builders")
	}
}

func TestBanChatMember_DefaultChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	// Test without setting ChatID (should use EffectiveChat.Id)
	result := testCtx.BanChatMember(userID)
	if result == nil {
		t.Error("BanChatMember should work without explicit ChatID")
	}

	// The Send() method should use EffectiveChat.Id when no ChatID is set
	// We can't test the actual API call with mockBot, but we can ensure the builder works
}

func TestBanChatMember_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.BanChatMember(userID).Send()

	if sendResult.IsErr() {
		t.Logf("BanChatMember Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.BanChatMember(userID).
		ChatID(789).
		RevokeMessages().
		Until(time.Now().Add(24 * time.Hour)).
		Timeout(30).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("BanChatMember configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
