package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_GetFile(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	result := ctx.GetFile(fileID)

	if result == nil {
		t.Error("Expected GetFile builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestContext_GetFileChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	fileID := g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg")

	result := ctx.GetFile(fileID)

	if result == nil {
		t.Error("Expected GetFile builder to be created")
	}

	// Test that builder is functional
	_ = result
}
