package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type SendVenue struct {
	ctx         *Context
	latitude    float64
	longitude   float64
	title       String
	address     String
	opts        *gotgbot.SendVenueOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the venue to be sent after the specified duration.
func (c *SendVenue) After(duration time.Duration) *SendVenue {
	c.after = Some(duration)
	return c
}

// DeleteAfter schedules the venue message to be deleted after the specified duration.
func (c *SendVenue) DeleteAfter(duration time.Duration) *SendVenue {
	c.deleteAfter = Some(duration)
	return c
}

// Silent disables notification for the venue message.
func (c *SendVenue) Silent() *SendVenue {
	c.opts.DisableNotification = true
	return c
}

// Protect enables content protection for the venue message.
func (c *SendVenue) Protect() *SendVenue {
	c.opts.ProtectContent = true
	return c
}

// Markup sets the reply markup keyboard for the venue message.
func (c *SendVenue) Markup(kb keyboard.KeyboardBuilder) *SendVenue {
	c.opts.ReplyMarkup = kb.Markup()
	return c
}

// FoursquareID sets the Foursquare identifier of the venue.
func (c *SendVenue) FoursquareID(id String) *SendVenue {
	c.opts.FoursquareId = id.Std()
	return c
}

// FoursquareType sets the Foursquare type of the venue.
func (c *SendVenue) FoursquareType(venueType String) *SendVenue {
	c.opts.FoursquareType = venueType.Std()
	return c
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (c *SendVenue) GooglePlaceID(id String) *SendVenue {
	c.opts.GooglePlaceId = id.Std()
	return c
}

// GooglePlaceType sets the Google Places type of the venue.
func (c *SendVenue) GooglePlaceType(placeType String) *SendVenue {
	c.opts.GooglePlaceType = placeType.Std()
	return c
}

// ReplyTo sets the message ID to reply to.
func (c *SendVenue) ReplyTo(messageID int64) *SendVenue {
	c.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return c
}

// Timeout sets a custom timeout for this request.
func (c *SendVenue) Timeout(duration time.Duration) *SendVenue {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *SendVenue) APIURL(url String) *SendVenue {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Business sets the business connection ID for the venue message.
func (c *SendVenue) Business(id String) *SendVenue {
	c.opts.BusinessConnectionId = id.Std()
	return c
}

// Thread sets the message thread ID for the venue message.
func (c *SendVenue) Thread(id int64) *SendVenue {
	c.opts.MessageThreadId = id
	return c
}

// To sets the target chat ID for the venue message.
func (c *SendVenue) To(chatID int64) *SendVenue {
	c.chatID = Some(chatID)
	return c
}

// Send sends the venue message to Telegram and returns the result.
func (c *SendVenue) Send() Result[*gotgbot.Message] {
	return c.ctx.timers(c.after, c.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := c.chatID.UnwrapOr(c.ctx.EffectiveChat.Id)
		return ResultOf(
			c.ctx.Bot.Raw().SendVenue(chatID, c.latitude, c.longitude, c.title.Std(), c.address.Std(), c.opts),
		)
	})
}
