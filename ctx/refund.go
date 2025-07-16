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

func (r *RefundStarPayment) UserID(id int64) *RefundStarPayment {
	r.userID = Some(id)
	return r
}

func (r *RefundStarPayment) Timeout(duration time.Duration) *RefundStarPayment {
	r.opts.RequestOpts = &gotgbot.RequestOpts{Timeout: duration}
	return r
}

func (r *RefundStarPayment) Send() Result[bool] {
	userID := r.userID.UnwrapOr(r.ctx.EffectiveUser.Id)
	return ResultOf(r.ctx.Bot.Raw().RefundStarPayment(userID, r.chargeID.Std(), r.opts))
}
