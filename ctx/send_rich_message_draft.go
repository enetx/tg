package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/rich"
)

// SendRichMessageDraft represents a request to send a rich message draft.
// Message drafts allow partial messages to be streamed to a user while being generated.
type SendRichMessageDraft struct {
	ctx     *Context
	message *rich.Message
	chatID  g.Option[int64]
	draftID int64
	opts    *gotgbot.SendRichMessageDraftOpts
}

// Message sets the rich message content of the draft.
func (srmd *SendRichMessageDraft) Message(message *rich.Message) *SendRichMessageDraft {
	srmd.message = message
	return srmd
}

// To sets the target chat ID for the message draft.
func (srmd *SendRichMessageDraft) To(chatID int64) *SendRichMessageDraft {
	srmd.chatID = g.Some(chatID)
	return srmd
}

// Thread sets the message thread ID for the message draft.
func (srmd *SendRichMessageDraft) Thread(id int64) *SendRichMessageDraft {
	srmd.opts.MessageThreadId = id
	return srmd
}

// Timeout sets a custom timeout for this request.
func (srmd *SendRichMessageDraft) Timeout(duration time.Duration) *SendRichMessageDraft {
	if srmd.opts.RequestOpts == nil {
		srmd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	srmd.opts.RequestOpts.Timeout = duration

	return srmd
}

// APIURL sets a custom API URL for this request.
func (srmd *SendRichMessageDraft) APIURL(url g.String) *SendRichMessageDraft {
	if srmd.opts.RequestOpts == nil {
		srmd.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	srmd.opts.RequestOpts.APIURL = url.Std()

	return srmd
}

// Send sends the rich message draft to Telegram and returns the result.
func (srmd *SendRichMessageDraft) Send() g.Result[bool] {
	chatID := srmd.chatID.UnwrapOr(srmd.ctx.EffectiveChat.Id)
	return g.ResultOf(srmd.ctx.Bot.Raw().SendRichMessageDraft(
		chatID,
		srmd.draftID,
		srmd.message.Std(),
		srmd.opts,
	))
}
