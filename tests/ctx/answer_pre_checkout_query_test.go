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

func TestAnswerPreCheckoutQuery_TimeoutAPIURL(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			PreCheckoutQuery: &gotgbot.PreCheckoutQuery{
				Id:          "timeout_test",
				Currency:    "USD",
				TotalAmount: 200,
			},
		},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Timeout method
	result := ctx.AnswerPreCheckoutQuery().Ok().Timeout(30)
	if result == nil {
		t.Error("Expected Timeout method to return builder")
	}

	// Test APIURL method
	urlResult := ctx.AnswerPreCheckoutQuery().Ok().APIURL(g.String("https://api.example.com"))
	if urlResult == nil {
		t.Error("Expected APIURL method to return builder")
	}
}

func TestAnswerPreCheckoutQuery_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			PreCheckoutQuery: &gotgbot.PreCheckoutQuery{
				Id:          "send_test",
				Currency:    "USD",
				TotalAmount: 300,
			},
		},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with Ok
	sendResult := ctx.AnswerPreCheckoutQuery().Ok().Send()
	if sendResult.IsErr() {
		t.Logf("AnswerPreCheckoutQuery Ok Send failed as expected: %v", sendResult.Err())
	}

	// Test Send method with Error
	errorSendResult := ctx.AnswerPreCheckoutQuery().Error(g.String("Payment error")).Send()
	if errorSendResult.IsErr() {
		t.Logf("AnswerPreCheckoutQuery Error Send failed as expected: %v", errorSendResult.Err())
	}
}

func TestAnswerPreCheckoutQuery_Send_NoQuery(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update: &gotgbot.Update{
			UpdateId: 1,
			// No PreCheckoutQuery - this will test the nil case
		},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method when PreCheckoutQuery is nil
	sendResult := ctx.AnswerPreCheckoutQuery().Ok().Send()
	if !sendResult.IsErr() {
		t.Error("Expected Send to return error when PreCheckoutQuery is nil")
	}

	expectedError := "no precheckout query"
	if !g.SliceOf(sendResult.Err().Error()).Contains(expectedError) {
		t.Errorf("Expected error to contain '%s', got: %v", expectedError, sendResult.Err())
	}
}
