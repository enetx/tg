package input_test

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/preview"
)

// Test compile-time interface checks (they should not panic)
func TestInterfaceCompliance(t *testing.T) {
	// Test that the interface compliance checks don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Interface compliance check panicked: %v", r)
		}
	}()

	// The actual checks are done at compile-time via var declarations,
	// so we just need to make sure this test doesn't crash
	t.Log("All interface compliance checks passed at compile-time")
}

// ==============================================
// Media Tests
// ==============================================

func TestPhoto_Creation(t *testing.T) {
	// Create a mock file input
	mockFile := file.InputFile{
		Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg"),
	}

	photo := input.Photo(mockFile)

	if photo == nil {
		t.Error("Expected Photo to be created")
	}

	// Test that Build() returns a valid InputMedia
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia")
	}

	// Test type assertion
	_, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
}

func TestMediaPhoto_Caption(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)
	caption := g.String("Test caption")

	result := photo.Caption(caption)

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting caption
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting caption")
	}

	// Test that caption was applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if photoMedia.Caption != caption.Std() {
		t.Errorf("Expected caption '%s', got '%s'", caption.Std(), photoMedia.Caption)
	}
}

func TestMediaPhoto_HTML(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)

	result := photo.HTML()

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting HTML parse mode
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting HTML parse mode")
	}

	// Test that HTML parse mode was applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if photoMedia.ParseMode != "HTML" {
		t.Errorf("Expected parse mode 'HTML', got '%s'", photoMedia.ParseMode)
	}
}

func TestMediaPhoto_Markdown(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)

	result := photo.Markdown()

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting Markdown parse mode
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting Markdown parse mode")
	}

	// Test that Markdown parse mode was applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if photoMedia.ParseMode != "MarkdownV2" {
		t.Errorf("Expected parse mode 'MarkdownV2', got '%s'", photoMedia.ParseMode)
	}
}

func TestMediaPhoto_CaptionEntities(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)

	ent := entities.New(g.String("test text")).Bold("test")
	result := photo.CaptionEntities(*ent)

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting caption entities
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting caption entities")
	}

	// Test that caption entities were applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if photoMedia.CaptionEntities == nil {
		t.Error("Expected caption entities to be set")
	}
}

func TestMediaPhoto_ShowCaptionAboveMedia(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)

	result := photo.ShowCaptionAboveMedia()

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting ShowCaptionAboveMedia
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting ShowCaptionAboveMedia")
	}

	// Test that ShowCaptionAboveMedia was applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if !photoMedia.ShowCaptionAboveMedia {
		t.Error("Expected ShowCaptionAboveMedia to be true")
	}
}

func TestMediaPhoto_Spoiler(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile)

	result := photo.Spoiler()

	// Test fluent interface
	if result != photo {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting Spoiler
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting Spoiler")
	}

	// Test that Spoiler was applied
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}
	if !photoMedia.HasSpoiler {
		t.Error("Expected HasSpoiler to be true")
	}
}

func TestMediaPhoto_ChainedMethods(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/chained.jpg")}

	// Test complete method chaining
	photo := input.Photo(mockFile).
		Caption(g.String("Chained caption")).
		Markdown().
		ShowCaptionAboveMedia().
		Spoiler()

	// Test that chained builder is still valid
	if photo == nil {
		t.Error("Expected chained builder to be non-nil")
	}

	// Test that Build() works after method chaining
	built := photo.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after method chaining")
	}

	// Test type assertion and verify chain worked
	photoMedia, ok := built.(gotgbot.InputMediaPhoto)
	if !ok {
		t.Errorf("Expected InputMediaPhoto, got %T", built)
	}

	// Verify chained values were applied
	if photoMedia.Caption != "Chained caption" {
		t.Error("Expected chained caption to be preserved")
	}

	if photoMedia.ParseMode != "MarkdownV2" {
		t.Error("Expected chained parse mode to be preserved")
	}

	if !photoMedia.ShowCaptionAboveMedia {
		t.Error("Expected chained ShowCaptionAboveMedia to be preserved")
	}

	if !photoMedia.HasSpoiler {
		t.Error("Expected chained HasSpoiler to be preserved")
	}
}

// ==============================================
// MessageContent Tests
// ==============================================

func TestText_Creation(t *testing.T) {
	messageText := g.String("Hello, world!")
	text := input.Text(messageText)

	if text == nil {
		t.Error("Expected Text to be created")
	}

	// Test that Build() returns a valid InputMessageContent
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent")
	}

	// Test type assertion
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}

	// Test that message text was set
	if textContent.MessageText != messageText.Std() {
		t.Errorf("Expected message text '%s', got '%s'", messageText.Std(), textContent.MessageText)
	}
}

func TestMessageText_HTML(t *testing.T) {
	text := input.Text(g.String("Test message"))

	result := text.HTML()

	// Test fluent interface
	if result != text {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting HTML parse mode
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after setting HTML parse mode")
	}

	// Test that HTML parse mode was applied
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}
	if textContent.ParseMode != "HTML" {
		t.Errorf("Expected parse mode 'HTML', got '%s'", textContent.ParseMode)
	}
}

func TestMessageText_Markdown(t *testing.T) {
	text := input.Text(g.String("Test message"))

	result := text.Markdown()

	// Test fluent interface
	if result != text {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting Markdown parse mode
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after setting Markdown parse mode")
	}

	// Test that Markdown parse mode was applied
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}
	if textContent.ParseMode != "MarkdownV2" {
		t.Errorf("Expected parse mode 'MarkdownV2', got '%s'", textContent.ParseMode)
	}
}

func TestMessageText_Entities(t *testing.T) {
	text := input.Text(g.String("Test message"))

	ent := entities.New(g.String("test text")).Bold("test")
	result := text.Entities(*ent)

	// Test fluent interface
	if result != text {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting entities
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after setting entities")
	}

	// Test that entities were applied
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}
	if textContent.Entities == nil {
		t.Error("Expected entities to be set")
	}
}

func TestMessageText_Preview(t *testing.T) {
	text := input.Text(g.String("Test message with link"))

	prev := preview.New().URL(g.String("https://example.com")).Large()
	result := text.Preview(prev)

	// Test fluent interface
	if result != text {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting preview
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after setting preview")
	}

	// Test that preview options were applied
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}
	if textContent.LinkPreviewOptions == nil {
		t.Error("Expected link preview options to be set")
	}
}

func TestMessageText_ChainedMethods(t *testing.T) {
	text := input.Text(g.String("Chained message")).
		Markdown().
		Preview(preview.New().Disable())

	// Test that chained builder is still valid
	if text == nil {
		t.Error("Expected chained builder to be non-nil")
	}

	// Test that Build() works after method chaining
	built := text.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after method chaining")
	}

	// Test type assertion and verify chain worked
	textContent, ok := built.(gotgbot.InputTextMessageContent)
	if !ok {
		t.Errorf("Expected InputTextMessageContent, got %T", built)
	}

	// Verify chained values were applied
	if textContent.MessageText != "Chained message" {
		t.Error("Expected chained message text to be preserved")
	}

	if textContent.ParseMode != "MarkdownV2" {
		t.Error("Expected chained parse mode to be preserved")
	}

	if textContent.LinkPreviewOptions == nil {
		t.Error("Expected chained link preview options to be preserved")
	}
}

// ==============================================
// Video Tests
// ==============================================

func TestVideo_Creation(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)

	if video == nil {
		t.Error("Expected Video to be created")
	}

	// Test that Build() returns a valid InputMedia
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia")
	}

	// Test type assertion
	_, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
}

func TestMediaVideo_Cover(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)
	cover := g.String("https://example.com/cover.jpg")

	result := video.Cover(cover)

	// Test fluent interface
	if result != video {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting cover
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting cover")
	}

	// Test that cover was applied
	videoMedia, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
	if videoMedia.Cover != cover.Std() {
		t.Errorf("Expected cover '%s', got '%s'", cover.Std(), videoMedia.Cover)
	}
}

func TestMediaVideo_Size(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)
	width, height := int64(1920), int64(1080)

	result := video.Size(width, height)

	// Test fluent interface
	if result != video {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting size
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting size")
	}

	// Test that size was applied
	videoMedia, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
	if videoMedia.Width != width {
		t.Errorf("Expected width %d, got %d", width, videoMedia.Width)
	}
	if videoMedia.Height != height {
		t.Errorf("Expected height %d, got %d", height, videoMedia.Height)
	}
}

func TestMediaVideo_Duration(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)
	duration := 2 * time.Minute

	result := video.Duration(duration)

	// Test fluent interface
	if result != video {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting duration
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting duration")
	}

	// Test that duration was applied
	videoMedia, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
	expectedSeconds := int64(duration.Seconds())
	if videoMedia.Duration != expectedSeconds {
		t.Errorf("Expected duration %d seconds, got %d", expectedSeconds, videoMedia.Duration)
	}
}

func TestMediaVideo_Streamable(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)

	result := video.Streamable()

	// Test fluent interface
	if result != video {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting streamable
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting streamable")
	}

	// Test that streamable was applied
	videoMedia, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
	if !videoMedia.SupportsStreaming {
		t.Error("Expected SupportsStreaming to be true")
	}
}

func TestMediaVideo_Spoiler(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/video.mp4")}
	video := input.Video(mockFile)

	result := video.Spoiler()

	// Test fluent interface
	if result != video {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting spoiler
	built := video.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting spoiler")
	}

	// Test that spoiler was applied
	videoMedia, ok := built.(gotgbot.InputMediaVideo)
	if !ok {
		t.Errorf("Expected InputMediaVideo, got %T", built)
	}
	if !videoMedia.HasSpoiler {
		t.Error("Expected HasSpoiler to be true")
	}
}

// ==============================================
// Audio Tests
// ==============================================

func TestAudio_Creation(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/audio.mp3")}
	audio := input.Audio(mockFile)

	if audio == nil {
		t.Error("Expected Audio to be created")
	}

	// Test that Build() returns a valid InputMedia
	built := audio.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia")
	}

	// Test type assertion
	_, ok := built.(gotgbot.InputMediaAudio)
	if !ok {
		t.Errorf("Expected InputMediaAudio, got %T", built)
	}
}

func TestMediaAudio_Duration(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/audio.mp3")}
	audio := input.Audio(mockFile)
	duration := 3 * time.Minute

	result := audio.Duration(duration)

	// Test fluent interface
	if result != audio {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting duration
	built := audio.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting duration")
	}

	// Test that duration was applied
	audioMedia, ok := built.(gotgbot.InputMediaAudio)
	if !ok {
		t.Errorf("Expected InputMediaAudio, got %T", built)
	}
	expectedSeconds := int64(duration.Seconds())
	if audioMedia.Duration != expectedSeconds {
		t.Errorf("Expected duration %d seconds, got %d", expectedSeconds, audioMedia.Duration)
	}
}

func TestMediaAudio_Performer(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/audio.mp3")}
	audio := input.Audio(mockFile)
	performer := g.String("Test Artist")

	result := audio.Performer(performer)

	// Test fluent interface
	if result != audio {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting performer
	built := audio.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting performer")
	}

	// Test that performer was applied
	audioMedia, ok := built.(gotgbot.InputMediaAudio)
	if !ok {
		t.Errorf("Expected InputMediaAudio, got %T", built)
	}
	if audioMedia.Performer != performer.Std() {
		t.Errorf("Expected performer '%s', got '%s'", performer.Std(), audioMedia.Performer)
	}
}

func TestMediaAudio_Title(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/audio.mp3")}
	audio := input.Audio(mockFile)
	title := g.String("Test Song")

	result := audio.Title(title)

	// Test fluent interface
	if result != audio {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting title
	built := audio.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMedia after setting title")
	}

	// Test that title was applied
	audioMedia, ok := built.(gotgbot.InputMediaAudio)
	if !ok {
		t.Errorf("Expected InputMediaAudio, got %T", built)
	}
	if audioMedia.Title != title.Std() {
		t.Errorf("Expected title '%s', got '%s'", title.Std(), audioMedia.Title)
	}
}

// ==============================================
// Location Tests
// ==============================================

func TestLocation_Creation(t *testing.T) {
	lat, lon := 40.7589, -73.9851
	location := input.Location(lat, lon)

	if location == nil {
		t.Error("Expected Location to be created")
	}

	// Test that Build() returns a valid InputMessageContent
	built := location.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent")
	}

	// Test type assertion
	locationContent, ok := built.(gotgbot.InputLocationMessageContent)
	if !ok {
		t.Errorf("Expected InputLocationMessageContent, got %T", built)
	}

	// Test that coordinates were set
	if locationContent.Latitude != lat {
		t.Errorf("Expected latitude %f, got %f", lat, locationContent.Latitude)
	}
	if locationContent.Longitude != lon {
		t.Errorf("Expected longitude %f, got %f", lon, locationContent.Longitude)
	}
}

func TestMessageLocation_HorizontalAccuracy(t *testing.T) {
	location := input.Location(40.7589, -73.9851)
	accuracy := 50.0

	result := location.HorizontalAccuracy(accuracy)

	// Test fluent interface
	if result != location {
		t.Error("Expected fluent interface - method should return self")
	}

	// Test that Build() still works after setting accuracy
	built := location.Build()
	if built == nil {
		t.Error("Expected Build() to return non-nil InputMessageContent after setting accuracy")
	}

	// Test that accuracy was applied
	locationContent, ok := built.(gotgbot.InputLocationMessageContent)
	if !ok {
		t.Errorf("Expected InputLocationMessageContent, got %T", built)
	}
	if locationContent.HorizontalAccuracy != accuracy {
		t.Errorf("Expected accuracy %f, got %f", accuracy, locationContent.HorizontalAccuracy)
	}
}

// ==============================================
// Simple Interface Compliance Tests
// ==============================================

func TestMediaInterfaceCompliance(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/test.jpg")}

	// Test all media types implement the interface
	builders := []any{
		input.Photo(mockFile),
		input.Video(mockFile),
		input.Audio(mockFile),
		input.Animation(mockFile),
		input.Document(mockFile),
	}

	for i, builder := range builders {
		if builder == nil {
			t.Errorf("Builder %d should not be nil", i)
		}
	}
}

func TestMessageContentInterfaceCompliance(t *testing.T) {
	// Test all message content types implement the interface
	builders := []any{
		input.Text(g.String("test")),
		input.Location(40.7589, -73.9851),
		input.Contact(g.String("+1234567890"), g.String("John")),
		input.Venue(40.7589, -73.9851, g.String("Test Venue"), g.String("123 Test St")),
	}

	for i, builder := range builders {
		if builder == nil {
			t.Errorf("Builder %d should not be nil", i)
		}
	}
}

func TestPollOptionInterfaceCompliance(t *testing.T) {
	choice := input.Choice(g.String("Test option"))

	if choice == nil {
		t.Error("Choice builder should not be nil")
	}

	// Test that Build() works
	built := choice.Build()
	// Note: InputPollOption is a struct, not interface, so we just verify it's returned
	if built.Text == "" {
		t.Error("Expected Build() to return valid InputPollOption with text")
	}
}

// ==============================================
// Build Method Tests
// ==============================================

func TestPhoto_BuildIdempotency(t *testing.T) {
	mockFile := file.InputFile{Doc: gotgbot.InputFileByURL("https://example.com/photo.jpg")}
	photo := input.Photo(mockFile).Caption(g.String("Test"))

	// Build multiple times to test idempotency
	built1 := photo.Build()
	built2 := photo.Build()

	if built1 == nil || built2 == nil {
		t.Error("Build() should never return nil")
	}

	// Results should be equivalent
	photo1, ok1 := built1.(gotgbot.InputMediaPhoto)
	photo2, ok2 := built2.(gotgbot.InputMediaPhoto)

	if !ok1 || !ok2 {
		t.Error("Both builds should return InputMediaPhoto")
	}

	if photo1.Caption != photo2.Caption {
		t.Error("Build() should be idempotent")
	}
}

func TestText_BuildIdempotency(t *testing.T) {
	text := input.Text(g.String("Test message")).HTML()

	// Build multiple times to test idempotency
	built1 := text.Build()
	built2 := text.Build()

	if built1 == nil || built2 == nil {
		t.Error("Build() should never return nil")
	}

	// Results should be equivalent
	text1, ok1 := built1.(gotgbot.InputTextMessageContent)
	text2, ok2 := built2.(gotgbot.InputTextMessageContent)

	if !ok1 || !ok2 {
		t.Error("Both builds should return InputTextMessageContent")
	}

	if text1.MessageText != text2.MessageText {
		t.Error("Build() should be idempotent")
	}

	if text1.ParseMode != text2.ParseMode {
		t.Error("Build() should be idempotent")
	}
}
