package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type SendDice struct {
	ctx         *Context
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
	opts        *gotgbot.SendDiceOpts
}

// After schedules the dice to be sent after the specified duration.
func (c *SendDice) After(duration time.Duration) *SendDice {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the dice message to be deleted after the specified duration.
func (c *SendDice) DeleteAfter(duration time.Duration) *SendDice {
	c.deleteAfter = Some(duration)
	return c
}

// Emoji sets a custom emoji for the dice.
func (c *SendDice) Emoji(e String) *SendDice {
	c.opts.Emoji = e.Std()
	return c
}

// Dart sets the dice emoji to dart.
func (c *SendDice) Dart() *SendDice {
	c.opts.Emoji = "🎯"
	return c
}

// Slot sets the dice emoji to slot machine.
func (c *SendDice) Slot() *SendDice {
	c.opts.Emoji = "🎰"
	return c
}

// Ball sets the dice emoji to basketball.
func (c *SendDice) Ball() *SendDice {
	c.opts.Emoji = "🏀"
	return c
}

// Soccer sets the dice emoji to soccer ball.
func (c *SendDice) Soccer() *SendDice {
	c.opts.Emoji = "⚽"
	return c
}

// Bowling sets the dice emoji to bowling.
func (c *SendDice) Bowling() *SendDice {
	c.opts.Emoji = "🎳"
	return c
}

// Silent disables notification for the dice message.
func (c *SendDice) Silent() *SendDice {
	c.opts.DisableNotification = true
	return c
}

// Thread sets the message thread ID for the dice message.
func (c *SendDice) Thread(id int64) *SendDice {
	c.opts.MessageThreadId = id
	return c
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (c *SendDice) AllowPaidBroadcast() *SendDice {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Effect sets a message effect for the dice message.
func (c *SendDice) Effect(effect effects.EffectType) *SendDice {
	c.opts.MessageEffectId = effect.String()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendDice) ReplyTo(id int64) *SendDice {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: id}
	return c
}

// Markup sets the reply markup keyboard for the dice message.
func (c *SendDice) Markup(kb keyboard.KeyboardBuilder) *SendDice {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// Business sets the business connection ID for the dice message.
func (c *SendDice) Business(id String) *SendDice {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Protect enables content protection for the dice message.
func (c *SendDice) Protect() *SendDice {
	c.opts.ProtectContent = true
	return c
}

// To sets the target chat ID for the dice message.
func (c *SendDice) To(chatID int64) *SendDice {
	c.chatID = Some(chatID)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendDice) Timeout(duration time.Duration) *SendDice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendDice) APIURL(url String) *SendDice {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the dice message to Telegram and returns the result.
func (c *SendDice) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendDice(chatID, c.opts))
	})
}
