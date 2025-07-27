package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Photo creates a new MediaPhoto builder.
func Photo(media String) *MediaPhoto {
	return NewMediaPhoto(media)
}

// Video creates a new MediaVideo builder.
func Video(media String) *MediaVideo {
	return NewMediaVideo(media)
}

// Animation creates a new MediaAnimation builder.
func Animation(media String) *MediaAnimation {
	return NewMediaAnimation(media)
}

// Audio creates a new MediaAudio builder.
func Audio(media String) *MediaAudio {
	return NewMediaAudio(media)
}

// Document creates a new MediaDocument builder.
func Document(media String) *MediaDocument {
	return NewMediaDocument(media)
}

// Text creates a new MessageText builder.
func Text(messageText String) *MessageText {
	return NewMessageText(messageText)
}

// Location creates a new MessageLocation builder.
func Location(latitude, longitude float64) *MessageLocation {
	return NewMessageLocation(latitude, longitude)
}

// Venue creates a new MessageVenue builder.
func Venue(latitude, longitude float64, title, address String) *MessageVenue {
	return NewMessageVenue(latitude, longitude, title, address)
}

// Contact creates a new MessageContact builder.
func Contact(phoneNumber, firstName String) *MessageContact {
	return NewMessageContact(phoneNumber, firstName)
}

// Invoice creates a new MessageInvoice builder.
func Invoice(title, description, payload, currency String, prices []gotgbot.LabeledPrice) *MessageInvoice {
	return NewMessageInvoice(title, description, payload, currency, prices)
}

// PaidPhoto creates a new PaidMediaPhoto builder.
func PaidPhoto(media String) *PaidMediaPhoto {
	return NewPaidMediaPhoto(media)
}

// PaidVideo creates a new PaidMediaVideo builder.
func PaidVideo(media String) *PaidMediaVideo {
	return NewPaidMediaVideo(media)
}

// StaticPhoto creates a new ProfilePhotoStatic builder.
func StaticPhoto(photo String) *ProfilePhotoStatic {
	return NewProfilePhotoStatic(photo)
}

// AnimatedPhoto creates a new ProfilePhotoAnimated builder.
func AnimatedPhoto(video String) *ProfilePhotoAnimated {
	return NewProfilePhotoAnimated(video)
}

// Story content builders - factory functions for InputStoryContent types

// StoryPhoto creates a new StoryContentPhoto builder.
func StoryPhoto(photo String) *StoryContentPhoto {
	return NewStoryContentPhoto(photo)
}

// StoryVideo creates a new StoryContentVideo builder.
func StoryVideo(video String) *StoryContentVideo {
	return NewStoryContentVideo(video)
}

// Compile-time interface checks
var (
	// Media checks
	_ Media = (*MediaPhoto)(nil)
	_ Media = (*MediaVideo)(nil)
	_ Media = (*MediaAnimation)(nil)
	_ Media = (*MediaAudio)(nil)
	_ Media = (*MediaDocument)(nil)

	// MessageContent checks
	_ MessageContent = (*MessageText)(nil)
	_ MessageContent = (*MessageLocation)(nil)
	_ MessageContent = (*MessageVenue)(nil)
	_ MessageContent = (*MessageContact)(nil)
	_ MessageContent = (*MessageInvoice)(nil)

	// PaidMedia checks
	_ PaidMedia = (*PaidMediaPhoto)(nil)
	_ PaidMedia = (*PaidMediaVideo)(nil)

	// ProfilePhoto checks
	_ ProfilePhoto = (*ProfilePhotoStatic)(nil)
	_ ProfilePhoto = (*ProfilePhotoAnimated)(nil)

	// StoryContent checks
	_ StoryContent = (*StoryContentPhoto)(nil)
	_ StoryContent = (*StoryContentVideo)(nil)
)
