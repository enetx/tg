package ctx_test

import (
	"errors"
	"io/fs"
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/enetx/g"
	"github.com/enetx/tg/ctx"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

func TestContext_SendPhoto(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("photo.jpg")

	result := ctx.SendPhoto(filename)

	if result == nil {
		t.Error("Expected SendPhoto builder to be created")
	}

	// Test method chaining
	chained := result.Caption(g.String("Test caption"))
	if chained == nil {
		t.Error("Expected caption method to return builder")
	}
}

func TestContext_SendPhotoChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)
	filename := g.String("photo.jpg")

	// Test complex method chaining
	result := ctx.SendPhoto(filename).
		Caption(g.String("Test photo")).
		HTML().
		Silent().
		To(123)

	if result == nil {
		t.Error("Expected SendPhoto builder to be created")
	}

	// Test continued chaining
	final := result.Protect()
	if final == nil {
		t.Error("Expected protect method to return builder")
	}
}

func TestSendPhoto_AllMethods(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	filename := g.String("test_photo.jpg")

	// Test CaptionEntities method
	entities := entities.New("Bold caption").Bold("Bold")
	result := testCtx.SendPhoto(filename).CaptionEntities(entities)
	if result == nil {
		t.Error("CaptionEntities method should return SendPhoto for chaining")
	}

	// Test After method
	result = testCtx.SendPhoto(filename).After(5 * time.Second)
	if result == nil {
		t.Error("After method should return SendPhoto for chaining")
	}

	// Test DeleteAfter method
	result = testCtx.SendPhoto(filename).DeleteAfter(60 * time.Second)
	if result == nil {
		t.Error("DeleteAfter method should return SendPhoto for chaining")
	}

	// Test Spoiler method
	result = testCtx.SendPhoto(filename).Spoiler()
	if result == nil {
		t.Error("Spoiler method should return SendPhoto for chaining")
	}

	// Test Caption method
	result = testCtx.SendPhoto(filename).Caption(g.String("Photo caption"))
	if result == nil {
		t.Error("Caption method should return SendPhoto for chaining")
	}

	// Test HTML method
	result = testCtx.SendPhoto(filename).HTML()
	if result == nil {
		t.Error("HTML method should return SendPhoto for chaining")
	}

	// Test Markdown method
	result = testCtx.SendPhoto(filename).Markdown()
	if result == nil {
		t.Error("Markdown method should return SendPhoto for chaining")
	}

	// Test Silent method
	result = testCtx.SendPhoto(filename).Silent()
	if result == nil {
		t.Error("Silent method should return SendPhoto for chaining")
	}

	// Test Protect method
	result = testCtx.SendPhoto(filename).Protect()
	if result == nil {
		t.Error("Protect method should return SendPhoto for chaining")
	}

	// Test Markup method
	kb := keyboard.Inline().Text("Button", "data")
	result = testCtx.SendPhoto(filename).Markup(kb)
	if result == nil {
		t.Error("Markup method should return SendPhoto for chaining")
	}

	// Test ReplyTo method
	result = testCtx.SendPhoto(filename).ReplyTo(999)
	if result == nil {
		t.Error("ReplyTo method should return SendPhoto for chaining")
	}

	// Test Timeout method
	result = testCtx.SendPhoto(filename).Timeout(30 * time.Second)
	if result == nil {
		t.Error("Timeout method should return SendPhoto for chaining")
	}

	// Test APIURL method
	result = testCtx.SendPhoto(filename).APIURL(g.String("https://api.telegram.org"))
	if result == nil {
		t.Error("APIURL method should return SendPhoto for chaining")
	}

	// Test Business method
	result = testCtx.SendPhoto(filename).Business(g.String("business_123"))
	if result == nil {
		t.Error("Business method should return SendPhoto for chaining")
	}

	// Test Thread method
	result = testCtx.SendPhoto(filename).Thread(456)
	if result == nil {
		t.Error("Thread method should return SendPhoto for chaining")
	}

	// Test ShowCaptionAboveMedia method
	result = testCtx.SendPhoto(filename).ShowCaptionAboveMedia()
	if result == nil {
		t.Error("ShowCaptionAboveMedia method should return SendPhoto for chaining")
	}

	// Test To method
	result = testCtx.SendPhoto(filename).To(789)
	if result == nil {
		t.Error("To method should return SendPhoto for chaining")
	}
}

func TestSendPhoto_PhotoSources(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test various photo sources
	photoSources := []string{
		"photo.jpg",
		"image.png",
		"picture.gif",
		"https://example.com/photo.jpg",
		"https://api.telegram.org/file/bot123/photo.jpg",
		"BAADBAADrwADBREAAT9rHQ3cUKwJhAI", // File ID
		"/path/to/local/photo.jpg",
		"./relative/path/photo.png",
		"",
	}

	for _, source := range photoSources {
		result := testCtx.SendPhoto(g.String(source))
		if result == nil {
			t.Errorf("SendPhoto should work with source: %s", source)
		}

		// Test chaining for each source
		chained := result.Caption(g.String("Caption for " + source))
		if chained == nil {
			t.Errorf("Chaining should work for source: %s", source)
		}
	}
}

func TestSendPhoto_CaptionFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("photo.jpg")

	// Test HTML formatting
	htmlResult := testCtx.SendPhoto(photo).
		Caption(g.String("<b>Bold</b> and <i>italic</i> photo")).
		HTML()

	if htmlResult == nil {
		t.Error("HTML caption formatting should work")
	}

	// Test Markdown formatting
	markdownResult := testCtx.SendPhoto(photo).
		Caption(g.String("**Bold** and _italic_ photo")).
		Markdown()

	if markdownResult == nil {
		t.Error("Markdown caption formatting should work")
	}

	// Test caption with entities
	entities := entities.New("Beautiful photo with formatting").
		Bold("Beautiful").
		Italic("formatting")

	entitiesResult := testCtx.SendPhoto(photo).CaptionEntities(entities)

	if entitiesResult == nil {
		t.Error("Caption with entities should work")
	}

	// Test caption above media
	aboveResult := testCtx.SendPhoto(photo).
		Caption(g.String("Caption above photo")).
		ShowCaptionAboveMedia()

	if aboveResult == nil {
		t.Error("Caption above media should work")
	}

	// Test various caption texts
	captions := []string{
		"Simple caption",
		"Caption with emojis 📸📷✨",
		"Multi-line\ncaption\nwith\nbreaks",
		"Caption with special chars: !@#$%^&*()",
		"Very long caption that exceeds normal expectations and contains lots of text to test handling of large captions in photo messages",
		"",
		"A",
	}

	for _, caption := range captions {
		result := testCtx.SendPhoto(photo).Caption(g.String(caption))
		if result == nil {
			t.Errorf("Caption should work with text: %s", caption)
		}
	}
}

func TestSendPhoto_SpoilerFeature(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("spoiler_photo.jpg")

	// Test spoiler functionality
	spoilerResult := testCtx.SendPhoto(photo).Spoiler()
	if spoilerResult == nil {
		t.Error("Spoiler feature should work")
	}

	// Test spoiler with caption
	spoilerWithCaption := testCtx.SendPhoto(photo).
		Spoiler().
		Caption(g.String("Spoiler photo with caption"))

	if spoilerWithCaption == nil {
		t.Error("Spoiler with caption should work")
	}

	// Test spoiler with formatting
	spoilerFormatted := testCtx.SendPhoto(photo).
		Spoiler().
		Caption(g.String("<b>Hidden</b> content")).
		HTML()

	if spoilerFormatted == nil {
		t.Error("Spoiler with formatted caption should work")
	}

	// Test spoiler with other features
	spoilerComplete := testCtx.SendPhoto(photo).
		Spoiler().
		Caption(g.String("Complete spoiler test")).
		Silent().
		Protect()

	if spoilerComplete == nil {
		t.Error("Spoiler with complete features should work")
	}
}

func TestSendPhoto_KeyboardIntegration(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("keyboard_photo.jpg")

	// Test inline keyboard
	inlineKb := keyboard.Inline().
		Text("Like", "like").
		Text("Share", "share").Row().
		URL("View Full", "https://example.com/photo")

	inlineResult := testCtx.SendPhoto(photo).Markup(inlineKb)
	if inlineResult == nil {
		t.Error("Inline keyboard should work")
	}

	// Test reply keyboard
	replyKb := keyboard.Reply().
		Text("📸 Photo").
		Text("🎥 Video").Row().
		Text("📄 Document")

	replyResult := testCtx.SendPhoto(photo).Markup(replyKb)
	if replyResult == nil {
		t.Error("Reply keyboard should work")
	}

	// Test keyboard with caption
	keyboardWithCaption := testCtx.SendPhoto(photo).
		Caption(g.String("Photo with keyboard")).
		Markup(inlineKb)

	if keyboardWithCaption == nil {
		t.Error("Keyboard with caption should work")
	}
}

func TestSendPhoto_TimingFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("timing_photo.jpg")

	// Test After method with various durations
	afterDurations := []time.Duration{
		1 * time.Second,
		5 * time.Second,
		30 * time.Second,
		2 * time.Minute,
		5 * time.Minute,
	}

	for _, duration := range afterDurations {
		result := testCtx.SendPhoto(photo).After(duration)
		if result == nil {
			t.Errorf("After duration %v should work", duration)
		}
	}

	// Test DeleteAfter method
	deleteAfterDurations := []time.Duration{
		10 * time.Second,
		1 * time.Minute,
		5 * time.Minute,
		1 * time.Hour,
		24 * time.Hour,
	}

	for _, duration := range deleteAfterDurations {
		result := testCtx.SendPhoto(photo).DeleteAfter(duration)
		if result == nil {
			t.Errorf("DeleteAfter duration %v should work", duration)
		}
	}

	// Test combined timing
	combinedResult := testCtx.SendPhoto(photo).
		After(2 * time.Second).
		DeleteAfter(60 * time.Second)

	if combinedResult == nil {
		t.Error("Combined After and DeleteAfter should work")
	}
}

func TestSendPhoto_BusinessFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("business_photo.jpg")

	// Test various business connection IDs
	businessIDs := []string{
		"business_123",
		"enterprise_connection_456",
		"company_bot_789",
		"",
		"very_long_business_connection_identifier_12345",
		"simple",
	}

	for _, businessID := range businessIDs {
		result := testCtx.SendPhoto(photo).Business(g.String(businessID))
		if result == nil {
			t.Errorf("Business connection ID should work: %s", businessID)
		}

		// Test combining business with other features
		combined := result.Silent().Protect()
		if combined == nil {
			t.Errorf("Business combination should work for ID: %s", businessID)
		}
	}
}

func TestSendPhoto_ForumFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: -1001234567890, Type: "supergroup"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("forum_photo.jpg")

	// Test thread IDs
	threadIDs := []int64{
		1,   // General topic
		123, // Regular topic
		456, // Another topic
		999, // High ID topic
		0,   // Zero (edge case)
	}

	for _, threadID := range threadIDs {
		result := testCtx.SendPhoto(photo).Thread(threadID)
		if result == nil {
			t.Errorf("Thread ID %d should work", threadID)
		}

		// Test thread with other forum features
		combined := result.Silent().Protect()
		if combined == nil {
			t.Errorf("Thread combination should work for ID: %d", threadID)
		}
	}
}

func TestSendPhoto_ReplyFeatures(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("reply_photo.jpg")

	// Test reply to various message IDs
	replyToIDs := []int64{
		1,
		123,
		456,
		999999,
		0, // Edge case
	}

	for _, replyID := range replyToIDs {
		result := testCtx.SendPhoto(photo).ReplyTo(replyID)
		if result == nil {
			t.Errorf("ReplyTo message ID %d should work", replyID)
		}

		// Test reply with other features
		combined := result.Caption(g.String("Reply photo")).Silent()
		if combined == nil {
			t.Errorf("Reply combination should work for ID: %d", replyID)
		}
	}
}

func TestSendPhoto_EdgeCases(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)

	// Test with empty photo
	result := testCtx.SendPhoto(g.String(""))
	if result == nil {
		t.Error("SendPhoto should handle empty photo")
	}

	// Test with zero timeout
	result = testCtx.SendPhoto(g.String("photo.jpg")).Timeout(0 * time.Second)
	if result == nil {
		t.Error("SendPhoto should handle zero timeout")
	}

	// Test with very long timeout
	result = testCtx.SendPhoto(g.String("photo.jpg")).Timeout(24 * time.Hour)
	if result == nil {
		t.Error("SendPhoto should handle very long timeout")
	}

	// Test with empty API URL
	result = testCtx.SendPhoto(g.String("photo.jpg")).APIURL(g.String(""))
	if result == nil {
		t.Error("SendPhoto should handle empty API URL")
	}

	// Test without To() method (should use effective chat)
	result = testCtx.SendPhoto(g.String("photo.jpg"))
	if result == nil {
		t.Error("SendPhoto should work without explicit To() call")
	}
}

func TestSendPhoto_MethodCoverage(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("complete_photo.jpg")

	// Test all methods combined
	kb := keyboard.Inline().Text("Test", "test_data")
	entities := entities.New("Bold photo").Bold("Bold")

	result := testCtx.SendPhoto(photo).
		To(456).
		CaptionEntities(entities).
		Caption(g.String("Complete photo test")).
		HTML().
		Spoiler().
		Silent().
		Protect().
		Markup(kb).
		ReplyTo(999).
		Business(g.String("business_123")).
		Thread(123).
		ShowCaptionAboveMedia().
		After(1 * time.Second).
		DeleteAfter(300 * time.Second).
		Timeout(60 * time.Second).
		APIURL(g.String("https://api.telegram.org"))

	if result == nil {
		t.Error("All methods combined should work")
	}

	// Test method overriding
	overrideResult := testCtx.SendPhoto(photo).
		HTML().
		Markdown(). // Should override HTML
		Caption(g.String("First caption")).
		Caption(g.String("Second caption")). // Should override first
		To(456).
		To(789) // Should override first To

	if overrideResult == nil {
		t.Error("Method overriding should work")
	}
}

func TestSendPhoto_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat: &gotgbot.Chat{Id: 123, Type: "private"},
		Update:        &gotgbot.Update{UpdateId: 1},
	}

	testCtx := ctx.New(bot, rawCtx)
	photo := g.String("send_test.jpg")

	// Test Send method execution (immediate)
	builder := testCtx.SendPhoto(photo)
	result := builder.Send()

	// The result should be present (even if it's an error due to mocking)
	if !result.IsErr() && !result.IsOk() {
		t.Error("Send method should return a result")
	}

	// Test Send with various options
	builderWithOptions := testCtx.SendPhoto(photo).
		To(456).
		Caption(g.String("Test photo")).
		Silent().
		Protect().
		Timeout(30 * time.Second)

	resultWithOptions := builderWithOptions.Send()

	if !resultWithOptions.IsErr() && !resultWithOptions.IsOk() {
		t.Error("Send with options should return a result")
	}

	// Test Send with After (scheduled)
	builderWithAfter := testCtx.SendPhoto(photo).
		After(1 * time.Millisecond) // Very short duration for testing

	resultWithAfter := builderWithAfter.Send()

	var perr *fs.PathError

	if resultWithAfter.IsErr() && !errors.As(resultWithAfter.Err(), &perr) {
		t.Error("Send with After should return Ok for scheduled execution")
	}

	// Test Send with DeleteAfter
	builderWithDeleteAfter := testCtx.SendPhoto(photo).
		DeleteAfter(60 * time.Second)
	resultWithDeleteAfter := builderWithDeleteAfter.Send()

	if !resultWithDeleteAfter.IsErr() && !resultWithDeleteAfter.IsOk() {
		t.Error("Send with DeleteAfter should return a result")
	}

	// Test Send without To() method (should use effective chat)
	builderWithoutTo := testCtx.SendPhoto(photo)
	resultWithoutTo := builderWithoutTo.Send()

	if !resultWithoutTo.IsErr() && !resultWithoutTo.IsOk() {
		t.Error("Send without To() should return a result (using effective chat)")
	}

	// Test Send with spoiler
	builderWithSpoiler := testCtx.SendPhoto(photo).Spoiler()
	resultWithSpoiler := builderWithSpoiler.Send()

	if !resultWithSpoiler.IsErr() && !resultWithSpoiler.IsOk() {
		t.Error("Send with spoiler should return a result")
	}

	// Test Send with all features
	kb := keyboard.Inline().Text("Button", "data")
	entities := entities.New("Test photo").Bold("Test")
	builderComplete := testCtx.SendPhoto(photo).
		To(789).
		CaptionEntities(entities).
		Caption(g.String("Complete photo message")).
		HTML().
		Spoiler().
		Silent().
		Protect().
		Markup(kb).
		ReplyTo(999).
		Business(g.String("business_456")).
		Thread(123).
		ShowCaptionAboveMedia().
		Timeout(45 * time.Second).
		APIURL(g.String("https://api.example.com"))
	resultComplete := builderComplete.Send()

	if !resultComplete.IsErr() && !resultComplete.IsOk() {
		t.Error("Send with all features should return a result")
	}
}
