package business

import (
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/enetx/g"
)

// SetBio is a request builder for setting the business account bio.
type SetBio struct {
	account *Account
	opts    *gotgbot.SetBusinessAccountBioOpts
}

// Timeout sets a custom timeout for this request.
func (sb *SetBio) Timeout(duration time.Duration) *SetBio {
	if sb.opts.RequestOpts == nil {
		sb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sb.opts.RequestOpts.Timeout = duration

	return sb
}

// APIURL sets a custom API URL for this request.
func (sb *SetBio) APIURL(url g.String) *SetBio {
	if sb.opts.RequestOpts == nil {
		sb.opts.RequestOpts = new(gotgbot.RequestOpts)
	}

	sb.opts.RequestOpts.APIURL = url.Std()

	return sb
}

// Send executes the SetBio request.
func (sb *SetBio) Send() g.Result[bool] {
	return g.ResultOf(sb.account.bot.Raw().SetBusinessAccountBio(
		sb.account.connID.Std(),
		sb.opts,
	))
}
