package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Checklist represents an input checklist builder.
type Checklist struct {
	input *gotgbot.InputChecklist
}

// NewChecklist creates a new Checklist builder.
func NewChecklist(title g.String, tasks g.Slice[gotgbot.InputChecklistTask]) *Checklist {
	return &Checklist{
		input: &gotgbot.InputChecklist{
			Title: title.Std(),
			Tasks: tasks,
		},
	}
}

// Build returns the gotgbot.InputChecklist directly as it's not an interface.
func (c *Checklist) Build() gotgbot.InputChecklist {
	return *c.input
}
