package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
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

// Entities sets custom entities for the edited text.
func (et *EditText) Entities(e *entities.Entities) *EditText {
	et.opts.Entities = e.Std()
	return et
}

// HTML sets the text parse mode to HTML.
func (et *EditText) HTML() *EditText {
	et.opts.ParseMode = "HTML"
	return et
}

// Markdown sets the text parse mode to MarkdownV2.
func (et *EditText) Markdown() *EditText {
	et.opts.ParseMode = "MarkdownV2"
	return et
}

// ChatID sets the target chat ID for the text edit.
func (et *EditText) ChatID(id int64) *EditText {
	et.chatID = Some(id)
	return et
}

// MessageID sets the target message ID to edit.
func (et *EditText) MessageID(id int64) *EditText {
	et.messageID = Some(id)
	return et
}

// InlineMessageID sets the inline message ID to edit.
func (et *EditText) InlineMessageID(id String) *EditText {
	et.opts.InlineMessageId = id.Std()
	return et
}

// Business sets the business connection ID for the text edit.
func (et *EditText) Business(id String) *EditText {
	et.opts.BusinessConnectionId = id.Std()
	return et
}

// Timeout sets a custom timeout for this request.
func (et *EditText) Timeout(duration time.Duration) *EditText {
	if et.opts.RequestOpts == nil {
		et.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	et.opts.RequestOpts.Timeout = duration

	return et
}

// APIURL sets a custom API URL for this request.
func (et *EditText) APIURL(url String) *EditText {
	if et.opts.RequestOpts == nil {
		et.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	et.opts.RequestOpts.APIURL = url.Std()

	return et
}

// Markup sets the reply markup keyboard for the edited message.
func (et *EditText) Markup(kb keyboard.KeyboardBuilder) *EditText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		et.opts.ReplyMarkup = markup
	}

	return et
}

// Preview sets link preview options for the edited text.
func (et *EditText) Preview(p *preview.Preview) *EditText {
	et.opts.LinkPreviewOptions = p.Std()
	return et
}

// Send edits the message text and returns the result.
func (et *EditText) Send() Result[*gotgbot.Message] {
	et.opts.ChatId = et.chatID.UnwrapOr(et.ctx.EffectiveChat.Id)
	et.opts.MessageId = et.messageID.UnwrapOr(et.ctx.EffectiveMessage.MessageId)
	msg, _, err := et.ctx.Bot.Raw().EditMessageText(et.text.Std(), et.opts)

	return ResultOf(msg, err)
}
