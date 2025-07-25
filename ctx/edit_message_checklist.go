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
	messageID            Option[int64]
	checklist            gotgbot.InputChecklist
	opts                 *gotgbot.EditMessageChecklistOpts
	businessConnectionID String
	chatID               Option[int64]
	taskIDCounter        int64
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

// AddTask adds a task to the checklist.
func (emc *EditMessageChecklist) AddTask(text String) *EditMessageChecklist {
	emc.taskIDCounter++

	task := gotgbot.InputChecklistTask{
		Id:   emc.taskIDCounter,
		Text: text.Std(),
	}

	emc.checklist.Tasks = append(emc.checklist.Tasks, task)
	return emc
}

// Title sets the checklist title.
func (emc *EditMessageChecklist) Title(title String) *EditMessageChecklist {
	emc.checklist.Title = title.Std()
	return emc
}

// Tasks sets the entire task list at once.
func (emc *EditMessageChecklist) Tasks(tasks []gotgbot.InputChecklistTask) *EditMessageChecklist {
	emc.checklist.Tasks = tasks
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
func (emc *EditMessageChecklist) Markup(kb keyboard.KeyboardBuilder) *EditMessageChecklist {
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
