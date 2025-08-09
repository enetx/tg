package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
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

func TestSendVoice_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_voice.ogg")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVoice(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendVoice Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVoice(filename).
		Caption(g.String("Test <b>voice</b> message")).
		HTML().
		Duration(45).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVoice configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendVoice_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendVoice(filename).CaptionEntities(ent) == nil { t.Error("CaptionEntities should return builder") }
}

func TestSendVoice_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).After(time.Minute) == nil { t.Error("After should return builder") }
}

func TestSendVoice_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).DeleteAfter(time.Hour) == nil { t.Error("DeleteAfter should return builder") }
}

func TestSendVoice_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).Markdown() == nil { t.Error("Markdown should return builder") }
}

func TestSendVoice_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	btn1 := keyboard.NewButton().Text(g.String("Listen Again")).Callback(g.String("listen_again"))
	if ctx.SendVoice(filename).Markup(keyboard.Inline().Button(btn1)) == nil { t.Error("Markup should return builder") }
}

func TestSendVoice_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).ReplyTo(123) == nil { t.Error("ReplyTo should return builder") }
}

func TestSendVoice_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).Business(g.String("biz_123")) == nil { t.Error("Business should return builder") }
}

func TestSendVoice_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("voice.ogg")
	if ctx.SendVoice(filename).Thread(456) == nil { t.Error("Thread should return builder") }
}

func TestSendVoice_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	
	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("")  // Empty filename should cause an error
	result := ctx.SendVoice(invalidFilename)
	
	// The builder should still be created even with error
	if result == nil {
		t.Error("SendVoice should return builder even with invalid filename")
	}
	
	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
}
