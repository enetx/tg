package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/suggested"
	"github.com/enetx/tg/types/effects"
)

type ForwardMessage struct {
	ctx         *Context
	fromChatID  int64
	messageID   int64
	opts        *gotgbot.ForwardMessageOpts
	toChatID    g.Option[int64]
	after       g.Option[time.Duration]
	deleteAfter g.Option[time.Duration]
}

// After schedules the forward to be sent after the specified duration.
func (fm *ForwardMessage) After(duration time.Duration) *ForwardMessage {
	fm.after = g.Some(duration)
	return fm
}

// DeleteAfter schedules the forwarded message to be deleted after the specified duration.
func (fm *ForwardMessage) DeleteAfter(duration time.Duration) *ForwardMessage {
	fm.deleteAfter = g.Some(duration)
	return fm
}

// Silent disables notification for the forwarded message.
func (fm *ForwardMessage) Silent() *ForwardMessage {
	fm.opts.DisableNotification = true
	return fm
}

// Protect enables content protection for the forwarded message.
func (fm *ForwardMessage) Protect() *ForwardMessage {
	fm.opts.ProtectContent = true
	return fm
}

// Effect sets a message effect for the forwarded message.
func (fm *ForwardMessage) Effect(effect effects.EffectType) *ForwardMessage {
	fm.opts.MessageEffectId = effect.String()
	return fm
}

// Timeout sets a custom timeout for this request.
func (fm *ForwardMessage) Timeout(duration time.Duration) *ForwardMessage {
	if fm.opts.RequestOpts == nil {
		fm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	fm.opts.RequestOpts.Timeout = duration

	return fm
}

// APIURL sets a custom API URL for this request.
func (fm *ForwardMessage) APIURL(url g.String) *ForwardMessage {
	if fm.opts.RequestOpts == nil {
		fm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	fm.opts.RequestOpts.APIURL = url.Std()

	return fm
}

// SuggestedPost sets suggested post parameters for direct messages chats.
func (fm *ForwardMessage) SuggestedPost(params *suggested.PostParameters) *ForwardMessage {
	if params != nil {
		fm.opts.SuggestedPostParameters = params.Std()
	}
	return fm
}

// To sets the target chat ID for the forwarded message.
func (fm *ForwardMessage) To(chatID int64) *ForwardMessage {
	fm.toChatID = g.Some(chatID)
	return fm
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (fm *ForwardMessage) DirectMessagesTopic(topicID int64) *ForwardMessage {
	fm.opts.DirectMessagesTopicId = topicID
	return fm
}

// Send forwards the message to the target chat and returns the result.
func (fm *ForwardMessage) Send() g.Result[*gotgbot.Message] {
	return fm.ctx.timers(fm.after, fm.deleteAfter, func() g.Result[*gotgbot.Message] {
		chatID := fm.toChatID.UnwrapOr(fm.ctx.EffectiveChat.Id)
		return g.ResultOf(fm.ctx.Bot.Raw().ForwardMessage(chatID, fm.fromChatID, fm.messageID, fm.opts))
	})
}
