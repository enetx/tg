package input_test

import (
	"testing"

	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestBuilder_InterfaceImplementations(t *testing.T) {
	// Test that concrete types properly implement their interfaces

	// Test Media implementations
	photo := input.Photo(file.Input(testURL).Ok())
	if !assertMedia(photo) {
		t.Error("Photo should implement Media interface correctly")
	}

	video := input.Video(file.Input(testURL).Ok())
	if !assertMedia(video) {
		t.Error("Video should implement Media interface correctly")
	}

	// Test MessageContent implementations
	text := input.Text(testText)
	if !assertMessageContent(text) {
		t.Error("Text should implement MessageContent interface correctly")
	}

	location := input.Location(testLatitude, testLongitude)
	if !assertMessageContent(location) {
		t.Error("Location should implement MessageContent interface correctly")
	}

	contact := input.Contact(testPhoneNumber, testFirstName)
	if !assertMessageContent(contact) {
		t.Error("Contact should implement MessageContent interface correctly")
	}
}

func TestBuilder_InterfacePolymorphism(t *testing.T) {
	// Test that interfaces can be used polymorphically

	// Test Media interface polymorphism
	var mediaItems []input.Media
	mediaItems = append(mediaItems, input.Photo(file.Input(testURL).Ok()))
	mediaItems = append(mediaItems, input.Video(file.Input(testURL).Ok()))

	for i, media := range mediaItems {
		if media == nil {
			t.Errorf("Media item %d should not be nil", i)
		}
		if media.Build() == nil {
			t.Errorf("Media item %d should build correctly", i)
		}
	}

	// Test MessageContent interface polymorphism
	var messageContents []input.MessageContent
	messageContents = append(messageContents, input.Text(testText))
	messageContents = append(messageContents, input.Location(testLatitude, testLongitude))

	for i, content := range messageContents {
		if content == nil {
			t.Errorf("MessageContent item %d should not be nil", i)
		}
		if content.Build() == nil {
			t.Errorf("MessageContent item %d should build correctly", i)
		}
	}
}
