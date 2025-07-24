package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type RefundStarPayment struct {
	ctx      *Context
	userID   Option[int64]
	chargeID String
	opts     *gotgbot.RefundStarPaymentOpts
}

// UserID sets the user ID for the refund.
func (rsp *RefundStarPayment) UserID(id int64) *RefundStarPayment {
	rsp.userID = Some(id)
	return rsp
}

// Timeout sets a custom timeout for this request.
func (rsp *RefundStarPayment) Timeout(duration time.Duration) *RefundStarPayment {
	if rsp.opts.RequestOpts == nil {
		rsp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rsp.opts.RequestOpts.Timeout = duration

	return rsp
}

// APIURL sets a custom API URL for this request.
func (rsp *RefundStarPayment) APIURL(url String) *RefundStarPayment {
	if rsp.opts.RequestOpts == nil {
		rsp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rsp.opts.RequestOpts.APIURL = url.Std()

	return rsp
}

// Send processes the star payment refund and returns the result.
func (rsp *RefundStarPayment) Send() Result[bool] {
	userID := rsp.userID.UnwrapOr(rsp.ctx.EffectiveUser.Id)
	return ResultOf(rsp.ctx.Bot.Raw().RefundStarPayment(userID, rsp.chargeID.Std(), rsp.opts))
}
