package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// SendChecklist represents a request to send a checklist.
type SendChecklist struct {
	ctx                  *Context
	checklist            gotgbot.InputChecklist
	opts                 *gotgbot.SendChecklistOpts
	businessConnectionID String
	chatID               Option[int64]
	after                Option[time.Duration]
	deleteAfter          Option[time.Duration]
	taskIDCounter        int64
}

// After schedules the checklist to be sent after the specified duration.
func (sc *SendChecklist) After(duration time.Duration) *SendChecklist {
	sc.after = Some(duration)
	return sc
}

// DeleteAfter schedules the checklist message to be deleted after the specified duration.
func (sc *SendChecklist) DeleteAfter(duration time.Duration) *SendChecklist {
	sc.deleteAfter = Some(duration)
	return sc
}

// AddTask adds a task to the checklist.
func (sc *SendChecklist) AddTask(text String) *SendChecklist {
	sc.taskIDCounter++

	task := gotgbot.InputChecklistTask{
		Id:   sc.taskIDCounter,
		Text: text.Std(),
	}

	sc.checklist.Tasks = append(sc.checklist.Tasks, task)
	return sc
}

// Tasks sets the entire task list at once.
func (sc *SendChecklist) Tasks(tasks []gotgbot.InputChecklistTask) *SendChecklist {
	sc.checklist.Tasks = tasks
	return sc
}

// OthersCanAddTasks allows other users to add tasks to the checklist.
func (sc *SendChecklist) OthersCanAddTasks() *SendChecklist {
	sc.checklist.OthersCanAddTasks = true
	return sc
}

// OthersCanMarkTasksAsDone allows other users to mark tasks as done.
func (sc *SendChecklist) OthersCanMarkTasksAsDone() *SendChecklist {
	sc.checklist.OthersCanMarkTasksAsDone = true
	return sc
}

// Silent disables notification for the checklist message.
func (sc *SendChecklist) Silent() *SendChecklist {
	sc.opts.DisableNotification = true
	return sc
}

// Protect enables content protection for the checklist message.
func (sc *SendChecklist) Protect() *SendChecklist {
	sc.opts.ProtectContent = true
	return sc
}

// Markup sets the reply markup keyboard for the checklist message.
func (sc *SendChecklist) Markup(kb keyboard.KeyboardBuilder) *SendChecklist {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		sc.opts.ReplyMarkup = markup
	}

	return sc
}

// ReplyTo sets the message ID to reply to.
func (sc *SendChecklist) ReplyTo(messageID int64) *SendChecklist {
	sc.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sc
}

// To sets the target chat ID for the checklist message.
func (sc *SendChecklist) To(chatID int64) *SendChecklist {
	sc.chatID = Some(chatID)
	return sc
}

// Timeout sets a custom timeout for this request.
func (sc *SendChecklist) Timeout(duration time.Duration) *SendChecklist {
	if sc.opts.RequestOpts == nil {
		sc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sc.opts.RequestOpts.Timeout = duration

	return sc
}

// APIURL sets a custom API URL for this request.
func (sc *SendChecklist) APIURL(url String) *SendChecklist {
	if sc.opts.RequestOpts == nil {
		sc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sc.opts.RequestOpts.APIURL = url.Std()

	return sc
}

// Send sends the checklist message to Telegram and returns the result.
func (sc *SendChecklist) Send() Result[*gotgbot.Message] {
	if len(sc.checklist.Tasks) == 0 {
		return Err[*gotgbot.Message](Errorf("no tasks added to checklist"))
	}

	if len(sc.checklist.Tasks) > 100 {
		return Err[*gotgbot.Message](Errorf("too many tasks: {} (maximum 100)", len(sc.checklist.Tasks)))
	}

	return sc.ctx.timers(sc.after, sc.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sc.chatID.UnwrapOr(sc.ctx.EffectiveChat.Id)
		return ResultOf(sc.ctx.Bot.Raw().SendChecklist(sc.businessConnectionID.Std(), chatID, sc.checklist, sc.opts))
	})
}

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
