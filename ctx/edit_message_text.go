package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

type EditMessageText struct {
	ctx       *Context
	text      g.String
	chatID    g.Option[int64]
	messageID g.Option[int64]
	opts      *gotgbot.EditMessageTextOpts
}

// Entities sets custom entities for the edited text.
func (emt *EditMessageText) Entities(e *entities.Entities) *EditMessageText {
	emt.opts.Entities = e.Std()
	return emt
}

// HTML sets the text parse mode to HTML.
func (emt *EditMessageText) HTML() *EditMessageText {
	emt.opts.ParseMode = "HTML"
	return emt
}

// Markdown sets the text parse mode to MarkdownV2.
func (emt *EditMessageText) Markdown() *EditMessageText {
	emt.opts.ParseMode = "MarkdownV2"
	return emt
}

// ChatID sets the target chat ID for the text edit.
func (emt *EditMessageText) ChatID(id int64) *EditMessageText {
	emt.chatID = g.Some(id)
	return emt
}

// MessageID sets the target message ID to edit.
func (emt *EditMessageText) MessageID(id int64) *EditMessageText {
	emt.messageID = g.Some(id)
	return emt
}

// InlineMessageID sets the inline message ID to edit.
func (emt *EditMessageText) InlineMessageID(id g.String) *EditMessageText {
	emt.opts.InlineMessageId = id.Std()
	return emt
}

// Business sets the business connection ID for the text edit.
func (emt *EditMessageText) Business(id g.String) *EditMessageText {
	emt.opts.BusinessConnectionId = id.Std()
	return emt
}

// Timeout sets a custom timeout for this request.
func (emt *EditMessageText) Timeout(duration time.Duration) *EditMessageText {
	if emt.opts.RequestOpts == nil {
		emt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emt.opts.RequestOpts.Timeout = duration

	return emt
}

// APIURL sets a custom API URL for this request.
func (emt *EditMessageText) APIURL(url g.String) *EditMessageText {
	if emt.opts.RequestOpts == nil {
		emt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emt.opts.RequestOpts.APIURL = url.Std()

	return emt
}

// Markup sets the reply markup keyboard for the edited message.
func (emt *EditMessageText) Markup(kb keyboard.Keyboard) *EditMessageText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		emt.opts.ReplyMarkup = markup
	}

	return emt
}

// Preview sets link preview options for the edited text.
func (emt *EditMessageText) Preview(p *preview.Preview) *EditMessageText {
	emt.opts.LinkPreviewOptions = p.Std()
	return emt
}

// Send edits the message text and returns the result.
func (emt *EditMessageText) Send() g.Result[*gotgbot.Message] {
	emt.opts.ChatId = emt.chatID.UnwrapOr(emt.ctx.EffectiveChat.Id)
	emt.opts.MessageId = emt.messageID.UnwrapOr(emt.ctx.EffectiveMessage.MessageId)
	msg, _, err := emt.ctx.Bot.Raw().EditMessageText(emt.text.Std(), emt.opts)

	return g.ResultOf(msg, err)
}
