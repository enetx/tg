package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// ChecklistTask represents an input checklist task builder.
type ChecklistTask struct {
	input *gotgbot.InputChecklistTask
}

// NewChecklistTask creates a new ChecklistTask builder.
func NewChecklistTask(id int64, text g.String) *ChecklistTask {
	return &ChecklistTask{
		input: &gotgbot.InputChecklistTask{
			Id:   id,
			Text: text.Std(),
		},
	}
}

// HTML sets parse mode to HTML.
func (ct *ChecklistTask) HTML() *ChecklistTask {
	ct.input.ParseMode = "HTML"
	return ct
}

// Markdown sets parse mode to MarkdownV2.
func (ct *ChecklistTask) Markdown() *ChecklistTask {
	ct.input.ParseMode = "MarkdownV2"
	return ct
}

// Entities sets the text entities for the task text.
func (ct *ChecklistTask) Entities(e *entities.Entities) *ChecklistTask {
	ct.input.TextEntities = e.Std()
	return ct
}

// Build returns the gotgbot.InputChecklistTask directly as it's not an interface.
func (ct *ChecklistTask) Build() gotgbot.InputChecklistTask {
	return *ct.input
}
