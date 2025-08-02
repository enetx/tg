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
