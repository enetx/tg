package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_RepostStory_Builder(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   -1001234567890,
			Type: "supergroup",
		},
		Update: &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	businessConnectionID := g.String("business_conn_123")
	fromChatID := int64(123456789)
	fromStoryID := int64(42)

	builder := c.RepostStory(businessConnectionID, fromChatID, fromStoryID)
	if builder == nil {
		t.Fatal("Expected RepostStory builder to be created")
	}

	builder = builder.
		PostToChatPage().
		Protect().
		ActiveFor(24 * time.Hour).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if builder == nil {
		t.Error("Chained methods should return the builder for chaining")
	}

	res := builder.Send()
	if res.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", res.Err())
	}
}

func TestContext_RepostStory_PostToChatPage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   123456789,
			Type: "private",
		},
		Update: &gotgbot.Update{UpdateId: 2},
	}

	c := ctx.New(bot, rawCtx)

	builder := c.RepostStory(g.String("conn_1"), 123, 1).
		PostToChatPage()

	if builder == nil {
		t.Error("PostToChatPage() should return the builder for chaining")
	}
}

func TestContext_RepostStory_Protect(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   123456789,
			Type: "private",
		},
		Update: &gotgbot.Update{UpdateId: 3},
	}

	c := ctx.New(bot, rawCtx)

	builder := c.RepostStory(g.String("conn_2"), 456, 2).
		Protect()

	if builder == nil {
		t.Error("Protect() should return the builder for chaining")
	}
}
