package bot

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetWebhookInfo represents a request to get current webhook status.
type GetWebhookInfo struct {
	bot  *Bot
	opts *gotgbot.GetWebhookInfoOpts
}

// Timeout sets a custom timeout for this request.
func (gwi *GetWebhookInfo) Timeout(duration time.Duration) *GetWebhookInfo {
	if gwi.opts.RequestOpts == nil {
		gwi.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gwi.opts.RequestOpts.Timeout = duration

	return gwi
}

// APIURL sets a custom API URL for this request.
func (gwi *GetWebhookInfo) APIURL(url g.String) *GetWebhookInfo {
	if gwi.opts.RequestOpts == nil {
		gwi.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gwi.opts.RequestOpts.APIURL = url.Std()

	return gwi
}

// Send gets the current webhook status.
func (gwi *GetWebhookInfo) Send() g.Result[*gotgbot.WebhookInfo] {
	return g.ResultOf(gwi.bot.Raw().GetWebhookInfo(gwi.opts))
}
