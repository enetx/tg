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

func TestSendDocument_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_document.pdf")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendDocument(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendDocument Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendDocument(filename).
		Caption(g.String("Test <b>document</b> with HTML")).
		HTML().
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendDocument configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendDocument_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendDocument(filename).CaptionEntities(ent) == nil {
		t.Error("CaptionEntities should return builder")
	}
}

func TestSendDocument_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendDocument_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendDocument_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).Markdown() == nil {
		t.Error("Markdown should return builder")
	}
}

func TestSendDocument_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	btn1 := keyboard.NewButton().Text(g.String("Test")).Callback(g.String("test"))
	if ctx.SendDocument(filename).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendDocument_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).Thumbnail(g.String("thumb.jpg")) == nil {
		t.Error("Thumbnail should return builder")
	}
}

func TestSendDocument_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).ReplyTo(123) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendDocument_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendDocument_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendDocument_DisableContentTypeDetection(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")
	if ctx.SendDocument(filename).DisableContentTypeDetection() == nil {
		t.Error("DisableContentTypeDetection should return builder")
	}
}
