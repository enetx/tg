package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
)

func TestContext_AnswerGuestQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("guest_query_abc")
	messageContent := input.Text(g.String("Guest response content"))
	result := inline.NewArticle(g.String("1"), g.String("Guest Article"), messageContent)

	answerQuery := testCtx.AnswerGuestQuery(queryID, result)
	if answerQuery == nil {
		t.Error("Expected AnswerGuestQuery builder to be created")
	}

	timeoutResult := answerQuery.Timeout(30 * time.Second)
	if timeoutResult == nil {
		t.Error("Timeout method should return AnswerGuestQuery for chaining")
	}

	apiURLResult := testCtx.AnswerGuestQuery(queryID, result).APIURL(g.String("https://api.telegram.org"))
	if apiURLResult == nil {
		t.Error("APIURL method should return AnswerGuestQuery for chaining")
	}
}

func TestContext_AnswerGuestQueryChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 2},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("guest_chaining_456")
	messageContent := input.Text(g.String("Chaining content"))
	result := inline.NewArticle(g.String("2"), g.String("Chaining Article"), messageContent)

	chained := testCtx.AnswerGuestQuery(queryID, result).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if chained == nil {
		t.Error("Complete method chaining should work and return AnswerGuestQuery")
	}
}

func TestAnswerGuestQuery_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	queryID := g.String("send_guest_query")
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("send1"), g.String("Send Test Article"), messageContent)

	sendResult := testCtx.AnswerGuestQuery(queryID, result).Send()
	if sendResult.IsErr() {
		t.Logf("AnswerGuestQuery Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configuredSendResult := testCtx.AnswerGuestQuery(queryID, result).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("AnswerGuestQuery configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
