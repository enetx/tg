package ctx

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// ConvertGiftToStars is a request builder for converting gifts to stars.
type ConvertGiftToStars struct {
	ctx                  *Context
	businessConnectionID g.String
	ownedGiftID          g.String
	opts                 *gotgbot.ConvertGiftToStarsOpts
}

// Timeout sets a custom timeout for this request.
func (cgts *ConvertGiftToStars) Timeout(duration time.Duration) *ConvertGiftToStars {
	if cgts.opts.RequestOpts == nil {
		cgts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgts.opts.RequestOpts.Timeout = duration

	return cgts
}

// APIURL sets a custom API URL for this request.
func (cgts *ConvertGiftToStars) APIURL(url g.String) *ConvertGiftToStars {
	if cgts.opts.RequestOpts == nil {
		cgts.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	cgts.opts.RequestOpts.APIURL = url.Std()

	return cgts
}

// Send executes the ConvertGiftToStars request.
func (cgts *ConvertGiftToStars) Send() g.Result[bool] {
	return g.ResultOf(cgts.ctx.Bot.Raw().ConvertGiftToStars(
		cgts.businessConnectionID.Std(),
		cgts.ownedGiftID.Std(),
		cgts.opts,
	))
}
