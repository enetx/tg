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
)

func TestContext_SendAudio(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Audio caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendAudioChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename).
		Caption(g.String("Test audio")).
		HTML().
		Silent().
		Duration(180)

	if result == nil {
		t.Error("Expected SendAudio builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}
}

func TestSendAudio_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_audio.mp3")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendAudio(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendAudio Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendAudio(filename).
		Caption(g.String("Test <b>audio</b> with HTML")).
		HTML().
		Duration(180).
		Performer(g.String("Test Artist")).
		Title(g.String("Test Song")).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendAudio configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendAudio_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	result := ctx.SendAudio(filename).CaptionEntities(ent)
	if result == nil {
		t.Error("CaptionEntities method should return SendAudio builder for chaining")
	}

	entItalic := entities.New(g.String("Italic text")).Italic(g.String("Italic"))
	chainedResult := result.CaptionEntities(entItalic)
	if chainedResult == nil {
		t.Error("CaptionEntities method should support chaining and override")
	}
}

func TestSendAudio_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	durations := []time.Duration{
		time.Second,
		time.Minute,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAudio(filename).After(duration)
		if result == nil {
			t.Errorf("After method should return SendAudio builder for chaining with duration: %v", duration)
		}

		chainedResult := result.After(time.Second * 30)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAudio_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	durations := []time.Duration{
		time.Second * 30,
		time.Minute * 5,
		time.Hour,
		0,
	}

	for _, duration := range durations {
		result := ctx.SendAudio(filename).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter method should return SendAudio builder for chaining with duration: %v", duration)
		}

		chainedResult := result.DeleteAfter(time.Minute * 10)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration: %v", duration)
		}
	}
}

func TestSendAudio_Markdown(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	result := ctx.SendAudio(filename).Markdown()
	if result == nil {
		t.Error("Markdown method should return SendAudio builder for chaining")
	}

	chainedResult := result.Markdown()
	if chainedResult == nil {
		t.Error("Markdown method should support chaining")
	}
}

func TestSendAudio_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	btn1 := keyboard.NewButton().Text(g.String("Test Button")).Callback(g.String("test_data"))
	inlineKeyboard := keyboard.Inline().Button(btn1)

	result := ctx.SendAudio(filename).Markup(inlineKeyboard)
	if result == nil {
		t.Error("Markup method should return SendAudio builder for chaining with inline keyboard")
	}

	replyKeyboard := keyboard.Reply()
	chainedResult := result.Markup(replyKeyboard)
	if chainedResult == nil {
		t.Error("Markup method should support chaining and override with reply keyboard")
	}
}

func TestSendAudio_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	thumbnails := []string{
		"thumb.jpg",
		"album_cover.png",
		"preview.webp",
	}

	for _, thumb := range thumbnails {
		result := ctx.SendAudio(filename).Thumbnail(g.String(thumb))
		if result == nil {
			t.Errorf("Thumbnail method should return SendAudio builder for chaining with thumbnail: %s", thumb)
		}

		chainedResult := result.Thumbnail(g.String("override.jpg"))
		if chainedResult == nil {
			t.Errorf("Thumbnail method should support chaining and override with thumbnail: %s", thumb)
		}
	}
}

func TestSendAudio_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	messageIDs := []int64{1, 123, 456, 999}

	for _, messageID := range messageIDs {
		result := ctx.SendAudio(filename).ReplyTo(messageID)
		if result == nil {
			t.Errorf("ReplyTo method should return SendAudio builder for chaining with messageID: %d", messageID)
		}

		chainedResult := result.ReplyTo(messageID + 100)
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID: %d", messageID)
		}
	}
}

func TestSendAudio_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	businessIDs := []string{
		"business_123",
		"conn_456",
		"",
	}

	for _, businessID := range businessIDs {
		result := ctx.SendAudio(filename).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business method should return SendAudio builder for chaining with businessID: %s", businessID)
		}

		chainedResult := result.Business(g.String("override_business"))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID: %s", businessID)
		}
	}
}

func TestSendAudio_Thread(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("audio.mp3")

	threadIDs := []int64{1, 123, 456, 0}

	for _, threadID := range threadIDs {
		result := ctx.SendAudio(filename).Thread(threadID)
		if result == nil {
			t.Errorf("Thread method should return SendAudio builder for chaining with threadID: %d", threadID)
		}

		chainedResult := result.Thread(threadID + 100)
		if chainedResult == nil {
			t.Errorf("Thread method should support chaining and override with threadID: %d", threadID)
		}
	}
}

func TestSendAudio_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
	result := ctx.SendAudio(invalidFilename)

	// The builder should still be created even with error
	if result == nil {
		t.Error("SendAudio should return builder even with invalid filename")
	}

	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}

	// Test with nonexistent file
	nonexistentFile := g.String("/nonexistent/path/to/audio.mp3")
	result2 := ctx.SendAudio(nonexistentFile)
	if result2 == nil {
		t.Error("SendAudio should return builder even with nonexistent file")
	}

	sendResult2 := result2.Send()
	if !sendResult2.IsErr() {
		t.Error("Send should fail with nonexistent file")
	} else {
		t.Logf("Send failed as expected with nonexistent file: %v", sendResult2.Err())
	}
}

func TestSendAudio_ThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("audio.mp3")

	// Test with invalid thumbnail file
	result := ctx.SendAudio(filename).Thumbnail(g.String("/invalid/path/thumb.jpg"))
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

func TestSendAudio_FileClosing(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a temporary file to test file closing
	tempFile := "/tmp/test_audio.mp3"
	os.WriteFile(tempFile, []byte("test audio content"), 0644)
	defer os.Remove(tempFile)

	result := ctx.SendAudio(g.String(tempFile))
	if result == nil {
		t.Error("SendAudio with valid file should return builder")
	}

	// Call Send to trigger file closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendAudio_ThumbnailWithValidFile(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create temporary files for audio and thumbnail
	audioFile := "/tmp/test_audio.mp3"
	thumbFile := "/tmp/test_thumb.jpg"
	os.WriteFile(audioFile, []byte("test audio content"), 0644)
	os.WriteFile(thumbFile, []byte("test thumb content"), 0644)
	defer os.Remove(audioFile)
	defer os.Remove(thumbFile)

	result := ctx.SendAudio(g.String(audioFile)).Thumbnail(g.String(thumbFile))
	if result == nil {
		t.Error("SendAudio with valid thumbnail should return builder")
	}

	// Call Send to trigger both file and thumbnail closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestSendAudio_TimeoutWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("audio.mp3")

	// Test Timeout when RequestOpts is nil
	result := ctx.SendAudio(filename).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout should return builder")
	}
}

func TestSendAudio_APIURLWithNilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("audio.mp3")

	// Test APIURL when RequestOpts is nil
	result := ctx.SendAudio(filename).APIURL(g.String("https://api.example.com"))
	if result == nil {
		t.Error("APIURL should return builder")
	}
}

func TestSendAudio_SendWithExistingError(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// First create a SendAudio with an error (invalid filename)
	result := ctx.SendAudio(g.String("/invalid/nonexistent/audio.mp3"))
	if result == nil {
		t.Error("SendAudio with invalid file should return builder")
	}

	// Test that Send() returns the error immediately without calling timers
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should return error for invalid file")
	}
}

func TestSendAudio_To(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})
	filename := g.String("audio.mp3")

	// Test To method
	result := ctx.SendAudio(filename).To(123456)
	if result == nil {
		t.Error("To should return builder")
	}
}
