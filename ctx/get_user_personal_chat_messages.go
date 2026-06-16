package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetUserPersonalChatMessages represents a request to fetch the last messages from
// the personal chat (i.e., the chat currently added to their profile) of a given user.
type GetUserPersonalChatMessages struct {
	ctx    *Context
	userID int64
	limit  int64
	opts   *gotgbot.GetUserPersonalChatMessagesOpts
}

// Timeout sets a custom timeout for this request.
func (gupcm *GetUserPersonalChatMessages) Timeout(duration time.Duration) *GetUserPersonalChatMessages {
	if gupcm.opts.RequestOpts == nil {
		gupcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupcm.opts.RequestOpts.Timeout = duration

	return gupcm
}

// APIURL sets a custom API URL for this request.
func (gupcm *GetUserPersonalChatMessages) APIURL(url g.String) *GetUserPersonalChatMessages {
	if gupcm.opts.RequestOpts == nil {
		gupcm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gupcm.opts.RequestOpts.APIURL = url.Std()

	return gupcm
}

// Send retrieves the personal chat messages and returns the result.
func (gupcm *GetUserPersonalChatMessages) Send() g.Result[g.Slice[gotgbot.Message]] {
	return g.ResultOf[g.Slice[gotgbot.Message]](
		gupcm.ctx.Bot.Raw().GetUserPersonalChatMessages(gupcm.userID, gupcm.limit, gupcm.opts),
	)
}
