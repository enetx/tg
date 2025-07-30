package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// PollChoice represents an input poll option builder.
type PollChoice struct {
	input *gotgbot.InputPollOption
}

// Choice creates a new PollOption builder.
func Choice(text g.String) *PollChoice {
	return &PollChoice{
		input: &gotgbot.InputPollOption{
			Text: text.Std(),
		},
	}
}

// HTML sets parse mode to HTML for the option text.
func (po *PollChoice) HTML() *PollChoice {
	po.input.TextParseMode = "HTML"
	return po
}

// Markdown sets parse mode to MarkdownV2 for the option text.
func (po *PollChoice) Markdown() *PollChoice {
	po.input.TextParseMode = "MarkdownV2"
	return po
}

// TextEntities sets the message entities for the option text.
func (po *PollChoice) TextEntities(e entities.Entities) *PollChoice {
	po.input.TextEntities = e.Std()
	return po
}

// Build creates the  gotgbot.InputPollOption.
func (po *PollChoice) Build() gotgbot.InputPollOption {
	return *po.input
}
