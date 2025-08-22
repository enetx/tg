package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ForwardMessages represents a request to forward multiple messages.
type ForwardMessages struct {
	ctx        *Context
	chatID     g.Option[int64]
	fromChatID g.Option[int64]
	messageIDs g.Slice[int64]
	opts       *gotgbot.ForwardMessagesOpts
}

// To sets the target chat ID for forwarding messages.
func (fms *ForwardMessages) To(chatID int64) *ForwardMessages {
	fms.chatID = g.Some(chatID)
	return fms
}

// From sets the source chat ID where messages are forwarded from.
func (fms *ForwardMessages) From(fromChatID int64) *ForwardMessages {
	fms.fromChatID = g.Some(fromChatID)
	return fms
}

// MessageIDs sets the message IDs to forward (up to 100 messages, must be in increasing order).
func (fms *ForwardMessages) MessageIDs(ids []int64) *ForwardMessages {
	fms.messageIDs = ids
	return fms
}

// AddMessages adds multiple message IDs to the forward list.
func (fms *ForwardMessages) AddMessages(ids ...int64) *ForwardMessages {
	fms.messageIDs.Push(ids...)
	return fms
}

// Thread sets the message thread ID for forum supergroups.
func (fms *ForwardMessages) Thread(threadID int64) *ForwardMessages {
	fms.opts.MessageThreadId = threadID
	return fms
}

// Silent sends the messages silently (no notification sound).
func (fms *ForwardMessages) Silent() *ForwardMessages {
	fms.opts.DisableNotification = true
	return fms
}

// Protect protects the forwarded messages from forwarding and saving.
func (fms *ForwardMessages) Protect() *ForwardMessages {
	fms.opts.ProtectContent = true
	return fms
}

// Timeout sets a custom timeout for this request.
func (fms *ForwardMessages) Timeout(duration time.Duration) *ForwardMessages {
	if fms.opts.RequestOpts == nil {
		fms.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	fms.opts.RequestOpts.Timeout = duration

	return fms
}

// APIURL sets a custom API URL for this request.
func (fms *ForwardMessages) APIURL(url g.String) *ForwardMessages {
	if fms.opts.RequestOpts == nil {
		fms.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	fms.opts.RequestOpts.APIURL = url.Std()

	return fms
}

// DirectMessagesTopic sets the direct messages topic ID for the message.
func (fms *ForwardMessages) DirectMessagesTopic(topicID int64) *ForwardMessages {
	fms.opts.DirectMessagesTopicId = topicID
	return fms
}

// Send forwards the messages and returns the array of sent message IDs.
func (fms *ForwardMessages) Send() g.Result[g.Slice[gotgbot.MessageId]] {
	if fms.messageIDs.Empty() {
		return g.Err[g.Slice[gotgbot.MessageId]](g.Errorf("no message IDs specified for forwarding"))
	}

	if fms.messageIDs.Len() > 100 {
		return g.Err[g.Slice[gotgbot.MessageId]](
			g.Errorf("too many message IDs: {} (maximum 100)", fms.messageIDs.Len()),
		)
	}

	if fms.fromChatID.IsNone() {
		return g.Err[g.Slice[gotgbot.MessageId]](g.Errorf("source chat ID must be specified"))
	}

	chatID := fms.chatID.UnwrapOr(fms.ctx.EffectiveChat.Id)
	fromChatID := fms.fromChatID.Some()

	result, err := fms.ctx.Bot.Raw().ForwardMessages(chatID, fromChatID, fms.messageIDs, fms.opts)

	return g.ResultOf[g.Slice[gotgbot.MessageId]](result, err)
}
