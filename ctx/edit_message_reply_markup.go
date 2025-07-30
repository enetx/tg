package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type EditMessageReplyMarkup struct {
	ctx       *Context
	opts      *gotgbot.EditMessageReplyMarkupOpts
	kb        keyboard.Keyboard
	chatID    g.Option[int64]
	messageID g.Option[int64]
}

// ChatID sets the target chat ID for the markup edit.
func (emrm *EditMessageReplyMarkup) ChatID(id int64) *EditMessageReplyMarkup {
	emrm.chatID = g.Some(id)
	return emrm
}

// MessageID sets the target message ID to edit.
func (emrm *EditMessageReplyMarkup) MessageID(id int64) *EditMessageReplyMarkup {
	emrm.messageID = g.Some(id)
	return emrm
}

// InlineMessageID sets the inline message ID to edit.
func (emrm *EditMessageReplyMarkup) InlineMessageID(id g.String) *EditMessageReplyMarkup {
	emrm.opts.InlineMessageId = id.Std()
	return emrm
}

// Business sets the business connection ID for the markup edit.
func (emrm *EditMessageReplyMarkup) Business(id g.String) *EditMessageReplyMarkup {
	emrm.opts.BusinessConnectionId = id.Std()
	return emrm
}

// Timeout sets a custom timeout for this request.
func (emrm *EditMessageReplyMarkup) Timeout(duration time.Duration) *EditMessageReplyMarkup {
	if emrm.opts.RequestOpts == nil {
		emrm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emrm.opts.RequestOpts.Timeout = duration

	return emrm
}

// APIURL sets a custom API URL for this request.
func (emrm *EditMessageReplyMarkup) APIURL(url g.String) *EditMessageReplyMarkup {
	if emrm.opts.RequestOpts == nil {
		emrm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emrm.opts.RequestOpts.APIURL = url.Std()

	return emrm
}

// Send edits the message reply markup and returns the result.
func (emrm *EditMessageReplyMarkup) Send() g.Result[*gotgbot.Message] {
	if emrm.kb != nil {
		if markup, ok := emrm.kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
			emrm.opts.ReplyMarkup = markup
		}
	}

	emrm.opts.ChatId = emrm.chatID.UnwrapOr(emrm.ctx.EffectiveChat.Id)
	emrm.opts.MessageId = emrm.messageID.UnwrapOr(emrm.ctx.EffectiveMessage.MessageId)
	msg, _, err := emrm.ctx.Bot.Raw().EditMessageReplyMarkup(emrm.opts)

	return g.ResultOf(msg, err)
}
