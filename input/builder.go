package input

import "github.com/PaulSonOfLars/gotgbot/v2"

// Media represents an interface for all input media builders.
type Media interface {
	Build() gotgbot.InputMedia
}

// MessageContent represents an interface for all input message content builders.
type MessageContent interface {
	Build() gotgbot.InputMessageContent
}

// PaidMedia represents an interface for all input paid media builders.
type PaidMedia interface {
	Build() gotgbot.InputPaidMedia
}

// ProfilePhoto represents an interface for all input profile photo builders.
type ProfilePhoto interface {
	Build() gotgbot.InputProfilePhoto
}

// StoryContent represents an interface for all input story content builders.
type StoryContent interface {
	Build() gotgbot.InputStoryContent
}
