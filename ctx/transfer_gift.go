package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

// TransferGift is a request builder for transferring gifts.
type TransferGift struct {
	ctx                  *Context
	businessConnectionID String
	ownedGiftID          String
	newOwnerChatID       int64
	opts                 *gotgbot.TransferGiftOpts
}

// StarCount sets the amount of stars to pay for transfer from business balance.
func (tg *TransferGift) StarCount(count int64) *TransferGift {
	tg.opts.StarCount = count
	return tg
}

// Timeout sets a custom timeout for this request.
func (tg *TransferGift) Timeout(duration time.Duration) *TransferGift {
	if tg.opts.RequestOpts == nil {
		tg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	tg.opts.RequestOpts.Timeout = duration

	return tg
}

// APIURL sets a custom API URL for this request.
func (tg *TransferGift) APIURL(url String) *TransferGift {
	if tg.opts.RequestOpts == nil {
		tg.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	tg.opts.RequestOpts.APIURL = url.Std()

	return tg
}

// Send executes the TransferGift request.
func (tg *TransferGift) Send() Result[bool] {
	return ResultOf(tg.ctx.Bot.Raw().TransferGift(
		tg.businessConnectionID.Std(),
		tg.ownedGiftID.Std(),
		tg.newOwnerChatID,
		tg.opts,
	))
}
