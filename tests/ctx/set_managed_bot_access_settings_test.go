package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SetManagedBotAccessSettings(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SetManagedBotAccessSettings(123456, true)
	if result == nil {
		t.Error("Expected SetManagedBotAccessSettings builder to be created")
	}

	if r := testCtx.SetManagedBotAccessSettings(123456, false); r == nil {
		t.Error("SetManagedBotAccessSettings should accept isAccessRestricted=false")
	}

	if r := testCtx.SetManagedBotAccessSettings(123456, true).AddedUserIDs(1, 2, 3); r == nil {
		t.Error("AddedUserIDs method should return SetManagedBotAccessSettings for chaining")
	}

	if r := testCtx.SetManagedBotAccessSettings(123456, true).Timeout(30 * time.Second); r == nil {
		t.Error("Timeout method should return SetManagedBotAccessSettings for chaining")
	}

	if r := testCtx.SetManagedBotAccessSettings(123456, true).APIURL(g.String("https://api.telegram.org")); r == nil {
		t.Error("APIURL method should return SetManagedBotAccessSettings for chaining")
	}
}

func TestSetManagedBotAccessSettings_AddedUserIDsMax(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Per Bot API 10.0 up to 10 IDs are supported; the builder accepts more and
	// defers validation to the server.
	manyIDs := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := testCtx.SetManagedBotAccessSettings(123456, true).AddedUserIDs(manyIDs...)
	if result == nil {
		t.Error("AddedUserIDs should accept up to 10 IDs")
	}

	// Variadic with no args is a valid call (matches the field's nil semantics).
	if r := testCtx.SetManagedBotAccessSettings(123456, true).AddedUserIDs(); r == nil {
		t.Error("AddedUserIDs with no IDs should still return SetManagedBotAccessSettings")
	}
}

func TestContext_SetManagedBotAccessSettings_Chaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	result := testCtx.SetManagedBotAccessSettings(123456, true).
		AddedUserIDs(111, 222, 333).
		Timeout(45 * time.Second).
		APIURL(g.String("https://custom-api.telegram.org"))

	if result == nil {
		t.Error("Complete method chaining should work and return SetManagedBotAccessSettings")
	}
}

func TestSetManagedBotAccessSettings_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	sendResult := testCtx.SetManagedBotAccessSettings(456, true).Send()
	if sendResult.IsErr() {
		t.Logf("SetManagedBotAccessSettings Send failed as expected with mock bot: %v", sendResult.Err())
	}

	configured := testCtx.SetManagedBotAccessSettings(456, true).
		AddedUserIDs(1, 2).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configured.IsErr() {
		t.Logf("SetManagedBotAccessSettings configured Send failed as expected: %v", configured.Err())
	}

	unrestricted := testCtx.SetManagedBotAccessSettings(456, false).Send()
	if unrestricted.IsErr() {
		t.Logf("SetManagedBotAccessSettings unrestricted Send failed as expected: %v", unrestricted.Err())
	}
}
