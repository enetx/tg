package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetChatAdministratorCustomTitle(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	userID := int64(456)
	customTitle := g.String("Custom Admin Title")

	result := ctx.SetChatAdministratorCustomTitle(userID, customTitle)

	if result == nil {
		t.Error("Expected SetChatAdministratorCustomTitle builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestSetChatAdministratorCustomTitle_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	customTitle := g.String("Custom Admin Title")
	if ctx.SetChatAdministratorCustomTitle(userID, customTitle).Timeout(time.Minute) == nil { t.Error("Timeout should return builder") }
}

func TestSetChatAdministratorCustomTitle_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	customTitle := g.String("Custom Admin Title")
	if ctx.SetChatAdministratorCustomTitle(userID, customTitle).APIURL(g.String("https://api.example.com")) == nil { t.Error("APIURL should return builder") }
}

func TestSetChatAdministratorCustomTitle_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	userID := int64(456)
	customTitle := g.String("Custom Admin Title")
	
	sendResult := ctx.SetChatAdministratorCustomTitle(userID, customTitle).Send()
	
	if sendResult.IsErr() {
		t.Logf("SetChatAdministratorCustomTitle Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
