package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// DeleteWebhook represents a request to remove webhook integration.
type DeleteWebhook struct {
	bot  *Bot
	opts *gotgbot.DeleteWebhookOpts
}

// DropPendingUpdates instructs to drop all pending updates.
func (dw *DeleteWebhook) DropPendingUpdates() *DeleteWebhook {
	dw.opts.DropPendingUpdates = true
	return dw
}

// Timeout sets a custom timeout for this request.
func (dw *DeleteWebhook) Timeout(duration time.Duration) *DeleteWebhook {
	if dw.opts.RequestOpts == nil {
		dw.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dw.opts.RequestOpts.Timeout = duration

	return dw
}

// APIURL sets a custom API URL for this request.
func (dw *DeleteWebhook) APIURL(url g.String) *DeleteWebhook {
	if dw.opts.RequestOpts == nil {
		dw.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	dw.opts.RequestOpts.APIURL = url.Std()

	return dw
}

// Send removes the webhook integration.
func (dw *DeleteWebhook) Send() g.Result[bool] {
	return g.ResultOf(dw.bot.Raw().DeleteWebhook(dw.opts))
}
