package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetMe represents a request to get basic information about the bot.
type GetMe struct {
	bot  *Bot
	opts *gotgbot.GetMeOpts
}

// Timeout sets a custom timeout for this request.
func (gm *GetMe) Timeout(duration time.Duration) *GetMe {
	if gm.opts.RequestOpts == nil {
		gm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gm.opts.RequestOpts.Timeout = duration

	return gm
}

// APIURL sets a custom API URL for this request.
func (gm *GetMe) APIURL(url g.String) *GetMe {
	if gm.opts.RequestOpts == nil {
		gm.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gm.opts.RequestOpts.APIURL = url.Std()

	return gm
}

// Send gets basic information about the bot.
func (gm *GetMe) Send() g.Result[*gotgbot.User] {
	return g.ResultOf(gm.bot.Raw().GetMe(gm.opts))
}
