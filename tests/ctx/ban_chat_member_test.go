package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestContext_BanChatMember(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.BanChatMember(userID)

	if result == nil {
		t.Error("Expected BanChatMember builder to be created")
	}

	// Test method chaining
	chained := result.Until(time.Now().Add(24 * time.Hour))
	if chained == nil {
		t.Error("Expected Until method to return builder")
	}
}

func TestContext_BanChatMemberChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(123456)

	result := ctx.BanChatMember(userID).
		Until(time.Now().Add(24 * time.Hour)).
		RevokeMessages()

	if result == nil {
		t.Error("Expected BanChatMember builder to be created")
	}

	// Test that builder is functional
	_ = result
}
