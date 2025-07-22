package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendLocation struct {
	ctx         *Context
	latitude    float64
	longitude   float64
	opts        *gotgbot.SendLocationOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the location to be sent after the specified duration.
func (c *SendLocation) After(duration time.Duration) *SendLocation {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the location message to be deleted after the specified duration.
func (c *SendLocation) DeleteAfter(duration time.Duration) *SendLocation {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the location message.
func (c *SendLocation) Silent() *SendLocation {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the location message.
func (c *SendLocation) Protect() *SendLocation {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the location message.
func (c *SendLocation) Markup(kb keyboard.KeyboardBuilder) *SendLocation {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// LivePeriod sets the period in seconds for which the location will be updated.
func (c *SendLocation) LivePeriod(seconds int64) *SendLocation {
	c.opts.LivePeriod = seconds
	return c
}

// Heading sets the direction in which the user is moving, in degrees.
func (c *SendLocation) Heading(heading int64) *SendLocation {
	c.opts.Heading = heading
	return c
}

// ProximityAlertRadius sets the maximum distance for proximity alerts about approaching another chat member.
func (c *SendLocation) ProximityAlertRadius(radius int64) *SendLocation {
	c.opts.ProximityAlertRadius = radius
	return c
}

// HorizontalAccuracy sets the radius of uncertainty for the location, measured in meters.
func (c *SendLocation) HorizontalAccuracy(accuracy float64) *SendLocation {
	c.opts.HorizontalAccuracy = accuracy
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendLocation) ReplyTo(messageID int64) *SendLocation {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendLocation) Timeout(duration time.Duration) *SendLocation {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendLocation) APIURL(url String) *SendLocation {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the location message.
func (c *SendLocation) Business(id String) *SendLocation {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the location message.
func (c *SendLocation) Thread(id int64) *SendLocation {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the location message.
func (c *SendLocation) To(chatID int64) *SendLocation {
	c.chatID = Some(chatID)
	return c
}

// Send sends the location message to Telegram and returns the result.
func (c *SendLocation) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(c.ctx.Bot.Raw().SendLocation(chatID, c.latitude, c.longitude, c.opts))
	})
}
