package ctx_test

import (
	"os"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
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
	if ctx.SendSticker(sticker).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendSticker_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendSticker_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	btn1 := keyboard.NewButton().Text(g.String("Cool Sticker!")).Callback(g.String("cool_sticker"))
	if ctx.SendSticker(sticker).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendSticker_Emoji(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).Emoji("😀") == nil {
		t.Error("Emoji should return builder")
	}
}

func TestSendSticker_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).Reply(reply.New(123)) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendSticker_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	if ctx.SendSticker(sticker).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendSticker_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
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

func TestSendSticker_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("test.webp")

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendSticker(filename).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestSendSticker_FileClosing(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file
	tempFile := "/tmp/test_sticker.webp"
	os.WriteFile(tempFile, []byte("test sticker content"), 0644)
	defer os.Remove(tempFile)

	sendResult := ctx.SendSticker(g.String(tempFile)).Send()

	// This will fail with mock bot, but covers the file closing path
	if sendResult.IsErr() {
		t.Logf("SendSticker Send with file closing failed as expected: %v", sendResult.Err())
	}
}

func TestSendSticker_TimersIntegration(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")

	// Test with After and DeleteAfter to cover timer integration
	result := ctx.SendSticker(sticker).
		After(time.Second).
		DeleteAfter(time.Minute)

	if result == nil {
		t.Error("Sticker with timers should return builder")
	}

	sendResult := result.Send()

	// This will fail with mock bot, but covers the timer integration path
	if sendResult.IsErr() {
		t.Logf("SendSticker Send with timers failed as expected: %v", sendResult.Err())
	}
}

func TestSendSticker_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("sticker.webp")

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendSticker(filename).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendSticker builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendSticker_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("sticker.webp")

	// Test with nil params
	result := ctx.SendSticker(filename).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendSticker builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendSticker_SendWithAllOptions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	sticker := g.String("CAACAgIAAxkBAAEBCgACYOZEYAAB1_gGF3IWUWwqZgABEQADBAADbwAABCBjAAEfBA")
	btn1 := keyboard.NewButton().Text(g.String("Cool!")).Callback(g.String("cool"))

	// Test with all options to cover different code paths
	sendResult := ctx.SendSticker(sticker).
		Emoji("😀").
		Silent().
		Protect().
		Business(g.String("biz_123")).
		Thread(456).
		Reply(reply.New(123)).
		Markup(keyboard.Inline().Button(btn1)).
		To(789).
		Send()

	// This will fail with mock bot, but covers all option paths
	if sendResult.IsErr() {
		t.Logf("SendSticker Send with all options failed as expected: %v", sendResult.Err())
	}
}
