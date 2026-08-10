package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

// EditEphemeralMessageCaption represents a request to edit the caption of an ephemeral message.
type EditEphemeralMessageCaption struct {
	ctx        *Context
	chatID     g.Option[int64]
	receiverID g.Option[int64]
	messageID  g.Option[int64]
	opts       *gotgbot.EditEphemeralMessageCaptionOpts
}

// ChatID sets the target chat ID for the caption edit.
func (eemc *EditEphemeralMessageCaption) ChatID(id int64) *EditEphemeralMessageCaption {
	eemc.chatID = g.Some(id)
	return eemc
}

// ReceiverUser sets the identifier of the user the ephemeral message was sent to.
func (eemc *EditEphemeralMessageCaption) ReceiverUser(id int64) *EditEphemeralMessageCaption {
	eemc.receiverID = g.Some(id)
	return eemc
}

// MessageID sets the target ephemeral message ID to edit.
func (eemc *EditEphemeralMessageCaption) MessageID(id int64) *EditEphemeralMessageCaption {
	eemc.messageID = g.Some(id)
	return eemc
}

// HTML sets the caption parse mode to HTML.
func (eemc *EditEphemeralMessageCaption) HTML() *EditEphemeralMessageCaption {
	eemc.opts.ParseMode = gotgbot.ParseModeHTML
	return eemc
}

// Markdown sets the caption parse mode to MarkdownV2.
func (eemc *EditEphemeralMessageCaption) Markdown() *EditEphemeralMessageCaption {
	eemc.opts.ParseMode = gotgbot.ParseModeMarkdownV2
	return eemc
}

// Entities sets custom entities for the caption.
func (eemc *EditEphemeralMessageCaption) Entities(e *entities.Entities) *EditEphemeralMessageCaption {
	eemc.opts.CaptionEntities = e.Std()
	return eemc
}

// Markup sets the inline keyboard markup for the message.
func (eemc *EditEphemeralMessageCaption) Markup(kb keyboard.Keyboard) *EditEphemeralMessageCaption {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		eemc.opts.ReplyMarkup = markup
	}

	return eemc
}

// Timeout sets a custom timeout for this request.
func (eemc *EditEphemeralMessageCaption) Timeout(duration time.Duration) *EditEphemeralMessageCaption {
	if eemc.opts.RequestOpts == nil {
		eemc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemc.opts.RequestOpts.Timeout = duration

	return eemc
}

// APIURL sets a custom API URL for this request.
func (eemc *EditEphemeralMessageCaption) APIURL(url g.String) *EditEphemeralMessageCaption {
	if eemc.opts.RequestOpts == nil {
		eemc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemc.opts.RequestOpts.APIURL = url.Std()

	return eemc
}

// Send edits the ephemeral message caption and returns the result.
func (eemc *EditEphemeralMessageCaption) Send() g.Result[bool] {
	return g.ResultOf(eemc.ctx.Bot.Raw().EditEphemeralMessageCaption(
		eemc.chatID.UnwrapOr(eemc.ctx.EffectiveChat.Id),
		eemc.receiverID.UnwrapOr(eemc.ctx.ephemeralReceiverID()),
		eemc.messageID.UnwrapOr(eemc.ctx.ephemeralMessageID()),
		eemc.opts,
	))
}
