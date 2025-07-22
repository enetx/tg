package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendSticker struct {
	ctx         *Context
	doc         gotgbot.InputFileOrString
	opts        *gotgbot.SendStickerOpts
	file        *File
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	err         error
}

// After schedules the sticker to be sent after the specified duration.
func (c *SendSticker) After(duration time.Duration) *SendSticker {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the sticker message to be deleted after the specified duration.
func (c *SendSticker) DeleteAfter(duration time.Duration) *SendSticker {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the sticker message.
func (c *SendSticker) Silent() *SendSticker {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the sticker message.
func (c *SendSticker) Protect() *SendSticker {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the sticker message.
func (c *SendSticker) Markup(kb keyboard.KeyboardBuilder) *SendSticker {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Emoji sets the emoji associated with the sticker.
func (c *SendSticker) Emoji(emoji String) *SendSticker {
	c.opts.Emoji = emoji.Std()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendSticker) ReplyTo(messageID int64) *SendSticker {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendSticker) Timeout(duration time.Duration) *SendSticker {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendSticker) APIURL(url String) *SendSticker {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the sticker message.
func (c *SendSticker) Business(id String) *SendSticker {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the sticker message.
func (c *SendSticker) Thread(id int64) *SendSticker {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the sticker message.
func (c *SendSticker) To(chatID int64) *SendSticker {
	c.chatID = Some(chatID)
	return c
}

// Send sends the sticker message to Telegram and returns the result.
func (c *SendSticker) Send() Result[*gotgbot.Message] {
	if c.err != nil {
		return Err[*gotgbot.Message](c.err)
	}

	if c.file != nil {
		defer c.file.Close()
	}

	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendSticker(chatID, c.doc, c.opts))
	})
}
