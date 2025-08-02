package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/input"
)

func TestContext_AnswerWebAppQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	queryID := g.String("web_app_query_123")
	messageContent := input.Text(g.String("Test content"))
	result := inline.NewArticle(g.String("1"), g.String("Test Article"), messageContent)

	answerQuery := ctx.AnswerWebAppQuery(queryID, result)

	if answerQuery == nil {
		t.Error("Expected AnswerWebAppQuery builder to be created")
	}

	// Test method chaining
	withTimeout := answerQuery.Timeout(30)
	if withTimeout == nil {
		t.Error("Expected Timeout method to return builder")
	}

	withAPI := answerQuery.APIURL(g.String("https://api.custom.url"))
	if withAPI == nil {
		t.Error("Expected APIURL method to return builder")
	}
}
