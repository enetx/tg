package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Read is a request builder for marking a business message as read.
type Read struct {
	bot       Bot
	connID    String
	chatID    int64
	messageID int64
	opts      *gotgbot.ReadBusinessMessageOpts
}

// Timeout sets a custom timeout for this request.
func (r *Read) Timeout(duration time.Duration) *Read {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *Read) APIURL(url String) *Read {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send executes the Read request.
func (r *Read) Send() Result[bool] {
	return ResultOf(r.bot.Raw().ReadBusinessMessage(
		r.connID.Std(),
		r.chatID,
		r.messageID,
		r.opts,
	))
}
