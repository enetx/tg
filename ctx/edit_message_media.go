package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/input"
	"github.com/enetx/tg/keyboard"
)

// EditMessageMedia represents a request to edit message media.
type EditMessageMedia struct {
	ctx       *Context
	media     input.Media
	opts      *gotgbot.EditMessageMediaOpts
	chatID    g.Option[int64]
	messageID g.Option[int64]
}

// ChatID sets the target chat ID for the media edit.
func (emm *EditMessageMedia) ChatID(id int64) *EditMessageMedia {
	emm.chatID = g.Some(id)
	return emm
}

// MessageID sets the target message ID to edit.
func (emm *EditMessageMedia) MessageID(id int64) *EditMessageMedia {
	emm.messageID = g.Some(id)
	return emm
}

// InlineMessageID sets the inline message ID to edit.
func (emm *EditMessageMedia) InlineMessageID(id g.String) *EditMessageMedia {
	emm.opts.InlineMessageId = id.Std()
	return emm
}

// Business sets the business connection ID for the media edit.
func (emm *EditMessageMedia) Business(id g.String) *EditMessageMedia {
	emm.opts.BusinessConnectionId = id.Std()
	return emm
}

// Markup sets the inline keyboard markup for the message.
func (emm *EditMessageMedia) Markup(kb keyboard.Keyboard) *EditMessageMedia {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		emm.opts.ReplyMarkup = markup
	}

	return emm
}

// Timeout sets a custom timeout for this request.
func (emm *EditMessageMedia) Timeout(duration time.Duration) *EditMessageMedia {
	if emm.opts.RequestOpts == nil {
		emm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emm.opts.RequestOpts.Timeout = duration

	return emm
}

// APIURL sets a custom API URL for this request.
func (emm *EditMessageMedia) APIURL(url g.String) *EditMessageMedia {
	if emm.opts.RequestOpts == nil {
		emm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emm.opts.RequestOpts.APIURL = url.Std()

	return emm
}

// Send edits the message media and returns the result.
func (emm *EditMessageMedia) Send() g.Result[*gotgbot.Message] {
	emm.opts.ChatId = emm.chatID.UnwrapOr(emm.ctx.EffectiveChat.Id)
	emm.opts.MessageId = emm.messageID.UnwrapOr(emm.ctx.EffectiveMessage.MessageId)

	msg, _, err := emm.ctx.Bot.Raw().EditMessageMedia(emm.media.Build(), emm.opts)
	return g.ResultOf(msg, err)
}
