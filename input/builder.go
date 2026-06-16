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

// PollOption represents an interface for input poll option builders.
type PollOption interface {
	Build() gotgbot.InputPollOption
}

// PollMedia represents an interface for media that can be attached to a poll
// question or quiz explanation.
type PollMedia interface {
	BuildPollMedia() gotgbot.InputPollMedia
}

// PollOptionMedia represents an interface for media that can be attached to a poll option.
type PollOptionMedia interface {
	BuildPollOptionMedia() gotgbot.InputPollOptionMedia
}
