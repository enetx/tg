package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/input"
)

func TestContext_SendPoll(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	question := g.String("What's your favorite color?")

	result := ctx.SendPoll(question)

	if result == nil {
		t.Error("Expected SendPoll builder to be created")
	}

	// Test method chaining
	chained := result.Anonymous()
	if chained == nil {
		t.Error("Expected Anonymous method to return builder")
	}
}

func TestContext_SendPollChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	question := g.String("What's your favorite color?")

	result := ctx.SendPoll(question).
		Option(input.Choice("Red")).
		Option(input.Choice("Blue")).
		Anonymous().
		MultipleAnswers().
		Silent()

	if result == nil {
		t.Error("Expected SendPoll builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendPoll_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	question := g.String("What's your favorite color?")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendPoll(question).Send()

	if sendResult.IsErr() {
		t.Logf("SendPoll Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendPoll(question).
		Option(input.Choice("Red")).
		Option(input.Choice("Blue")).
		Option(input.Choice("Green")).
		Anonymous().
		MultipleAnswers().
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendPoll configured Send failed as expected: %v", configuredSendResult.Err())
	}
}
