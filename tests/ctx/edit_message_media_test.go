package ctx_test

import (
	"testing"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestContext_EditMessageMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	media := input.Photo(mockFile)
	result := ctx.EditMessageMedia(media)

	if result == nil {
		t.Error("Expected EditMessageMedia builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_EditMessageMediaChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	media := input.Photo(mockFile)
	result := ctx.EditMessageMedia(media).
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageMedia builder to be created")
	}

	// Test that builder is functional
	_ = result
}
