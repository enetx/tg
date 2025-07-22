package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/types/effects"
)

type Game struct {
	ctx           *Context
	gameShortName String
	opts          *gotgbot.SendGameOpts
	chatID        Option[int64]
	after         Option[time.Duration]
	deleteAfter   Option[time.Duration]
}

// After schedules the game to be sent after the specified duration.
func (c *Game) After(duration time.Duration) *Game {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the game message to be deleted after the specified duration.
func (c *Game) DeleteAfter(duration time.Duration) *Game {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the game message.
func (c *Game) Silent() *Game {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the game message.
func (c *Game) Protect() *Game {
	c.opts.ProtectContent = true
	return c
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (c *Game) AllowPaidBroadcast() *Game {
	c.opts.AllowPaidBroadcast = true
	return c
}

// Thread sets the message thread ID for the game message.
func (c *Game) Thread(id int64) *Game {
	c.opts.MessageThreadId = id
	return c
}

// Effect sets a message effect for the game message.
func (c *Game) Effect(effect effects.EffectType) *Game {
	c.opts.MessageEffectId = effect.String()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *Game) ReplyTo(messageID int64) *Game {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Markup sets the reply markup keyboard for the game message.
func (c *Game) Markup(kb keyboard.KeyboardBuilder) *Game {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		c.opts.ReplyMarkup = markup
	}

	return c
}

// Business sets the business connection ID for the game message.
func (c *Game) Business(id String) *Game {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// To sets the target chat ID for the game message.
func (c *Game) To(chatID int64) *Game {
	c.chatID = Some(chatID)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *Game) Timeout(duration time.Duration) *Game {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *Game) APIURL(url String) *Game {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send sends the game message to Telegram and returns the result.
func (c *Game) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendGame(chatID, c.gameShortName.Std(), c.opts))
	})
}
