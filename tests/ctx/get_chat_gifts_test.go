package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetChatGifts_Builder(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   -1001234567890,
			Type: "supergroup",
		},
		Update: &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)

	builder := c.GetChatGifts()
	if builder == nil {
		t.Fatal("Expected GetChatGifts builder to be created")
	}

	builder = builder.
		ExcludeUnsaved().
		ExcludeSaved().
		ExcludeUnlimited().
		ExcludeLimitedUpgradable().
		ExcludeLimitedNonUpgradable().
		ExcludeUnique().
		ExcludeFromBlockchain().
		SortByPrice().
		Offset(g.String("offset_test")).
		Limit(50).
		Timeout(15 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if builder == nil {
		t.Error("Chained methods should return the builder for chaining")
	}

	res := builder.Send()
	if res.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", res.Err())
	}
}
