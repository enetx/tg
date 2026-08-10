package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/keyboard"
	"github.com/enetx/tg/reply"
	"github.com/enetx/tg/rich"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

// SendRichMessage represents a request to send a rich formatted message.
type SendRichMessage struct {
	ctx         *Context
	message     *rich.Message
	chatID      g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
	opts        *gotgbot.SendRichMessageOpts
}

// After schedules the message to be sent after the specified duration.
func (srm *SendRichMessage) After(duration time.Duration) *SendRichMessage {
	srm.after = g.Some(duration)
	return srm
}

// DeleteAfter schedules the message to be deleted after the specified duration.
func (srm *SendRichMessage) DeleteAfter(duration time.Duration) *SendRichMessage {
	srm.deleteAfter = g.Some(duration)
	return srm
}

// To sets the target chat ID for the message.
func (srm *SendRichMessage) To(chatID int64) *SendRichMessage {
	srm.chatID = g.Some(chatID)
	return srm
}

// Silent disables notification for the message.
func (srm *SendRichMessage) Silent() *SendRichMessage {
	srm.opts.DisableNotification = true
	return srm
}

// Protect enables content protection for the message.
func (srm *SendRichMessage) Protect() *SendRichMessage {
	srm.opts.ProtectContent = true
	return srm
}

// Effect sets a message effect for the message.
func (srm *SendRichMessage) Effect(effect effects.EffectType) *SendRichMessage {
	srm.opts.MessageEffectId = effect.String()
	return srm
}

// Reply sets reply parameters using the reply builder.
func (srm *SendRichMessage) Reply(params *reply.Parameters) *SendRichMessage {
	if params != nil {
		srm.opts.ReplyParameters = params.Std()
	}

	return srm
}

// Markup sets the reply markup keyboard for the message.
func (srm *SendRichMessage) Markup(kb keyboard.Keyboard) *SendRichMessage {
	srm.opts.ReplyMarkup = kb.Markup()
	return srm
}

// AllowPaidBroadcast allows the message to be sent in paid broadcast channels.
func (srm *SendRichMessage) AllowPaidBroadcast() *SendRichMessage {
	srm.opts.AllowPaidBroadcast = true
	return srm
}

// Thread sets the message thread ID for the message.
func (srm *SendRichMessage) Thread(id int64) *SendRichMessage {
	srm.opts.MessageThreadId = id
	return srm
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (srm *SendRichMessage) DirectMessagesTopic(topicID int64) *SendRichMessage {
	srm.opts.DirectMessagesTopicId = topicID
	return srm
}

// Business sets the business connection ID for the message.
func (srm *SendRichMessage) Business(id g.String) *SendRichMessage {
	srm.opts.BusinessConnectionId = id.Std()
	return srm
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (srm *SendRichMessage) SuggestedPost(params *suggested.PostParameters) *SendRichMessage {
	if params != nil {
		srm.opts.SuggestedPostParameters = params.Std()
	}

	return srm
}

// Timeout sets a custom timeout for this request.
func (srm *SendRichMessage) Timeout(duration time.Duration) *SendRichMessage {
	if srm.opts.RequestOpts == nil {
		srm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	srm.opts.RequestOpts.Timeout = duration

	return srm
}

// APIURL sets a custom API URL for this request.
func (srm *SendRichMessage) APIURL(url g.String) *SendRichMessage {
	if srm.opts.RequestOpts == nil {
		srm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	srm.opts.RequestOpts.APIURL = url.Std()

	return srm
}

// Send sends the rich message to Telegram and returns the result.
func (srm *SendRichMessage) Send() g.Result[*gotgbot.Message] {
	return srm.ctx.timers(srm.after, srm.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := srm.chatID.UnwrapOr(srm.ctx.EffectiveChat.Id)
		return g.ResultOf(srm.ctx.Bot.Raw().SendRichMessage(chatID, srm.message.Std(), srm.opts))
	})
}
