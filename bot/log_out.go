package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// LogOut represents a request to log out from the cloud Bot API server.
type LogOut struct {
	bot  *Bot
	opts *gotgbot.LogOutOpts
}

// Timeout sets a custom timeout for this request.
func (lo *LogOut) Timeout(duration time.Duration) *LogOut {
	if lo.opts.RequestOpts == nil {
		lo.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	lo.opts.RequestOpts.Timeout = duration

	return lo
}

// APIURL sets a custom API URL for this request.
func (lo *LogOut) APIURL(url String) *LogOut {
	if lo.opts.RequestOpts == nil {
		lo.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	lo.opts.RequestOpts.APIURL = url.Std()

	return lo
}

// Send logs out from the cloud Bot API server and returns the result.
func (lo *LogOut) Send() Result[bool] {
	return ResultOf(lo.bot.raw.LogOut(lo.opts))
}
