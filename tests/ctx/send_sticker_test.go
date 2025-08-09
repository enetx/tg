package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
)

func TestContext_SendSticker(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	result := ctx.SendSticker(sticker)

	if result == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendStickerChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	result := ctx.SendSticker(sticker).
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected SendSticker builder to be created")
	}

	// Test continued chaining
	final := result.Thread(456)
	if final == nil {
		t.Error("Expected Thread method to return builder")
	}
}

func TestSendSticker_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendSticker(sticker).Send()

	if sendResult.IsErr() {
		t.Logf("SendSticker Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendSticker(sticker).
		Silent().
		Protect().
		To(123).
		Thread(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendSticker configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendSticker_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).After(time.Minute) == nil { t.Error("After should return builder") }
}

func TestSendSticker_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).DeleteAfter(time.Hour) == nil { t.Error("DeleteAfter should return builder") }
}

func TestSendSticker_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	btn1 := keyboard.NewButton().Text(g.String("Cool Sticker!")).Callback(g.String("cool_sticker"))
	if ctx.SendSticker(sticker).Markup(keyboard.Inline().Button(btn1)) == nil { t.Error("Markup should return builder") }
}

func TestSendSticker_Emoji(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).Emoji("😀") == nil { t.Error("Emoji should return builder") }
}

func TestSendSticker_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).ReplyTo(123) == nil { t.Error("ReplyTo should return builder") }
}

func TestSendSticker_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).Business(g.String("biz_123")) == nil { t.Error("Business should return builder") }
}

func TestSendSticker_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("")  // Empty filename should cause an error
	result := ctx.SendSticker(invalidFilename)
	
	// The builder should still be created even with error
	if result == nil {
		t.Error("SendSticker should return builder even with invalid filename")
	}
	
	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
}
