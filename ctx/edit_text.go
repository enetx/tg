package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/entities"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/preview"
)

type EditText struct {
	ctx       *Context
	text      String
	chatID    Option[int64]
	messageID Option[int64]
	opts      *gotgbot.EditMessageTextOpts
}

// Entities sets custom entities for the edited text.
func (c *EditText) Entities(e *entities.Entities) *EditText {
	c.opts.Entities = e.Std()
	return c
}

// HTML sets the text parse mode to HTML.
func (c *EditText) HTML() *EditText {
	c.opts.ParseMode = "HTML"
	return c
}

// Markdown sets the text parse mode to MarkdownV2.
func (c *EditText) Markdown() *EditText {
	c.opts.ParseMode = "MarkdownV2"
	return c
}

// ChatID sets the target chat ID for the text edit.
func (c *EditText) ChatID(id int64) *EditText {
	c.chatID = Some(id)
	return c
}

// MessageID sets the target message ID to edit.
func (c *EditText) MessageID(id int64) *EditText {
	c.messageID = Some(id)
	return c
}

// InlineMessageID sets the inline message ID to edit.
func (c *EditText) InlineMessageID(id String) *EditText {
	c.opts.InlineMessageId = id.Std()
	return c
}

// Business sets the business connection ID for the text edit.
func (c *EditText) Business(id String) *EditText {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Timeout sets a custom timeout for this request.
func (c *EditText) Timeout(duration time.Duration) *EditText {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *EditText) APIURL(url String) *EditText {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Markup sets the reply markup keyboard for the edited message.
func (c *EditText) Markup(kb keyboard.KeyboardBuilder) *EditText {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		c.opts.ReplyMarkup = markup
	}

	return c
}

// Preview sets link preview options for the edited text.
func (c *EditText) Preview(p *preview.Preview) *EditText {
	c.opts.LinkPreviewOptions = p.Std()
	return c
}

// Send edits the message text and returns the result.
func (c *EditText) Send() Result[*gotgbot.Message] {
	c.opts.ChatId = c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	c.opts.MessageId = c.messageID.UnwrapOr(c.ctx.EffectiveMessage.MessageId)
	msg, _, err := c.ctx.Bot.Raw().EditMessageText(c.text.Std(), c.opts)

	return ResultOf(msg, err)
}
