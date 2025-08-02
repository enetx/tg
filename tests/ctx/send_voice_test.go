package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
)

func TestContext_SendVoice(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("voice.ogg")

	result := ctx.SendVoice(filename)

	if result == nil {
		t.Error("Expected SendVoice builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Voice message"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendVoiceChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("voice.ogg")

	result := ctx.SendVoice(filename).
		Caption(g.String("Test voice")).
		Duration(30).
		Silent()

	if result == nil {
		t.Error("Expected SendVoice builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}
