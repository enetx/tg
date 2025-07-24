package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	. "github.com/enetx/g"
)

type PreCheckout struct {
	ctx  *Context
	ok   bool
	err  String
	opts *gotgbot.AnswerPreCheckoutQueryOpts
}

// Ok marks the pre-checkout query as successful.
func (pc *PreCheckout) Ok() *PreCheckout {
	pc.ok = true
	return pc
}

// Error marks the pre-checkout query as failed with the specified error message.
func (pc *PreCheckout) Error(text String) *PreCheckout {
	pc.ok = false
	pc.err = text

	return pc
}

// Timeout sets a custom timeout for this request.
func (pc *PreCheckout) Timeout(duration time.Duration) *PreCheckout {
	if pc.opts == nil {
		pc.opts = new(gotgbot.AnswerPreCheckoutQueryOpts)
	}

	if pc.opts.RequestOpts == nil {
		pc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pc.opts.RequestOpts.Timeout = duration

	return pc
}

// APIURL sets a custom API URL for this request.
func (pc *PreCheckout) APIURL(url String) *PreCheckout {
	if pc.opts == nil {
		pc.opts = new(gotgbot.AnswerPreCheckoutQueryOpts)
	}

	if pc.opts.RequestOpts == nil {
		pc.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	pc.opts.RequestOpts.APIURL = url.Std()

	return pc
}

// Send answers the pre-checkout query and returns the result.
func (pc *PreCheckout) Send() Result[bool] {
	query := pc.ctx.Update.PreCheckoutQuery
	if query == nil {
		return Err[bool](Errorf("no precheckout query"))
	}

	if !pc.ok {
		pc.opts = &gotgbot.AnswerPreCheckoutQueryOpts{ErrorMessage: pc.err.Std()}
	}

	return ResultOf(query.Answer(pc.ctx.Bot.Raw(), pc.ok, pc.opts))
}
