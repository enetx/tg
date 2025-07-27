package input

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

	// PollChoice checks
	_ PollOption = (*PollChoice)(nil)
)
