package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
)

func TestApproveSuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	asp := testCtx.ApproveSuggestedPost()
	if asp == nil {
		t.Error("Expected ApproveSuggestedPost to return a valid instance")
	}
}

func TestApproveSuggestedPost_ChatID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	asp := testCtx.ApproveSuggestedPost().ChatID(123456)
	if asp == nil {
		t.Error("Expected ChatID method to return ApproveSuggestedPost for chaining")
	}
}

func TestApproveSuggestedPost_MessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	asp := testCtx.ApproveSuggestedPost().MessageID(789)
	if asp == nil {
		t.Error("Expected MessageID method to return ApproveSuggestedPost for chaining")
	}
}

func TestApproveSuggestedPost_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 123},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.ApproveSuggestedPost().
		ChatID(123456).
		MessageID(789).
		Send()

	if result.IsOk() {
		t.Error("Expected Send to fail with mock bot")
	}
}
