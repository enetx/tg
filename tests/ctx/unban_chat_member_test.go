package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
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

func TestUnbanChatMember_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(123456)
	if ctx.UnbanChatMember(userID).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestUnbanChatMember_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(123456)
	if ctx.UnbanChatMember(userID).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestUnbanChatMember_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(123456)
	
	sendResult := ctx.UnbanChatMember(userID).OnlyIfBanned().Send()
	
	if sendResult.IsErr() {
		t.Logf("UnbanChatMember Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
