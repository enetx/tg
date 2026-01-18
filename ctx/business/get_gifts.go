package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetGifts is a request builder for retrieving gifts from a business account.
// Use fluent methods to configure options before calling Send().
type GetGifts struct {
	bot    Bot
	connID g.String
	opts   *gotgbot.GetBusinessAccountGiftsOpts
}

// ExcludeUnsaved sets the request to exclude gifts that aren't saved to the account's profile page.
func (ggs *GetGifts) ExcludeUnsaved() *GetGifts {
	ggs.opts.ExcludeUnsaved = true
	return ggs
}

// ExcludeSaved sets the request to exclude gifts that are saved to the account's profile page.
func (ggs *GetGifts) ExcludeSaved() *GetGifts {
	ggs.opts.ExcludeSaved = true
	return ggs
}

// ExcludeUnlimited sets the request to exclude gifts that can be purchased an unlimited number of times.
func (ggs *GetGifts) ExcludeUnlimited() *GetGifts {
	ggs.opts.ExcludeUnlimited = true
	return ggs
}

// ExcludeLimitedUpgradable sets the request to exclude gifts that are limited but can be upgraded to unique.
func (ggs *GetGifts) ExcludeLimitedUpgradable() *GetGifts {
	ggs.opts.ExcludeLimitedUpgradable = true
	return ggs
}

// ExcludeLimitedNonUpgradable sets the request to exclude gifts that are limited and cannot be upgraded to unique.
func (ggs *GetGifts) ExcludeLimitedNonUpgradable() *GetGifts {
	ggs.opts.ExcludeLimitedNonUpgradable = true
	return ggs
}

// ExcludeUnique sets the request to exclude unique gifts.
func (ggs *GetGifts) ExcludeUnique() *GetGifts {
	ggs.opts.ExcludeUnique = true
	return ggs
}

// ExcludeFromBlockchain sets the request to exclude gifts that were assigned from the TON blockchain
// and cannot be resold or transferred in Telegram.
func (ggs *GetGifts) ExcludeFromBlockchain() *GetGifts {
	ggs.opts.ExcludeFromBlockchain = true
	return ggs
}

// SortByPrice sets the request to sort results by gift price instead of send date.
// Sorting is applied before pagination.
func (ggs *GetGifts) SortByPrice() *GetGifts {
	ggs.opts.SortByPrice = true
	return ggs
}

// Offset sets the pagination offset for the first gift entry to return.
// Use an empty string to get the first chunk of results.
func (ggs *GetGifts) Offset(offset g.String) *GetGifts {
	ggs.opts.Offset = offset.Std()
	return ggs
}

// Limit sets the maximum number of gifts to be returned (1–100, defaults to 100).
func (ggs *GetGifts) Limit(limit int64) *GetGifts {
	ggs.opts.Limit = limit
	return ggs
}

// Timeout sets a custom timeout for this request.
func (ggs *GetGifts) Timeout(duration time.Duration) *GetGifts {
	if ggs.opts.RequestOpts == nil {
		ggs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ggs.opts.RequestOpts.Timeout = duration
	return ggs
}

// APIURL sets a custom Telegram Bot API URL for this request.
func (ggs *GetGifts) APIURL(url g.String) *GetGifts {
	if ggs.opts.RequestOpts == nil {
		ggs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ggs.opts.RequestOpts.APIURL = url.Std()
	return ggs
}

// Send executes the request to retrieve gifts from the business account.
// Returns OwnedGifts wrapped in g.Result.
func (ggs *GetGifts) Send() g.Result[*gotgbot.OwnedGifts] {
	return g.ResultOf(ggs.bot.Raw().GetBusinessAccountGifts(
		ggs.connID.Std(),
		ggs.opts,
	))
}
