package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetUserGifts represents a request to get user gifts.
type GetUserGifts struct {
	ctx    *Context
	userID int64
	opts   *gotgbot.GetUserGiftsOpts
}

// ExcludeUnlimited excludes gifts that can be purchased an unlimited number of times.
func (gug *GetUserGifts) ExcludeUnlimited() *GetUserGifts {
	gug.opts.ExcludeUnlimited = true
	return gug
}

// ExcludeLimitedUpgradable excludes limited gifts that can be upgraded to unique.
func (gug *GetUserGifts) ExcludeLimitedUpgradable() *GetUserGifts {
	gug.opts.ExcludeLimitedUpgradable = true
	return gug
}

// ExcludeLimitedNonUpgradable excludes limited gifts that can't be upgraded to unique.
func (gug *GetUserGifts) ExcludeLimitedNonUpgradable() *GetUserGifts {
	gug.opts.ExcludeLimitedNonUpgradable = true
	return gug
}

// ExcludeFromBlockchain excludes gifts assigned from the TON blockchain.
func (gug *GetUserGifts) ExcludeFromBlockchain() *GetUserGifts {
	gug.opts.ExcludeFromBlockchain = true
	return gug
}

// ExcludeUnique excludes unique gifts.
func (gug *GetUserGifts) ExcludeUnique() *GetUserGifts {
	gug.opts.ExcludeUnique = true
	return gug
}

// SortByPrice sorts results by gift price instead of send date.
func (gug *GetUserGifts) SortByPrice() *GetUserGifts {
	gug.opts.SortByPrice = true
	return gug
}

// Offset sets the pagination offset.
func (gug *GetUserGifts) Offset(offset g.String) *GetUserGifts {
	gug.opts.Offset = offset.Std()
	return gug
}

// Limit sets the maximum number of gifts to return (1–100).
func (gug *GetUserGifts) Limit(limit int64) *GetUserGifts {
	gug.opts.Limit = limit
	return gug
}

// Timeout sets a custom timeout for this request.
func (gug *GetUserGifts) Timeout(duration time.Duration) *GetUserGifts {
	if gug.opts.RequestOpts == nil {
		gug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gug.opts.RequestOpts.Timeout = duration

	return gug
}

// APIURL sets a custom API URL for this request.
func (gug *GetUserGifts) APIURL(url g.String) *GetUserGifts {
	if gug.opts.RequestOpts == nil {
		gug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gug.opts.RequestOpts.APIURL = url.Std()

	return gug
}

// Send executes the GetUserGifts request and returns the user's gifts.
func (gug *GetUserGifts) Send() g.Result[*gotgbot.OwnedGifts] {
	return g.ResultOf(gug.ctx.Bot.Raw().GetUserGifts(gug.userID, gug.opts))
}
