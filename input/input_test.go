package input

import (
	"testing"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/file"
	"github.com/enetx/tg/preview"
)

// Test constants for reuse across tests
const (
	testURL       = "https://example.com/test.jpg"
	testText      = "Test message text"
	testLatitude  = 40.7128
	testLongitude = -74.0060
	testPhone     = "+1234567890"
	testFirstName = "John"
	testLastName  = "Doe"
	testTitle     = "Test Title"
	testAddress   = "New York, NY"
)

// Helper function to create test file input
func createTestFile() file.InputFile {
	return file.Input(g.String(testURL)).Ok()
}

// Helper function to create test entities
func createTestEntities() entities.Entities {
	return *entities.New(g.String("Test text with bold")).Bold("bold")
}

// ==============================================
// Interface Compliance Tests
// ==============================================

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

func TestPhoto(t *testing.T) {
	file := createTestFile()
	photo := Photo(file)

	if photo == nil {
		t.Fatal("Photo() should not return nil")
	}

	// Test Build method
	built := photo.Build()
	if built == nil {
		t.Error("Photo.Build() should not return nil")
	}

	// Test method chaining with Markdown
	result := photo.Caption(g.String("test caption")).Markdown().Spoiler()
	if result != photo {
		t.Error("Photo methods should return the same instance for chaining")
	}
}

func TestVideo(t *testing.T) {
	file := createTestFile()
	video := Video(file)

	if video == nil {
		t.Fatal("Video() should not return nil")
	}

	// Test Build method
	built := video.Build()
	if built == nil {
		t.Error("Video.Build() should not return nil")
	}

	// Test all methods for proper chaining
	result := video.
		Cover(g.String("cover")).
		Thumbnail(file).
		Caption(g.String("caption")).
		HTML().
		Size(1920, 1080).
		Duration(time.Minute).
		StartAt(time.Second * 30).
		Streamable().
		Spoiler()

	if result != video {
		t.Error("Video methods should return the same instance for chaining")
	}
}

func TestAnimation(t *testing.T) {
	file := createTestFile()
	animation := Animation(file)

	if animation == nil {
		t.Fatal("Animation() should not return nil")
	}

	// Test Build method
	built := animation.Build()
	if built == nil {
		t.Error("Animation.Build() should not return nil")
	}

	// Test method chaining
	entities := createTestEntities()
	result := animation.
		Thumbnail(file).
		Caption(g.String("caption")).
		HTML().
		CaptionEntities(entities).
		ShowCaptionAboveMedia().
		Markdown().
		Size(640, 480).
		Duration(time.Second * 5).
		Spoiler()

	if result != animation {
		t.Error("Animation methods should return the same instance for chaining")
	}
}

func TestAudio(t *testing.T) {
	file := createTestFile()
	audio := Audio(file)

	if audio == nil {
		t.Fatal("Audio() should not return nil")
	}

	// Test Build method
	built := audio.Build()
	if built == nil {
		t.Error("Audio.Build() should not return nil")
	}

	// Test method chaining
	entities := createTestEntities()
	result := audio.
		Thumbnail(file).
		Caption(g.String("caption")).
		HTML().
		Markdown().
		CaptionEntities(entities).
		Duration(time.Minute * 3).
		Performer(g.String("Artist")).
		Title(g.String("Song Title"))

	if result != audio {
		t.Error("Audio methods should return the same instance for chaining")
	}
}

func TestDocument(t *testing.T) {
	file := createTestFile()
	document := Document(file)

	if document == nil {
		t.Fatal("Document() should not return nil")
	}

	// Test Build method
	built := document.Build()
	if built == nil {
		t.Error("Document.Build() should not return nil")
	}

	// Test method chaining
	entities := createTestEntities()
	result := document.
		Thumbnail(file).
		Caption(g.String("caption")).
		HTML().
		CaptionEntities(entities).
		Markdown().
		DisableContentTypeDetection()

	if result != document {
		t.Error("Document methods should return the same instance for chaining")
	}
}

// ==============================================
// Message Content Tests
// ==============================================

func TestText(t *testing.T) {
	text := Text(g.String(testText))

	if text == nil {
		t.Fatal("Text() should not return nil")
	}

	// Test Build method
	built := text.Build()
	if built == nil {
		t.Error("Text.Build() should not return nil")
	}

	// Test method chaining
	p := &preview.Preview{}
	entities := createTestEntities()

	result := text.
		HTML().
		Entities(entities).
		Preview(p)

	if result != text {
		t.Error("Text methods should return the same instance for chaining")
	}

	// Test Markdown method separately
	text2 := Text(g.String(testText))
	result2 := text2.Markdown()
	if result2 != text2 {
		t.Error("Text.Markdown() should return the same instance for chaining")
	}
}

func TestLocation(t *testing.T) {
	location := Location(testLatitude, testLongitude)

	if location == nil {
		t.Fatal("Location() should not return nil")
	}

	// Test Build method
	built := location.Build()
	if built == nil {
		t.Error("Location.Build() should not return nil")
	}

	// Test method chaining
	result := location.
		HorizontalAccuracy(10.5).
		LivePeriod(3600).
		Heading(90).
		ProximityAlertRadius(100)

	if result != location {
		t.Error("Location methods should return the same instance for chaining")
	}
}

func TestContact(t *testing.T) {
	contact := Contact(g.String(testPhone), g.String(testFirstName))

	if contact == nil {
		t.Fatal("Contact() should not return nil")
	}

	// Test Build method
	built := contact.Build()
	if built == nil {
		t.Error("Contact.Build() should not return nil")
	}

	// Test method chaining
	result := contact.
		LastName(g.String(testLastName)).
		Vcard(g.String("BEGIN:VCARD\nVERSION:3.0\nEND:VCARD"))

	if result != contact {
		t.Error("Contact methods should return the same instance for chaining")
	}
}

func TestVenue(t *testing.T) {
	venue := Venue(testLatitude, testLongitude, g.String(testTitle), g.String(testAddress))

	if venue == nil {
		t.Fatal("Venue() should not return nil")
	}

	// Test Build method
	built := venue.Build()
	if built == nil {
		t.Error("Venue.Build() should not return nil")
	}

	// Test method chaining
	result := venue.
		FoursquareID(g.String("4sq123")).
		FoursquareType(g.String("restaurant")).
		GooglePlaceID(g.String("ChIJ123")).
		GooglePlaceType(g.String("restaurant"))

	if result != venue {
		t.Error("Venue methods should return the same instance for chaining")
	}
}

func TestInvoice(t *testing.T) {
	invoice := Invoice(
		g.String("Test Product"),
		g.String("Test Description"),
		g.String("test_payload"),
		g.String("USD"),
	)

	if invoice == nil {
		t.Fatal("Invoice() should not return nil")
	}

	// Test Build method
	built := invoice.Build()
	if built == nil {
		t.Error("Invoice.Build() should not return nil")
	}

	// Test method chaining
	result := invoice.
		Price(g.String("Product"), 1000).
		ProviderToken(g.String("provider_token")).
		MaxTip(500).
		SuggestedTips(100, 200, 300).
		ProviderData(g.String(`{"key":"value"}`)).
		Photo(g.String(testURL), 1024, 512, 256).
		NeedName().
		NeedPhone().
		NeedEmail().
		NeedShipping().
		SendPhone().
		SendEmail().
		Flexible()

	if result != invoice {
		t.Error("Invoice methods should return the same instance for chaining")
	}
}

// ==============================================
// Paid Media Tests
// ==============================================

func TestPaidPhoto(t *testing.T) {
	file := createTestFile()
	paidPhoto := PaidPhoto(file)

	if paidPhoto == nil {
		t.Fatal("PaidPhoto() should not return nil")
	}

	// Test Build method
	built := paidPhoto.Build()
	if built == nil {
		t.Error("PaidPhoto.Build() should not return nil")
	}
}

func TestPaidVideo(t *testing.T) {
	file := createTestFile()
	paidVideo := PaidVideo(file)

	if paidVideo == nil {
		t.Fatal("PaidVideo() should not return nil")
	}

	// Test Build method
	built := paidVideo.Build()
	if built == nil {
		t.Error("PaidVideo.Build() should not return nil")
	}

	// Test method chaining
	result := paidVideo.
		Cover(g.String("cover")).
		Thumbnail(file).
		Width(1920).
		Height(1080).
		Duration(time.Minute).
		StartAt(time.Second * 30).
		Streamable()

	if result != paidVideo {
		t.Error("PaidVideo methods should return the same instance for chaining")
	}
}

// ==============================================
// Profile Photo Tests
// ==============================================

func TestStaticPhoto(t *testing.T) {
	staticPhoto := StaticPhoto(g.String(testURL))

	if staticPhoto == nil {
		t.Fatal("StaticPhoto() should not return nil")
	}

	// Test Build method
	built := staticPhoto.Build()
	if built == nil {
		t.Error("StaticPhoto.Build() should not return nil")
	}
}

func TestAnimatedPhoto(t *testing.T) {
	animatedPhoto := AnimatedPhoto(g.String(testURL))

	if animatedPhoto == nil {
		t.Fatal("AnimatedPhoto() should not return nil")
	}

	// Test Build method
	built := animatedPhoto.Build()
	if built == nil {
		t.Error("AnimatedPhoto.Build() should not return nil")
	}

	// Test method chaining
	result := animatedPhoto.MainFrameTimestamp(2.5)
	if result != animatedPhoto {
		t.Error("AnimatedPhoto.MainFrameTimestamp() should return the same instance for chaining")
	}
}

// ==============================================
// Story Content Tests
// ==============================================

func TestStoryPhoto(t *testing.T) {
	storyPhoto := StoryPhoto(g.String(testURL))

	if storyPhoto == nil {
		t.Fatal("StoryPhoto() should not return nil")
	}

	// Test Build method
	built := storyPhoto.Build()
	if built == nil {
		t.Error("StoryPhoto.Build() should not return nil")
	}
}

func TestStoryVideo(t *testing.T) {
	storyVideo := StoryVideo(g.String(testURL))

	if storyVideo == nil {
		t.Fatal("StoryVideo() should not return nil")
	}

	// Test Build method
	built := storyVideo.Build()
	if built == nil {
		t.Error("StoryVideo.Build() should not return nil")
	}

	// Test method chaining
	result := storyVideo.
		Duration(time.Minute).
		CoverFrameTimestamp(time.Second * 30).
		Animation()

	if result != storyVideo {
		t.Error("StoryVideo methods should return the same instance for chaining")
	}
}

// ==============================================
// Poll Option Tests
// ==============================================

func TestChoice(t *testing.T) {
	choice := Choice(g.String("Option A"))

	if choice == nil {
		t.Fatal("Choice() should not return nil")
	}

	// Test Build method
	built := choice.Build()
	if built.Text == "" {
		t.Error("Choice.Build() should return valid InputPollOption")
	}

	// Test method chaining
	entities := createTestEntities()
	result := choice.
		HTML().
		TextEntities(entities)

	if result != choice {
		t.Error("Choice methods should return the same instance for chaining")
	}

	// Test Markdown method separately
	choice2 := Choice(g.String("Option B"))
	result2 := choice2.Markdown()
	if result2 != choice2 {
		t.Error("Choice.Markdown() should return the same instance for chaining")
	}
}

// ==============================================
// Sticker Tests
// ==============================================

func TestSticker(t *testing.T) {
	file := createTestFile()
	emojiList := g.SliceOf(g.String("😀"), g.String("😃"))
	sticker := Sticker(file, g.String("static"), emojiList)

	if sticker == nil {
		t.Fatal("Sticker() should not return nil")
	}

	// Test Build method
	built := sticker.Build()
	if built.Sticker == nil {
		t.Error("Sticker.Build() should return valid InputSticker")
	}

	// Test method chaining
	maskPos := &gotgbot.MaskPosition{
		Point:  "forehead",
		XShift: 0.0,
		YShift: 0.0,
		Scale:  1.0,
	}
	keywords := g.SliceOf(g.String("happy"), g.String("smile"))

	result := sticker.
		MaskPosition(maskPos).
		Keywords(keywords)

	if result != sticker {
		t.Error("Sticker methods should return the same instance for chaining")
	}
}

// ==============================================
// Checklist Tests
// ==============================================

func TestNewChecklist(t *testing.T) {
	tasks := g.Slice[gotgbot.InputChecklistTask]{}
	checklist := NewChecklist(g.String("Test Checklist"), tasks)

	if checklist == nil {
		t.Fatal("NewChecklist() should not return nil")
	}

	// Test Build method
	built := checklist.Build()
	if built.Title == "" {
		t.Error("Checklist.Build() should return valid InputChecklist")
	}
}

func TestNewChecklistTask(t *testing.T) {
	task := NewChecklistTask(1, g.String("Task 1"))

	if task == nil {
		t.Fatal("NewChecklistTask() should not return nil")
	}

	// Test Build method
	built := task.Build()
	if built.Text == "" {
		t.Error("ChecklistTask.Build() should return valid InputChecklistTask")
	}

	// Test method chaining
	entities := createTestEntities()
	result := task.
		HTML().
		Entities(&entities)

	if result != task {
		t.Error("ChecklistTask methods should return the same instance for chaining")
	}

	// Test Markdown method separately
	task2 := NewChecklistTask(2, g.String("Task 2"))
	result2 := task2.Markdown()
	if result2 != task2 {
		t.Error("ChecklistTask.Markdown() should return the same instance for chaining")
	}
}

// ==============================================
// Edge Cases and Error Scenarios
// ==============================================

func TestEmptyInputs(t *testing.T) {
	// Test with empty strings
	text := Text(g.String(""))
	if text == nil {
		t.Error("Text() should handle empty string input")
	}

	contact := Contact(g.String(""), g.String(""))
	if contact == nil {
		t.Error("Contact() should handle empty string inputs")
	}

	choice := Choice(g.String(""))
	if choice == nil {
		t.Error("Choice() should handle empty string input")
	}
}

func TestZeroValues(t *testing.T) {
	// Test with zero coordinates
	location := Location(0, 0)
	if location == nil {
		t.Error("Location() should handle zero coordinates")
	}

	venue := Venue(0, 0, g.String(""), g.String(""))
	if venue == nil {
		t.Error("Venue() should handle zero coordinates")
	}
}

func TestBuildConsistency(t *testing.T) {
	// Test that multiple Build() calls return consistent results
	file := createTestFile()
	photo := Photo(file).Caption(g.String("test"))

	built1 := photo.Build()
	built2 := photo.Build()

	if built1 == nil || built2 == nil {
		t.Error("Build() should always return non-nil values")
	}

	// Build() should be idempotent (calling it multiple times should not change state)
	// We can't easily test equality of gotgbot types, but we can ensure they're not nil
}

// ==============================================
// Complex Chaining Tests
// ==============================================

func TestComplexMethodChaining(t *testing.T) {
	// Test complex method chaining scenarios
	file := createTestFile()
	entities := createTestEntities()

	// Complex photo chain
	photo := Photo(file).
		Caption(g.String("Complex caption")).
		HTML().
		CaptionEntities(entities).
		ShowCaptionAboveMedia().
		Spoiler()

	built := photo.Build()
	if built == nil {
		t.Error("Complex photo chain should build successfully")
	}

	// Complex video chain
	video := Video(file).
		Cover(g.String("cover")).
		Thumbnail(file).
		Caption(g.String("video caption")).
		Markdown().
		CaptionEntities(entities).
		ShowCaptionAboveMedia().
		Size(1920, 1080).
		Duration(time.Hour).
		StartAt(time.Minute * 30).
		Streamable().
		Spoiler()

	builtVideo := video.Build()
	if builtVideo == nil {
		t.Error("Complex video chain should build successfully")
	}
}

// ==============================================
// Interface Type Assertions
// ==============================================

func TestInterfaceTypeAssertions(t *testing.T) {
	file := createTestFile()

	// Test Media interface implementations
	var media Media

	media = Photo(file)
	if media.Build() == nil {
		t.Error("Photo should implement Media interface correctly")
	}

	media = Video(file)
	if media.Build() == nil {
		t.Error("Video should implement Media interface correctly")
	}

	media = Animation(file)
	if media.Build() == nil {
		t.Error("Animation should implement Media interface correctly")
	}

	media = Audio(file)
	if media.Build() == nil {
		t.Error("Audio should implement Media interface correctly")
	}

	media = Document(file)
	if media.Build() == nil {
		t.Error("Document should implement Media interface correctly")
	}

	// Test MessageContent interface implementations
	var msgContent MessageContent

	msgContent = Text(g.String(testText))
	if msgContent.Build() == nil {
		t.Error("Text should implement MessageContent interface correctly")
	}

	msgContent = Location(testLatitude, testLongitude)
	if msgContent.Build() == nil {
		t.Error("Location should implement MessageContent interface correctly")
	}

	msgContent = Contact(g.String(testPhone), g.String(testFirstName))
	if msgContent.Build() == nil {
		t.Error("Contact should implement MessageContent interface correctly")
	}

	msgContent = Venue(testLatitude, testLongitude, g.String(testTitle), g.String(testAddress))
	if msgContent.Build() == nil {
		t.Error("Venue should implement MessageContent interface correctly")
	}

	msgContent = Invoice(g.String("title"), g.String("desc"), g.String("payload"), g.String("USD"))
	if msgContent.Build() == nil {
		t.Error("Invoice should implement MessageContent interface correctly")
	}

	// Test PaidMedia interface implementations
	var paidMedia PaidMedia

	paidMedia = PaidPhoto(file)
	if paidMedia.Build() == nil {
		t.Error("PaidPhoto should implement PaidMedia interface correctly")
	}

	paidMedia = PaidVideo(file)
	if paidMedia.Build() == nil {
		t.Error("PaidVideo should implement PaidMedia interface correctly")
	}

	// Test ProfilePhoto interface implementations
	var profilePhoto ProfilePhoto

	profilePhoto = StaticPhoto(g.String(testURL))
	if profilePhoto.Build() == nil {
		t.Error("StaticPhoto should implement ProfilePhoto interface correctly")
	}

	profilePhoto = AnimatedPhoto(g.String(testURL))
	if profilePhoto.Build() == nil {
		t.Error("AnimatedPhoto should implement ProfilePhoto interface correctly")
	}

	// Test StoryContent interface implementations
	var storyContent StoryContent

	storyContent = StoryPhoto(g.String(testURL))
	if storyContent.Build() == nil {
		t.Error("StoryPhoto should implement StoryContent interface correctly")
	}

	storyContent = StoryVideo(g.String(testURL))
	if storyContent.Build() == nil {
		t.Error("StoryVideo should implement StoryContent interface correctly")
	}

	// Test PollOption interface implementation
	var pollOption PollOption

	pollOption = Choice(g.String("Option"))
	if pollOption.Build().Text == "" {
		t.Error("Choice should implement PollOption interface correctly")
	}
}
