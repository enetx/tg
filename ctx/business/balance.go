package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// Balance provides balance management
type Balance struct {
	bot    Bot
	connID String
}

// Get creates a request to retrieve the current star balance.
func (b *Balance) Get() *GetBalance {
	return &GetBalance{
		bot:    b.bot,
		connID: b.connID,
		opts:   new(gotgbot.GetBusinessAccountStarBalanceOpts),
	}
}

// GetBalance request builder for star balance.
type GetBalance struct {
	bot    Bot
	connID String
	opts   *gotgbot.GetBusinessAccountStarBalanceOpts
}

// Timeout sets a custom timeout for this request.
func (gb *GetBalance) Timeout(duration time.Duration) *GetBalance {
	if gb.opts.RequestOpts == nil {
		gb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gb.opts.RequestOpts.Timeout = duration

	return gb
}

// APIURL sets a custom API URL for this request.
func (gb *GetBalance) APIURL(url String) *GetBalance {
	if gb.opts.RequestOpts == nil {
		gb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	gb.opts.RequestOpts.APIURL = url.Std()

	return gb
}

// Send executes the GetBalance request.
func (gb *GetBalance) Send() Result[*gotgbot.StarAmount] {
	return ResultOf(gb.bot.Raw().GetBusinessAccountStarBalance(
		gb.connID.Std(),
		gb.opts,
	))
}

// Transfer creates a request to transfer a specific amount of stars.
func (b *Balance) Transfer(amount int64) *Transfer {
	return &Transfer{
		bot:    b.bot,
		connID: b.connID,
		amount: amount,
		opts:   new(gotgbot.TransferBusinessAccountStarsOpts),
	}
}

// Transfer request builder for sending stars.
type Transfer struct {
	bot    Bot
	connID String
	amount int64
	opts   *gotgbot.TransferBusinessAccountStarsOpts
}

// Timeout sets a custom timeout for this request.
func (t *Transfer) Timeout(duration time.Duration) *Transfer {
	if t.opts.RequestOpts == nil {
		t.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	t.opts.RequestOpts.Timeout = duration

	return t
}

// APIURL sets a custom API URL for this request.
func (t *Transfer) APIURL(url String) *Transfer {
	if t.opts.RequestOpts == nil {
		t.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	t.opts.RequestOpts.APIURL = url.Std()

	return t
}

// Send executes the Transfer request.
func (t *Transfer) Send() Result[bool] {
	return ResultOf(t.bot.Raw().TransferBusinessAccountStars(
		t.connID.Std(),
		t.amount,
		t.opts,
	))
}

// GetGifts creates a request to retrieve owned gifts.
func (b *Balance) GetGifts() *GetGifts {
	return &GetGifts{
		bot:    b.bot,
		connID: b.connID,
		opts:   new(gotgbot.GetBusinessAccountGiftsOpts),
	}
}

// GetGifts request builder for retrieving business account gifts.
type GetGifts struct {
	bot    Bot
	connID String
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
func (ggs *GetGifts) Offset(offset String) *GetGifts {
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
func (ggs *GetGifts) APIURL(url String) *GetGifts {
	if ggs.opts.RequestOpts == nil {
		ggs.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ggs.opts.RequestOpts.APIURL = url.Std()

	return ggs
}

// Send executes the Gifts request.
func (ggs *GetGifts) Send() Result[*gotgbot.OwnedGifts] {
	return ResultOf(ggs.bot.Raw().GetBusinessAccountGifts(
		ggs.connID.Std(),
		ggs.opts,
	))
}
