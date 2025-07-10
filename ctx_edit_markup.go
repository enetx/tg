package tg

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type EditMarkup struct {
	ctx       *Context
	opts      *gotgbot.EditMessageReplyMarkupOpts
	kb        keyboard.KeyboardBuilder
	chatID    Option[int64]
	messageID Option[int64]
}

func (e *EditMarkup) ChatID(id int64) *EditMarkup {
	e.chatID = Some(id)
	return e
}

func (e *EditMarkup) MessageID(id int64) *EditMarkup {
	e.messageID = Some(id)
	return e
}

func (e *EditMarkup) InlineMessageID(id String) *EditMarkup {
	e.opts.InlineMessageId = id.Std()
	return e
}

func (e *EditMarkup) BusinessID(id String) *EditMarkup {
	e.opts.BusinessConnectionId = id.Std()
	return e
}

func (e *EditMarkup) Send() Result[*gotgbot.Message] {
	if e.kb != nil {
		if markup, ok := e.kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
			e.opts.ReplyMarkup = markup
		}
	}

	e.opts.ChatId = e.chatID.UnwrapOr(e.ctx.EffectiveChat.Id)
	e.opts.MessageId = e.messageID.UnwrapOr(e.ctx.EffectiveMessage.MessageId)
	msg, _, err := e.ctx.Bot.Raw.EditMessageReplyMarkup(e.opts)

	return ResultOf(msg, err)
}
