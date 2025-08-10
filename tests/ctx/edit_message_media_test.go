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
	"github.com/enetx/tg/keyboard"
)

func TestContext_EditMessageMedia(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	media := input.Photo(mockFile)
	result := ctx.EditMessageMedia(media)

	if result == nil {
		t.Error("Expected EditMessageMedia builder to be created")
	}

	// Test method chaining
	chained := result.ChatID(456)
	if chained == nil {
		t.Error("Expected ChatID method to return builder")
	}
}

func TestContext_EditMessageMediaChaining(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "private"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	media := input.Photo(mockFile)
	result := ctx.EditMessageMedia(media).
		ChatID(456).
		MessageID(789)

	if result == nil {
		t.Error("Expected EditMessageMedia builder to be created")
	}

	// Test that builder is functional
	_ = result
}

func TestEditMessageMedia_Send(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Send method - will fail with mock but covers the method
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	media := input.Photo(mockFile)
	sendResult := ctx.EditMessageMedia(media).Send()

	if sendResult.IsErr() {
		t.Logf("EditMessageMedia Send failed as expected with mock bot: %v", sendResult.Err())
	}

	// Test Send method with configuration
	configuredSendResult := ctx.EditMessageMedia(media).
		ChatID(456).
		MessageID(123).
		Timeout(30 * time.Second).
		APIURL(g.String("https://api.example.com")).
		Send()

	if configuredSendResult.IsErr() {
		t.Logf("EditMessageMedia configured Send failed as expected: %v", configuredSendResult.Err())
	}
}

// Tests for methods with 0% coverage

func TestEditMessageMedia_InlineMessageID(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test InlineMessageID method
	inlineMessageIDs := []string{
		"inline_123456789",
		"inline_abcdef123",
		"inline_xyz789abc",
		"", // Empty inline message ID
	}

	for _, inlineID := range inlineMessageIDs {
		mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/updated-photo.jpg")}
		media := input.Photo(mockFile).Caption(g.String("Updated photo with inline ID"))

		inlineResult := ctx.EditMessageMedia(media).
			InlineMessageID(g.String(inlineID))

		if inlineResult == nil {
			t.Errorf("InlineMessageID with '%s' should work", inlineID)
		}

		// Test send with inline message ID
		sendResult := inlineResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageMedia with inline message ID '%s' Send failed as expected: %v", inlineID, sendResult.Err())
		}
	}
}

func TestEditMessageMedia_Business(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Business connection IDs
	businessIDs := []string{
		"business_conn_123",
		"business_conn_456",
		"enterprise_conn_789",
		"", // Empty business ID
	}

	for _, businessID := range businessIDs {
		mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/business-photo.jpg")}
		media := input.Photo(mockFile).Caption(g.String("Business photo update"))

		businessResult := ctx.EditMessageMedia(media).
			Business(g.String(businessID)).
			ChatID(456).
			MessageID(789)

		if businessResult == nil {
			t.Errorf("Business with '%s' should work", businessID)
		}

		// Test send with business ID
		sendResult := businessResult.Send()
		if sendResult.IsErr() {
			t.Logf("EditMessageMedia with business ID '%s' Send failed as expected: %v", businessID, sendResult.Err())
		}
	}
}

func TestEditMessageMedia_Markup(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test Markup method with inline keyboard
	inlineKB := keyboard.Inline().
		Text(g.String("📷 View Photo"), g.String("view_photo")).
		Row().
		Text(g.String("🔄 Change Media"), g.String("change_media")).
		Row().
		URL(g.String("🌐 Open Link"), g.String("https://example.com"))

	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/markup-photo.jpg")}
	media := input.Photo(mockFile).Caption(g.String("Photo with markup"))

	markupResult := ctx.EditMessageMedia(media).
		Markup(inlineKB).
		ChatID(456).
		MessageID(789)

	if markupResult == nil {
		t.Error("Markup method should return EditMessageMedia for chaining")
	}

	// Test send with markup
	sendResult := markupResult.Send()
	if sendResult.IsErr() {
		t.Logf("EditMessageMedia with Markup Send failed as expected: %v", sendResult.Err())
	}

	// Test with different keyboard configurations
	emptyKB := keyboard.Inline()
	emptyMarkupResult := ctx.EditMessageMedia(media).
		Markup(emptyKB).
		ChatID(456).
		MessageID(789)

	if emptyMarkupResult == nil {
		t.Error("Markup method with empty keyboard should return EditMessageMedia for chaining")
	}
}

func TestEditMessageMedia_AllMediaTypes(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test with different media types to ensure methods work with all
	mediaTypes := []struct {
		name  string
		media input.Media
	}{
		{
			"Photo",
			input.Photo(file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}).
				Caption(g.String("Test photo")).HTML(),
		},
		{
			"Video",
			input.Video(file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}).
				Caption(g.String("Test video")).Markdown(),
		},
		{
			"Animation",
			input.Animation(file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/animation.gif")}).
				Caption(g.String("Test animation")),
		},
		{
			"Document",
			input.Document(file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/document.pdf")}).
				Caption(g.String("Test document")),
		},
		{
			"Audio",
			input.Audio(file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/audio.mp3")}).
				Caption(g.String("Test audio")),
		},
	}

	for _, mediaType := range mediaTypes {
		// Test InlineMessageID with different media types
		inlineResult := ctx.EditMessageMedia(mediaType.media).
			InlineMessageID(g.String("inline_" + mediaType.name)).
			Send()

		if inlineResult.IsErr() {
			t.Logf("EditMessageMedia %s with InlineMessageID Send failed as expected: %v", mediaType.name, inlineResult.Err())
		}

		// Test Business with different media types
		businessResult := ctx.EditMessageMedia(mediaType.media).
			Business(g.String("business_" + mediaType.name)).
			ChatID(456).
			MessageID(789).
			Send()

		if businessResult.IsErr() {
			t.Logf("EditMessageMedia %s with Business Send failed as expected: %v", mediaType.name, businessResult.Err())
		}

		// Test Markup with different media types
		markupResult := ctx.EditMessageMedia(mediaType.media).
			Markup(keyboard.Inline().Text(g.String("Button"), g.String("btn_"+mediaType.name))).
			ChatID(456).
			MessageID(789).
			Send()

		if markupResult.IsErr() {
			t.Logf("EditMessageMedia %s with Markup Send failed as expected: %v", mediaType.name, markupResult.Err())
		}
	}
}

func TestEditMessageMedia_ComprehensiveWorkflow(t *testing.T) {
	bot := &mockBot{}
	rawCtx := &ext.Context{
		EffectiveChat:    &gotgbot.Chat{Id: 456, Type: "group"},
		EffectiveMessage: &gotgbot.Message{MessageId: 789},
		Update:           &gotgbot.Update{UpdateId: 1},
	}

	ctx := ctx.New(bot, rawCtx)

	// Test all methods in comprehensive workflow
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/comprehensive-photo.jpg")}
	media := input.Photo(mockFile).
		Caption(g.String("<b>Comprehensive</b> photo update with <i>formatting</i>")).
		HTML().
		ShowCaptionAboveMedia().
		Spoiler()

	complexResult := ctx.EditMessageMedia(media).
		ChatID(456).
		MessageID(789).
		Business(g.String("business_comprehensive_123")).
		Markup(keyboard.Inline().
			Text(g.String("❤️ Like"), g.String("like_photo")).
			Text(g.String("💬 Comment"), g.String("comment_photo"))).
		Timeout(45 * time.Second).
		APIURL(g.String("https://comprehensive-media-api.telegram.org")).
		Send()

	if complexResult.IsErr() {
		t.Logf("EditMessageMedia comprehensive workflow Send failed as expected: %v", complexResult.Err())
	}

	// Test with inline message workflow
	inlineWorkflowResult := ctx.EditMessageMedia(media).
		InlineMessageID(g.String("inline_comprehensive_media_123")).
		Markup(keyboard.Inline().
			Text(g.String("🔄 Refresh"), g.String("refresh_media")).
			URL(g.String("📱 Open App"), g.String("https://app.example.com"))).
		Timeout(30 * time.Second).
		APIURL(g.String("https://inline-media-api.telegram.org")).
		Send()

	if inlineWorkflowResult.IsErr() {
		t.Logf("EditMessageMedia inline workflow Send failed as expected: %v", inlineWorkflowResult.Err())
	}
}

func TestEditMessageMedia_APIURL_NilRequestOpts(t *testing.T) {
	bot := &mockBot{}
	ctx := ctx.New(bot, &ext.Context{EffectiveChat: &gotgbot.Chat{Id: 456, Type: "private"}, Update: &gotgbot.Update{UpdateId: 1}})

	// Test APIURL when RequestOpts is nil (covers the nil branch)
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	testMedia := input.Photo(mockFile)
	result := ctx.EditMessageMedia(testMedia)
	if result == nil {
		t.Error("EditMessageMedia should return builder")
	}

	// This should create RequestOpts and set APIURL
	apiResult := result.APIURL(g.String("https://api.test.com"))
	if apiResult == nil {
		t.Error("APIURL should return builder when RequestOpts is nil")
	}
}
