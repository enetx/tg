package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// UpgradeGift is a request builder for upgrading gifts.
type UpgradeGift struct {
	ctx                  *Context
	businessConnectionID g.String
	ownedGiftID          g.String
	opts                 *gotgbot.UpgradeGiftOpts
}

// KeepOriginalDetails preserves original gift text, sender, and receiver.
func (ug *UpgradeGift) KeepOriginalDetails() *UpgradeGift {
	ug.opts.KeepOriginalDetails = true
	return ug
}

// StarCount sets the amount of stars to pay for upgrade from business balance.
func (ug *UpgradeGift) StarCount(count int64) *UpgradeGift {
	ug.opts.StarCount = count
	return ug
}

// Timeout sets a custom timeout for this request.
func (ug *UpgradeGift) Timeout(duration time.Duration) *UpgradeGift {
	if ug.opts.RequestOpts == nil {
		ug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ug.opts.RequestOpts.Timeout = duration

	return ug
}

// APIURL sets a custom API URL for this request.
func (ug *UpgradeGift) APIURL(url g.String) *UpgradeGift {
	if ug.opts.RequestOpts == nil {
		ug.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	ug.opts.RequestOpts.APIURL = url.Std()

	return ug
}

// Send executes the UpgradeGift request.
func (ug *UpgradeGift) Send() g.Result[bool] {
	return g.ResultOf(ug.ctx.Bot.Raw().UpgradeGift(
		ug.businessConnectionID.Std(),
		ug.ownedGiftID.Std(),
		ug.opts,
	))
}
