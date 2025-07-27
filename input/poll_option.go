package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// PollOption represents an input poll option builder.
type PollOption struct {
	input *gotgbot.InputPollOption
}

// NewPollOption creates a new PollOption builder.
func NewPollOption(text String) *PollOption {
	return &PollOption{
		input: &gotgbot.InputPollOption{
			Text: text.Std(),
		},
	}
}

// HTML sets parse mode to HTML for the option text.
func (po *PollOption) HTML() *PollOption {
	po.input.TextParseMode = "HTML"
	return po
}

// Markdown sets parse mode to MarkdownV2 for the option text.
func (po *PollOption) Markdown() *PollOption {
	po.input.TextParseMode = "MarkdownV2"
	return po
}

// TextEntities sets the message entities for the option text.
func (po *PollOption) TextEntities(e entities.Entities) *PollOption {
	po.input.TextEntities = e.Std()
	return po
}

// Build returns the gotgbot.InputPollOption directly as it's not an interface.
func (po *PollOption) BuildOption() *gotgbot.InputPollOption {
	return po.input
}
