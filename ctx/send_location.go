package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type SendLocation struct {
	ctx         *Context
	latitude    float64
	longitude   float64
	opts        *gotgbot.SendLocationOpts
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// After schedules the location to be sent after the specified duration.
func (sl *SendLocation) After(duration time.Duration) *SendLocation {
	sl.after = g.Some(duration)
	return sl
}

// DeleteAfter schedules the location message to be deleted after the specified duration.
func (sl *SendLocation) DeleteAfter(duration time.Duration) *SendLocation {
	sl.deleteAfter = g.Some(duration)
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

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sl *SendLocation) AllowPaidBroadcast() *SendLocation {
	sl.opts.AllowPaidBroadcast = true
	return sl
}

// Effect sets a message effect for the message.
func (sl *SendLocation) Effect(effect effects.EffectType) *SendLocation {
	sl.opts.MessageEffectId = effect.String()
	return sl
}

// Markup sets the reply markup keyboard for the location message.
func (sl *SendLocation) Markup(kb keyboard.Keyboard) *SendLocation {
	sl.opts.ReplyMarkup = kb.Markup()
	return sl
}

// LiveFor sets the period in seconds for which the location will be updated.
func (sl *SendLocation) LiveFor(duration time.Duration) *SendLocation {
	sl.opts.LivePeriod = int64(duration.Seconds())
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

// Reply sets reply parameters using the reply builder.
func (sl *SendLocation) Reply(params *reply.Parameters) *SendLocation {
	sl.opts.ReplyParameters = params.Std()
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
func (sl *SendLocation) APIURL(url g.String) *SendLocation {
	if sl.opts.RequestOpts == nil {
		sl.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sl.opts.RequestOpts.APIURL = url.Std()

	return sl
}

// Business sets the business connection ID for the location message.
func (sl *SendLocation) Business(id g.String) *SendLocation {
	sl.opts.BusinessConnectionId = id.Std()
	return sl
}

// Thread sets the message thread ID for the location message.
func (sl *SendLocation) Thread(id int64) *SendLocation {
	sl.opts.MessageThreadId = id
	return sl
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sl *SendLocation) SuggestedPost(params *suggested.PostParameters) *SendLocation {
	if params != nil {
		sl.opts.SuggestedPostParameters = params.Std()
	}
	return sl
}

// To sets the target chat ID for the location message.
func (sl *SendLocation) To(chatID int64) *SendLocation {
	sl.chatID = g.Some(chatID)
	return sl
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sl *SendLocation) DirectMessagesTopic(topicID int64) *SendLocation {
	sl.opts.DirectMessagesTopicId = topicID
	return sl
}

// Send sends the location message to Telegram and returns the result.
func (sl *SendLocation) Send() g.Result[*gotgbot.Message] {
	return sl.ctx.timers(sl.after, sl.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sl.chatID.UnwrapOr(sl.ctx.EffectiveChat.Id)
		return g.ResultOf(sl.ctx.Bot.Raw().SendLocation(chatID, sl.latitude, sl.longitude, sl.opts))
	})
}
