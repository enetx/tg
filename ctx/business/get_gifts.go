package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// GetGifts request builder for retrieving business account gifts.
type GetGifts struct {
	bot    Bot
	connID g.String
	opts   *gotgbot.GetBusinessAccountGiftsOpts
}

// ExcludeUnsaved excludes gifts not saved to profile page.
func (ggs *GetGifts) ExcludeUnsaved() *GetGifts {
	ggs.opts.ExcludeUnsaved = true
	return ggs
}

// ExcludeSaved excludes gifts saved to profile page.
func (ggs *GetGifts) ExcludeSaved() *GetGifts {
	ggs.opts.ExcludeSaved = true
	return ggs
}

// ExcludeUnlimited excludes unlimited gifts.
func (ggs *GetGifts) ExcludeUnlimited() *GetGifts {
	ggs.opts.ExcludeUnlimited = true
	return ggs
}

// ExcludeLimited excludes limited gifts.
func (ggs *GetGifts) ExcludeLimited() *GetGifts {
	ggs.opts.ExcludeLimited = true
	return ggs
}

// ExcludeUnique excludes unique gifts.
func (ggs *GetGifts) ExcludeUnique() *GetGifts {
	ggs.opts.ExcludeUnique = true
	return ggs
}

// SortByPrice sorts gifts by price instead of send date.
func (ggs *GetGifts) SortByPrice() *GetGifts {
	ggs.opts.SortByPrice = true
	return ggs
}

// Offset sets pagination offset.
func (ggs *GetGifts) Offset(offset g.String) *GetGifts {
	ggs.opts.Offset = offset.Std()
	return ggs
}

// Limit sets maximum gifts to return (1-100, defaults to 100).
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

// APIURL sets a custom API URL for this request.
func (ggs *GetGifts) APIURL(url g.String) *GetGifts {
	if ggs.opts.RequestOpts == nil {
		ggs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ggs.opts.RequestOpts.APIURL = url.Std()

	return ggs
}

// Send executes the Gifts request.
func (ggs *GetGifts) Send() g.Result[*gotgbot.OwnedGifts] {
	return g.ResultOf(ggs.bot.Raw().GetBusinessAccountGifts(
		ggs.connID.Std(),
		ggs.opts,
	))
}
