package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
)

type Venue struct {
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

func (v *Venue) After(duration time.Duration) *Venue {
	v.after = Some(duration)
	return v
}

func (v *Venue) DeleteAfter(duration time.Duration) *Venue {
	v.deleteAfter = Some(duration)
	return v
}

func (v *Venue) Silent() *Venue {
	v.opts.DisableNotification = true
	return v
}

func (v *Venue) Protect() *Venue {
	v.opts.ProtectContent = true
	return v
}

func (v *Venue) Markup(kb keyboard.KeyboardBuilder) *Venue {
	v.opts.ReplyMarkup = kb.Markup()
	return v
}

func (v *Venue) FoursquareID(id String) *Venue {
	v.opts.FoursquareId = id.Std()
	return v
}

func (v *Venue) FoursquareType(venueType String) *Venue {
	v.opts.FoursquareType = venueType.Std()
	return v
}

func (v *Venue) GooglePlaceID(id String) *Venue {
	v.opts.GooglePlaceId = id.Std()
	return v
}

func (v *Venue) GooglePlaceType(placeType String) *Venue {
	v.opts.GooglePlaceType = placeType.Std()
	return v
}

func (v *Venue) ReplyTo(messageID int64) *Venue {
	v.opts.ReplyParameters = &gotgbot.ReplyParameters{MessageId: messageID}
	return v
}

func (v *Venue) Timeout(duration time.Duration) *Venue {
	v.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return v
}

func (v *Venue) Business(id String) *Venue {
	v.opts.BusinessConnectionId = id.Std()
	return v
}

func (v *Venue) Thread(id int64) *Venue {
	v.opts.MessageThreadId = id
	return v
}

func (v *Venue) To(chatID int64) *Venue {
	v.chatID = Some(chatID)
	return v
}

func (v *Venue) Send() Result[*gotgbot.Message] {
	return v.ctx.timers(v.after, v.deleteAfter, func() Result[*gotgbot.Message] {
		chatID := v.chatID.UnwrapOr(v.ctx.EffectiveChat.Id)
		return ResultOf(
			v.ctx.Bot.Raw().SendVenue(chatID, v.latitude, v.longitude, v.title.Std(), v.address.Std(), v.opts),
		)
	})
}
