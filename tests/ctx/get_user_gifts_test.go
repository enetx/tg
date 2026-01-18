package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetUserGifts_Builder(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   -1001234567890,
			Type: "supergroup",
		},
		Update: &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	userID := int64(98765)

	builder := c.GetUserGifts(userID)
	if builder == nil {
		t.Fatal("Expected GetUserGifts builder to be created")
	}

	builder = builder.
		ExcludeUnlimited().
		ExcludeLimitedUpgradable().
		ExcludeLimitedNonUpgradable().
		ExcludeUnique().
		ExcludeFromBlockchain().
		SortByPrice().
		Offset(g.String("offset_test_user")).
		Limit(50).
		Timeout(20 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if builder == nil {
		t.Error("Chained methods should return the builder for chaining")
	}

	res := builder.Send()
	if res.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", res.Err())
	}
}
