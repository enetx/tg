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

type SendContact struct {
	ctx         *Context
	phoneNumber g.String
	firstName   g.String
	opts        *gotgbot.SendContactOpts
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// After schedules the contact to be sent after the specified duration.
func (sc *SendContact) After(duration time.Duration) *SendContact {
	sc.after = g.Some(duration)
	return sc
}

// DeleteAfter schedules the contact message to be deleted after the specified duration.
func (sc *SendContact) DeleteAfter(duration time.Duration) *SendContact {
	sc.deleteAfter = g.Some(duration)
	return sc
}

// Silent disables notification for the contact message.
func (sc *SendContact) Silent() *SendContact {
	sc.opts.DisableNotification = true
	return sc
}

// Protect enables content protection for the contact message.
func (sc *SendContact) Protect() *SendContact {
	sc.opts.ProtectContent = true
	return sc
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sc *SendContact) AllowPaidBroadcast() *SendContact {
	sc.opts.AllowPaidBroadcast = true
	return sc
}

// Effect sets a message effect for the message.
func (sc *SendContact) Effect(effect effects.EffectType) *SendContact {
	sc.opts.MessageEffectId = effect.String()
	return sc
}

// Markup sets the reply markup keyboard for the contact message.
func (sc *SendContact) Markup(kb keyboard.Keyboard) *SendContact {
	sc.opts.ReplyMarkup = kb.Markup()
	return sc
}

// LastName sets the contact's last name.
func (sc *SendContact) LastName(lastName g.String) *SendContact {
	sc.opts.LastName = lastName.Std()
	return sc
}

// VCard sets additional contact information in vCard format.
func (sc *SendContact) VCard(vcard g.String) *SendContact {
	sc.opts.Vcard = vcard.Std()
	return sc
}

// Reply sets reply parameters using the reply builder.
func (sc *SendContact) Reply(params *reply.Parameters) *SendContact {
	if params != nil {
		sc.opts.ReplyParameters = params.Std()
	}
	return sc
}

// Timeout sets a custom timeout for this request.
func (sc *SendContact) Timeout(duration time.Duration) *SendContact {
	if sc.opts.RequestOpts == nil {
		sc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sc.opts.RequestOpts.Timeout = duration

	return sc
}

// APIURL sets a custom API URL for this request.
func (sc *SendContact) APIURL(url g.String) *SendContact {
	if sc.opts.RequestOpts == nil {
		sc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sc.opts.RequestOpts.APIURL = url.Std()

	return sc
}

// Business sets the business connection ID for the contact message.
func (sc *SendContact) Business(id g.String) *SendContact {
	sc.opts.BusinessConnectionId = id.Std()
	return sc
}

// Thread sets the message thread ID for the contact message.
func (sc *SendContact) Thread(id int64) *SendContact {
	sc.opts.MessageThreadId = id
	return sc
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sc *SendContact) SuggestedPost(params *suggested.PostParameters) *SendContact {
	if params != nil {
		sc.opts.SuggestedPostParameters = params.Std()
	}
	return sc
}

// To sets the target chat ID for the contact message.
func (sc *SendContact) To(chatID int64) *SendContact {
	sc.chatID = g.Some(chatID)
	return sc
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sc *SendContact) DirectMessagesTopic(topicID int64) *SendContact {
	sc.opts.DirectMessagesTopicId = topicID
	return sc
}

// Send sends the contact message to Telegram and returns the result.
func (sc *SendContact) Send() g.Result[*gotgbot.Message] {
	return sc.ctx.timers(sc.after, sc.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sc.chatID.UnwrapOr(sc.ctx.EffectiveChat.Id)
		return g.ResultOf(sc.ctx.Bot.Raw().SendContact(chatID, sc.phoneNumber.Std(), sc.firstName.Std(), sc.opts))
	})
}
