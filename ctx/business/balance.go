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
func (r *GetBalance) Timeout(duration time.Duration) *GetBalance {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.Timeout = duration

	return r
}

// APIURL sets a custom API URL for this request.
func (r *GetBalance) APIURL(url String) *GetBalance {
	if r.opts.RequestOpts == nil {
		r.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	r.opts.RequestOpts.APIURL = url.Std()

	return r
}

// Send executes the GetBalance request.
func (r *GetBalance) Send() Result[*gotgbot.StarAmount] {
	return ResultOf(r.bot.Raw().GetBusinessAccountStarBalance(
		r.connID.Std(),
		r.opts,
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
func (g *GetGifts) ExcludeUnsaved() *GetGifts {
	g.opts.ExcludeUnsaved = true
	return g
}

// ExcludeSaved excludes gifts saved to profile page.
func (g *GetGifts) ExcludeSaved() *GetGifts {
	g.opts.ExcludeSaved = true
	return g
}

// ExcludeUnlimited excludes unlimited gifts.
func (g *GetGifts) ExcludeUnlimited() *GetGifts {
	g.opts.ExcludeUnlimited = true
	return g
}

// ExcludeLimited excludes limited gifts.
func (g *GetGifts) ExcludeLimited() *GetGifts {
	g.opts.ExcludeLimited = true
	return g
}

// ExcludeUnique excludes unique gifts.
func (g *GetGifts) ExcludeUnique() *GetGifts {
	g.opts.ExcludeUnique = true
	return g
}

// SortByPrice sorts gifts by price instead of send date.
func (g *GetGifts) SortByPrice() *GetGifts {
	g.opts.SortByPrice = true
	return g
}

// Offset sets pagination offset.
func (g *GetGifts) Offset(offset String) *GetGifts {
	g.opts.Offset = offset.Std()
	return g
}

// Limit sets maximum gifts to return (1-100, defaults to 100).
func (g *GetGifts) Limit(limit int64) *GetGifts {
	g.opts.Limit = limit
	return g
}

// Timeout sets a custom timeout for this request.
func (g *GetGifts) Timeout(duration time.Duration) *GetGifts {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.Timeout = duration

	return g
}

// APIURL sets a custom API URL for this request.
func (g *GetGifts) APIURL(url String) *GetGifts {
	if g.opts.RequestOpts == nil {
		g.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	g.opts.RequestOpts.APIURL = url.Std()

	return g
}

// Send executes the Gifts request.
func (g *GetGifts) Send() Result[*gotgbot.OwnedGifts] {
	return ResultOf(g.bot.Raw().GetBusinessAccountGifts(
		g.connID.Std(),
		g.opts,
	))
}
