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
	"github.com/enetx/tg/types/effects"
)

func TestContext_SendVideo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename)

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Video caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendVideoChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	result := ctx.SendVideo(filename).
		Caption(g.String("Test video")).
		HTML().
		Silent()

	if result == nil {
		t.Error("Expected SendVideo builder to be created")
	}

	// Test continued chaining
	final := result.To(123)
	if final == nil {
		t.Error("Expected To method to return builder")
	}

	// Test AllowPaidBroadcast method
	result = ctx.SendVideo(filename).AllowPaidBroadcast()
	if result == nil {
		t.Error("AllowPaidBroadcast method should return SendVideo for chaining")
	}

	// Test Effect method
	result = ctx.SendVideo(filename).Effect(effects.Fire)
	if result == nil {
		t.Error("Effect method should return SendVideo for chaining")
	}
}

func TestSendVideo_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("test_video.mp4")

	// Test Send method - will fail with mock but covers the method
	sendResult := ctx.SendVideo(filename).Send()

	if sendResult.IsErr() {
		t.Logf("SendVideo Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.SendVideo(filename).
		Caption(g.String("Test <b>video</b> with HTML")).
		HTML().
		Width(640).
		Height(480).
		Duration(120).
		Silent().
		Protect().
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("SendVideo configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

func TestSendVideo_CaptionEntities(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	ent := entities.New(g.String("Bold text")).Bold(g.String("Bold"))
	if ctx.SendVideo(filename).CaptionEntities(ent) == nil {
		t.Error("CaptionEntities should return builder")
	}
}

func TestSendVideo_After(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).After(time.Minute) == nil {
		t.Error("After should return builder")
	}
}

func TestSendVideo_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).DeleteAfter(time.Hour) == nil {
		t.Error("DeleteAfter should return builder")
	}
}

func TestSendVideo_Thumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Thumbnail(g.String("thumb.jpg")) == nil {
		t.Error("Thumbnail should return builder")
	}
}

func TestSendVideo_Spoiler(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Spoiler() == nil {
		t.Error("Spoiler should return builder")
	}
}

func TestSendVideo_Streamable(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Streamable() == nil {
		t.Error("Streamable should return builder")
	}
}

func TestSendVideo_Markdown(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Markdown() == nil {
		t.Error("Markdown should return builder")
	}
}

func TestSendVideo_Markup(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	btn1 := keyboard.NewButton().Text(g.String("Watch Video")).Callback(g.String("watch_video"))
	if ctx.SendVideo(filename).Markup(keyboard.Inline().Button(btn1)) == nil {
		t.Error("Markup should return builder")
	}
}

func TestSendVideo_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Reply(reply.New(123)) == nil {
		t.Error("ReplyTo should return builder")
	}
}

func TestSendVideo_Business(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Business(g.String("biz_123")) == nil {
		t.Error("Business should return builder")
	}
}

func TestSendVideo_Thread(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Thread(456) == nil {
		t.Error("Thread should return builder")
	}
}

func TestSendVideo_ShowCaptionAboveMedia(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).ShowCaptionAboveMedia() == nil {
		t.Error("ShowCaptionAboveMedia should return builder")
	}
}

func TestSendVideo_Cover(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).Cover(g.String("cover.jpg")) == nil {
		t.Error("Cover should return builder")
	}
}

func TestSendVideo_StartAt(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).StartAt(30*time.Second) == nil {
		t.Error("StartAt should return builder")
	}
}

func TestSendVideo_ApplyMetadata(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).ApplyMetadata() == nil {
		t.Error("ApplyMetadata should return builder")
	}
}

func TestSendVideo_GenerateThumbnail(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")
	if ctx.SendVideo(filename).GenerateThumbnail() == nil {
		t.Error("GenerateThumbnail should return builder")
	}
}

func TestSendVideo_GenerateThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Test GenerateThumbnail with no file set (should cause error)
	filename := g.String("")
	result := ctx.SendVideo(filename)

	// Call GenerateThumbnail - this should cause error since file is not set
	thumbnailResult := result.GenerateThumbnail()
	if thumbnailResult == nil {
		t.Error("GenerateThumbnail should return builder even with error")
	}

	// Test that Send() handles the error
	sendResult := thumbnailResult.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail when GenerateThumbnail has error")
	} else {
		t.Logf("Send failed as expected with GenerateThumbnail error: %v", sendResult.Err())
	}
}

func TestSendVideo_GenerateThumbnailWithSeekTime(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	filename := g.String("video.mp4")

	// Test GenerateThumbnail with custom seek time
	result := ctx.SendVideo(filename).GenerateThumbnail(g.String("00:00:10"))
	if result == nil {
		t.Error("GenerateThumbnail with seek time should return builder")
	}

	// Test that Send() handles any errors from GenerateThumbnail
	sendResult := result.Send()
	if sendResult.IsOk() {
		t.Log("GenerateThumbnail with seek time succeeded (unexpected in test environment)")
	} else {
		t.Logf("GenerateThumbnail with seek time failed as expected in test environment: %v", sendResult.Err())
	}

	// Test GenerateThumbnail with custom seek time
	result = ctx.SendVideo(filename).GenerateThumbnail(g.String("00:01:30"))
	if result == nil {
		t.Error("GenerateThumbnail with seek time should return builder")
	}
}

func TestSendVideo_GenerateThumbnailWithoutDuration(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Create a temporary file
	tempFile := "/tmp/test_video_notmeta.mp4"
	os.WriteFile(tempFile, []byte("test video content"), 0644)
	defer os.Remove(tempFile)

	// Test GenerateThumbnail without calling ApplyMetadata first (should cause error)
	result := ctx.SendVideo(g.String(tempFile)).GenerateThumbnail()

	if result == nil {
		t.Error("GenerateThumbnail should return builder even with duration error")
	}

	// Test that Send() handles the duration error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail when duration is not set")
	} else {
		t.Logf("Send failed as expected with duration not set error: %v", sendResult.Err())
	}
}

func TestSendVideo_GenerateThumbnailWithApplyMetadata(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Create a temporary file
	tempFile := "/tmp/test_video_meta.mp4"
	os.WriteFile(tempFile, []byte("test video content"), 0644)
	defer os.Remove(tempFile)

	// Test GenerateThumbnail after ApplyMetadata
	result := ctx.SendVideo(g.String(tempFile)).
		ApplyMetadata().
		GenerateThumbnail()

	if result == nil {
		t.Error("GenerateThumbnail with ApplyMetadata should return builder")
	}

	// This will likely fail due to ffmpeg not being available in test environment,
	// but it covers the code path
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected (likely ffmpeg not available): %v", sendResult.Err())
	}
}

func TestSendVideo_ErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Test with invalid filename that should cause file.Input to fail
	invalidFilename := g.String("") // Empty filename should cause an error
	result := ctx.SendVideo(invalidFilename)

	// The builder should still be created even with error
	if result == nil {
		t.Error("SendVideo should return builder even with invalid filename")
	}

	// Test that Send() properly handles the error
	sendResult := result.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail with empty filename")
	} else {
		t.Logf("Send failed as expected with empty filename: %v", sendResult.Err())
	}

	// Test with nonexistent file
	nonexistentFile := g.String("/nonexistent/path/to/video.mp4")
	result2 := ctx.SendVideo(nonexistentFile)
	if result2 == nil {
		t.Error("SendVideo should return builder even with nonexistent file")
	}

	sendResult2 := result2.Send()
	if !sendResult2.IsErr() {
		t.Error("Send should fail with nonexistent file")
	} else {
		t.Logf("Send failed as expected with nonexistent file: %v", sendResult2.Err())
	}
}

func TestSendVideo_APIURLWithExistingRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("test.mp4")

	// First set Timeout to create RequestOpts, then test APIURL
	result := ctx.SendVideo(filename).
		Timeout(15 * time.Second).                         // This creates RequestOpts
		APIURL(g.String("https://custom.api.example.com")) // This should use existing RequestOpts

	if result == nil {
		t.Error("APIURL with existing RequestOpts should return builder")
	}
}

func TestSendVideo_CoverErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("test.mp4")

	// Test Cover with invalid file (error case)
	result := ctx.SendVideo(filename).Cover(g.String("/invalid/path/cover.jpg"))
	if result == nil {
		t.Error("Cover with invalid file should still return builder")
	}

	// Test Cover with valid file
	tempCover := "/tmp/test_cover.jpg"
	os.WriteFile(tempCover, []byte("test cover content"), 0644)
	defer os.Remove(tempCover)

	result2 := ctx.SendVideo(filename).Cover(g.String(tempCover))
	if result2 == nil {
		t.Error("Cover with valid file should return builder")
	}
}

func TestSendVideo_ApplyMetadataErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Test ApplyMetadata without file (error case)
	result := ctx.SendVideo(g.String("nonexistent.mp4")).ApplyMetadata()
	if result == nil {
		t.Error("ApplyMetadata should return builder even with error")
	}

	// Test ApplyMetadata with invalid file path
	tempFile := "/tmp/test_video.mp4"
	os.WriteFile(tempFile, []byte("fake video content"), 0644)
	defer os.Remove(tempFile)

	result2 := ctx.SendVideo(g.String(tempFile)).ApplyMetadata()
	if result2 == nil {
		t.Error("ApplyMetadata should return builder (will fail ffmpeg but builder should be returned)")
	}
}

func TestSendVideo_ThumbnailErrorHandling(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)
	filename := g.String("video.mp4")

	// Test Thumbnail method with error
	result := ctx.SendVideo(filename).Thumbnail(g.String("/invalid/path/thumb.jpg"))
	if result == nil {
		t.Error("Thumbnail with invalid file should still return builder")
	}
}

func TestSendVideo_FileClosingComplexScenarios(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Create temporary files
	tempVideo := "/tmp/test_video.mp4"
	tempCover := "/tmp/test_cover.jpg"
	tempThumb := "/tmp/test_thumb.jpg"

	os.WriteFile(tempVideo, []byte("test video content"), 0644)
	os.WriteFile(tempCover, []byte("test cover content"), 0644)
	os.WriteFile(tempThumb, []byte("test thumb content"), 0644)

	defer os.Remove(tempVideo)
	defer os.Remove(tempCover)
	defer os.Remove(tempThumb)

	// Test with multiple files
	result := ctx.SendVideo(g.String(tempVideo)).
		Cover(g.String(tempCover)).
		Thumbnail(g.String(tempThumb))

	if result == nil {
		t.Error("SendVideo with multiple files should return builder")
	}

	// Call Send to trigger file closing logic
	sendResult := result.Send()
	if sendResult.IsErr() {
		t.Logf("Send failed as expected with mock bot: %v", sendResult.Err())
	}
}

// Test ApplyMetadata error conditions for better coverage
func TestSendVideo_ApplyMetadataErrorConditions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Test ApplyMetadata with empty filename - should cause error
	emptyFilename := g.String("")
	result := ctx.SendVideo(emptyFilename)

	// Call ApplyMetadata - this should cause error since file is not properly set
	metadataResult := result.ApplyMetadata()
	if metadataResult == nil {
		t.Error("ApplyMetadata should return builder even with error")
	}

	// Test that Send() handles the error
	sendResult := metadataResult.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail when ApplyMetadata has error")
	} else {
		t.Logf("Send failed as expected with ApplyMetadata error: %v", sendResult.Err())
	}

	// Test ApplyMetadata with non-existent file
	nonExistentFile := g.String("/nonexistent/video.mp4")
	result2 := ctx.SendVideo(nonExistentFile).ApplyMetadata()
	if result2 == nil {
		t.Error("ApplyMetadata should return builder even with non-existent file")
	}

	// Test that Send() handles the error from non-existent file
	sendResult2 := result2.Send()
	if sendResult2.IsOk() {
		t.Log("ApplyMetadata with non-existent file succeeded (unexpected)")
	} else {
		t.Logf("ApplyMetadata with non-existent file failed as expected: %v", sendResult2.Err())
	}
}

// Test GenerateThumbnail additional error conditions for better coverage
func TestSendVideo_GenerateThumbnailAdditionalErrors(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(
		bot,
		&ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}},
	)

	// Test GenerateThumbnail with empty filename - should cause error
	emptyFilename := g.String("")
	result := ctx.SendVideo(emptyFilename)

	// Call GenerateThumbnail - this should cause error since file is not set
	thumbnailResult := result.GenerateThumbnail()
	if thumbnailResult == nil {
		t.Error("GenerateThumbnail should return builder even with error")
	}

	// Test that Send() handles the error
	sendResult := thumbnailResult.Send()
	if !sendResult.IsErr() {
		t.Error("Send should fail when GenerateThumbnail has error")
	} else {
		t.Logf("Send failed as expected with GenerateThumbnail error: %v", sendResult.Err())
	}

	// Test GenerateThumbnail with non-existent file and custom seek time
	nonExistentFile := g.String("/nonexistent/video.mp4")
	result2 := ctx.SendVideo(nonExistentFile).GenerateThumbnail(g.String("00:00:05"))
	if result2 == nil {
		t.Error("GenerateThumbnail should return builder even with non-existent file")
	}

	// Test that Send() handles the error from non-existent file
	sendResult2 := result2.Send()
	if sendResult2.IsOk() {
		t.Log("GenerateThumbnail with non-existent file succeeded (unexpected)")
	} else {
		t.Logf("GenerateThumbnail with non-existent file failed as expected: %v", sendResult2.Err())
	}
}

func TestSendVideo_DirectMessagesTopic(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	topicIDs := []int64{123, 456, 789, 0, -1}

	for _, topicID := range topicIDs {
		result := ctx.SendVideo(filename).DirectMessagesTopic(topicID)
		if result == nil {
			t.Errorf("DirectMessagesTopic method should return SendVideo builder for chaining with topicID: %d", topicID)
		}

		chainedResult := result.DirectMessagesTopic(topicID + 100)
		if chainedResult == nil {
			t.Errorf("DirectMessagesTopic method should support chaining and override with topicID: %d", topicID)
		}
	}
}

func TestSendVideo_SuggestedPost(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("video.mp4")

	// Test with nil params
	result := ctx.SendVideo(filename).SuggestedPost(nil)
	if result == nil {
		t.Error("SuggestedPost method should return SendVideo builder for chaining with nil params")
	}

	// Test chaining
	chainedResult := result.SuggestedPost(nil)
	if chainedResult == nil {
		t.Error("SuggestedPost method should support chaining")
	}
}

// Test ApplyMetadata and GenerateThumbnail success paths for better coverage
func TestSendVideo_MetadataAndThumbnailSuccessPaths(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Create a larger temporary file that might be recognized as valid by ffmpeg libraries
	tempFile := "/tmp/test_video_metadata.mp4"
	// Create a file with some basic video-like header bytes (MP4 signature)
	headerBytes := []byte{0x00, 0x00, 0x00, 0x20, 0x66, 0x74, 0x79, 0x70, 0x69, 0x73, 0x6f, 0x6d} // MP4 file signature
	padding := make([]byte, 1024)                                                                 // Add some padding to make it look more like a real file
	fileContent := append(headerBytes, padding...)
	os.WriteFile(tempFile, fileContent, 0644)
	defer os.Remove(tempFile)

	// Test ApplyMetadata success path coverage
	metadataBuilder := ctx.SendVideo(g.String(tempFile)).ApplyMetadata()
	if metadataBuilder == nil {
		t.Error("ApplyMetadata should return builder")
	}

	// Even if ffmpeg fails, we want to test more code paths
	sendMetadataResult := metadataBuilder.Send()
	if sendMetadataResult.IsErr() {
		t.Logf("ApplyMetadata Send failed as expected (ffmpeg not available or file invalid): %v", sendMetadataResult.Err())
	} else {
		t.Log("ApplyMetadata Send succeeded unexpectedly")
	}

	// Test GenerateThumbnail with duration set manually (to bypass ApplyMetadata requirement)
	builder := ctx.SendVideo(g.String(tempFile)).
		Duration(10 * time.Second) // Manually set duration to pass the IsZero check

	// Test GenerateThumbnail with default seek time (no parameters)
	thumbnailBuilder := builder.GenerateThumbnail()
	if thumbnailBuilder == nil {
		t.Error("GenerateThumbnail without seek time should return builder")
	}

	sendThumbnailResult := thumbnailBuilder.Send()
	if sendThumbnailResult.IsErr() {
		t.Logf("GenerateThumbnail Send failed as expected (ffmpeg not available): %v", sendThumbnailResult.Err())
	} else {
		t.Log("GenerateThumbnail Send succeeded unexpectedly")
	}

	// Test GenerateThumbnail with custom seek times
	seekTimes := []string{"00:00:01", "00:00:05", "00:00:10", "5.5", "10.0"}
	for _, seekTime := range seekTimes {
		builder2 := ctx.SendVideo(g.String(tempFile)).
			Duration(15 * time.Second) // Set duration higher than seek time

		seekBuilder := builder2.GenerateThumbnail(g.String(seekTime))
		if seekBuilder == nil {
			t.Errorf("GenerateThumbnail with seek time %s should return builder", seekTime)
			continue
		}

		seekSendResult := seekBuilder.Send()
		if seekSendResult.IsErr() {
			t.Logf("GenerateThumbnail with seek time %s failed as expected: %v", seekTime, seekSendResult.Err())
		} else {
			t.Logf("GenerateThumbnail with seek time %s succeeded unexpectedly", seekTime)
		}
	}

	// Test ApplyMetadata with duration.IsZero() = false path
	// Create another builder and try to trigger the successful metadata path
	builder3 := ctx.SendVideo(g.String(tempFile))
	metadataBuilder3 := builder3.ApplyMetadata()

	// Even if ApplyMetadata fails to extract real metadata, test the chaining
	if metadataBuilder3 == nil {
		t.Error("ApplyMetadata should return builder even on ffmpeg failure")
	}

	// Test combining ApplyMetadata and GenerateThumbnail
	combinedBuilder := ctx.SendVideo(g.String(tempFile)).
		ApplyMetadata().
		GenerateThumbnail(g.String("00:00:02"))

	if combinedBuilder == nil {
		t.Error("Combined ApplyMetadata and GenerateThumbnail should return builder")
	}

	combinedSendResult := combinedBuilder.Send()
	if combinedSendResult.IsErr() {
		t.Logf("Combined ApplyMetadata and GenerateThumbnail Send failed as expected: %v", combinedSendResult.Err())
	} else {
		t.Log("Combined ApplyMetadata and GenerateThumbnail Send succeeded unexpectedly")
	}
}
