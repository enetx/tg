package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

type RefundStarPayment struct {
	ctx      *Context
	userID   g.Option[int64]
	chargeID g.String
	opts     *gotgbot.RefundStarPaymentOpts
}

// UserID sets the user ID for the refund.
func (rsp *RefundStarPayment) UserID(id int64) *RefundStarPayment {
	rsp.userID = g.Some(id)
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
func (rsp *RefundStarPayment) APIURL(url g.String) *RefundStarPayment {
	if rsp.opts.RequestOpts == nil {
		rsp.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	rsp.opts.RequestOpts.APIURL = url.Std()

	return rsp
}

// Send processes the star payment refund and returns the result.
func (rsp *RefundStarPayment) Send() g.Result[bool] {
	userID := rsp.userID.UnwrapOr(rsp.ctx.EffectiveUser.Id)
	return g.ResultOf(rsp.ctx.Bot.Raw().RefundStarPayment(userID, rsp.chargeID.Std(), rsp.opts))
}
