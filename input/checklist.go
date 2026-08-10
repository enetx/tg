package input

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
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

// TitleHTML sets the checklist title parse mode to HTML.
// Ignored if TitleEntities are set.
func (c *Checklist) TitleHTML() *Checklist {
	c.input.ParseMode = gotgbot.ParseModeHTML
	return c
}

// TitleMarkdown sets the checklist title parse mode to MarkdownV2.
// Ignored if TitleEntities are set.
func (c *Checklist) TitleMarkdown() *Checklist {
	c.input.ParseMode = gotgbot.ParseModeMarkdownV2
	return c
}

// TitleEntities sets custom entities for the checklist title.
// Overrides any title parse mode if set.
func (c *Checklist) TitleEntities(e entities.Entities) *Checklist {
	c.input.TitleEntities = e.Std()
	return c
}

// OthersCanAddTasks allows other users to add tasks to the checklist.
func (c *Checklist) OthersCanAddTasks() *Checklist {
	c.input.OthersCanAddTasks = true
	return c
}

// OthersCanMarkTasksAsDone allows other users to mark tasks as done.
func (c *Checklist) OthersCanMarkTasksAsDone() *Checklist {
	c.input.OthersCanMarkTasksAsDone = true
	return c
}

// Build returns the gotgbot.InputChecklist directly as it's not an interface.
func (c *Checklist) Build() gotgbot.InputChecklist {
	return *c.input
}
