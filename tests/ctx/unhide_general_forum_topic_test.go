package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_UnhideGeneralForumTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.UnhideGeneralForumTopic()

	if result == nil {
		t.Error("Expected UnhideGeneralForumTopic builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(-1001234567890)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestUnhideGeneralForumTopic_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.UnhideGeneralForumTopic().Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestUnhideGeneralForumTopic_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.UnhideGeneralForumTopic().APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestUnhideGeneralForumTopic_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"}, Update: &gotgbot.Update{UpdateId: 1}})

	sendResult := ctx.UnhideGeneralForumTopic().Send()

	if sendResult.IsErr() {
		t.Logf("UnhideGeneralForumTopic Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
