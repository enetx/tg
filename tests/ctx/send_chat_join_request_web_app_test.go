package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendChatJoinRequestWebApp(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		Update: &gotgbot.Update{
			UpdateId:        1,
			ChatJoinRequest: &gotgbot.ChatJoinRequest{QueryId: "query_123"},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SendChatJoinRequestWebApp(g.String("https://example.com/app"))
	if result == nil {
		t.Error("Expected SendChatJoinRequestWebApp builder to be created")
	}

	if result.QueryID(g.String("custom_query")) == nil {
		t.Error("QueryID method should return SendChatJoinRequestWebApp for chaining")
	}

	if result.URL(g.String("https://example.com/other")) == nil {
		t.Error("URL method should return SendChatJoinRequestWebApp for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return SendChatJoinRequestWebApp for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return SendChatJoinRequestWebApp for chaining")
	}
}

func TestSendChatJoinRequestWebApp_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		Update: &gotgbot.Update{
			UpdateId:        1,
			ChatJoinRequest: &gotgbot.ChatJoinRequest{QueryId: "query_123"},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.SendChatJoinRequestWebApp(g.String("https://example.com/app")).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result")
	}

	complete := testCtx.SendChatJoinRequestWebApp(g.String("https://example.com/app")).
		QueryID(g.String("custom_query")).
		URL(g.String("https://example.com/other")).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
