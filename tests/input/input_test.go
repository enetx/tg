package input_test

import (
	"testing"

	"github.com/enetx/tg/file"
	"github.com/enetx/tg/input"
)

func TestInput_InterfaceCompliance(t *testing.T) {
	// Test that all types implement their respective interfaces

	// Media interface implementations
	var _ input.Media = &input.MediaPhoto{}
	var _ input.Media = &input.MediaVideo{}
	var _ input.Media = &input.MediaAnimation{}
	var _ input.Media = &input.MediaAudio{}
	var _ input.Media = &input.MediaDocument{}

	// MessageContent interface implementations
	var _ input.MessageContent = &input.MessageText{}
	var _ input.MessageContent = &input.MessageLocation{}
	var _ input.MessageContent = &input.MessageVenue{}
	var _ input.MessageContent = &input.MessageContact{}
	var _ input.MessageContent = &input.MessageInvoice{}

	// PaidMedia interface implementations
	var _ input.PaidMedia = &input.PaidMediaPhoto{}
	var _ input.PaidMedia = &input.PaidMediaVideo{}

	// ProfilePhoto interface implementations
	var _ input.ProfilePhoto = &input.ProfilePhotoStatic{}
	var _ input.ProfilePhoto = &input.ProfilePhotoAnimated{}

	// StoryContent interface implementations
	var _ input.StoryContent = &input.StoryContentPhoto{}
	var _ input.StoryContent = &input.StoryContentVideo{}

	// PollOption interface implementation
	var _ input.PollOption = &input.PollChoice{}
}

func TestInput_BasicFactoryFunctions(t *testing.T) {
	// Test basic factory functions work

	// Test Text creation
	text := input.Text(testText)
	if !assertMessageContent(text) {
		t.Error("Text should implement MessageContent correctly")
	}

	// Test Photo creation
	photo := input.Photo(file.Input(testURL).Ok())
	if !assertMedia(photo) {
		t.Error("Photo should implement Media correctly")
	}

	// Test Video creation
	video := input.Video(file.Input(testURL).Ok())
	if !assertMedia(video) {
		t.Error("Video should implement Media correctly")
	}

	// Test Location creation
	location := input.Location(testLatitude, testLongitude)
	if !assertMessageContent(location) {
		t.Error("Location should implement MessageContent correctly")
	}

	// Test Contact creation
	contact := input.Contact(testPhoneNumber, testFirstName)
	if !assertMessageContent(contact) {
		t.Error("Contact should implement MessageContent correctly")
	}
}
