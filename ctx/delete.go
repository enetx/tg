package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/core"
)

type Delete struct {
	ctx       *Context
	chatID    Option[int64]
	messageID Option[int64]
	after     Option[time.Duration]
	opts      *gotgbot.DeleteMessageOpts
}

// After schedules the message deletion after the specified duration.
func (c *Delete) After(duration time.Duration) *Delete {
	c.after = Some(duration)
	return c
}

// ChatID sets the target chat ID for the delete action.
func (c *Delete) ChatID(id int64) *Delete {
	c.chatID = Some(id)
	return c
}

// MessageID sets the target message ID to delete.
func (c *Delete) MessageID(id int64) *Delete {
	c.messageID = Some(id)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Delete) Timeout(duration time.Duration) *Delete {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Delete) APIURL(url String) *Delete {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send deletes the message and returns the result.
func (c *Delete) Send() Result[bool] {
	chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
	messageID := c.messageID.UnwrapOr(c.ctx.EffectiveMessage.MessageId)

	if c.after.IsSome() {
		delay := c.after.Some()
		c.after = None[time.Duration]()

		bot := c.ctx.Bot

		var opts *gotgbot.DeleteMessageOpts
		if c.opts != nil {
			ocp := *c.opts
			opts = &ocp
		}

		go func(bot core.BotAPI, chatID, messageID int64, opts *gotgbot.DeleteMessageOpts, delay time.Duration) {
			<-time.After(delay)
			bot.Raw().DeleteMessage(chatID, messageID, opts)
		}(bot, chatID, messageID, opts, delay)

		return Ok(true)
	}

	return ResultOf(c.ctx.Bot.Raw().DeleteMessage(chatID, messageID, c.opts))
}
