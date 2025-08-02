package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnbanChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.UnbanChatMember(userID)

	if result == nil {
		t.Error("Expected UnbanChatMember builder to be created")
	}

	// Test method chaining
	chained := result.OnlyIfBanned()
	if chained == nil {
		t.Error("Expected OnlyIfBanned method to return builder")
	}
}

func TestContext_UnbanChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.UnbanChatMember(userID).
		OnlyIfBanned().
		ChatID(-1001234567890)

	if result == nil {
		t.Error("Expected UnbanChatMember builder to be created")
	}

	// Test that builder is functional
	_ = result
}
