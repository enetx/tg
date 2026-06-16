package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
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

func TestSendPoll_QuestionEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendPoll(question).QuestionEntities(ent) == nil {
		t.Error("QuestionEntities should return builder")
	}
}

func TestSendPoll_ExplanationEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendPoll(question).ExplanationEntities(ent) == nil {
		t.Error("ExplanationEntities should return builder")
	}
}

func TestSendPoll_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendPoll_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendPoll_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendPoll_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendPoll_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).AllowPaidBroadcast() == nil {
		t.Error("AllowPaidBroadcast should return builder")
	}
}

func TestSendPoll_Effect(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Effect(effects.Fire) == nil {
		t.Error("Effect should return builder")
	}
}

func TestSendPoll_Quiz(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Quiz(0) == nil {
		t.Error("Quiz should return builder")
	}
}

func TestSendPoll_Explanation(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Explanation(g.String("This is a test poll")) == nil {
		t.Error("Explanation should return builder")
	}
}

func TestSendPoll_ExplanationHTML(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).ExplanationHTML() == nil {
		t.Error("ExplanationHTML should return builder")
	}
}

func TestSendPoll_ExplanationMarkdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).ExplanationMarkdown() == nil {
		t.Error("ExplanationMarkdown should return builder")
	}
}

func TestSendPoll_ClosesIn(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).ClosesIn(time.Hour) == nil {
		t.Error("ClosesIn should return builder")
	}
}

func TestSendPoll_ClosesAt(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	closeTime := time.Now().Add(time.Hour)
	if ctx.SendPoll(question).ClosesAt(closeTime) == nil {
		t.Error("ClosesAt should return builder")
	}
}

func TestSendPoll_Closed(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Closed() == nil {
		t.Error("Closed should return builder")
	}
}

func TestSendPoll_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	btn1 := keyboard.NewButton().Text(g.String("Vote Now")).Callback(g.String("vote_now"))
	if ctx.SendPoll(question).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendPoll_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	question := g.String("What's your favorite color?")
	if ctx.SendPoll(question).Reply(reply.New(123)) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func newPollCtx() *ctx.Context {
	bot := &mockBot{}
	return ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
}

func TestSendPoll_QuizMultipleCorrect(t *testing.T) {
	question := g.String("Select all primes")
	if newPollCtx().SendPoll(question).Quiz(0, 2, 4) == nil {
		t.Error("Quiz with multiple correct options should return builder")
	}
}

func TestSendPoll_Media(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).Media(input.LocationMedia(40.0, -74.0)) == nil {
		t.Error("Media should return builder")
	}
}

func TestSendPoll_ExplanationMedia(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).ExplanationMedia(input.VenueMedia(40.0, -74.0, g.String("T"), g.String("A"))) == nil {
		t.Error("ExplanationMedia should return builder")
	}
}

func TestSendPoll_Description(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).Description(g.String("desc")) == nil {
		t.Error("Description should return builder")
	}
}

func TestSendPoll_DescriptionHTML(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).DescriptionHTML() == nil {
		t.Error("DescriptionHTML should return builder")
	}
}

func TestSendPoll_DescriptionMarkdown(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).DescriptionMarkdown() == nil {
		t.Error("DescriptionMarkdown should return builder")
	}
}

func TestSendPoll_DescriptionEntities(t *testing.T) {
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if newPollCtx().SendPoll(g.String("Q")).DescriptionEntities(ent) == nil {
		t.Error("DescriptionEntities should return builder")
	}
}

func TestSendPoll_AllowRevoting(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).AllowRevoting() == nil {
		t.Error("AllowRevoting should return builder")
	}
}

func TestSendPoll_QuestionHTML(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).QuestionHTML() == nil {
		t.Error("QuestionHTML should return builder")
	}
}

func TestSendPoll_QuestionMarkdown(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).QuestionMarkdown() == nil {
		t.Error("QuestionMarkdown should return builder")
	}
}

func TestSendPoll_ShuffleOptions(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).ShuffleOptions() == nil {
		t.Error("ShuffleOptions should return builder")
	}
}

func TestSendPoll_AllowAddingOptions(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).AllowAddingOptions() == nil {
		t.Error("AllowAddingOptions should return builder")
	}
}

func TestSendPoll_HideResultsUntilClosed(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).HideResultsUntilClosed() == nil {
		t.Error("HideResultsUntilClosed should return builder")
	}
}

func TestSendPoll_MembersOnly(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).MembersOnly() == nil {
		t.Error("MembersOnly should return builder")
	}
}

func TestSendPoll_CountryCodes(t *testing.T) {
	if newPollCtx().SendPoll(g.String("Q")).CountryCodes(g.String("US"), g.String("GB")) == nil {
		t.Error("CountryCodes should return builder")
	}
}

func TestSendPoll_OptionMediaChaining(t *testing.T) {
	question := g.String("Pick a place")
	result := newPollCtx().SendPoll(question).
		Option(input.Choice("Home").Media(input.StickerMedia(file.Input(g.String("CAACAgID")).Ok()))).
		Option(input.Choice("Work").Media(input.LocationMedia(1.0, 2.0)))
	if result == nil {
		t.Error("Option with media should chain correctly")
	}
}
