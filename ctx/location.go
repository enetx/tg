package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Location struct {
	ctx         *Context
	latitude    float64
	longitude   float64
	opts        *gotgbot.SendLocationOpts
	chatID      Option[int64]
	after       Option[time.Duration]
	deleteAfter Option[time.Duration]
}

// After schedules the location to be sent after the specified duration.
func (l *Location) After(duration time.Duration) *Location {
	l.after = Some(duration)
	return l
}

// DeleteAfter schedules the location message to be deleted after the specified duration.
func (l *Location) DeleteAfter(duration time.Duration) *Location {
	l.deleteAfter = Some(duration)
	return l
}

// Silent disables notification for the location message.
func (l *Location) Silent() *Location {
	l.opts.DisableNotification = true
	return l
}

// Protect enables content protection for the location message.
func (l *Location) Protect() *Location {
	l.opts.ProtectContent = true
	return l
}

// Markup sets the reply markup keyboard for the location message.
func (l *Location) Markup(kb keyboard.KeyboardBuilder) *Location {
	l.opts.ReplyMarkup = kb.Markup()
	return l
}

// LivePeriod sets the period in seconds for which the location will be updated.
func (l *Location) LivePeriod(seconds int64) *Location {
	l.opts.LivePeriod = seconds
	return l
}

// Heading sets the direction in which the user is moving, in degrees.
func (l *Location) Heading(heading int64) *Location {
	l.opts.Heading = heading
	return l
}

// ProximityAlertRadius sets the maximum distance for proximity alerts about approaching another chat member.
func (l *Location) ProximityAlertRadius(radius int64) *Location {
	l.opts.ProximityAlertRadius = radius
	return l
}

// HorizontalAccuracy sets the radius of uncertainty for the location, measured in meters.
func (l *Location) HorizontalAccuracy(accuracy float64) *Location {
	l.opts.HorizontalAccuracy = accuracy
	return l
}

// ReplyTo sets the message ID to reply to.
func (l *Location) ReplyTo(messageID int64) *Location {
	l.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return l
}

// Timeout sets the request timeout duration.
func (l *Location) Timeout(duration time.Duration) *Location {
	l.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return l
}

// Business sets the business connection ID for the location message.
func (l *Location) Business(id String) *Location {
	l.opts.BusinessConnectionId = id.Std()
	return l
}

// Thread sets the message thread ID for the location message.
func (l *Location) Thread(id int64) *Location {
	l.opts.MessageThreadId = id
	return l
}

// To sets the target chat ID for the location message.
func (l *Location) To(chatID int64) *Location {
	l.chatID = Some(chatID)
	return l
}

// Send sends the location message to Telegram and returns the result.
func (l *Location) Send() Result[*gotgbot.Message] {
	return l.ctx.timers(l.after, l.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := l.chatID.UnwrapOr(l.ctx.EffectiveChat.Id)
		return ResultOf(l.ctx.Bot.Raw().SendLocation(chatID, l.latitude, l.longitude, l.opts))
	})
}
