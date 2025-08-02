package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendDocument(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("nonexistent.pdf")

	result := ctx.SendDocument(filename)

	// Test that the builder is created properly
	if result == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("test"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendDocumentChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("document.pdf")

	result := ctx.SendDocument(filename).
		Caption(g.String("Test document")).
		HTML().
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendDocument builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected protect method to return builder")
	}
}
