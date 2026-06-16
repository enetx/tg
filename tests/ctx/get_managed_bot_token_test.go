package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetManagedBotToken(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.GetManagedBotToken(123456)
	if result == nil {
		t.Error("Expected GetManagedBotToken builder to be created")
	}

	if r := testCtx.GetManagedBotToken(123456).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return GetManagedBotToken for chaining")
	}

	if r := testCtx.GetManagedBotToken(123456).APIURL(g.String("https://api.telegram.org")); r == nil {
		t.Error("APIURL method should return GetManagedBotToken for chaining")
	}
}

func TestContext_GetManagedBotToken_Chaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.GetManagedBotToken(123456).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return GetManagedBotToken")
	}
}

func TestGetManagedBotToken_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.GetManagedBotToken(456).Send()
	if sendResult.IsErr() {
		t.Logf("GetManagedBotToken Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configured := testCtx.GetManagedBotToken(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configured.IsErr() {
		t.Logf("GetManagedBotToken configured Send failed as expected: %v", configured.Err())
	}
}
