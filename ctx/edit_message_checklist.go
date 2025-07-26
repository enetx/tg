package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// EditMessageChecklist represents a request to edit a checklist message.
type EditMessageChecklist struct {
	ctx                  *Context
	checklist            gotgbot.InputChecklist
	opts                 *gotgbot.EditMessageChecklistOpts
	businessConnectionID String
	chatID               Option[int64]
	messageID            Option[int64]
	taskIDCounter        int64
}

// Task starts building a new checklist task.
// Returns a builder allowing you to set formatting (HTML, Markdown, Entities) and add the task.
// After calling .Add(), the task is added to the checklist, and you can continue the chain (e.g., call .Send()).
func (emc *EditMessageChecklist) Task(text String) *TaskBuilder[*EditMessageChecklist] {
	return &TaskBuilder[*EditMessageChecklist]{
		target: emc,
		text:   text,
		add: func(t *EditMessageChecklist, task gotgbot.InputChecklistTask) {
			t.checklist.Tasks = append(t.checklist.Tasks, task)
		},
		next: func(t *EditMessageChecklist) int64 {
			t.taskIDCounter++
			return t.taskIDCounter
		},
	}
}

// MessageID sets the message ID to edit.
func (emc *EditMessageChecklist) MessageID(messageID int64) *EditMessageChecklist {
	emc.messageID = Some(messageID)
	return emc
}

// ChatID sets the chat ID where the message is located.
func (emc *EditMessageChecklist) ChatID(chatID int64) *EditMessageChecklist {
	emc.chatID = Some(chatID)
	return emc
}

// Title sets the checklist title.
func (emc *EditMessageChecklist) Title(title String) *EditMessageChecklist {
	emc.checklist.Title = title.Std()
	return emc
}

// OthersCanAddTasks allows other users to add tasks to the checklist.
func (emc *EditMessageChecklist) OthersCanAddTasks() *EditMessageChecklist {
	emc.checklist.OthersCanAddTasks = true
	return emc
}

// OthersCanMarkTasksAsDone allows other users to mark tasks as done.
func (emc *EditMessageChecklist) OthersCanMarkTasksAsDone() *EditMessageChecklist {
	emc.checklist.OthersCanMarkTasksAsDone = true
	return emc
}

// Markup sets the reply markup keyboard for the checklist message.
func (emc *EditMessageChecklist) Markup(kb keyboard.Keyboard) *EditMessageChecklist {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		emc.opts.ReplyMarkup = markup
	}

	return emc
}

// Timeout sets a custom timeout for this request.
func (emc *EditMessageChecklist) Timeout(duration time.Duration) *EditMessageChecklist {
	if emc.opts.RequestOpts == nil {
		emc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emc.opts.RequestOpts.Timeout = duration

	return emc
}

// APIURL sets a custom API URL for this request.
func (emc *EditMessageChecklist) APIURL(url String) *EditMessageChecklist {
	if emc.opts.RequestOpts == nil {
		emc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emc.opts.RequestOpts.APIURL = url.Std()

	return emc
}

// Send edits the checklist message and returns the result.
func (emc *EditMessageChecklist) Send() Result[*gotgbot.Message] {
	if len(emc.checklist.Tasks) == 0 {
		return Err[*gotgbot.Message](Errorf("no tasks in checklist"))
	}

	if len(emc.checklist.Tasks) > 100 {
		return Err[*gotgbot.Message](Errorf("too many tasks: {} (maximum 100)", len(emc.checklist.Tasks)))
	}

	// Handle regular message editing
	chatID := emc.chatID.UnwrapOr(emc.ctx.EffectiveChat.Id)
	messageID := emc.messageID.UnwrapOr(emc.ctx.EffectiveMessage.MessageId)

	return ResultOf(emc.ctx.Bot.Raw().
		EditMessageChecklist(emc.businessConnectionID.Std(), chatID, messageID, emc.checklist, emc.opts))
}
