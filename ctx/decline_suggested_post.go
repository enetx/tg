package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeclineSuggestedPost represents a request to decline a suggested post.
type DeclineSuggestedPost struct {
	ctx       *Context
	chatID    g.Option[int64]
	messageID g.Option[int64]
	opts      *gotgbot.DeclineSuggestedPostOpts
}

// Comment sets the comment for the creator of the suggested post, 0-128 characters.
func (dsp *DeclineSuggestedPost) Comment(comment g.String) *DeclineSuggestedPost {
	dsp.opts.Comment = comment.Std()
	return dsp
}

// ChatID sets the target direct messages chat ID.
func (dsp *DeclineSuggestedPost) ChatID(chatID int64) *DeclineSuggestedPost {
	dsp.chatID = g.Some(chatID)
	return dsp
}

// MessageID sets the suggested post message ID to decline.
func (dsp *DeclineSuggestedPost) MessageID(messageID int64) *DeclineSuggestedPost {
	dsp.messageID = g.Some(messageID)
	return dsp
}

// Timeout sets a custom timeout for this request.
func (dsp *DeclineSuggestedPost) Timeout(duration time.Duration) *DeclineSuggestedPost {
	if dsp.opts.RequestOpts == nil {
		dsp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dsp.opts.RequestOpts.Timeout = duration

	return dsp
}

// APIURL sets a custom API URL for this request.
func (dsp *DeclineSuggestedPost) APIURL(url g.String) *DeclineSuggestedPost {
	if dsp.opts.RequestOpts == nil {
		dsp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dsp.opts.RequestOpts.APIURL = url.Std()

	return dsp
}

// Send declines the suggested post.
func (dsp *DeclineSuggestedPost) Send() g.Result[bool] {
	chatID := dsp.chatID.UnwrapOr(dsp.ctx.EffectiveChat.Id)
	messageID := dsp.messageID.UnwrapOr(dsp.ctx.EffectiveMessage.MessageId)

	return g.ResultOf(dsp.ctx.Bot.Raw().DeclineSuggestedPost(chatID, messageID, dsp.opts))
}
