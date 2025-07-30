package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
	"github.com/enetx/tg/types/passport"
)

// SetPassportDataErrors represents a request to set passport data errors.
type SetPassportDataErrors struct {
	ctx    *Context
	userID int64
	errors g.Slice[gotgbot.PassportElementError]
	opts   *gotgbot.SetPassportDataErrorsOpts
}

// Errors sets the passport element errors using our passport error builders.
func (spde *SetPassportDataErrors) Errors(errors ...*passport.PassportError) *SetPassportDataErrors {
	spde.errors = passport.Errors(errors...)
	return spde
}

// Timeout sets a custom timeout for this request.
func (spde *SetPassportDataErrors) Timeout(duration time.Duration) *SetPassportDataErrors {
	if spde.opts.RequestOpts == nil {
		spde.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spde.opts.RequestOpts.Timeout = duration

	return spde
}

// APIURL sets a custom API URL for this request.
func (spde *SetPassportDataErrors) APIURL(url g.String) *SetPassportDataErrors {
	if spde.opts.RequestOpts == nil {
		spde.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	spde.opts.RequestOpts.APIURL = url.Std()

	return spde
}

// Send sets the passport data errors.
func (spde *SetPassportDataErrors) Send() g.Result[bool] {
	return g.ResultOf(spde.ctx.Bot.Raw().SetPassportDataErrors(spde.userID, spde.errors, spde.opts))
}
