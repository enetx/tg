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
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendVideoNote(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("videonote.mp4")

	result := ctx.SendVideoNote(filename)

	if result == nil {
		t.Error("Expected SendVideoNote builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_SendVideoNoteChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("videonote.mp4")

	result := ctx.SendVideoNote(filename).
		Duration(30).
		Length(240).
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendVideoNote builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected Protect method to return builder")
	}

	// Test AllowPaidBroadcast method
	result = ctx.SendVideoNote(filename).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendVideoNote for chaining")
	}

	// Test Effect method
	result = ctx.SendVideoNote(filename).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendVideoNote for chaining")
	}
}

func TestSendVideoNote_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_videonote.mp4")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVideoNote(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendVideoNote Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVideoNote(filename).
		Duration(60).
		Length(240).
		Silent().
		Protect().
		To(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVideoNote configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendVideoNote_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendVideoNote_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendVideoNote_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	btn1 := keyboard.NewButton().Text(g.String("Play Video")).Callback(g.String("play_video"))
	if ctx.SendVideoNote(filename).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendVideoNote_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).Thumbnail(g.String("thumb.jpg")) == nil {
		t.Error("Thumbnail should return builder")
	}
}

func TestSendVideoNote_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).Reply(reply.New(123)) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendVideoNote_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendVideoNote_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("videonote.mp4")
	if ctx.SendVideoNote(filename).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendVideoNote_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
	result := ctx.SendVideoNote(invalidFilename)

	// The builder should still be created even with error
	if result == nil {
		t.Error("SendVideoNote should return builder even with invalid filename")
	}

	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}
}

func TestSendVideoNote_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("test.mp4")

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendVideoNote(filename).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestSendVideoNote_FileClosing(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file
	tempFile := "/tmp/test_video_note.mp4"
	os.WriteFile(tempFile, []byte("test video note content"), 0644)
	defer os.Remove(tempFile)

	sendResult := ctx.SendVideoNote(g.String(tempFile)).Send()

	// This will fail with mock bot, but covers the file closing path
	if sendResult.IsErr() {
		t.Logf("SendVideoNote Send with file closing failed as expected: %v", sendResult.Err())
	}
}

func TestSendVideoNote_ThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid thumbnail path that causes error
	result := ctx.SendVideoNote(g.String("test.mp4")).Thumbnail(g.String("/invalid/nonexistent/thumb.jpg"))

	if result == nil {
		t.Error("Thumbnail with error should still return builder")
	}

	// Test that Send() properly handles the thumbnail error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with invalid thumbnail")
	} else {
		t.Logf("Send failed as expected with invalid thumbnail: %v", sendResult.Err())
	}
}

func TestSendVideoNote_ThumbnailWithValidFile(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary thumbnail file
	thumbFile := "/tmp/test_thumb.jpg"
	os.WriteFile(thumbFile, []byte("test thumbnail content"), 0644)
	defer os.Remove(thumbFile)

	// Create a temporary video note file
	videoFile := "/tmp/test_video_note.mp4"
	os.WriteFile(videoFile, []byte("test video note content"), 0644)
	defer os.Remove(videoFile)

	result := ctx.SendVideoNote(g.String(videoFile)).Thumbnail(g.String(thumbFile))

	if result == nil {
		t.Error("Thumbnail with valid file should return builder")
	}

	sendResult := result.Send()

	// This will fail with mock bot, but covers the thumbnail file closing path
	if sendResult.IsErr() {
		t.Logf("SendVideoNote Send with valid thumbnail failed as expected: %v", sendResult.Err())
	}
}

func TestSendVideoNote_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video_note.mp4")

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendVideoNote(filename).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendVideoNote builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendVideoNote_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video_note.mp4")

	// Test with nil params
	result := ctx.SendVideoNote(filename).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendVideoNote builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

func TestSendVideoNote_TimersIntegration(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with After and DeleteAfter to cover timer integration
	result := ctx.SendVideoNote(g.String("test.mp4")).
		After(time.Second).
		DeleteAfter(time.Minute)

	if result == nil {
		t.Error("VideoNote with timers should return builder")
	}

	sendResult := result.Send()

	// This will fail with mock bot, but covers the timer integration path
	if sendResult.IsErr() {
		t.Logf("SendVideoNote Send with timers failed as expected: %v", sendResult.Err())
	}
}
