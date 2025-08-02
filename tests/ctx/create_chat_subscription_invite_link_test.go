package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_CreateChatSubscriptionInviteLink(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	subscriptionPeriod := int64(2592000) // 30 days
	subscriptionPrice := int64(100)      // 100 stars

	result := ctx.CreateChatSubscriptionInviteLink(subscriptionPeriod, subscriptionPrice)

	if result == nil {
		t.Error("Expected CreateChatSubscriptionInviteLink builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}

	named := chained.Name(g.String("Premium Subscription"))
	if named == nil {
		t.Error("Expected Name method to return builder")
	}
}
