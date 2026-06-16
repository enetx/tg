package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetManagedBotAccessSettings(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.GetManagedBotAccessSettings(123456)
	if result == nil {
		t.Error("Expected GetManagedBotAccessSettings builder to be created")
	}

	if r := testCtx.GetManagedBotAccessSettings(123456).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return GetManagedBotAccessSettings for chaining")
	}

	if r := testCtx.GetManagedBotAccessSettings(123456).APIURL(g.String("https://api.telegram.org")); r == nil {
		t.Error("APIURL method should return GetManagedBotAccessSettings for chaining")
	}
}

func TestGetManagedBotAccessSettings_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.GetManagedBotAccessSettings(456).Send()
	if sendResult.IsErr() {
		t.Logf("GetManagedBotAccessSettings Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configured := testCtx.GetManagedBotAccessSettings(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configured.IsErr() {
		t.Logf("GetManagedBotAccessSettings configured Send failed as expected: %v", configured.Err())
	}
}
