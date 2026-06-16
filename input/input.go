package input

// Compile-time interface checks
var (
	// Media checks
	_ Media = (*MediaPhoto)(nil)
	_ Media = (*MediaVideo)(nil)
	_ Media = (*MediaAnimation)(nil)
	_ Media = (*MediaAudio)(nil)
	_ Media = (*MediaDocument)(nil)
	_ Media = (*MediaLivePhoto)(nil)

	// PollMedia checks (poll question / quiz explanation)
	_ PollMedia = (*MediaPhoto)(nil)
	_ PollMedia = (*MediaVideo)(nil)
	_ PollMedia = (*MediaAnimation)(nil)
	_ PollMedia = (*MediaAudio)(nil)
	_ PollMedia = (*MediaDocument)(nil)
	_ PollMedia = (*MediaLivePhoto)(nil)
	_ PollMedia = (*MediaLocation)(nil)
	_ PollMedia = (*MediaVenue)(nil)

	// PollOptionMedia checks
	_ PollOptionMedia = (*MediaPhoto)(nil)
	_ PollOptionMedia = (*MediaVideo)(nil)
	_ PollOptionMedia = (*MediaAnimation)(nil)
	_ PollOptionMedia = (*MediaLivePhoto)(nil)
	_ PollOptionMedia = (*MediaLocation)(nil)
	_ PollOptionMedia = (*MediaVenue)(nil)
	_ PollOptionMedia = (*MediaSticker)(nil)

	// MessageContent checks
	_ MessageContent = (*MessageText)(nil)
	_ MessageContent = (*MessageLocation)(nil)
	_ MessageContent = (*MessageVenue)(nil)
	_ MessageContent = (*MessageContact)(nil)
	_ MessageContent = (*MessageInvoice)(nil)

	// PaidMedia checks
	_ PaidMedia = (*PaidMediaPhoto)(nil)
	_ PaidMedia = (*PaidMediaVideo)(nil)
	_ PaidMedia = (*PaidMediaLivePhoto)(nil)

	// ProfilePhoto checks
	_ ProfilePhoto = (*ProfilePhotoStatic)(nil)
	_ ProfilePhoto = (*ProfilePhotoAnimated)(nil)

	// StoryContent checks
	_ StoryContent = (*StoryContentPhoto)(nil)
	_ StoryContent = (*StoryContentVideo)(nil)

	// PollChoice checks
	_ PollOption = (*PollChoice)(nil)
)
