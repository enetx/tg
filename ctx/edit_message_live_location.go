package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
	"github.com/enetx/g/ref"
	"github.com/enetx/tg/keyboard"
)

// EditMessageLiveLocation represents a request to edit a live location message.
type EditMessageLiveLocation struct {
	ctx       *Context
	latitude  float64
	longitude float64
	opts      *gotgbot.EditMessageLiveLocationOpts
	chatID    Option[int64]
	messageID Option[int64]
}

// ChatID sets the target chat ID.
func (emll *EditMessageLiveLocation) ChatID(chatID int64) *EditMessageLiveLocation {
	emll.chatID = Some(chatID)
	return emll
}

// MessageID sets the target message ID.
func (emll *EditMessageLiveLocation) MessageID(messageID int64) *EditMessageLiveLocation {
	emll.messageID = Some(messageID)
	return emll
}

// InlineMessageID sets the inline message ID to edit.
func (emll *EditMessageLiveLocation) InlineMessageID(id String) *EditMessageLiveLocation {
	emll.opts.InlineMessageId = id.Std()
	return emll
}

// Business sets the business connection ID for the location edit.
func (emll *EditMessageLiveLocation) Business(id String) *EditMessageLiveLocation {
	emll.opts.BusinessConnectionId = id.Std()
	return emll
}

// LiveFor sets how long the location can be updated.
func (emll *EditMessageLiveLocation) LiveFor(duration time.Duration) *EditMessageLiveLocation {
	emll.opts.LivePeriod = ref.Of(int64(duration.Seconds()))
	return emll
}

// HorizontalAccuracy sets the radius of uncertainty for the location.
func (emll *EditMessageLiveLocation) HorizontalAccuracy(accuracy float64) *EditMessageLiveLocation {
	emll.opts.HorizontalAccuracy = accuracy
	return emll
}

// Heading sets the direction in which the user is moving.
func (emll *EditMessageLiveLocation) Heading(heading int64) *EditMessageLiveLocation {
	emll.opts.Heading = heading
	return emll
}

// ProximityAlertRadius sets the proximity alert radius.
func (emll *EditMessageLiveLocation) ProximityAlertRadius(radius int64) *EditMessageLiveLocation {
	emll.opts.ProximityAlertRadius = radius
	return emll
}

// Markup sets the reply markup keyboard for the message.
func (emll *EditMessageLiveLocation) Markup(kb *keyboard.InlineKeyboard) *EditMessageLiveLocation {
	if markup, ok := kb.Markup().(gotgbot.InlineKeyboardMarkup); ok {
		emll.opts.ReplyMarkup = markup
	}

	return emll
}

// Timeout sets a custom timeout for this request.
func (emll *EditMessageLiveLocation) Timeout(duration time.Duration) *EditMessageLiveLocation {
	if emll.opts.RequestOpts == nil {
		emll.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emll.opts.RequestOpts.Timeout = duration

	return emll
}

// APIURL sets a custom API URL for this request.
func (emll *EditMessageLiveLocation) APIURL(url String) *EditMessageLiveLocation {
	if emll.opts.RequestOpts == nil {
		emll.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	emll.opts.RequestOpts.APIURL = url.Std()

	return emll
}

// Send edits the live location message.
func (emll *EditMessageLiveLocation) Send() Result[*gotgbot.Message] {
	emll.opts.ChatId = emll.chatID.UnwrapOr(emll.ctx.EffectiveChat.Id)
	emll.opts.MessageId = emll.messageID.UnwrapOr(emll.ctx.EffectiveMessage.MessageId)
	msg, _, err := emll.ctx.Bot.Raw().EditMessageLiveLocation(emll.latitude, emll.longitude, emll.opts)

	return ResultOf(msg, err)
}
