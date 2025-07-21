package poll

// PollType represents the type of Telegram poll: regular or quiz.
type PollType int

const (
	Regular PollType = iota // Regular poll (multiple options allowed)
	Quiz                    // Quiz with correct answer
)

// String returns the canonical string representation used by Telegram Bot API.
func (p PollType) String() string {
	switch p {
	case Regular:
		return "regular"
	case Quiz:
		return "quiz"
	default:
		return "unknown"
	}
}
