package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SavePreparedKeyboardButton(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SavePreparedKeyboardButton(123456)
	if result == nil {
		t.Error("Expected SavePreparedKeyboardButton builder to be created")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).Text(g.String("Pick users")); r == nil {
		t.Error("Text method should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).RequestUsers(7); r == nil {
		t.Error("RequestUsers method should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).RequestChat(7, true); r == nil {
		t.Error("RequestChat(channel=true) should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).RequestChat(7, false); r == nil {
		t.Error("RequestChat(channel=false) should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).RequestManagedBot(7); r == nil {
		t.Error("RequestManagedBot method should return SavePreparedKeyboardButton for chaining")
	}

	// SuggestedName/SuggestedUsername should be safe even before RequestManagedBot —
	// they lazily initialize the underlying gotgbot struct.
	if r := testCtx.SavePreparedKeyboardButton(123456).SuggestedName(g.String("MyBot")); r == nil {
		t.Error("SuggestedName method should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).SuggestedUsername(g.String("my_bot")); r == nil {
		t.Error("SuggestedUsername method should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return SavePreparedKeyboardButton for chaining")
	}

	if r := testCtx.SavePreparedKeyboardButton(123456).APIURL(g.String("https://api.telegram.org")); r == nil {
		t.Error("APIURL method should return SavePreparedKeyboardButton for chaining")
	}
}

func TestSavePreparedKeyboardButton_FullChain(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SavePreparedKeyboardButton(123456).
		Text(g.String("Create managed bot")).
		RequestManagedBot(42).
		SuggestedName(g.String("Helper")).
		SuggestedUsername(g.String("helper_bot")).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return SavePreparedKeyboardButton")
	}
}

func TestSavePreparedKeyboardButton_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.SavePreparedKeyboardButton(456).
		Text(g.String("Share")).
		RequestUsers(1).
		Send()

	if sendResult.IsErr() {
		t.Logf("SavePreparedKeyboardButton Send failed as expected with mock bot: %v", sendResult.Err())
	}
}
