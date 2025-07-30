package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// PinChatMessage represents a request to pin a message.
type PinChatMessage struct {
	ctx       *Context
	messageID int64
	opts      *gotgbot.PinChatMessageOpts
	chatID    g.Option[int64]
}

// ChatID sets the target chat ID for this request.
func (pcm *PinChatMessage) ChatID(id int64) *PinChatMessage {
	pcm.chatID = g.Some(id)
	return pcm
}

// Business sets the business connection ID for the pin action.
func (pcm *PinChatMessage) Business(id g.String) *PinChatMessage {
	pcm.opts.BusinessConnectionId = id.Std()
	return pcm
}

// Silent sets whether to disable notification.
func (pcm *PinChatMessage) Silent() *PinChatMessage {
	pcm.opts.DisableNotification = true
	return pcm
}

// Timeout sets a custom timeout for this request.
func (pcm *PinChatMessage) Timeout(duration time.Duration) *PinChatMessage {
	if pcm.opts.RequestOpts == nil {
		pcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pcm.opts.RequestOpts.Timeout = duration

	return pcm
}

// APIURL sets a custom API URL for this request.
func (pcm *PinChatMessage) APIURL(url g.String) *PinChatMessage {
	if pcm.opts.RequestOpts == nil {
		pcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pcm.opts.RequestOpts.APIURL = url.Std()

	return pcm
}

// Send executes the PinChatMessage request.
func (pcm *PinChatMessage) Send() g.Result[bool] {
	chatID := pcm.chatID.UnwrapOr(pcm.ctx.EffectiveChat.Id)
	return g.ResultOf(pcm.ctx.Bot.Raw().PinChatMessage(chatID, pcm.messageID, pcm.opts))
}
