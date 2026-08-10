package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

// EditEphemeralMessageReplyMarkup represents a request to edit the reply markup
// of an ephemeral message.
type EditEphemeralMessageReplyMarkup struct {
	ctx        *Context
	chatID     g.Option[int64]
	receiverID g.Option[int64]
	messageID  g.Option[int64]
	opts       *gotgbot.EditEphemeralMessageReplyMarkupOpts
}

// ChatID sets the target chat ID for the reply markup edit.
func (eemrm *EditEphemeralMessageReplyMarkup) ChatID(id int64) *EditEphemeralMessageReplyMarkup {
	eemrm.chatID = g.Some(id)
	return eemrm
}

// ReceiverUser sets the identifier of the user the ephemeral message was sent to.
func (eemrm *EditEphemeralMessageReplyMarkup) ReceiverUser(id int64) *EditEphemeralMessageReplyMarkup {
	eemrm.receiverID = g.Some(id)
	return eemrm
}

// MessageID sets the target ephemeral message ID to edit.
func (eemrm *EditEphemeralMessageReplyMarkup) MessageID(id int64) *EditEphemeralMessageReplyMarkup {
	eemrm.messageID = g.Some(id)
	return eemrm
}

// Markup sets the inline keyboard markup for the message.
func (eemrm *EditEphemeralMessageReplyMarkup) Markup(kb keyboard.Keyboard) *EditEphemeralMessageReplyMarkup {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		eemrm.opts.ReplyMarkup = markup
	}

	return eemrm
}

// Timeout sets a custom timeout for this request.
func (eemrm *EditEphemeralMessageReplyMarkup) Timeout(duration time.Duration) *EditEphemeralMessageReplyMarkup {
	if eemrm.opts.RequestOpts == nil {
		eemrm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemrm.opts.RequestOpts.Timeout = duration

	return eemrm
}

// APIURL sets a custom API URL for this request.
func (eemrm *EditEphemeralMessageReplyMarkup) APIURL(url g.String) *EditEphemeralMessageReplyMarkup {
	if eemrm.opts.RequestOpts == nil {
		eemrm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemrm.opts.RequestOpts.APIURL = url.Std()

	return eemrm
}

// Send edits the ephemeral message reply markup and returns the result.
func (eemrm *EditEphemeralMessageReplyMarkup) Send() g.Result[bool] {
	return g.ResultOf(eemrm.ctx.Bot.Raw().EditEphemeralMessageReplyMarkup(
		eemrm.chatID.UnwrapOr(eemrm.ctx.EffectiveChat.Id),
		eemrm.receiverID.UnwrapOr(eemrm.ctx.ephemeralReceiverID()),
		eemrm.messageID.UnwrapOr(eemrm.ctx.ephemeralMessageID()),
		eemrm.opts,
	))
}
