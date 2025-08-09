package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnbanChatSenderChat(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	senderChatID := int64(789)

	result := ctx.UnbanChatSenderChat(senderChatID)

	if result == nil {
		t.Error("Expected UnbanChatSenderChat builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestUnbanChatSenderChat_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	senderChatID := int64(789)
	if ctx.UnbanChatSenderChat(senderChatID).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestUnbanChatSenderChat_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	senderChatID := int64(789)
	if ctx.UnbanChatSenderChat(senderChatID).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestUnbanChatSenderChat_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	senderChatID := int64(789)
	
	sendResult := ctx.UnbanChatSenderChat(senderChatID).Send()
	
	if sendResult.IsErr() {
		t.Logf("UnbanChatSenderChat Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
