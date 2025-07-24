package ctx

import (
	"time"

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

// ChatID sets the target chat ID for the markup edit.
func (em *EditMarkup) ChatID(id int64) *EditMarkup {
	em.chatID = Some(id)
	return em
}

// MessageID sets the target message ID to edit.
func (em *EditMarkup) MessageID(id int64) *EditMarkup {
	em.messageID = Some(id)
	return em
}

// InlineMessageID sets the inline message ID to edit.
func (em *EditMarkup) InlineMessageID(id String) *EditMarkup {
	em.opts.InlineMessageId = id.Std()
	return em
}

// Business sets the business connection ID for the markup edit.
func (em *EditMarkup) Business(id String) *EditMarkup {
	em.opts.BusinessConnectionId = id.Std()
	return em
}

// Timeout sets a custom timeout for this request.
func (em *EditMarkup) Timeout(duration time.Duration) *EditMarkup {
	if em.opts.RequestOpts == nil {
		em.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	em.opts.RequestOpts.Timeout = duration

	return em
}

// APIURL sets a custom API URL for this request.
func (em *EditMarkup) APIURL(url String) *EditMarkup {
	if em.opts.RequestOpts == nil {
		em.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	em.opts.RequestOpts.APIURL = url.Std()

	return em
}

// Send edits the message reply markup and returns the result.
func (em *EditMarkup) Send() Result[*gotgbot.Message] {
	if em.kb != nil {
		if markup, ok := em.kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
			em.opts.ReplyMarkup = markup
		}
	}

	em.opts.ChatId = em.chatID.UnwrapOr(em.ctx.EffectiveChat.Id)
	em.opts.MessageId = em.messageID.UnwrapOr(em.ctx.EffectiveMessage.MessageId)
	msg, _, err := em.ctx.Bot.Raw().EditMessageReplyMarkup(em.opts)

	return ResultOf(msg, err)
}
