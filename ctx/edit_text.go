package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

type EditText struct {
	ctx       *Context
	text      String
	chatID    Option[int64]
	messageID Option[int64]
	opts      *gotgbot.EditMessageTextOpts
}

// HTML sets the text parse mode to HTML.
func (e *EditText) HTML() *EditText {
	e.opts.ParseMode = "HTML"
	return e
}

// Markdown sets the text parse mode to MarkdownV2.
func (e *EditText) Markdown() *EditText {
	e.opts.ParseMode = "MarkdownV2"
	return e
}

// ChatID sets the target chat ID for the text edit.
func (e *EditText) ChatID(id int64) *EditText {
	e.chatID = Some(id)
	return e
}

// MessageID sets the target message ID to edit.
func (e *EditText) MessageID(id int64) *EditText {
	e.messageID = Some(id)
	return e
}

// InlineMessageID sets the inline message ID to edit.
func (e *EditText) InlineMessageID(id String) *EditText {
	e.opts.InlineMessageId = id.Std()
	return e
}

// Business sets the business connection ID for the text edit.
func (e *EditText) Business(id String) *EditText {
	e.opts.BusinessConnectionId = id.Std()
	return e
}

// Timeout sets the request timeout duration.
func (e *EditText) Timeout(duration time.Duration) *EditText {
	e.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return e
}

// Markup sets the reply markup keyboard for the edited message.
func (e *EditText) Markup(kb keyboard.KeyboardBuilder) *EditText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		e.opts.ReplyMarkup = markup
	}

	return e
}

// Preview sets link preview options for the edited text.
func (e *EditText) Preview(p *preview.Preview) *EditText {
	e.opts.LinkPreviewOptions = p.Std()
	return e
}

// Send edits the message text and returns the result.
func (e *EditText) Send() Result[*gotgbot.Message] {
	e.opts.ChatId = e.chatID.UnwrapOr(e.ctx.EffectiveChat.Id)
	e.opts.MessageId = e.messageID.UnwrapOr(e.ctx.EffectiveMessage.MessageId)
	msg, _, err := e.ctx.Bot.Raw().EditMessageText(e.text.Std(), e.opts)

	return ResultOf(msg, err)
}
