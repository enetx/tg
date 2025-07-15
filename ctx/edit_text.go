package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"

	. "github.com/enetx/g"
)

type EditText struct {
	ctx       *Context
	text      String
	chatID    Option[int64]
	messageID Option[int64]
	opts      *gotgbot.EditMessageTextOpts
}

func (e *EditText) HTML() *EditText {
	e.opts.ParseMode = "HTML"
	return e
}

func (e *EditText) Markdown() *EditText {
	e.opts.ParseMode = "Markdown"
	return e
}

func (e *EditText) ChatID(id int64) *EditText {
	e.chatID = Some(id)
	return e
}

func (e *EditText) MessageID(id int64) *EditText {
	e.messageID = Some(id)
	return e
}

func (e *EditText) InlineMessageID(id String) *EditText {
	e.opts.InlineMessageId = id.Std()
	return e
}

func (e *EditText) BusinessID(id String) *EditText {
	e.opts.BusinessConnectionId = id.Std()
	return e
}

func (e *EditText) Timeout(duration time.Duration) *EditText {
	e.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return e
}

func (e *EditText) Markup(kb keyboard.KeyboardBuilder) *EditText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		e.opts.ReplyMarkup = markup
	}

	return e
}

func (e *EditText) Preview(p *preview.Preview) *EditText {
	e.opts.LinkPreviewOptions = p.Std()
	return e
}

func (e *EditText) Send() Result[*gotgbot.Message] {
	e.opts.ChatId = e.chatID.UnwrapOr(e.ctx.EffectiveChat.Id)
	e.opts.MessageId = e.messageID.UnwrapOr(e.ctx.EffectiveMessage.MessageId)
	msg, _, err := e.ctx.Bot.Raw().EditMessageText(e.text.Std(), e.opts)

	return ResultOf(msg, err)
}
