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
func (c *EditMarkup) ChatID(id int64) *EditMarkup {
	c.chatID = Some(id)
	return c
}

// MessageID sets the target message ID to edit.
func (c *EditMarkup) MessageID(id int64) *EditMarkup {
	c.messageID = Some(id)
	return c
}

// InlineMessageID sets the inline message ID to edit.
func (c *EditMarkup) InlineMessageID(id String) *EditMarkup {
	c.opts.InlineMessageId = id.Std()
	return c
}

// Business sets the business connection ID for the markup edit.
func (c *EditMarkup) Business(id String) *EditMarkup {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Timeout sets a custom timeout for this request.
func (c *EditMarkup) Timeout(duration time.Duration) *EditMarkup {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *EditMarkup) APIURL(url String) *EditMarkup {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send edits the message reply markup and returns the result.
func (c *EditMarkup) Send() Result[*gotgbot.Message] {
	if c.kb != nil {
		if markup, ok := c.kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
			c.opts.ReplyMarkup = markup
		}
	}

	c.opts.ChatId = c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	c.opts.MessageId = c.messageID.UnwrapOr(c.ctx.EffectiveMessage.MessageId)
	msg, _, err := c.ctx.Bot.Raw().EditMessageReplyMarkup(c.opts)

	return ResultOf(msg, err)
}
