package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerChatJoinRequestQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		Update: &gotgbot.Update{
			UpdateId:        1,
			ChatJoinRequest: &gotgbot.ChatJoinRequest{QueryId: "query_123"},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.AnswerChatJoinRequestQuery()
	if result == nil {
		t.Error("Expected AnswerChatJoinRequestQuery builder to be created")
	}

	if result.QueryID(g.String("custom_query")) == nil {
		t.Error("QueryID method should return AnswerChatJoinRequestQuery for chaining")
	}

	if result.Approve() == nil {
		t.Error("Approve method should return AnswerChatJoinRequestQuery for chaining")
	}

	if result.Decline() == nil {
		t.Error("Decline method should return AnswerChatJoinRequestQuery for chaining")
	}

	if result.Queue() == nil {
		t.Error("Queue method should return AnswerChatJoinRequestQuery for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return AnswerChatJoinRequestQuery for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return AnswerChatJoinRequestQuery for chaining")
	}
}

func TestAnswerChatJoinRequestQuery_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		Update: &gotgbot.Update{
			UpdateId:        1,
			ChatJoinRequest: &gotgbot.ChatJoinRequest{QueryId: "query_123"},
		},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.AnswerChatJoinRequestQuery().Approve().Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result")
	}

	overridden := testCtx.AnswerChatJoinRequestQuery().
		QueryID(g.String("custom_query")).
		Decline().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !overridden.IsErr() && !overridden.IsOk() {
		t.Error("Send with all features should return a result")
	}
}

func TestAnswerChatJoinRequestQuery_SendWithoutQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.AnswerChatJoinRequestQuery().Queue().Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send without a join request in the update should still return a result")
	}
}
