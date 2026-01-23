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

type SendDice struct {
	ctx         *Context
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	opts        *gotgbot.SendDiceOpts
}

// After schedules the dice to be sent after the specified duration.
func (sd *SendDice) After(duration time.Duration) *SendDice {
	sd.after = g.Some(duration)
	return sd
}

// DeleteAfter schedules the dice message to be deleted after the specified duration.
func (sd *SendDice) DeleteAfter(duration time.Duration) *SendDice {
	sd.deleteAfter = g.Some(duration)
	return sd
}

// Emoji sets a custom emoji for the dice.
func (sd *SendDice) Emoji(e g.String) *SendDice {
	sd.opts.Emoji = e.Std()
	return sd
}

// Dart sets the dice emoji to dart.
func (sd *SendDice) Dart() *SendDice {
	sd.opts.Emoji = "🎯"
	return sd
}

// Slot sets the dice emoji to slot machine.
func (sd *SendDice) Slot() *SendDice {
	sd.opts.Emoji = "🎰"
	return sd
}

// Ball sets the dice emoji to basketball.
func (sd *SendDice) Ball() *SendDice {
	sd.opts.Emoji = "🏀"
	return sd
}

// Soccer sets the dice emoji to soccer ball.
func (sd *SendDice) Soccer() *SendDice {
	sd.opts.Emoji = "⚽"
	return sd
}

// Bowling sets the dice emoji to bowling.
func (sd *SendDice) Bowling() *SendDice {
	sd.opts.Emoji = "🎳"
	return sd
}

// Silent disables notification for the dice message.
func (sd *SendDice) Silent() *SendDice {
	sd.opts.DisableNotification = true
	return sd
}

// Thread sets the message thread ID for the dice message.
func (sd *SendDice) Thread(id int64) *SendDice {
	sd.opts.MessageThreadId = id
	return sd
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (sd *SendDice) AllowPaidBroadcast() *SendDice {
	sd.opts.AllowPaidBroadcast = true
	return sd
}

// Effect sets a message effect for the dice message.
func (sd *SendDice) Effect(effect effects.EffectType) *SendDice {
	sd.opts.MessageEffectId = effect.String()
	return sd
}

// Reply sets reply parameters using the reply builder.
func (sd *SendDice) Reply(params *reply.Parameters) *SendDice {
	if params != nil {
		sd.opts.ReplyParameters = params.Std()
	}
	return sd
}

// Markup sets the reply markup keyboard for the dice message.
func (sd *SendDice) Markup(kb keyboard.Keyboard) *SendDice {
	sd.opts.ReplyMarkup = kb.Markup()
	return sd
}

// Business sets the business connection ID for the dice message.
func (sd *SendDice) Business(id g.String) *SendDice {
	sd.opts.BusinessConnectionId = id.Std()
	return sd
}

// Protect enables content protection for the dice message.
func (sd *SendDice) Protect() *SendDice {
	sd.opts.ProtectContent = true
	return sd
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (sd *SendDice) SuggestedPost(params *suggested.PostParameters) *SendDice {
	if params != nil {
		sd.opts.SuggestedPostParameters = params.Std()
	}
	return sd
}

// To sets the target chat ID for the dice message.
func (sd *SendDice) To(chatID int64) *SendDice {
	sd.chatID = g.Some(chatID)
	return sd
}

// Timeout sets a custom timeout for this request.
func (sd *SendDice) Timeout(duration time.Duration) *SendDice {
	if sd.opts.RequestOpts == nil {
		sd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sd.opts.RequestOpts.Timeout = duration

	return sd
}

// APIURL sets a custom API URL for this request.
func (sd *SendDice) APIURL(url g.String) *SendDice {
	if sd.opts.RequestOpts == nil {
		sd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sd.opts.RequestOpts.APIURL = url.Std()

	return sd
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (sd *SendDice) DirectMessagesTopic(topicID int64) *SendDice {
	sd.opts.DirectMessagesTopicId = topicID
	return sd
}

// Send sends the dice message to Telegram and returns the result.
func (sd *SendDice) Send() g.Result[*gotgbot.Message] {
	return sd.ctx.timers(sd.after, sd.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := sd.chatID.UnwrapOr(sd.ctx.EffectiveChat.Id)
		return g.ResultOf(sd.ctx.Bot.Raw().SendDice(chatID, sd.opts))
	})
}
