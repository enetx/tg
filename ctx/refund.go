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
func (c *RefundStarPayment) UserID(id int64) *RefundStarPayment {
	c.userID = Some(id)
	return c
}

// Timeout sets a custom timeout for this request.
func (c *RefundStarPayment) Timeout(duration time.Duration) *RefundStarPayment {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.Timeout = duration

	return c
}

// APIURL sets a custom API URL for this request.
func (c *RefundStarPayment) APIURL(url String) *RefundStarPayment {
	if c.opts.RequestOpts == nil {
		c.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	c.opts.RequestOpts.APIURL = url.Std()

	return c
}

// Send processes the star payment refund and returns the result.
func (c *RefundStarPayment) Send() Result[bool] {
	userID := c.userID.UnwrapOr(c.ctx.EffectiveUser.Id)
	return ResultOf(c.ctx.Bot.Raw().RefundStarPayment(userID, c.chargeID.Std(), c.opts))
}
