package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

type AnswerPreCheckoutQuery struct {
	ctx  *Context
	ok   bool
	opts *gotgbot.AnswerPreCheckoutQueryOpts
}

// Ok marks the pre-checkout query as successful.
func (apcq *AnswerPreCheckoutQuery) Ok() *AnswerPreCheckoutQuery {
	apcq.ok = true
	return apcq
}

// Error marks the pre-checkout query as failed with the specified error message.
func (apcq *AnswerPreCheckoutQuery) Error(text g.String) *AnswerPreCheckoutQuery {
	apcq.ok = false
	apcq.opts.ErrorMessage = text.Std()

	return apcq
}

// Timeout sets a custom timeout for this request.
func (apcq *AnswerPreCheckoutQuery) Timeout(duration time.Duration) *AnswerPreCheckoutQuery {
	if apcq.opts.RequestOpts == nil {
		apcq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	apcq.opts.RequestOpts.Timeout = duration

	return apcq
}

// APIURL sets a custom API URL for this request.
func (apcq *AnswerPreCheckoutQuery) APIURL(url g.String) *AnswerPreCheckoutQuery {
	if apcq.opts.RequestOpts == nil {
		apcq.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	apcq.opts.RequestOpts.APIURL = url.Std()

	return apcq
}

// Send answers the pre-checkout query and returns the result.
func (apcq *AnswerPreCheckoutQuery) Send() g.Result[bool] {
	query := apcq.ctx.Update.PreCheckoutQuery
	if query == nil {
		return g.Err[bool](g.Errorf("no precheckout query"))
	}

	return g.ResultOf(query.Answer(apcq.ctx.Bot.Raw(), apcq.ok, apcq.opts))
}
