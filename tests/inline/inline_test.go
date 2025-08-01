package inline_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/inline"
	"github.com/enetx/tg/keyboard"
)

// Mock message content for testing
type mockMessageContent struct {
	text string
}

func (m *mockMessageContent) Build() gotgbot.InputMessageContent {
	return gotgbot.InputTextMessageContent{
		MessageText: m.text,
	}
}

// func TestArticle_Creation(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
//
// 	article := inline.Article(id, title, content)
//
// 	if article == nil {
// 		t.Error("Expected Article to be created")
// 	}
//
// 	if article.inline.Id != id.Std() {
// 		t.Errorf("Expected ID '%s', got '%s'", id.Std(), article.inline.Id)
// 	}
//
// 	if article.inline.Title != title.Std() {
// 		t.Errorf("Expected title '%s', got '%s'", title.Std(), article.inline.Title)
// 	}
//
// 	// Verify message content is set
// 	if article.inline.InputMessageContent == nil {
// 		t.Error("Expected InputMessageContent to be set")
// 	}
// }
//
// func TestInlineArticle_URL(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
// 	url := g.String("https://example.com")
//
// 	article := inline.Article(id, title, content)
// 	result := article.URL(url)
//
// 	if result.inline.Url != url.Std() {
// 		t.Errorf("Expected URL '%s', got '%s'", url.Std(), result.inline.Url)
// 	}
//
// 	if result != article {
// 		t.Error("Expected method chaining to work")
// 	}
// }
//
// func TestInlineArticle_Description(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
// 	desc := g.String("Test description")
//
// 	article := inline.Article(id, title, content)
// 	result := article.Description(desc)
//
// 	if result.inline.Description != desc.Std() {
// 		t.Errorf("Expected description '%s', got '%s'", desc.Std(), result.inline.Description)
// 	}
//
// 	if result != article {
// 		t.Error("Expected method chaining to work")
// 	}
// }
//
// func TestInlineArticle_ThumbnailURL(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
// 	thumbURL := g.String("https://example.com/thumb.jpg")
//
// 	article := inline.Article(id, title, content)
// 	result := article.ThumbnailURL(thumbURL)
//
// 	if result.inline.ThumbnailUrl != thumbURL.Std() {
// 		t.Errorf("Expected thumbnail URL '%s', got '%s'", thumbURL.Std(), result.inline.ThumbnailUrl)
// 	}
//
// 	if result != article {
// 		t.Error("Expected method chaining to work")
// 	}
// }
//
// func TestInlineArticle_ThumbnailSize(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
// 	width := int64(150)
// 	height := int64(100)
//
// 	article := inline.Article(id, title, content)
// 	result := article.ThumbnailSize(width, height)
//
// 	if result.inline.ThumbnailWidth != width {
// 		t.Errorf("Expected thumbnail width %d, got %d", width, result.inline.ThumbnailWidth)
// 	}
//
// 	if result.inline.ThumbnailHeight != height {
// 		t.Errorf("Expected thumbnail height %d, got %d", height, result.inline.ThumbnailHeight)
// 	}
//
// 	if result != article {
// 		t.Error("Expected method chaining to work")
// 	}
// }
//
// func TestInlineArticle_Markup(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
//
// 	kb := keyboard.Inline().
// 		Row().
// 		Text("Button", "callback_data")
//
// 	article := inline.Article(id, title, content)
// 	result := article.Markup(kb)
//
// 	if result.inline.ReplyMarkup == nil {
// 		t.Error("Expected reply markup to be set")
// 	}
//
// 	if result != article {
// 		t.Error("Expected method chaining to work")
// 	}
// }

func TestInlineArticle_Build(t *testing.T) {
	id := g.String("test_id")
	title := g.String("Test Title")
	content := &mockMessageContent{text: "Test content"}

	article := inline.NewArticle(id, title, content).
		URL(g.String("https://example.com")).
		Description(g.String("Test description")).
		ThumbnailURL(g.String("https://example.com/thumb.jpg")).
		ThumbnailSize(150, 100)

	built := article.Build()

	switch result := built.(type) {
	case gotgbot.InlineQueryResultArticle:
		if result.Id != id.Std() {
			t.Errorf("Expected ID '%s', got '%s'", id.Std(), result.Id)
		}
		if result.Title != title.Std() {
			t.Errorf("Expected title '%s', got '%s'", title.Std(), result.Title)
		}
		if result.Url != "https://example.com" {
			t.Errorf("Expected URL 'https://example.com', got '%s'", result.Url)
		}
		if result.Description != "Test description" {
			t.Errorf("Expected description 'Test description', got '%s'", result.Description)
		}
		if result.ThumbnailUrl != "https://example.com/thumb.jpg" {
			t.Errorf("Expected thumbnail URL 'https://example.com/thumb.jpg', got '%s'", result.ThumbnailUrl)
		}
		if result.ThumbnailWidth != 150 {
			t.Errorf("Expected thumbnail width 150, got %d", result.ThumbnailWidth)
		}
		if result.ThumbnailHeight != 100 {
			t.Errorf("Expected thumbnail height 100, got %d", result.ThumbnailHeight)
		}
	default:
		t.Errorf("Expected Inlineinline.QueryResultArticle, got %T", built)
	}
}

// func TestInlineArticle_ChainedMethods(t *testing.T) {
// 	id := g.String("chained_id")
// 	title := g.String("Chained Title")
// 	content := &mockMessageContent{text: "Chained content"}
//
// 	result := inline.Article(id, title, content).
// 		URL(g.String("https://chained.com")).
// 		Description(g.String("Chained description")).
// 		ThumbnailURL(g.String("https://chained.com/thumb.jpg")).
// 		ThumbnailSize(200, 150)
//
// 	if result.inline.Id != "chained_id" {
// 		t.Error("Expected chained ID to be set")
// 	}
//
// 	if result.inline.Title != "Chained Title" {
// 		t.Error("Expected chained title to be set")
// 	}
//
// 	if result.inline.Url != "https://chained.com" {
// 		t.Error("Expected chained URL to be set")
// 	}
//
// 	if result.inline.Description != "Chained description" {
// 		t.Error("Expected chained description to be set")
// 	}
//
// 	if result.inline.ThumbnailUrl != "https://chained.com/thumb.jpg" {
// 		t.Error("Expected chained thumbnail URL to be set")
// 	}
//
// 	if result.inline.ThumbnailWidth != 200 {
// 		t.Error("Expected chained thumbnail width to be set")
// 	}
//
// 	if result.inline.ThumbnailHeight != 150 {
// 		t.Error("Expected chained thumbnail height to be set")
// 	}
// }

// Test that InlineArticle implements the inline.QueryResult interface
func TestInlineArticle_ImplementsQueryResultInterface(t *testing.T) {
	id := g.String("test_id")
	title := g.String("Test Title")
	content := &mockMessageContent{text: "Test content"}

	var queryResult inline.QueryResult = inline.NewArticle(id, title, content)

	built := queryResult.Build()
	if built == nil {
		t.Error("Expected Build() to return a value")
	}

	// Verify it's actually an Inlineinline.QueryResultArticle
	if _, ok := built.(gotgbot.InlineQueryResultArticle); !ok {
		t.Errorf("Expected Inlineinline.QueryResultArticle, got %T", built)
	}
}

// func TestInlineArticle_DefaultValues(t *testing.T) {
// 	id := g.String("test_id")
// 	title := g.String("Test Title")
// 	content := &mockMessageContent{text: "Test content"}
//
// 	article := inline.Article(id, title, content)
//
// 	// Test default values
// 	if article.inline.Url != "" {
// 		t.Error("Expected default URL to be empty")
// 	}
//
// 	if article.inline.Description != "" {
// 		t.Error("Expected default description to be empty")
// 	}
//
// 	if article.inline.ThumbnailUrl != "" {
// 		t.Error("Expected default thumbnail URL to be empty")
// 	}
//
// 	if article.inline.ThumbnailWidth != 0 {
// 		t.Error("Expected default thumbnail width to be 0")
// 	}
//
// 	if article.inline.ThumbnailHeight != 0 {
// 		t.Error("Expected default thumbnail height to be 0")
// 	}
//
// 	if article.inline.ReplyMarkup != nil {
// 		t.Error("Expected default reply markup to be nil")
// 	}
// }

// func TestInlineArticle_EmptyStrings(t *testing.T) {
// 	id := g.String("")
// 	title := g.String("")
// 	content := &mockMessageContent{text: ""}
//
// 	article := inline.Article(id, title, content)
//
// 	if article.inline.Id != "" {
// 		t.Error("Expected empty ID to be preserved")
// 	}
//
// 	if article.inline.Title != "" {
// 		t.Error("Expected empty title to be preserved")
// 	}
// }

// Test Photo inline result
func TestPhoto_Creation(t *testing.T) {
	id := g.String("photo_id")
	photoURL := g.String("https://example.com/photo.jpg")
	thumbURL := g.String("https://example.com/thumb.jpg")

	photo := inline.NewPhoto(id, photoURL, thumbURL)

	if photo == nil {
		t.Error("Expected Photo to be created")
	}

	built := photo.Build()
	if result, ok := built.(gotgbot.InlineQueryResultPhoto); ok {
		if result.Id != id.Std() {
			t.Errorf("Expected ID '%s', got '%s'", id.Std(), result.Id)
		}
		if result.PhotoUrl != photoURL.Std() {
			t.Errorf("Expected photo URL '%s', got '%s'", photoURL.Std(), result.PhotoUrl)
		}
		if result.ThumbnailUrl != thumbURL.Std() {
			t.Errorf("Expected thumbnail URL '%s', got '%s'", thumbURL.Std(), result.ThumbnailUrl)
		}
	} else {
		t.Errorf("Expected Inlineinline.QueryResultPhoto, got %T", built)
	}
}

// Test basic inline result creation with New constructors
func TestInlineResults_BasicCreation(t *testing.T) {
	// Test NewPhoto
	photo := inline.NewPhoto(
		g.String("photo_id"),
		g.String("https://example.com/photo.jpg"),
		g.String("https://example.com/thumb.jpg"),
	)
	if photo == nil {
		t.Error("Expected NewPhoto to create a Photo")
	}

	photoResult := photo.Build()
	if _, ok := photoResult.(gotgbot.InlineQueryResultPhoto); !ok {
		t.Errorf("Expected Inlineinline.QueryResultPhoto, got %T", photoResult)
	}

	// Test NewVideo
	video := inline.NewVideo(
		g.String("video_id"),
		g.String("https://example.com/video.mp4"),
		g.String("video/mp4"),
		g.String("https://example.com/thumb.jpg"),
		g.String("Video Title"),
	)

	if video == nil {
		t.Error("Expected NewVideo to create a Video")
	}

	videoResult := video.Build()
	if _, ok := videoResult.(gotgbot.InlineQueryResultVideo); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVideo, got %T", videoResult)
	}

	// Test NewAudio
	audio := inline.NewAudio(g.String("audio_id"), g.String("https://example.com/audio.mp3"), g.String("Audio Title"))
	if audio == nil {
		t.Error("Expected NewAudio to create an Audio")
	}

	audioResult := audio.Build()
	if _, ok := audioResult.(gotgbot.InlineQueryResultAudio); !ok {
		t.Errorf("Expected Inlineinline.QueryResultAudio, got %T", audioResult)
	}

	// Test NewVoice
	voice := inline.NewVoice(g.String("voice_id"), g.String("https://example.com/voice.ogg"), g.String("Voice Title"))
	if voice == nil {
		t.Error("Expected NewVoice to create a Voice")
	}

	voiceResult := voice.Build()
	if _, ok := voiceResult.(gotgbot.InlineQueryResultVoice); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVoice, got %T", voiceResult)
	}

	// Test NewDocument
	doc := inline.NewDocument(
		g.String("doc_id"),
		g.String("Document Title"),
		g.String("https://example.com/doc.pdf"),
		g.String("application/pdf"),
	)
	if doc == nil {
		t.Error("Expected NewDocument to create a Document")
	}

	docResult := doc.Build()
	if _, ok := docResult.(gotgbot.InlineQueryResultDocument); !ok {
		t.Errorf("Expected Inlineinline.QueryResultDocument, got %T", docResult)
	}

	// Test NewLocation
	location := inline.NewLocation(g.String("location_id"), 40.7128, -74.0060, g.String("Location Title"))
	if location == nil {
		t.Error("Expected NewLocation to create a Location")
	}

	locationResult := location.Build()
	if _, ok := locationResult.(gotgbot.InlineQueryResultLocation); !ok {
		t.Errorf("Expected Inlineinline.QueryResultLocation, got %T", locationResult)
	}

	// Test NewVenue
	venue := inline.NewVenue(g.String("venue_id"), 40.7128, -74.0060, g.String("Venue Title"), g.String("123 Main St"))
	if venue == nil {
		t.Error("Expected NewVenue to create a Venue")
	}

	venueResult := venue.Build()
	if _, ok := venueResult.(gotgbot.InlineQueryResultVenue); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVenue, got %T", venueResult)
	}

	// Test NewContact
	contact := inline.NewContact(g.String("contact_id"), g.String("+1234567890"), g.String("John"))
	if contact == nil {
		t.Error("Expected NewContact to create a Contact")
	}

	contactResult := contact.Build()
	if _, ok := contactResult.(gotgbot.InlineQueryResultContact); !ok {
		t.Errorf("Expected Inlineinline.QueryResultContact, got %T", contactResult)
	}

	// Test NewGame
	game := inline.NewGame(g.String("game_id"), g.String("my_game"))
	if game == nil {
		t.Error("Expected NewGame to create a Game")
	}

	gameResult := game.Build()
	if _, ok := gameResult.(gotgbot.InlineQueryResultGame); !ok {
		t.Errorf("Expected Inlineinline.QueryResultGame, got %T", gameResult)
	}
}

// Test inline.QueryResult interface compliance
func TestInlineResults_inlineQueryResultInterface(t *testing.T) {
	results := []inline.QueryResult{
		inline.NewPhoto(
			g.String("photo_id"),
			g.String("https://example.com/photo.jpg"),
			g.String("https://example.com/thumb.jpg"),
		),
		inline.NewArticle(
			g.String("article_id"),
			g.String("Article Title"),
			&mockMessageContent{text: "Article content"},
		),
		inline.NewGame(g.String("game_id"), g.String("my_game")),
	}

	for i, result := range results {
		if result == nil {
			t.Errorf("Result %d is nil", i)
			continue
		}

		built := result.Build()
		if built == nil {
			t.Errorf("Result %d Build() returned nil", i)
		}
	}
}

// ==============================================
// Audio Tests - All untested methods (0% coverage)
// ==============================================

// func TestAudio_AllMethods(t *testing.T) {
// 	id := g.String("audio_id")
// 	audioURL := g.String("https://example.com/audio.mp3")
// 	title := g.String("Audio Title")
//
// 	audio := inline.NewAudio(id, audioURL, title)
//
// 	// Test Caption method
// 	result := audio.Caption(g.String("Audio caption"))
// 	if result != audio {
// 		t.Error("Expected fluent interface for Caption")
// 	}
// 	if audio.inline.Caption != "Audio caption" {
// 		t.Error("Expected caption to be set")
// 	}
//
// 	// Test HTML method
// 	result = audio.HTML()
// 	if result != audio {
// 		t.Error("Expected fluent interface for HTML")
// 	}
// 	if audio.inline.ParseMode != "HTML" {
// 		t.Error("Expected parse mode to be HTML")
// 	}
//
// 	// Test Markdown method
// 	result = audio.Markdown()
// 	if result != audio {
// 		t.Error("Expected fluent interface for Markdown")
// 	}
// 	if audio.inline.ParseMode != "MarkdownV2" {
// 		t.Error("Expected parse mode to be MarkdownV2")
// 	}
//
// 	// Test CaptionEntities method
// 	ent := entities.New("test text").Bold("test")
// 	result = audio.CaptionEntities(*ent)
// 	if result != audio {
// 		t.Error("Expected fluent interface for CaptionEntities")
// 	}
//
// 	// Test Performer method
// 	result = audio.Performer(g.String("Artist Name"))
// 	if result != audio {
// 		t.Error("Expected fluent interface for Performer")
// 	}
// 	if audio.inline.Performer != "Artist Name" {
// 		t.Error("Expected performer to be set")
// 	}
//
// 	// Test Duration method
// 	result = audio.Duration(time.Minute * 3)
// 	if result != audio {
// 		t.Error("Expected fluent interface for Duration")
// 	}
// 	if audio.inline.AudioDuration != 180 {
// 		t.Error("Expected duration to be 180 seconds")
// 	}
//
// 	// Test Markup method
// 	kb := keyboard.Inline().Row().Text("Button", "callback")
// 	result = audio.Markup(kb)
// 	if result != audio {
// 		t.Error("Expected fluent interface for Markup")
// 	}
// 	if audio.inline.ReplyMarkup == nil {
// 		t.Error("Expected reply markup to be set")
// 	}
//
// 	// Test InputMessageContent method
// 	content := &mockMessageContent{text: "Test message"}
// 	result = audio.InputMessageContent(content)
// 	if result != audio {
// 		t.Error("Expected fluent interface for Content")
// 	}
// 	if audio.inline.InputMessageContent == nil {
// 		t.Error("Expected input message content to be set")
// 	}
//
// 	// Test Build method
// 	built := audio.Build()
// 	if _, ok := built.(gotgbot.InlineQueryResultAudio); !ok {
// 		t.Errorf("Expected Inlineinline.QueryResultAudio, got %T", built)
// 	}
// }
//
// ==============================================
// Photo Tests - All untested methods
// ==============================================

func TestPhoto_AllMethods(t *testing.T) {
	id := g.String("photo_id")
	photoURL := g.String("https://example.com/photo.jpg")
	thumbURL := g.String("https://example.com/thumb.jpg")

	photo := inline.NewPhoto(id, photoURL, thumbURL)

	// Test Caption method
	result := photo.Caption(g.String("Photo caption"))
	if result != photo {
		t.Error("Expected fluent interface for Caption")
	}

	// Test HTML method
	result = photo.HTML()
	if result != photo {
		t.Error("Expected fluent interface for HTML")
	}

	// Test Markdown method
	result = photo.Markdown()
	if result != photo {
		t.Error("Expected fluent interface for Markdown")
	}

	// Test CaptionEntities method
	ent := entities.New("test text").Bold("test")
	result = photo.CaptionEntities(*ent)
	if result != photo {
		t.Error("Expected fluent interface for CaptionEntities")
	}

	// Test Size method
	result = photo.Size(800, 600)
	if result != photo {
		t.Error("Expected fluent interface for PhotoSize")
	}

	// Test Title method
	result = photo.Title(g.String("Photo Title"))
	if result != photo {
		t.Error("Expected fluent interface for Title")
	}

	// Test Description method
	result = photo.Description(g.String("Photo description"))
	if result != photo {
		t.Error("Expected fluent interface for Description")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = photo.Markup(kb)
	if result != photo {
		t.Error("Expected fluent interface for Markup")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = photo.InputMessageContent(content)
	if result != photo {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := photo.Build()
	if _, ok := built.(gotgbot.InlineQueryResultPhoto); !ok {
		t.Errorf("Expected Inlineinline.QueryResultPhoto, got %T", built)
	}
}

// ==============================================
// Video Tests - All untested methods
// ==============================================

func TestVideo_AllMethods(t *testing.T) {
	id := g.String("video_id")
	videoURL := g.String("https://example.com/video.mp4")
	mimeType := g.String("video/mp4")
	thumbURL := g.String("https://example.com/thumb.jpg")
	title := g.String("Video Title")

	video := inline.NewVideo(id, videoURL, mimeType, thumbURL, title)

	// Test Caption method
	result := video.Caption(g.String("Video caption"))
	if result != video {
		t.Error("Expected fluent interface for Caption")
	}

	// Test HTML method
	result = video.HTML()
	if result != video {
		t.Error("Expected fluent interface for HTML")
	}

	// Test Markdown method
	result = video.Markdown()
	if result != video {
		t.Error("Expected fluent interface for Markdown")
	}

	// Test CaptionEntities method
	ent := entities.New("test text").Bold("test")
	result = video.CaptionEntities(*ent)
	if result != video {
		t.Error("Expected fluent interface for CaptionEntities")
	}

	// Test Description method
	result = video.Description(g.String("Video description"))
	if result != video {
		t.Error("Expected fluent interface for Description")
	}

	// Test Size method
	result = video.Size(1920, 1080)
	if result != video {
		t.Error("Expected fluent interface for Size")
	}

	// Test Duration method
	result = video.Duration(time.Minute * 5)
	if result != video {
		t.Error("Expected fluent interface for Duration")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = video.Markup(kb)
	if result != video {
		t.Error("Expected fluent interface for Markup")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = video.InputMessageContent(content)
	if result != video {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := video.Build()
	if _, ok := built.(gotgbot.InlineQueryResultVideo); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVideo, got %T", built)
	}
}

// ==============================================
// Voice Tests - All untested methods
// ==============================================

func TestVoice_AllMethods(t *testing.T) {
	id := g.String("voice_id")
	voiceURL := g.String("https://example.com/voice.ogg")
	title := g.String("Voice Title")

	voice := inline.NewVoice(id, voiceURL, title)

	// Test Caption method
	result := voice.Caption(g.String("Voice caption"))
	if result != voice {
		t.Error("Expected fluent interface for Caption")
	}

	// Test HTML method
	result = voice.HTML()
	if result != voice {
		t.Error("Expected fluent interface for HTML")
	}

	// Test Markdown method
	result = voice.Markdown()
	if result != voice {
		t.Error("Expected fluent interface for Markdown")
	}

	// Test CaptionEntities method
	ent := entities.New("test text").Bold("test")
	result = voice.CaptionEntities(*ent)
	if result != voice {
		t.Error("Expected fluent interface for CaptionEntities")
	}

	// Test Duration method
	result = voice.Duration(time.Second * 30)
	if result != voice {
		t.Error("Expected fluent interface for Duration")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = voice.Markup(kb)
	if result != voice {
		t.Error("Expected fluent interface for Markup")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = voice.InputMessageContent(content)
	if result != voice {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := voice.Build()
	if _, ok := built.(gotgbot.InlineQueryResultVoice); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVoice, got %T", built)
	}
}

// ==============================================
// Document Tests - All untested methods
// ==============================================

func TestDocument_AllMethods(t *testing.T) {
	id := g.String("doc_id")
	title := g.String("Document Title")
	docURL := g.String("https://example.com/doc.pdf")
	mimeType := g.String("application/pdf")

	doc := inline.NewDocument(id, title, docURL, mimeType)

	// Test Caption method
	result := doc.Caption(g.String("Document caption"))
	if result != doc {
		t.Error("Expected fluent interface for Caption")
	}

	// Test HTML method
	result = doc.HTML()
	if result != doc {
		t.Error("Expected fluent interface for HTML")
	}

	// Test Markdown method
	result = doc.Markdown()
	if result != doc {
		t.Error("Expected fluent interface for Markdown")
	}

	// Test CaptionEntities method
	ent := entities.New("test text").Bold("test")
	result = doc.CaptionEntities(*ent)
	if result != doc {
		t.Error("Expected fluent interface for CaptionEntities")
	}

	// Test Description method
	result = doc.Description(g.String("Document description"))
	if result != doc {
		t.Error("Expected fluent interface for Description")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = doc.Markup(kb)
	if result != doc {
		t.Error("Expected fluent interface for Markup")
	}

	// Test ThumbnailURL method
	result = doc.ThumbnailURL(g.String("https://example.com/thumb.jpg"))
	if result != doc {
		t.Error("Expected fluent interface for ThumbnailURL")
	}

	// Test ThumbnailSize method
	result = doc.ThumbnailSize(150, 100)
	if result != doc {
		t.Error("Expected fluent interface for ThumbnailSize")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = doc.InputMessageContent(content)
	if result != doc {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := doc.Build()
	if _, ok := built.(gotgbot.InlineQueryResultDocument); !ok {
		t.Errorf("Expected Inlineinline.QueryResultDocument, got %T", built)
	}
}

// ==============================================
// Location Tests - All untested methods
// ==============================================

func TestLocation_AllMethods(t *testing.T) {
	id := g.String("location_id")
	lat := 40.7128
	lon := -74.0060
	title := g.String("Location Title")

	location := inline.NewLocation(id, lat, lon, title)

	// Test HorizontalAccuracy method
	result := location.HorizontalAccuracy(100)
	if result != location {
		t.Error("Expected fluent interface for HorizontalAccuracy")
	}

	// Test LiveFor method
	result = location.LiveFor(time.Hour)
	if result != location {
		t.Error("Expected fluent interface for LiveFor")
	}

	// Test Heading method
	result = location.Heading(90)
	if result != location {
		t.Error("Expected fluent interface for Heading")
	}

	// Skip ProximityAlert as it doesn't exist

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = location.Markup(kb)
	if result != location {
		t.Error("Expected fluent interface for Markup")
	}

	// Test ThumbnailURL method
	result = location.ThumbnailURL(g.String("https://example.com/thumb.jpg"))
	if result != location {
		t.Error("Expected fluent interface for ThumbnailURL")
	}

	// Test ThumbnailSize method
	result = location.ThumbnailSize(150, 100)
	if result != location {
		t.Error("Expected fluent interface for ThumbnailSize")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = location.InputMessageContent(content)
	if result != location {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := location.Build()
	if _, ok := built.(gotgbot.InlineQueryResultLocation); !ok {
		t.Errorf("Expected Inlineinline.QueryResultLocation, got %T", built)
	}
}

// ==============================================
// Venue Tests - All untested methods
// ==============================================

func TestVenue_AllMethods(t *testing.T) {
	id := g.String("venue_id")
	lat := 40.7128
	lon := -74.0060
	title := g.String("Venue Title")
	address := g.String("123 Main St")

	venue := inline.NewVenue(id, lat, lon, title, address)

	// Test FoursquareID method
	result := venue.FoursquareID(g.String("foursquare123"))
	if result != venue {
		t.Error("Expected fluent interface for FoursquareID")
	}

	// Test FoursquareType method
	result = venue.FoursquareType(g.String("restaurant"))
	if result != venue {
		t.Error("Expected fluent interface for FoursquareType")
	}

	// Test GooglePlaceID method
	result = venue.GooglePlaceID(g.String("google123"))
	if result != venue {
		t.Error("Expected fluent interface for GooglePlaceID")
	}

	// Test GooglePlaceType method
	result = venue.GooglePlaceType(g.String("establishment"))
	if result != venue {
		t.Error("Expected fluent interface for GooglePlaceType")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = venue.Markup(kb)
	if result != venue {
		t.Error("Expected fluent interface for Markup")
	}

	// Test ThumbnailURL method
	result = venue.ThumbnailURL(g.String("https://example.com/thumb.jpg"))
	if result != venue {
		t.Error("Expected fluent interface for ThumbnailURL")
	}

	// Test ThumbnailSize method
	result = venue.ThumbnailSize(150, 100)
	if result != venue {
		t.Error("Expected fluent interface for ThumbnailSize")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = venue.InputMessageContent(content)
	if result != venue {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := venue.Build()
	if _, ok := built.(gotgbot.InlineQueryResultVenue); !ok {
		t.Errorf("Expected Inlineinline.QueryResultVenue, got %T", built)
	}
}

// ==============================================
// Contact Tests - All untested methods
// ==============================================

func TestContact_AllMethods(t *testing.T) {
	id := g.String("contact_id")
	phone := g.String("+1234567890")
	firstName := g.String("John")

	contact := inline.NewContact(id, phone, firstName)

	// Test LastName method
	result := contact.LastName(g.String("Doe"))
	if result != contact {
		t.Error("Expected fluent interface for LastName")
	}

	// Test VCard method
	result = contact.VCard(g.String("BEGIN:VCARD\nVERSION:3.0\nEND:VCARD"))
	if result != contact {
		t.Error("Expected fluent interface for VCard")
	}

	// Test Markup method
	kb := keyboard.Inline().Row().Text("Button", "callback")
	result = contact.Markup(kb)
	if result != contact {
		t.Error("Expected fluent interface for Markup")
	}

	// Test ThumbnailURL method
	result = contact.ThumbnailURL(g.String("https://example.com/thumb.jpg"))
	if result != contact {
		t.Error("Expected fluent interface for ThumbnailURL")
	}

	// Test ThumbnailSize method
	result = contact.ThumbnailSize(150, 100)
	if result != contact {
		t.Error("Expected fluent interface for ThumbnailSize")
	}

	// Test InputMessageContent method
	content := &mockMessageContent{text: "Test message"}
	result = contact.InputMessageContent(content)
	if result != contact {
		t.Error("Expected fluent interface for Content")
	}

	// Test Build method
	built := contact.Build()
	if _, ok := built.(gotgbot.InlineQueryResultContact); !ok {
		t.Errorf("Expected Inlineinline.QueryResultContact, got %T", built)
	}
}

// ==============================================
// Game Tests - All untested methods
// ==============================================

// func TestGame_AllMethods(t *testing.T) {
// 	id := g.String("game_id")
// 	gameShortName := g.String("my_game")
//
// 	game := inline.NewGame(id, gameShortName)
//
// 	// Test Markup method
// 	kb := keyboard.Inline().Row().Text("Button", "callback")
// 	result := game.Markup(kb)
// 	if result != game {
// 		t.Error("Expected fluent interface for Markup")
// 	}
// 	if game.inline.ReplyMarkup == nil {
// 		t.Error("Expected reply markup to be set")
// 	}
//
// 	// Test Build method
// 	built := game.Build()
// 	if _, ok := built.(gotgbot.InlineQueryResultGame); !ok {
// 		t.Errorf("Expected Inlineinline.QueryResultGame, got %T", built)
// 	}
// }
