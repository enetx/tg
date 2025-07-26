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
func (sv *SendVenue) After(duration time.Duration) *SendVenue {
	sv.after = Some(duration)
	return sv
}

// DeleteAfter schedules the venue message to be deleted after the specified duration.
func (sv *SendVenue) DeleteAfter(duration time.Duration) *SendVenue {
	sv.deleteAfter = Some(duration)
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

// Markup sets the reply markup keyboard for the venue message.
func (sv *SendVenue) Markup(kb keyboard.Keyboard) *SendVenue {
	sv.opts.ReplyMarkup = kb.Markup()
	return sv
}

// FoursquareID sets the Foursquare identifier of the venue.
func (sv *SendVenue) FoursquareID(id String) *SendVenue {
	sv.opts.FoursquareId = id.Std()
	return sv
}

// FoursquareType sets the Foursquare type of the venue.
func (sv *SendVenue) FoursquareType(venueType String) *SendVenue {
	sv.opts.FoursquareType = venueType.Std()
	return sv
}

// GooglePlaceID sets the Google Places identifier of the venue.
func (sv *SendVenue) GooglePlaceID(id String) *SendVenue {
	sv.opts.GooglePlaceId = id.Std()
	return sv
}

// GooglePlaceType sets the Google Places type of the venue.
func (sv *SendVenue) GooglePlaceType(placeType String) *SendVenue {
	sv.opts.GooglePlaceType = placeType.Std()
	return sv
}

// ReplyTo sets the message ID to reply to.
func (sv *SendVenue) ReplyTo(messageID int64) *SendVenue {
	sv.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
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
func (sv *SendVenue) APIURL(url String) *SendVenue {
	if sv.opts.RequestOpts == nil {
		sv.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sv.opts.RequestOpts.APIURL = url.Std()

	return sv
}

// Business sets the business connection ID for the venue message.
func (sv *SendVenue) Business(id String) *SendVenue {
	sv.opts.BusinessConnectionId = id.Std()
	return sv
}

// Thread sets the message thread ID for the venue message.
func (sv *SendVenue) Thread(id int64) *SendVenue {
	sv.opts.MessageThreadId = id
	return sv
}

// To sets the target chat ID for the venue message.
func (sv *SendVenue) To(chatID int64) *SendVenue {
	sv.chatID = Some(chatID)
	return sv
}

// Send sends the venue message to Telegram and returns the result.
func (sv *SendVenue) Send() Result[*gotgbot.Message] {
	return sv.ctx.timers(sv.after, sv.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := sv.chatID.UnwrapOr(sv.ctx.EffectiveChat.Id)
		return ResultOf(
			sv.ctx.Bot.Raw().SendVenue(chatID, sv.latitude, sv.longitude, sv.title.Std(), sv.address.Std(), sv.opts),
		)
	})
}
