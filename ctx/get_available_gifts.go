package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// GetAvailableGifts is a request builder for getting available gifts.
type GetAvailableGifts struct {
	ctx  *Context
	opts *gotgbot.GetAvailableGiftsOpts
}

// Timeout sets a custom timeout for this request.
func (gags *GetAvailableGifts) Timeout(duration time.Duration) *GetAvailableGifts {
	if gags.opts.RequestOpts == nil {
		gags.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gags.opts.RequestOpts.Timeout = duration

	return gags
}

// APIURL sets a custom API URL for this request.
func (gags *GetAvailableGifts) APIURL(url String) *GetAvailableGifts {
	if gags.opts.RequestOpts == nil {
		gags.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gags.opts.RequestOpts.APIURL = url.Std()

	return gags
}

// Send executes the GetAvailableGifts request.
func (gags *GetAvailableGifts) Send() Result[*gotgbot.Gifts] {
	return ResultOf(gags.ctx.Bot.Raw().GetAvailableGifts(gags.opts))
}
