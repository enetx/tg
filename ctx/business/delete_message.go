package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// Delete is a request builder for deleting business messages.
type Delete struct {
	bot        Bot
	connID     g.String
	messageIDs g.Slice[int64]
	opts       *gotgbot.DeleteBusinessMessagesOpts
}

// Timeout sets a custom timeout for this request.
func (d *Delete) Timeout(duration time.Duration) *Delete {
	if d.opts.RequestOpts == nil {
		d.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	d.opts.RequestOpts.Timeout = duration

	return d
}

// APIURL sets a custom API URL for this request.
func (d *Delete) APIURL(url g.String) *Delete {
	if d.opts.RequestOpts == nil {
		d.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	d.opts.RequestOpts.APIURL = url.Std()

	return d
}

// Send executes the Delete request.
func (d *Delete) Send() g.Result[bool] {
	return g.ResultOf(d.bot.Raw().DeleteBusinessMessages(
		d.connID.Std(),
		d.messageIDs,
		d.opts,
	))
}
