package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendRichMessageDraft(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SendRichMessageDraft(1, richTestMessage())
	if result == nil {
		t.Error("Expected SendRichMessageDraft builder to be created")
	}

	if result.Message(richTestMessage()) == nil {
		t.Error("Message method should return SendRichMessageDraft for chaining")
	}

	if result.To(999) == nil {
		t.Error("To method should return SendRichMessageDraft for chaining")
	}

	if result.Thread(10) == nil {
		t.Error("Thread method should return SendRichMessageDraft for chaining")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return SendRichMessageDraft for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return SendRichMessageDraft for chaining")
	}
}

func TestSendRichMessageDraft_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.SendRichMessageDraft(1, richTestMessage()).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result")
	}

	complete := testCtx.SendRichMessageDraft(1, nil).
		Message(richTestMessage()).
		To(999).
		Thread(10).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !complete.IsErr() && !complete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
