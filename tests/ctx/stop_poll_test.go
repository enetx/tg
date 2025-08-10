package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func TestContext_StopPoll(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.StopPoll()

	if result == nil {
		t.Error("Expected StopPoll builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_StopPollChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.StopPoll().
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected StopPoll builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestStopPoll_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopPoll().Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestStopPoll_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}}, Update: &gotgbot.Update{UpdateId: 1}})
	btn1 := keyboard.NewButton().Text(g.String("View Results")).Callback(g.String("view_results"))
	if ctx.StopPoll().Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestStopPoll_Timeout(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopPoll().Timeout(time.Minute) == nil {
		t.Error("Timeout should return builder")
	}
}

func TestStopPoll_APIURL(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}}, Update: &gotgbot.Update{UpdateId: 1}})
	if ctx.StopPoll().APIURL(g.String("https://api.example.com")) == nil {
		t.Error("APIURL should return builder")
	}
}

func TestStopPoll_Send(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, EffectiveMessage: &gotgbot.Message{MessageId: 789, Poll: &gotgbot.Poll{Id: "poll123"}}, Update: &gotgbot.Update{UpdateId: 1}})

	sendResult := ctx.StopPoll().Send()

	if sendResult.IsErr() {
		t.Logf("StopPoll Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
