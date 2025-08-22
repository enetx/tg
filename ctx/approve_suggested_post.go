package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ApproveSuggestedPost represents a request to approve a suggested post.
type ApproveSuggestedPost struct {
	ctx       *Context
	chatID    g.Option[int64]
	messageID g.Option[int64]
	opts      *gotgbot.ApproveSuggestedPostOpts
}

// ChatID sets the target direct messages chat ID.
func (asp *ApproveSuggestedPost) ChatID(chatID int64) *ApproveSuggestedPost {
	asp.chatID = g.Some(chatID)
	return asp
}

// MessageID sets the suggested post message ID to approve.
func (asp *ApproveSuggestedPost) MessageID(messageID int64) *ApproveSuggestedPost {
	asp.messageID = g.Some(messageID)
	return asp
}

// SendDate sets when the post is expected to be published (Unix timestamp).
func (asp *ApproveSuggestedPost) SendDate(sendTime time.Time) *ApproveSuggestedPost {
	asp.opts.SendDate = sendTime.Unix()
	return asp
}

// SendAfter sets when the post should be published relative to now.
func (asp *ApproveSuggestedPost) SendAfter(duration time.Duration) *ApproveSuggestedPost {
	asp.opts.SendDate = time.Now().Add(duration).Unix()
	return asp
}

// Timeout sets a custom timeout for this request.
func (asp *ApproveSuggestedPost) Timeout(duration time.Duration) *ApproveSuggestedPost {
	if asp.opts.RequestOpts == nil {
		asp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	asp.opts.RequestOpts.Timeout = duration

	return asp
}

// APIURL sets a custom API URL for this request.
func (asp *ApproveSuggestedPost) APIURL(url g.String) *ApproveSuggestedPost {
	if asp.opts.RequestOpts == nil {
		asp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	asp.opts.RequestOpts.APIURL = url.Std()

	return asp
}

// Send approves the suggested post.
func (asp *ApproveSuggestedPost) Send() g.Result[bool] {
	chatID := asp.chatID.UnwrapOr(asp.ctx.EffectiveChat.Id)
	messageID := asp.messageID.UnwrapOr(asp.ctx.EffectiveMessage.MessageId)

	return g.ResultOf(asp.ctx.Bot.Raw().ApproveSuggestedPost(chatID, messageID, asp.opts))
}
