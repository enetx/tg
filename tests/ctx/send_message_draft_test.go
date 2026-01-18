package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
)

func TestContext_SendMessageDraft_Builder(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   -1001234567890,
			Type: "supergroup",
		},
		Update: &gotgbot.Update{UpdateId: 1},
	}

	c := ctx.New(bot, rawCtx)
	draftID := int64(12345)
	text := g.String("Draft message text")

	builder := c.SendMessageDraft(draftID, text)
	if builder == nil {
		t.Fatal("Expected SendMessageDraft builder to be created")
	}

	builder = builder.
		To(123456789).
		Thread(999).
		HTML().
		Timeout(10 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if builder == nil {
		t.Error("Chained methods should return the builder for chaining")
	}

	res := builder.Send()
	if res.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", res.Err())
	}
}

func TestContext_SendMessageDraft_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   123456789,
			Type: "private",
		},
		Update: &gotgbot.Update{UpdateId: 2},
	}

	c := ctx.New(bot, rawCtx)

	builder := c.SendMessageDraft(1, g.String("**bold** text")).
		Markdown()

	if builder == nil {
		t.Error("Markdown() should return the builder for chaining")
	}
}

func TestContext_SendMessageDraft_Entities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{
			Id:   123456789,
			Type: "private",
		},
		Update: &gotgbot.Update{UpdateId: 3},
	}

	c := ctx.New(bot, rawCtx)

	ent := entities.New(g.String("text with bold")).Bold(g.String("bold"))
	builder := c.SendMessageDraft(1, g.String("text with bold")).
		Entities(ent)

	if builder == nil {
		t.Error("Entities() should return the builder for chaining")
	}
}
