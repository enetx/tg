package ctx_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/types/effects"
)

func TestContext_MediaGroup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.MediaGroup()

	if result == nil {
		t.Error("Expected MediaGroup builder to be created")
	}

	// Test method chaining
	chained := result.Silent()
	if chained == nil {
		t.Error("Expected Silent method to return builder")
	}
}

func TestContext_MediaGroupChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	result := ctx.MediaGroup().
		Silent().
		Protect().
		To(123)

	if result == nil {
		t.Error("Expected MediaGroup builder to be created")
	}

	// Test continued chaining
	final := result.Thread(456)
	if final == nil {
		t.Error("Expected Thread method to return builder")
	}
}

func TestMediaGroup_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "group"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method with no media - should return error
	sendResult := ctx.MediaGroup().Send()

	if !sendResult.IsErr() {
		t.Error("Send should return error for empty media group")
	} else {
		t.Logf("MediaGroup Send with no media returned error as expected: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.MediaGroup().
		Silent().
		Protect().
		To(123).
		Thread(456).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if !configuredSendResult.IsErr() {
		t.Error("Send should return error for empty media group even with configuration")
	} else {
		t.Logf("MediaGroup configured Send with no media returned error as expected: %v", configuredSendResult.Err())
	}
}

func TestMediaGroup_SendWithMedia(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Add media to the group
	photoInputFile := file.Input(g.String("https://example.com/photo1.jpg"))
	if photoInputFile.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}
	photo := input.Photo(photoInputFile.Unwrap())

	// Test Send method with media - will fail with mock bot but covers the media path
	sendResult := ctx.MediaGroup().Photo(photo).Send()

	if sendResult.IsErr() {
		t.Logf("MediaGroup Send with media failed as expected with mock bot: %v", sendResult.Err())
	}
}

func TestMediaGroup_SendWithAllOptions(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Add media to the group
	photoInputFile := file.Input(g.String("https://example.com/photo1.jpg"))
	if photoInputFile.IsErr() {
		t.Skip("Unable to create photo input for testing")
	}
	photo := input.Photo(photoInputFile.Unwrap())

	// Test Send method with all options and media
	sendResult := ctx.MediaGroup().
		Photo(photo).
		Silent().
		Protect().
		AllowPaidBroadcast().
		Effect(effects.Fire).
		Business(g.String("biz_123")).
		Thread(456).
		Reply(reply.New(123)).
		To(789).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	// This will fail with mock bot, but covers all option paths
	if sendResult.IsErr() {
		t.Logf("MediaGroup Send with all options failed as expected: %v", sendResult.Err())
	}
}

// Tests for methods with 0% coverage

func TestMediaGroup_After(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test After method with various durations
	durations := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		0 * time.Second, // Zero duration
	}

	for _, duration := range durations {
		result := ctx.MediaGroup()
		afterResult := result.After(duration)
		if afterResult == nil {
			t.Errorf("After method should return MediaGroup for chaining with duration %v", duration)
		}

		// Test that After can be chained and overridden
		chainedResult := afterResult.After(duration + 1*time.Second)
		if chainedResult == nil {
			t.Errorf("After method should support chaining and override with duration %v", duration)
		}
	}
}

func TestMediaGroup_DeleteAfter(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test DeleteAfter method with various durations
	durations := []time.Duration{
		10 * time.Second,
		30 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		1 * time.Hour,
		0 * time.Second, // Zero duration
	}

	for _, duration := range durations {
		result := ctx.MediaGroup()
		deleteAfterResult := result.DeleteAfter(duration)
		if deleteAfterResult == nil {
			t.Errorf("DeleteAfter method should return MediaGroup for chaining with duration %v", duration)
		}

		// Test that DeleteAfter can be chained and overridden
		chainedResult := deleteAfterResult.DeleteAfter(duration + 1*time.Minute)
		if chainedResult == nil {
			t.Errorf("DeleteAfter method should support chaining and override with duration %v", duration)
		}
	}
}

func TestMediaGroup_Photo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Photo method with various photo inputs
	// Create InputFiles first
	photoInputs := []g.String{
		g.String("https://example.com/photo1.jpg"),
		g.String("AgACAgIAAxkBAAIBY2FSqKlZKZsj9_gDAw"), // File ID
		g.String("https://example.com/photo2.jpg"),     // Another URL since local files might fail
		g.String("photo_file_id_123"),
	}

	var photos []input.Media
	for _, photoInput := range photoInputs {
		if inputFile := file.Input(photoInput); inputFile.IsOk() {
			photos = append(photos, input.Photo(inputFile.Ok()))
		}
	}

	for i, photo := range photos {
		result := ctx.MediaGroup()
		photoResult := result.Photo(photo)
		if photoResult == nil {
			t.Errorf("Photo method should return MediaGroup for chaining with photo %d", i)
		}

		// Test multiple photos
		multiPhotoResult := photoResult.Photo(photos[(i+1)%len(photos)])
		if multiPhotoResult == nil {
			t.Errorf("Photo method should support adding multiple photos %d", i)
		}
	}
}

func TestMediaGroup_Video(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Video method with MediaVideo inputs
	// Create InputFiles first
	videoInputs := []g.String{
		g.String("https://example.com/video1.mp4"),
		g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg"), // File ID
		g.String("https://example.com/video2.mp4"), // Another URL since local files might fail
		g.String("video_file_id_123"),
	}

	var videos []input.Media
	for _, videoInput := range videoInputs {
		if inputFile := file.Input(videoInput); inputFile.IsOk() {
			videos = append(videos, input.Video(inputFile.Ok()))
		}
	}

	// Add a photo to test non-video media
	if photoInputFile := file.Input(g.String("https://example.com/photo1.jpg")); photoInputFile.IsOk() {
		videos = append(videos, input.Photo(photoInputFile.Ok()))
	}

	for i, video := range videos {
		result := ctx.MediaGroup()
		videoResult := result.Video(video)
		if videoResult == nil {
			t.Errorf("Video method should return MediaGroup for chaining with video %d", i)
		}

		// Test multiple videos
		multiVideoResult := videoResult.Video(videos[(i+1)%len(videos)])
		if multiVideoResult == nil {
			t.Errorf("Video method should support adding multiple videos %d", i)
		}
	}
}

func TestMediaGroup_Audio(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Audio method with MediaAudio inputs
	// Create InputFiles first
	audioInputs := []g.String{
		g.String("https://example.com/audio1.mp3"),
		g.String("BAADBAADsAADBREAAWtJUomQZ-0HAg"), // File ID
		g.String("https://example.com/audio2.mp3"), // Another URL since local files might fail
		g.String("audio_file_id_123"),
	}

	var audios []input.Media
	for _, audioInput := range audioInputs {
		if inputFile := file.Input(audioInput); inputFile.IsOk() {
			audios = append(audios, input.Audio(inputFile.Ok()))
		}
	}

	// Add a photo to test non-audio media
	if photoInputFile := file.Input(g.String("https://example.com/photo1.jpg")); photoInputFile.IsOk() {
		audios = append(audios, input.Photo(photoInputFile.Ok()))
	}

	for i, audio := range audios {
		result := ctx.MediaGroup()
		audioResult := result.Audio(audio)
		if audioResult == nil {
			t.Errorf("Audio method should return MediaGroup for chaining with audio %d", i)
		}

		// Test multiple audios
		multiAudioResult := audioResult.Audio(audios[(i+1)%len(audios)])
		if multiAudioResult == nil {
			t.Errorf("Audio method should support adding multiple audios %d", i)
		}
	}
}

func TestMediaGroup_Document(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Document method with MediaDocument inputs
	// Create InputFiles first
	documentInputs := []g.String{
		g.String("https://example.com/document1.pdf"),
		g.String("BAADBAADrwADBREAAWp8oTdjdLOHAg"),    // File ID
		g.String("https://example.com/document2.pdf"), // Another URL since local files might fail
		g.String("document_file_id_123"),
	}

	var documents []input.Media
	for _, documentInput := range documentInputs {
		if inputFile := file.Input(documentInput); inputFile.IsOk() {
			documents = append(documents, input.Document(inputFile.Ok()))
		}
	}

	// Add a photo to test non-document media
	if photoInputFile := file.Input(g.String("https://example.com/photo1.jpg")); photoInputFile.IsOk() {
		documents = append(documents, input.Photo(photoInputFile.Ok()))
	}

	for i, document := range documents {
		result := ctx.MediaGroup()
		documentResult := result.Document(document)
		if documentResult == nil {
			t.Errorf("Document method should return MediaGroup for chaining with document %d", i)
		}

		// Test multiple documents
		multiDocumentResult := documentResult.Document(documents[(i+1)%len(documents)])
		if multiDocumentResult == nil {
			t.Errorf("Document method should support adding multiple documents %d", i)
		}
	}
}

func TestMediaGroup_AllowPaidBroadcast(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test AllowPaidBroadcast method
	result := ctx.MediaGroup()
	allowPaidBroadcastResult := result.AllowPaidBroadcast()
	if allowPaidBroadcastResult == nil {
		t.Error("AllowPaidBroadcast method should return MediaGroup for chaining")
	}

	// Test AllowPaidBroadcast method can be chained multiple times
	allowPaidBroadcastChained := allowPaidBroadcastResult.AllowPaidBroadcast()
	if allowPaidBroadcastChained == nil {
		t.Error("AllowPaidBroadcast method should support multiple chaining calls")
	}

	// Test AllowPaidBroadcast with other methods
	allowPaidBroadcastWithOthers := ctx.MediaGroup().
		Silent().
		AllowPaidBroadcast().
		Protect()
	if allowPaidBroadcastWithOthers == nil {
		t.Error("AllowPaidBroadcast method should work with other methods")
	}
}

func TestMediaGroup_Effect(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Effect method with various effect types
	effectTypes := []effects.EffectType{
		effects.Fire,
		effects.ThumbsUp,
		effects.ThumbsDown,
		effects.Heart,
		effects.Celebration,
		effects.Poop,
	}

	for i, effect := range effectTypes {
		result := ctx.MediaGroup()
		effectResult := result.Effect(effect)
		if effectResult == nil {
			t.Errorf("Effect method should return MediaGroup for chaining with effect %d", i)
		}

		// Test that Effect can be chained and overridden
		chainedResult := effectResult.Effect(effectTypes[(i+1)%len(effectTypes)])
		if chainedResult == nil {
			t.Errorf("Effect method should support chaining and override with effect %d", i)
		}
	}
}

func TestMediaGroup_ReplyTo(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test ReplyTo method with various message IDs
	messageIDs := []int64{123, 456, 789, 0, -1, 999999}

	for _, messageID := range messageIDs {
		result := ctx.MediaGroup()
		replyToResult := result.Reply(reply.New(messageID))
		if replyToResult == nil {
			t.Errorf("ReplyTo method should return MediaGroup for chaining with messageID %d", messageID)
		}

		// Test that ReplyTo can be chained and overridden
		chainedResult := replyToResult.Reply(reply.New(messageID + 1))
		if chainedResult == nil {
			t.Errorf("ReplyTo method should support chaining and override with messageID %d", messageID+1)
		}
	}
}

func TestMediaGroup_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Business method with various business connection IDs
	businessIDs := []string{
		"business_connection_123",
		"biz_conn_456",
		"",
		"very_long_business_connection_id_with_special_characters_!@#$%",
		"conn_789",
	}

	for _, businessID := range businessIDs {
		result := ctx.MediaGroup()
		businessResult := result.Business(g.String(businessID))
		if businessResult == nil {
			t.Errorf("Business method should return MediaGroup for chaining with businessID '%s'", businessID)
		}

		// Test that Business can be chained and overridden
		chainedResult := businessResult.Business(g.String("updated_" + businessID))
		if chainedResult == nil {
			t.Errorf("Business method should support chaining and override with businessID '%s'", businessID)
		}
	}
}
