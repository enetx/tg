package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// EditEphemeralMessageMedia represents a request to edit the media of an ephemeral message.
type EditEphemeralMessageMedia struct {
	ctx        *Context
	media      input.Media
	chatID     g.Option[int64]
	receiverID g.Option[int64]
	messageID  g.Option[int64]
	opts       *gotgbot.EditEphemeralMessageMediaOpts
}

// ChatID sets the target chat ID for the media edit.
func (eemm *EditEphemeralMessageMedia) ChatID(id int64) *EditEphemeralMessageMedia {
	eemm.chatID = g.Some(id)
	return eemm
}

// ReceiverUser sets the identifier of the user the ephemeral message was sent to.
func (eemm *EditEphemeralMessageMedia) ReceiverUser(id int64) *EditEphemeralMessageMedia {
	eemm.receiverID = g.Some(id)
	return eemm
}

// MessageID sets the target ephemeral message ID to edit.
func (eemm *EditEphemeralMessageMedia) MessageID(id int64) *EditEphemeralMessageMedia {
	eemm.messageID = g.Some(id)
	return eemm
}

// Markup sets the inline keyboard markup for the message.
func (eemm *EditEphemeralMessageMedia) Markup(kb keyboard.Keyboard) *EditEphemeralMessageMedia {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		eemm.opts.ReplyMarkup = markup
	}

	return eemm
}

// Timeout sets a custom timeout for this request.
func (eemm *EditEphemeralMessageMedia) Timeout(duration time.Duration) *EditEphemeralMessageMedia {
	if eemm.opts.RequestOpts == nil {
		eemm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemm.opts.RequestOpts.Timeout = duration

	return eemm
}

// APIURL sets a custom API URL for this request.
func (eemm *EditEphemeralMessageMedia) APIURL(url g.String) *EditEphemeralMessageMedia {
	if eemm.opts.RequestOpts == nil {
		eemm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemm.opts.RequestOpts.APIURL = url.Std()

	return eemm
}

// Send edits the ephemeral message media and returns the result.
func (eemm *EditEphemeralMessageMedia) Send() g.Result[bool] {
	return g.ResultOf(eemm.ctx.Bot.Raw().EditEphemeralMessageMedia(
		eemm.chatID.UnwrapOr(eemm.ctx.EffectiveChat.Id),
		eemm.receiverID.UnwrapOr(eemm.ctx.ephemeralReceiverID()),
		eemm.messageID.UnwrapOr(eemm.ctx.ephemeralMessageID()),
		eemm.media.Build(),
		eemm.opts,
	))
}
