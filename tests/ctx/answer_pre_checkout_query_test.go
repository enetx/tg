package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_AnswerPreCheckoutQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			PreCheckoutQuery: &gotgbot.PreCheckoutQuery{
				Id:          "query_123",
				Currency:    "USD",
				TotalAmount: 100,
			},
		},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.AnswerPreCheckoutQuery()

	if result == nil {
		t.Error("Expected AnswerPreCheckoutQuery builder to be created")
	}

	// Test method chaining
	okResult := result.Ok()
	if okResult == nil {
		t.Error("Expected Ok method to return builder")
	}

	errorResult := result.Error(g.String("Payment failed"))
	if errorResult == nil {
		t.Error("Expected Error method to return builder")
	}
}
