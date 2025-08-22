package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestDeclineSuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	dsp := testCtx.DeclineSuggestedPost()
	if dsp == nil {
		t.Error("Expected DeclineSuggestedPost to return a valid instance")
	}
}

func TestDeclineSuggestedPost_ChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	dsp := testCtx.DeclineSuggestedPost().ChatID(123456)
	if dsp == nil {
		t.Error("Expected ChatID method to return DeclineSuggestedPost for chaining")
	}
}

func TestDeclineSuggestedPost_MessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	dsp := testCtx.DeclineSuggestedPost().MessageID(789)
	if dsp == nil {
		t.Error("Expected MessageID method to return DeclineSuggestedPost for chaining")
	}
}

func TestDeclineSuggestedPost_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.DeclineSuggestedPost().
		ChatID(123456).
		MessageID(789).
		Send()

	if result.IsOk() {
		t.Error("Expected Send to fail with mock bot")
	}
}
