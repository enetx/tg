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
func (sl *SendLocation) After(duration time.Duration) *SendLocation {
	sl.after = Some(duration)
	return sl
}

// DeleteAfter schedules the location message to be deleted after the specified duration.
func (sl *SendLocation) DeleteAfter(duration time.Duration) *SendLocation {
	sl.deleteAfter = Some(duration)
	return sl
}

// Silent disables notification for the location message.
func (sl *SendLocation) Silent() *SendLocation {
	sl.opts.DisableNotification = true
	return sl
}

// Protect enables content protection for the location message.
func (sl *SendLocation) Protect() *SendLocation {
	sl.opts.ProtectContent = true
	return sl
}

// Markup sets the reply markup keyboard for the location message.
func (sl *SendLocation) Markup(kb keyboard.KeyboardBuilder) *SendLocation {
	sl.opts.ReplyMarkup = kb.Markup()
	return sl
}

// LivePeriod sets the period in seconds for which the location will be updated.
func (sl *SendLocation) LivePeriod(seconds int64) *SendLocation {
	sl.opts.LivePeriod = seconds
	return sl
}

// Heading sets the direction in which the user is moving, in degrees.
func (sl *SendLocation) Heading(heading int64) *SendLocation {
	sl.opts.Heading = heading
	return sl
}

// ProximityAlertRadius sets the maximum distance for proximity alerts about approaching another chat member.
func (sl *SendLocation) ProximityAlertRadius(radius int64) *SendLocation {
	sl.opts.ProximityAlertRadius = radius
	return sl
}

// HorizontalAccuracy sets the radius of uncertainty for the location, measured in meters.
func (sl *SendLocation) HorizontalAccuracy(accuracy float64) *SendLocation {
	sl.opts.HorizontalAccuracy = accuracy
	return sl
}

// ReplyTo sets the message ID to reply to.
func (sl *SendLocation) ReplyTo(messageID int64) *SendLocation {
	sl.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return sl
}

// Timeout sets a custom timeout for this request.
func (sl *SendLocation) Timeout(duration time.Duration) *SendLocation {
	if sl.opts.RequestOpts == nil {
		sl.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sl.opts.RequestOpts.Timeout = duration

	return sl
}

// APIURL sets a custom API URL for this request.
func (sl *SendLocation) APIURL(url String) *SendLocation {
	if sl.opts.RequestOpts == nil {
		sl.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sl.opts.RequestOpts.APIURL = url.Std()

	return sl
}

// Business sets the business connection ID for the location message.
func (sl *SendLocation) Business(id String) *SendLocation {
	sl.opts.BusinessConnectionId = id.Std()
	return sl
}

// Thread sets the message thread ID for the location message.
func (sl *SendLocation) Thread(id int64) *SendLocation {
	sl.opts.MessageThreadId = id
	return sl
}

// To sets the target chat ID for the location message.
func (sl *SendLocation) To(chatID int64) *SendLocation {
	sl.chatID = Some(chatID)
	return sl
}

// Send sends the location message to Telegram and returns the result.
func (sl *SendLocation) Send() Result[*gotgbot.Message] {
	return sl.ctx.timers(sl.after, sl.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sl.chatID.UnwrapOr(sl.ctx.EffectiveChat.Id)
		return ResultOf(sl.ctx.Bot.Raw().SendLocation(chatID, sl.latitude, sl.longitude, sl.opts))
	})
}
