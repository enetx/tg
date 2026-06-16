package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_ReplaceManagedBotToken(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.ReplaceManagedBotToken(123456)
	if result == nil {
		t.Error("Expected ReplaceManagedBotToken builder to be created")
	}

	if r := testCtx.ReplaceManagedBotToken(123456).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return ReplaceManagedBotToken for chaining")
	}

	if r := testCtx.ReplaceManagedBotToken(123456).APIURL(g.String("https://api.telegram.org")); r == nil {
		t.Error("APIURL method should return ReplaceManagedBotToken for chaining")
	}
}

func TestContext_ReplaceManagedBotToken_Chaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.ReplaceManagedBotToken(123456).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return ReplaceManagedBotToken")
	}
}

func TestReplaceManagedBotToken_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.ReplaceManagedBotToken(456).Send()
	if sendResult.IsErr() {
		t.Logf("ReplaceManagedBotToken Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configured := testCtx.ReplaceManagedBotToken(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configured.IsErr() {
		t.Logf("ReplaceManagedBotToken configured Send failed as expected: %v", configured.Err())
	}
}
