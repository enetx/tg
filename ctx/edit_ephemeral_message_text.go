package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

// EditEphemeralMessageText represents a request to edit the text of an ephemeral message.
type EditEphemeralMessageText struct {
	ctx        *Context
	text       g.String
	chatID     g.Option[int64]
	receiverID g.Option[int64]
	messageID  g.Option[int64]
	opts       *gotgbot.EditEphemeralMessageTextOpts
}

// ChatID sets the target chat ID for the text edit.
func (eemt *EditEphemeralMessageText) ChatID(id int64) *EditEphemeralMessageText {
	eemt.chatID = g.Some(id)
	return eemt
}

// ReceiverUser sets the identifier of the user the ephemeral message was sent to.
func (eemt *EditEphemeralMessageText) ReceiverUser(id int64) *EditEphemeralMessageText {
	eemt.receiverID = g.Some(id)
	return eemt
}

// MessageID sets the target ephemeral message ID to edit.
func (eemt *EditEphemeralMessageText) MessageID(id int64) *EditEphemeralMessageText {
	eemt.messageID = g.Some(id)
	return eemt
}

// HTML sets the text parse mode to HTML.
func (eemt *EditEphemeralMessageText) HTML() *EditEphemeralMessageText {
	eemt.opts.ParseMode = gotgbot.ParseModeHTML
	return eemt
}

// Markdown sets the text parse mode to MarkdownV2.
func (eemt *EditEphemeralMessageText) Markdown() *EditEphemeralMessageText {
	eemt.opts.ParseMode = gotgbot.ParseModeMarkdownV2
	return eemt
}

// Entities sets custom entities for the edited text.
func (eemt *EditEphemeralMessageText) Entities(e *entities.Entities) *EditEphemeralMessageText {
	eemt.opts.Entities = e.Std()
	return eemt
}

// Preview sets link preview options for the edited text.
func (eemt *EditEphemeralMessageText) Preview(p *preview.Preview) *EditEphemeralMessageText {
	eemt.opts.LinkPreviewOptions = p.Std()
	return eemt
}

// Markup sets the reply markup keyboard for the edited message.
func (eemt *EditEphemeralMessageText) Markup(kb keyboard.Keyboard) *EditEphemeralMessageText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		eemt.opts.ReplyMarkup = markup
	}

	return eemt
}

// Timeout sets a custom timeout for this request.
func (eemt *EditEphemeralMessageText) Timeout(duration time.Duration) *EditEphemeralMessageText {
	if eemt.opts.RequestOpts == nil {
		eemt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemt.opts.RequestOpts.Timeout = duration

	return eemt
}

// APIURL sets a custom API URL for this request.
func (eemt *EditEphemeralMessageText) APIURL(url g.String) *EditEphemeralMessageText {
	if eemt.opts.RequestOpts == nil {
		eemt.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	eemt.opts.RequestOpts.APIURL = url.Std()

	return eemt
}

// Send edits the ephemeral message text and returns the result.
func (eemt *EditEphemeralMessageText) Send() g.Result[bool] {
	return g.ResultOf(eemt.ctx.Bot.Raw().EditEphemeralMessageText(
		eemt.chatID.UnwrapOr(eemt.ctx.EffectiveChat.Id),
		eemt.receiverID.UnwrapOr(eemt.ctx.ephemeralReceiverID()),
		eemt.messageID.UnwrapOr(eemt.ctx.ephemeralMessageID()),
		eemt.text.Std(),
		eemt.opts,
	))
}
