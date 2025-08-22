package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
)

// EditMessageCaption represents a request to edit a message caption.
type EditMessageCaption struct {
	ctx       *Context
	opts      *gotgbot.EditMessageCaptionOpts
	chatID    g.Option[int64]
	messageID g.Option[int64]
}

// ChatID sets the target chat ID for the caption edit.
func (emc *EditMessageCaption) ChatID(id int64) *EditMessageCaption {
	emc.chatID = g.Some(id)
	return emc
}

// MessageID sets the target message ID to edit.
func (emc *EditMessageCaption) MessageID(id int64) *EditMessageCaption {
	emc.messageID = g.Some(id)
	return emc
}

// InlineMessageID sets the inline message ID to edit.
func (emc *EditMessageCaption) InlineMessageID(id g.String) *EditMessageCaption {
	emc.opts.InlineMessageId = id.Std()
	return emc
}

// Business sets the business connection ID for the caption edit.
func (emc *EditMessageCaption) Business(id g.String) *EditMessageCaption {
	emc.opts.BusinessConnectionId = id.Std()
	return emc
}

// HTML sets the parse mode to HTML.
func (emc *EditMessageCaption) HTML() *EditMessageCaption {
	emc.opts.ParseMode = "HTML"
	return emc
}

// Markdown sets the parse mode to MarkdownV2.
func (emc *EditMessageCaption) Markdown() *EditMessageCaption {
	emc.opts.ParseMode = "MarkdownV2"
	return emc
}

// Entities sets custom entities for the caption.
func (emc *EditMessageCaption) Entities(e *entities.Entities) *EditMessageCaption {
	emc.opts.CaptionEntities = e.Std()
	return emc
}

// ShowCaptionAboveMedia sets whether the caption should be shown above the media.
func (emc *EditMessageCaption) ShowCaptionAboveMedia() *EditMessageCaption {
	emc.opts.ShowCaptionAboveMedia = true
	return emc
}

// Markup sets the inline keyboard markup for the message.
func (emc *EditMessageCaption) Markup(kb keyboard.Keyboard) *EditMessageCaption {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		emc.opts.ReplyMarkup = markup
	}

	return emc
}

// Timeout sets a custom timeout for this request.
func (emc *EditMessageCaption) Timeout(duration time.Duration) *EditMessageCaption {
	if emc.opts.RequestOpts == nil {
		emc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emc.opts.RequestOpts.Timeout = duration

	return emc
}

// APIURL sets a custom API URL for this request.
func (emc *EditMessageCaption) APIURL(url g.String) *EditMessageCaption {
	if emc.opts.RequestOpts == nil {
		emc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emc.opts.RequestOpts.APIURL = url.Std()

	return emc
}

// Send edits the message caption and returns the result.
func (emc *EditMessageCaption) Send() g.Result[*gotgbot.Message] {
	emc.opts.ChatId = emc.chatID.UnwrapOr(emc.ctx.EffectiveChat.Id)
	emc.opts.MessageId = emc.messageID.UnwrapOr(emc.ctx.EffectiveMessage.MessageId)

	msg, _, err := emc.ctx.Bot.Raw().EditMessageCaption(emc.opts)
	return g.ResultOf(msg, err)
}
