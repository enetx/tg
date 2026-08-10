package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/richtext"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

func richTestMessage() *rich.Message {
	return rich.New().Block(rich.Paragraph(richtext.New("Hello").Bold()))
}

func TestContext_SendRichMessage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SendRichMessage(richTestMessage())
	if result == nil {
		t.Error("Expected SendRichMessage builder to be created")
	}

	if result.To(999) == nil {
		t.Error("To method should return SendRichMessage for chaining")
	}

	if result.Silent() == nil {
		t.Error("Silent method should return SendRichMessage for chaining")
	}

	if result.Protect() == nil {
		t.Error("Protect method should return SendRichMessage for chaining")
	}

	if result.Effect(effects.Fire) == nil {
		t.Error("Effect method should return SendRichMessage for chaining")
	}

	if result.Reply(reply.New(789)) == nil {
		t.Error("Reply method should return SendRichMessage for chaining")
	}

	if result.Reply(nil) == nil {
		t.Error("Reply method should tolerate nil parameters")
	}

	if result.Markup(keyboard.Inline()) == nil {
		t.Error("Markup method should return SendRichMessage for chaining")
	}

	if result.AllowPaidBroadcast() == nil {
		t.Error("AllowPaidBroadcast method should return SendRichMessage for chaining")
	}

	if result.Thread(10) == nil {
		t.Error("Thread method should return SendRichMessage for chaining")
	}

	if result.DirectMessagesTopic(20) == nil {
		t.Error("DirectMessagesTopic method should return SendRichMessage for chaining")
	}

	if result.Business(g.String("business_conn_123")) == nil {
		t.Error("Business method should return SendRichMessage for chaining")
	}

	if result.SuggestedPost(suggested.New()) == nil {
		t.Error("SuggestedPost method should return SendRichMessage for chaining")
	}

	if result.SuggestedPost(nil) == nil {
		t.Error("SuggestedPost method should tolerate nil parameters")
	}

	if result.Timeout(30*time.Second) == nil {
		t.Error("Timeout method should return SendRichMessage for chaining")
	}

	if result.APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL method should return SendRichMessage for chaining")
	}

	if result.After(time.Millisecond) == nil {
		t.Error("After method should return SendRichMessage for chaining")
	}

	if result.DeleteAfter(time.Millisecond) == nil {
		t.Error("DeleteAfter method should return SendRichMessage for chaining")
	}
}

func TestSendRichMessage_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.SendRichMessage(richTestMessage()).Send()
	if !sendResult.IsErr() && !sendResult.IsOk() {
		t.Error("Send should return a result")
	}

	scheduled := testCtx.SendRichMessage(richTestMessage()).
		To(999).
		After(time.Millisecond).
		Send()

	if !scheduled.IsOk() {
		t.Error("Send with After should return Ok for scheduled execution")
	}
}
