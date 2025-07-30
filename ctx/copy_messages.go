package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// CopyMessages represents a request to copy multiple messages.
type CopyMessages struct {
	ctx        *Context
	chatID     g.Option[int64]
	fromChatID g.Option[int64]
	messageIDs g.Slice[int64]
	opts       *gotgbot.CopyMessagesOpts
}

// To sets the target chat ID for copying messages.
func (cm *CopyMessages) To(chatID int64) *CopyMessages {
	cm.chatID = g.Some(chatID)
	return cm
}

// From sets the source chat ID where messages are copied from.
func (cm *CopyMessages) From(fromChatID int64) *CopyMessages {
	cm.fromChatID = g.Some(fromChatID)
	return cm
}

// MessageIDs sets the message IDs to copy (up to 100 messages, must be in increasing order).
func (cm *CopyMessages) MessageIDs(ids []int64) *CopyMessages {
	cm.messageIDs = ids
	return cm
}

// AddMessages adds multiple message IDs to the copy list.
func (cm *CopyMessages) AddMessages(ids ...int64) *CopyMessages {
	cm.messageIDs.Push(ids...)
	return cm
}

// Thread sets the message thread ID for forum supergroups.
func (cm *CopyMessages) Thread(threadID int64) *CopyMessages {
	cm.opts.MessageThreadId = threadID
	return cm
}

// Silent sends the messages silently (no notification sound).
func (cm *CopyMessages) Silent() *CopyMessages {
	cm.opts.DisableNotification = true
	return cm
}

// Protect protects the copied messages from forwarding and saving.
func (cm *CopyMessages) Protect() *CopyMessages {
	cm.opts.ProtectContent = true
	return cm
}

// RemoveCaption copies the messages without their captions.
func (cm *CopyMessages) RemoveCaption() *CopyMessages {
	cm.opts.RemoveCaption = true
	return cm
}

// Timeout sets a custom timeout for this request.
func (cm *CopyMessages) Timeout(duration time.Duration) *CopyMessages {
	if cm.opts.RequestOpts == nil {
		cm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cm.opts.RequestOpts.Timeout = duration

	return cm
}

// APIURL sets a custom API URL for this request.
func (cm *CopyMessages) APIURL(url g.String) *CopyMessages {
	if cm.opts.RequestOpts == nil {
		cm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cm.opts.RequestOpts.APIURL = url.Std()

	return cm
}

// Send copies the messages and returns the array of sent message IDs.
func (cm *CopyMessages) Send() g.Result[g.Slice[gotgbot.MessageId]] {
	if cm.messageIDs.Empty() {
		return g.Err[g.Slice[gotgbot.MessageId]](g.Errorf("no message IDs specified for copying"))
	}

	if cm.messageIDs.Len() > 100 {
		return g.Err[g.Slice[gotgbot.MessageId]](
			g.Errorf("too many message IDs: {} (maximum 100)", cm.messageIDs.Len()),
		)
	}

	if cm.fromChatID.IsNone() {
		return g.Err[g.Slice[gotgbot.MessageId]](g.Errorf("source chat ID must be specified"))
	}

	chatID := cm.chatID.UnwrapOr(cm.ctx.EffectiveChat.Id)
	fromChatID := cm.fromChatID.Some()

	result, err := cm.ctx.Bot.Raw().CopyMessages(chatID, fromChatID, cm.messageIDs, cm.opts)

	return g.ResultOf[g.Slice[gotgbot.MessageId]](result, err)
}
