package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/entities"
)

// GiftPremiumSubscription represents a request to gift premium subscription to a user.
type GiftPremiumSubscription struct {
	ctx        *Context
	userID     int64
	monthCount int64
	starCount  int64
	opts       *gotgbot.GiftPremiumSubscriptionOpts
}

// Text sets the text that will be shown along with the service message.
func (gps *GiftPremiumSubscription) Text(text g.String) *GiftPremiumSubscription {
	gps.opts.Text = text.Std()
	return gps
}

// HTML sets the text parse mode to HTML.
func (gps *GiftPremiumSubscription) HTML() *GiftPremiumSubscription {
	gps.opts.TextParseMode = "HTML"
	return gps
}

// Markdown sets the text parse mode to MarkdownV2.
func (gps *GiftPremiumSubscription) Markdown() *GiftPremiumSubscription {
	gps.opts.TextParseMode = "MarkdownV2"
	return gps
}

// Entities sets custom entities for the gift text.
func (gps *GiftPremiumSubscription) Entities(e *entities.Entities) *GiftPremiumSubscription {
	gps.opts.TextEntities = e.Std()
	return gps
}

// Timeout sets a custom timeout for this request.
func (gps *GiftPremiumSubscription) Timeout(duration time.Duration) *GiftPremiumSubscription {
	if gps.opts.RequestOpts == nil {
		gps.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gps.opts.RequestOpts.Timeout = duration

	return gps
}

// APIURL sets a custom API URL for this request.
func (gps *GiftPremiumSubscription) APIURL(url g.String) *GiftPremiumSubscription {
	if gps.opts.RequestOpts == nil {
		gps.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gps.opts.RequestOpts.APIURL = url.Std()

	return gps
}

// Send gifts the premium subscription to the user.
func (gps *GiftPremiumSubscription) Send() g.Result[bool] {
	return g.ResultOf(gps.ctx.Bot.Raw().GiftPremiumSubscription(
		gps.userID,
		gps.monthCount,
		gps.starCount,
		gps.opts,
	))
}
