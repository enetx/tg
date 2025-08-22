package ctx_test

import (
	"os"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
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
	if ctx.SendDocument(filename).Reply(reply.New(123)) == nil {
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

func TestSendDocument_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
	result := ctx.SendDocument(invalidFilename)

	// The builder should still be created even with error
	if result == nil {
		t.Error("SendDocument should return builder even with invalid filename")
	}

	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
}

func TestSendDocument_ThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")

	// Test with invalid thumbnail file
	result := ctx.SendDocument(filename).Thumbnail(g.String("/invalid/path/thumb.jpg"))
	if result == nil {
		t.Error("Thumbnail with invalid file should still return builder")
	}

	// Test that Send() handles thumbnail error properly
	sendResult := result.Send()
	if sendResult.IsOk() {
		t.Logf("Send succeeded despite invalid thumbnail: %v", sendResult.Ok())
	} else {
		t.Logf("Send failed as expected with invalid thumbnail: %v", sendResult.Err())
	}
}

func TestSendDocument_FileClosing(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file to test file closing
	tempFile := "/tmp/test_document.txt"
	os.WriteFile(tempFile, []byte("test content"), 0644)
	defer os.Remove(tempFile)

	result := ctx.SendDocument(g.String(tempFile))
	if result == nil {
		t.Error("SendDocument with valid file should return builder")
	}

	// Call Send to trigger file closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendDocument_ThumbnailWithValidFile(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create temporary files for document and thumbnail
	docFile := "/tmp/test_document.txt"
	thumbFile := "/tmp/test_thumb.jpg"
	os.WriteFile(docFile, []byte("test document content"), 0644)
	os.WriteFile(thumbFile, []byte("test thumb content"), 0644)
	defer os.Remove(docFile)
	defer os.Remove(thumbFile)

	result := ctx.SendDocument(g.String(docFile)).Thumbnail(g.String(thumbFile))
	if result == nil {
		t.Error("SendDocument with valid thumbnail should return builder")
	}

	// Call Send to trigger both file and thumbnail closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendDocument_TimeoutWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")

	// Test Timeout when RequestOpts is nil
	result := ctx.SendDocument(filename).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSendDocument_APIURLWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("document.pdf")

	// Test APIURL when RequestOpts is nil
	result := ctx.SendDocument(filename).APIURL(g.String("https://api.example.com"))
	if result == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSendDocument_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("document.pdf")

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendDocument(filename).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendDocument builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendDocument_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("document.pdf")

	// Test with nil params
	result := ctx.SendDocument(filename).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendDocument builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendDocument_SendWithExistingError(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First create a SendDocument with an error (invalid filename)
	result := ctx.SendDocument(g.String("/invalid/nonexistent/file.pdf"))
	if result == nil {
		t.Error("SendDocument with invalid file should return builder")
	}

	// Test that Send() returns the error immediately without calling timers
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should return error for invalid file")
	}
}
