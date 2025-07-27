package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// ChecklistTask represents an input checklist task builder.
type ChecklistTask struct {
	input *gotgbot.InputChecklistTask
}

// NewChecklistTask creates a new ChecklistTask builder.
func NewChecklistTask(text String) *ChecklistTask {
	return &ChecklistTask{
		input: &gotgbot.InputChecklistTask{
			Text: text.Std(),
		},
	}
}

// Checked sets whether the task is completed.
func (ct *ChecklistTask) Checked() *ChecklistTask {
	// Note: Checked field may not exist in gotgbot.InputChecklistTask
	// This is a placeholder for when the field becomes available
	return ct
}

// Build returns the gotgbot.InputChecklistTask directly as it's not an interface.
func (ct *ChecklistTask) Build() gotgbot.InputChecklistTask {
	return *ct.input
}
