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

type SendVenue struct {
	ctx         *Context
	latitude    float64
	longitude   float64
	title       g.String
	address     g.String
	opts        *gotgbot.SendVenueOpts
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// After schedules the venue to be sent after the specified duration.
func (sv *SendVenue) After(duration time.Duration) *SendVenue {
	sv.after = g.Some(duration)
	return sv
}

// DeleteAfter schedules the venue message to be deleted after the specified duration.
func (sv *SendVenue) DeleteAfter(duration time.Duration) *SendVenue {
	sv.deleteAfter = g.Some(duration)
	return sv
}

// Silent disables notification for the venue message.
func (sv *SendVenue) Silent() *SendVenue {
	sv.opts.DisableNotification = true
	return sv
}

// Protect enables content protection for the venue message.
func (sv *SendVenue) Protect() *SendVenue {
	sv.opts.ProtectContent = true
	return sv
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sv *SendVenue) AllowPaidBroadcast() *SendVenue {
	sv.opts.AllowPaidBroadcast = true
	return sv
}

// Effect sets a message effect for the message.
func (sv *SendVenue) Effect(effect effects.EffectType) *SendVenue {
	sv.opts.MessageEffectId = effect.String()
	return sv
}

// Markup sets the reply markup keyboard for the venue message.
func (sv *SendVenue) Markup(kb keyboard.Keyboard) *SendVenue {
	sv.opts.ReplyMarkup = kb.Markup()
	return sv
}

// FoursquareID sets the Foursquare identifier of the venue.
func (sv *SendVenue) FoursquareID(id g.String) *SendVenue {
	sv.opts.FoursquareId = id.Std()
	return sv
}

// FoursquareType sets the Foursquare type of the venue.
func (sv *SendVenue) FoursquareType(venueType g.String) *SendVenue {
	sv.opts.FoursquareType = venueType.Std()
	return sv
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (sv *SendVenue) GooglePlaceID(id g.String) *SendVenue {
	sv.opts.GooglePlaceId = id.Std()
	return sv
}

// GooglePlaceType sets the Google Places type of the venue.
func (sv *SendVenue) GooglePlaceType(placeType g.String) *SendVenue {
	sv.opts.GooglePlaceType = placeType.Std()
	return sv
}

// Reply sets reply parameters using the reply builder.
func (sv *SendVenue) Reply(params *reply.Parameters) *SendVenue {
	if params != nil {
		sv.opts.ReplyParameters = params.Std()
	}
	return sv
}

// Timeout sets a custom timeout for this request.
func (sv *SendVenue) Timeout(duration time.Duration) *SendVenue {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.Timeout = duration

	return sv
}

// APIURL sets a custom API URL for this request.
func (sv *SendVenue) APIURL(url g.String) *SendVenue {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.APIURL = url.Std()

	return sv
}

// Business sets the business connection ID for the venue message.
func (sv *SendVenue) Business(id g.String) *SendVenue {
	sv.opts.BusinessConnectionId = id.Std()
	return sv
}

// Thread sets the message thread ID for the venue message.
func (sv *SendVenue) Thread(id int64) *SendVenue {
	sv.opts.MessageThreadId = id
	return sv
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sv *SendVenue) SuggestedPost(params *suggested.PostParameters) *SendVenue {
	if params != nil {
		sv.opts.SuggestedPostParameters = params.Std()
	}
	return sv
}

// To sets the target chat ID for the venue message.
func (sv *SendVenue) To(chatID int64) *SendVenue {
	sv.chatID = g.Some(chatID)
	return sv
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sv *SendVenue) DirectMessagesTopic(topicID int64) *SendVenue {
	sv.opts.DirectMessagesTopicId = topicID
	return sv
}

// Send sends the venue message to Telegram and returns the result.
func (sv *SendVenue) Send() g.Result[*gotgbot.Message] {
	return sv.ctx.timers(sv.after, sv.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sv.chatID.UnwrapOr(sv.ctx.EffectiveChat.Id)
		return g.ResultOf(
			sv.ctx.Bot.Raw().SendVenue(chatID, sv.latitude, sv.longitude, sv.title.Std(), sv.address.Std(), sv.opts),
		)
	})
}
